package service

import (
	"context"
	"database/sql"
	"errors"
	servicerepo "finals-be/app/services/repository"
	shared "finals-be/app/shared/service"
	"finals-be/app/ticket/dto"
	"finals-be/app/ticket/model"
	tiketrepo "finals-be/app/ticket/repository"
	userrepo "finals-be/app/user/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/rs/zerolog/log"
)

type TicketService struct {
	cfg               *config.Config
	db                *connection.MultiInstruction
	ticketRepository  tiketrepo.ITicketRepository
	userRepository    userrepo.IUserRepository
	serviceRepository servicerepo.IServiceRepository
	s3Client          *shared.S3Service
}

func NewTicketService(cfg *config.Config, conn *connection.SQLServerConnectionManager, userRepo userrepo.IUserRepository,
	serviceRepo servicerepo.IServiceRepository, s3Client *shared.S3Service) *TicketService {
	db := conn.GetTransaction()
	return &TicketService{
		cfg:               cfg,
		db:                db,
		ticketRepository:  tiketrepo.NewTicketRepository(db),
		userRepository:    userRepo,
		serviceRepository: serviceRepo,
		s3Client:          s3Client,
	}
}

func (s *TicketService) CreateTicket(ctx context.Context, req dto.InsertTicketRequest) error {
	userId := req.UserId
	customerDetail, err := s.userRepository.GetCustomerDetail(ctx, userId)
	if err != nil {
		return err
	}
	req.CustomerID = customerDetail.Customer.CustomerID

	latestTicket, err := s.ticketRepository.GetLatestTicketByServiceId(ctx, req.ServiceId)
	if err == nil && latestTicket != nil {
		if latestTicket.Status != constants.TICKET_STATUS_CLOSED {
			return fmt.Errorf("cannot create ticket: service already has an open ticket (%s)", latestTicket.NomorTiket)
		}
	}

	ticket := &model.Ticket{
		NomorTiket: helper.GenerateTicketNumber(req.ServiceId, req.GangguanId, req.CustomerID),
		ServiceId:  req.ServiceId,
		CustomerId: req.CustomerID,
		Status:     constants.TICKET_STATUS_OPEN,
		GangguanId: req.GangguanId,
		TeknisiId:  nil,
	}

	err = s.ticketRepository.CreateTicket(ctx, ticket)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}

	log.Info().Msgf("Ticket created successfully with number: %s", ticket.NomorTiket)
	return nil
}

func (s *TicketService) GetTickets(ctx context.Context) ([]dto.TicketsResponse, error) {
	userCtx := auth.GetUserContext(ctx)
	userRole := userCtx.Role
	userId := userCtx.ID

	var tickets []model.Ticket

	switch userRole {
	case constants.ROLE_CUSTOMER:
		customer, err := s.userRepository.GetCustomerDetail(ctx, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get customer details: %w", err)
		}
		customerId := customer.Customer.CustomerID
		tickets, err = s.ticketRepository.GetTickets(ctx, userRole, userId, nil, &customerId)
		if err != nil {
			return nil, fmt.Errorf("failed to get tickets: %w", err)
		}

	case constants.ROLE_TEKNISI:
		teknisi, err := s.userRepository.GetTechnicianByUserId(ctx, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get technician details: %w", err)
		}
		teknisiId := teknisi.ID
		tickets, err = s.ticketRepository.GetTickets(ctx, userRole, userId, &teknisiId, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get tickets: %w", err)
		}
	default:
		ticketsDef, err := s.ticketRepository.GetTickets(ctx, userRole, userId, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get tickets: %w", err)
		}
		tickets = ticketsDef
	}

	var response []dto.TicketsResponse

	for _, ticket := range tickets {
		var teknisiName, customerName, serviceName, address, problem string

		if ticket.TeknisiId != nil {
			if teknisi, err := s.userRepository.GetTechnicianByTeknisiID(ctx, *ticket.TeknisiId); err == nil {
				teknisiName = teknisi.Nama
			} else {
				log.Error().Err(err).Msg("Failed to get technician details")
			}
		}

		if ticket.CustomerId != 0 {
			if customer, err := s.userRepository.GetCustomerByCustomerID(ctx, ticket.CustomerId); err == nil {
				customerName = customer.Customer.NamaPerusahaan
			} else {
				log.Error().Err(err).Msg("Failed to get customer details")
			}
		}

		if ticket.ServiceId != 0 {
			if service, err := s.serviceRepository.GetServiceDetail(ctx, ticket.ServiceId); err == nil {
				serviceName = service.NamaService
				address = service.AddressLine
			} else {
				log.Error().Err(err).Msg("Failed to get service details")
			}
		}

		if ticket.GangguanId != 0 {
			if gangguan, err := s.serviceRepository.GetProblemById(ctx, ticket.GangguanId); err == nil {
				problem = gangguan.NamaGangguan
			} else {
				log.Error().Err(err).Msg("Failed to get gangguan details")
			}
		}

		response = append(response, dto.TicketsResponse{
			ID:             ticket.ID,
			Status:         ticket.Status,
			NomorTiket:     ticket.NomorTiket,
			NamaService:    serviceName,
			NamaPerusahaan: customerName,
			NamaTeknisi:    &teknisiName,
			AddressLine:    address,
			NamaGangguan:   problem,
			CreatedAt:      ticket.CreatedAt,
		})
	}

	return response, nil
}

func (s *TicketService) GetTicketsSummary(ctx context.Context) (*dto.TicketsSummaryResponse, error) {
	userCtx := auth.GetUserContext(ctx)
	userRole := userCtx.Role
	userId := userCtx.ID

	var (
		customerId *int64
		teknisiId  *int64
	)

	switch userRole {
	case constants.ROLE_CUSTOMER:
		customer, err := s.userRepository.GetCustomerDetail(ctx, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get customer details: %w", err)
		}
		id := customer.Customer.CustomerID
		customerId = &id

	case constants.ROLE_TEKNISI:
		teknisi, err := s.userRepository.GetTechnicianByUserId(ctx, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get technician details: %w", err)
		}
		id := teknisi.ID
		teknisiId = &id
	}

	summary, err := s.ticketRepository.GetTicketsSummary(ctx, userRole, userId, teknisiId, customerId)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket summary: %w", err)
	}

	return summary, nil
}

func (s *TicketService) GetTicketDetailById(ctx context.Context, ticketId int64) (*dto.TicketDetailResponse, error) {
	ticket, err := s.ticketRepository.GetTicketById(ctx, ticketId)
	if err != nil {
		return nil, fmt.Errorf("ticket not found: %w", err)
	}

	var (
		serviceName  string
		ipKit        string
		kitSN        string
		ssid         string
		customerName string
		address      string
		teknisiName  *string
		namaGangguan string
	)

	if ticket.ServiceId != 0 {
		service, err := s.serviceRepository.GetServiceDetail(ctx, ticket.ServiceId)
		if err == nil {
			serviceName = service.NamaService
			ipKit = service.Ipkit
			kitSN = service.KitSn
			ssid = service.SSID
			address = service.AddressLine
		}
	}

	if ticket.CustomerId != 0 {
		customer, err := s.userRepository.GetCustomerByCustomerID(ctx, ticket.CustomerId)
		if err == nil {
			customerName = customer.Customer.NamaPerusahaan
		}
	}

	if ticket.TeknisiId != nil {
		teknisi, err := s.userRepository.GetTechnicianByTeknisiID(ctx, *ticket.TeknisiId)
		if err == nil {
			name := teknisi.Nama
			teknisiName = &name
		}
	}

	if ticket.GangguanId != 0 {
		gangguan, err := s.serviceRepository.GetProblemById(ctx, ticket.GangguanId)
		if err == nil {
			namaGangguan = gangguan.NamaGangguan
		}
	}

	return &dto.TicketDetailResponse{
		ID:             ticket.ID,
		Status:         ticket.Status,
		NomorTiket:     ticket.NomorTiket,
		NamaService:    serviceName,
		Ipkit:          ipKit,
		KitSn:          kitSN,
		SSID:           ssid,
		NamaPerusahaan: customerName,
		NamaTeknisi:    teknisiName,
		AddressLine:    address,
		NamaGangguan:   namaGangguan,
		CreatedAt:      ticket.CreatedAt,
	}, nil
}

func (s *TicketService) AssignTicket(ctx context.Context, ticketId int64, teknisiId int64) error {
	if teknisiId == 0 {
		return errors.New("teknisi_id is required")
	}

	ticket, err := s.ticketRepository.GetTicketById(ctx, ticketId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("ticket with ID %d not found", ticketId)
		}
		return fmt.Errorf("failed to get ticket: %w", err)
	}

	if ticket.Status != constants.TICKET_STATUS_OPEN {
		return fmt.Errorf("ticket with ID %d is already assigned", ticketId)
	}

	err = s.ticketRepository.AssignTicket(ctx, ticketId, teknisiId)
	if err != nil {
		return fmt.Errorf("failed to assign ticket: %w", err)
	}

	log.Info().Msgf("Ticket with ID %d assigned to teknisi with ID %d", ticketId, teknisiId)
	return nil
}

func (s *TicketService) CreateBa(ctx context.Context, userID int64, req *dto.CreateBaRequest) error {
	// Start transaction
	if err := s.db.Begin(ctx); err != nil {
		return helper.NewErrInternalServer("failed to begin transaction: " + err.Error())
	}
	defer func() {
		_ = s.db.Rollback(ctx) // safe rollback if commit wasn't called
	}()

	// Upload main images
	gambarPerangkatURL, err := s.uploadSingleFile(ctx, "ba_gambar_perangkat", userID, req.GambarPerangkatHeader, req.GambarPerangkat)
	if err != nil {
		return err
	}
	gambarSpeedtestURL, err := s.uploadSingleFile(ctx, "ba_gambar_speedtest", userID, req.GambarSpeedtestHeader, req.GambarSpeedtest)
	if err != nil {
		return err
	}

	// Create BA
	ba := &model.Ba{
		FkTicketID:      req.TicketID,
		GambarPerangkat: gambarPerangkatURL,
		GambarSpeedtest: gambarSpeedtestURL,
		DetailBa:        req.DetailBa,
	}

	baID, err := s.ticketRepository.CreateBa(ctx, ba)
	if err != nil {
		return helper.NewErrInternalServer("failed to create BA: " + err.Error())
	}

	// Prepare biaya lainnya
	var biayaModels []*model.BiayaLainnya
	for _, biaya := range req.BiayaLainnya {
		lampiranURL, err := s.uploadSingleFile(ctx, "ba_biaya_lainnya", userID, biaya.LampiranHeader, biaya.Lampiran)
		if err != nil {
			return err
		}

		biayaModels = append(biayaModels, &model.BiayaLainnya{
			FkBaID:     baID,
			JenisBiaya: biaya.JenisBiaya,
			Jumlah:     biaya.Jumlah,
			Lampiran:   lampiranURL,
		})
	}

	// Bulk insert biaya lainnya
	if len(biayaModels) > 0 {
		if err := s.ticketRepository.BulkInsertBiayaLainnya(ctx, baID, biayaModels); err != nil {
			return helper.NewErrInternalServer("failed to insert biaya lainnya: " + err.Error())
		}
	}

	// Commit transaction
	if err := s.db.Commit(ctx); err != nil {
		return helper.NewErrInternalServer("failed to commit transaction: " + err.Error())
	}

	return nil
}

func (s *TicketService) uploadSingleFile(ctx context.Context, fileType string, userID int64, header *multipart.FileHeader, file multipart.File) (string, error) {
	if file == nil || header == nil {
		return "", helper.NewErrBadRequest(fmt.Sprintf("missing file for %s", fileType))
	}

	defer file.Close()

	// Read content type
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return "", helper.NewErrBadRequest("failed to read file: " + err.Error())
	}
	contentType := http.DetectContentType(buf)

	// Rewind file
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", helper.NewErrInternalServer("failed to rewind file: " + err.Error())
	}

	// Upload
	url, err := s.s3Client.UploadFile(ctx, header.Filename, fileType, userID, file, contentType)
	if err != nil {
		return "", helper.NewErrInternalServer("failed to upload to S3: " + err.Error())
	}

	return url, nil
}

func (s *TicketService) GetBaDetail(ctx context.Context, ticketID int64) (*dto.BaDetailResponse, error) {
	return s.ticketRepository.GetBaDetailByTicketID(ctx, ticketID)
}
