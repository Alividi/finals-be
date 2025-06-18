package service

import (
	"context"
	"finals-be/app/services/dto"
	servicerepo "finals-be/app/services/repository"
	userrepo "finals-be/app/user/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"

	"github.com/rs/zerolog/log"
)

type ServiceService struct {
	cfg               *config.Config
	db                *connection.MultiInstruction
	serviceRepository servicerepo.IServiceRepository
	userRepository    userrepo.IUserRepository
}

func NewServiceService(cfg *config.Config, conn *connection.SQLServerConnectionManager, userRepo userrepo.IUserRepository) *ServiceService {
	db := conn.GetTransaction()
	return &ServiceService{
		cfg:               cfg,
		db:                db,
		serviceRepository: servicerepo.NewServiceRepository(db),
		userRepository:    userRepo,
	}
}

func (s *ServiceService) GetServices(ctx context.Context, req dto.ServicesRequest) (response []dto.GetServicesResponse, err error) {
	userId := req.UserId
	customerDetail, err := s.userRepository.GetCustomerDetail(ctx, userId)
	if err != nil {
		return nil, err
	}
	req.CustomerID = customerDetail.Customer.CustomerID

	services, err := s.serviceRepository.GetServices(ctx, req)
	if err != nil {
		return nil, err
	}

	response = make([]dto.GetServicesResponse, len(services))
	for i, service := range services {
		response[i] = dto.GetServicesResponse{
			ID:             service.ID,
			NamaService:    service.NamaService,
			AddressLine:    service.AddressLine,
			Active:         service.Active,
			DataUsage:      service.DataUsage,
			ActivationDate: service.ActivationDate,
		}
	}

	return response, nil
}

func (s *ServiceService) GetServiceDetail(ctx context.Context, id int64) (dto.GetServiceDetailResponse, error) {
	service, err := s.serviceRepository.GetServiceDetail(ctx, id)
	if err != nil {
		return dto.GetServiceDetailResponse{}, err
	}

	return dto.GetServiceDetailResponse{
		ID:                service.ID,
		ProductId:         service.FkProductId,
		CustomerId:        service.FkCustomerId,
		GangguanId:        service.FkGangguanId,
		NamaService:       service.NamaService,
		AddressLine:       service.AddressLine,
		Locality:          service.Locality,
		Latitude:          service.Latitude,
		Longitude:         service.Longitude,
		ServiceLineNumber: service.ServiceLineNumber,
		Nickname:          service.Nickname,
		Active:            service.Active,
		Ipkit:             service.Ipkit,
		KitSn:             service.KitSn,
		SSID:              service.SSID,
		ActivationDate:    service.ActivationDate,
		IsProblem:         service.IsProblem,
		Device:            service.Device,
		CustomerName:      service.CustomerName,
		DataUsage:         service.DataUsage,
	}, nil
}

func (s *ServiceService) GetServiceTelemetry(ctx context.Context, serviceId int64, req dto.ServiceTelemetryRequest) ([]dto.GetServiceTelemetryResponse, error) {
	telemetries, err := s.serviceRepository.GetServiceTelemetry(ctx, serviceId, req)
	if err != nil {
		return nil, err
	}

	response := make([]dto.GetServiceTelemetryResponse, len(telemetries))
	for i, t := range telemetries {
		response[i] = dto.GetServiceTelemetryResponse{
			ID:                t.ID,
			ServiceId:         t.FkServiceId,
			Timestamp:         t.Timestamp,
			DownlinkTroughput: t.DownlinkTroughput,
			UplinkTroughput:   t.UplinkTroughput,
			PingDropRate:      t.PingDropRate,
			Latency:           t.Latency,
			Obstruction:       t.Obstruction,
			Uptime:            t.Uptime,
			SignalQuality:     t.SignalQuality,
		}
	}

	return response, nil
}

func (s *ServiceService) ChangeServiceCoordinate(ctx context.Context, req dto.ChangeCoordinateRequest) error {
	log := log.Ctx(ctx).With().Str("service", "ChangeServiceCoordinate").Logger()

	err := s.serviceRepository.ChangeCoordinate(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to change service coordinate")
		return err
	}

	return nil
}

func (s *ServiceService) GetTroubleshootSteps(ctx context.Context, gangguanId int64) (dto.GetSolutionResponse, error) {
	problem, err := s.serviceRepository.GetProblemById(ctx, gangguanId)
	if err != nil {
		return dto.GetSolutionResponse{}, err
	}

	steps, err := s.serviceRepository.GetStepsByProblemId(ctx, gangguanId)
	if err != nil {
		return dto.GetSolutionResponse{}, err
	}

	var responseSteps []dto.Step
	for _, step := range steps {
		substeps, err := s.serviceRepository.GetSubstepsByStepId(ctx, step.ID)
		if err != nil {
			return dto.GetSolutionResponse{}, err
		}

		var dtoSubsteps []dto.Substep
		for _, ss := range substeps {
			dtoSubsteps = append(dtoSubsteps, dto.Substep{
				SubstepId: ss.ID,
				Substep:   ss.Substep,
				Gambar:    ss.Gambar,
				Deskripsi: ss.Deskripsi,
			})
		}

		responseSteps = append(responseSteps, dto.Step{
			StepId:     step.ID,
			Step:       step.Step,
			StepNumber: step.StepNumber,
			Substeps:   dtoSubsteps,
		})
	}

	return dto.GetSolutionResponse{
		GangguanId:        problem.ID,
		NamaGangguan:      problem.NamaGangguan,
		DeskripsiGangguan: problem.DeskripsiGangguan,
		Steps:             responseSteps,
	}, nil
}

func (s *ServiceService) GetServiceRepository() servicerepo.IServiceRepository {
	return s.serviceRepository
}
