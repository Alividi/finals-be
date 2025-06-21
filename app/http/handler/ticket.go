package handler

import (
	dtoTicket "finals-be/app/ticket/dto"
	ticketService "finals-be/app/ticket/service"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type TicketHandler struct {
	ticketService *ticketService.TicketService
	validate      *validator.Validate
}

func NewTicketHandler(ticketService *ticketService.TicketService, validate *validator.Validate) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
		validate:      validate,
	}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request := dtoTicket.InsertTicketRequest{}

	userCtx := auth.GetUserContext(ctx)
	request.UserId = userCtx.ID

	err := helper.ReadRequest(r, &request)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	if err := h.validate.Struct(request); err != nil {
		helper.WriteResponse(r.Context(), w, helper.NewErrValidation(err), nil)
		return
	}

	err = h.ticketService.CreateTicket(ctx, request)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, nil, nil)
}

func (h *TicketHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tickets, err := h.ticketService.GetTickets(ctx)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	response := make([]dtoTicket.TicketsResponse, len(tickets))
	for i, ticket := range tickets {
		response[i] = dtoTicket.TicketsResponse{
			ID:             ticket.ID,
			Status:         ticket.Status,
			NomorTiket:     ticket.NomorTiket,
			NamaService:    ticket.NamaService,
			NamaPerusahaan: ticket.NamaPerusahaan,
			NamaTeknisi:    ticket.NamaTeknisi,
			AddressLine:    ticket.AddressLine,
			NamaGangguan:   ticket.NamaGangguan,
			CreatedAt:      ticket.CreatedAt,
		}
	}

	helper.WriteResponse(r.Context(), w, nil, response)
}

func (h *TicketHandler) GetTicketsSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	summary, err := h.ticketService.GetTicketsSummary(ctx)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, summary)
}

func (h *TicketHandler) GetTicketById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ticketId := helper.GetURLParamInt64(r, "ticketId")
	ticket, err := h.ticketService.GetTicketDetailById(ctx, ticketId)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, ticket)
}

func (h *TicketHandler) AssignTicket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request := dtoTicket.AsignTicketRequest{}

	err := helper.ReadRequest(r, &request)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	if err := h.validate.Struct(request); err != nil {
		helper.WriteResponse(r.Context(), w, helper.NewErrValidation(err), nil)
		return
	}

	err = h.ticketService.AssignTicket(ctx, request.TicketId, request.TeknisiId)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, nil, nil)
}

func (h *TicketHandler) CreateBa(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userCtx := auth.GetUserContext(ctx)

	// Parse form
	if err := r.ParseMultipartForm(32 * 1024 * 1024); err != nil {
		helper.WriteResponse(ctx, w, helper.NewErrBadRequest("failed to parse multipart form: "+err.Error()), nil)
		return
	}

	// Extract and validate basic fields
	ticketIDStr := helper.GetFormValue(r, "ticket_id")
	detailBa := helper.GetFormValue(r, "detail_ba")

	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		helper.WriteResponse(ctx, w, helper.NewErrBadRequest("invalid ticket_id"), nil)
		return
	}

	// Get main BA images
	gambarPerangkat, gambarPerangkatHeader, _ := r.FormFile("gambar_perangkat")
	gambarSpeedtest, gambarSpeedtestHeader, _ := r.FormFile("gambar_speedtest")

	// Handle biaya_lainnya
	biayaLainnya := []*dtoTicket.BiayaLainnyaRequest{}
	for i := 0; ; i++ {
		prefix := "biaya_lainnya[" + strconv.Itoa(i) + "]"
		if helper.GetFormValue(r, prefix+"[jenis_biaya]") == "" {
			break // Stop when there's no more entry
		}

		jumlahStr := helper.GetFormValue(r, prefix+"[jumlah]")
		jumlah, err := strconv.ParseInt(jumlahStr, 10, 64)
		if err != nil {
			helper.WriteResponse(ctx, w, helper.NewErrBadRequest("invalid jumlah at index "+strconv.Itoa(i)), nil)
			return
		}

		jenisBiaya := helper.GetFormValue(r, prefix+"[jenis_biaya]")
		file, fileHeader, _ := r.FormFile(prefix + "[lampiran]")

		biayaLainnya = append(biayaLainnya, &dtoTicket.BiayaLainnyaRequest{
			JenisBiaya:     jenisBiaya,
			Jumlah:         jumlah,
			Lampiran:       file,
			LampiranHeader: fileHeader,
		})
	}

	// Create payload
	payload := &dtoTicket.CreateBaRequest{
		TicketID:              ticketID,
		DetailBa:              detailBa,
		GambarPerangkat:       gambarPerangkat,
		GambarPerangkatHeader: gambarPerangkatHeader,
		GambarSpeedtest:       gambarSpeedtest,
		GambarSpeedtestHeader: gambarSpeedtestHeader,
		BiayaLainnya:          biayaLainnya,
	}

	// Validate required fields
	if err := h.validate.Struct(payload); err != nil {
		helper.WriteResponse(ctx, w, helper.NewErrValidation(err), nil)
		return
	}

	// Call service
	err = h.ticketService.CreateBa(ctx, userCtx.ID, payload)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, map[string]string{"message": "Berita Acara created successfully"})
}

func (h *TicketHandler) GetBaDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ticketId := helper.GetURLParamInt64(r, "ticketId")
	baDetail, err := h.ticketService.GetBaDetail(ctx, ticketId)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}
	helper.WriteResponse(ctx, w, nil, baDetail)
}
