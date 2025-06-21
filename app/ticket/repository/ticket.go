package repository

import (
	"context"
	"finals-be/app/ticket/dto"
	"finals-be/app/ticket/model"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"fmt"
	"strings"
)

type ITicketRepository interface {
	GetLatestTicketByServiceId(ctx context.Context, serviceId int64) (*model.Ticket, error)
	CreateTicket(ctx context.Context, ticket *model.Ticket) (err error)
	GetTickets(ctx context.Context, role string, userId int64, teknisiId *int64, customerId *int64) ([]model.Ticket, error)
	GetTicketsSummary(ctx context.Context, role string, userId int64, teknisiId *int64, customerId *int64) (*dto.TicketsSummaryResponse, error)
	GetTicketById(ctx context.Context, ticketId int64) (*model.Ticket, error)
	AssignTicket(ctx context.Context, ticketId, teknisiId int64) error
	CreateBa(ctx context.Context, ticket *model.Ba) (int64, error)
	BulkInsertBiayaLainnya(ctx context.Context, baID int64, biaya []*model.BiayaLainnya) error
	GetBaDetailByTicketID(ctx context.Context, ticketID int64) (*dto.BaDetailResponse, error)
}

type TicketRepository struct {
	db connection.Connection
}

func NewTicketRepository(db connection.Connection) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (r *TicketRepository) GetLatestTicketByServiceId(ctx context.Context, serviceId int64) (*model.Ticket, error) {
	query := fmt.Sprintf(`SELECT id, nomor_tiket, service_id, customer_id, status, 
		gangguan_id, teknisi_id, created_at, updated_at 
		FROM %s WHERE service_id = $1 
		ORDER BY created_at DESC LIMIT 1`, constants.TABLE_TICKET)

	row := r.db.QueryRow(ctx, query, serviceId)

	var ticket model.Ticket
	err := row.Scan(&ticket.ID, &ticket.NomorTiket, &ticket.ServiceId, &ticket.CustomerId,
		&ticket.Status, &ticket.GangguanId, &ticket.TeknisiId, &ticket.CreatedAt, &ticket.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (r *TicketRepository) CreateTicket(ctx context.Context, ticket *model.Ticket) (err error) {
	query := fmt.Sprintf(
		`INSERT INTO %s 
		(nomor_tiket, service_id, customer_id, status, 
		gangguan_id, teknisi_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id`,
		constants.TABLE_TICKET)

	_, err = r.db.Exec(ctx, query, ticket.NomorTiket, ticket.ServiceId, ticket.CustomerId, ticket.Status, ticket.GangguanId, ticket.TeknisiId)
	if err != nil {
		return err
	}

	return
}

func (r *TicketRepository) GetTickets(ctx context.Context, role string, userId int64, teknisiId *int64, customerId *int64) ([]model.Ticket, error) {
	baseQuery := fmt.Sprintf(`SELECT id, nomor_tiket, service_id, customer_id, status, 
		gangguan_id, teknisi_id, created_at, updated_at FROM %s`, constants.TABLE_TICKET)

	var (
		query string
		args  []interface{}
	)

	switch role {
	case constants.ROLE_TEKNISI:
		query = baseQuery + ` WHERE teknisi_id = $1`
		args = append(args, teknisiId)
	case constants.ROLE_CUSTOMER:
		query = baseQuery + ` WHERE customer_id = $1`
		args = append(args, customerId)
	default:
		query = baseQuery
	}

	query = r.db.Rebind(query)

	var tickets []model.Ticket
	err := r.db.Select(ctx, &tickets, query, args...)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepository) GetTicketsSummary(ctx context.Context, role string, userId int64, teknisiId *int64, customerId *int64) (*dto.TicketsSummaryResponse, error) {
	query := fmt.Sprintf(`
		SELECT 
			COUNT(*) FILTER (WHERE status = 'open') AS open_count,
			COUNT(*) FILTER (WHERE status = 'in_progress') AS in_progress_count,
			COUNT(*) FILTER (WHERE status = 'closed') AS closed_count,
			COUNT(*) AS total_count
		FROM %s
	`, constants.TABLE_TICKET)

	var args []interface{}

	switch role {
	case constants.ROLE_TEKNISI:
		query += ` WHERE teknisi_id = $1`
		args = append(args, teknisiId)
	case constants.ROLE_CUSTOMER:
		query += ` WHERE customer_id = $1`
		args = append(args, customerId)
	}

	row := r.db.QueryRow(ctx, query, args...)

	var summary dto.TicketsSummaryResponse
	err := row.Scan(&summary.OpenCount, &summary.InProgressCount, &summary.ClosedCount, &summary.TotalCount)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *TicketRepository) GetTicketById(ctx context.Context, ticketId int64) (*model.Ticket, error) {
	query := fmt.Sprintf(`SELECT id, nomor_tiket, service_id, customer_id, status, 
		gangguan_id, teknisi_id, created_at, updated_at 
		FROM %s WHERE id = $1`, constants.TABLE_TICKET)

	row := r.db.QueryRow(ctx, query, ticketId)

	var ticket model.Ticket
	err := row.Scan(&ticket.ID, &ticket.NomorTiket, &ticket.ServiceId, &ticket.CustomerId,
		&ticket.Status, &ticket.GangguanId, &ticket.TeknisiId, &ticket.CreatedAt, &ticket.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (r *TicketRepository) AssignTicket(ctx context.Context, ticketId, teknisiId int64) error {
	query := fmt.Sprintf(`
	UPDATE %s SET teknisi_id = $1, status = $2, 
	updated_at = NOW() WHERE id = $3`,
		constants.TABLE_TICKET)

	_, err := r.db.Exec(ctx, query, teknisiId, constants.TICKET_STATUS_IN_PROGRESS, ticketId)
	return err
}

func (r *TicketRepository) CreateBa(ctx context.Context, ba *model.Ba) (int64, error) {
	query := fmt.Sprintf(`
	INSERT INTO %s (ticket_id, gambar_perangkat, gambar_speedtest, detail_ba)
	VALUES ($1, $2, $3, $4)
	RETURNING id`,
		constants.TABLE_BA)

	query = r.db.RebindTxx(query)

	var id int64
	row := r.db.QueryRowTxx(
		ctx,
		query,
		ba.FkTicketID,
		ba.GambarPerangkat,
		ba.GambarSpeedtest,
		ba.DetailBa,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TicketRepository) BulkInsertBiayaLainnya(ctx context.Context, baID int64, biaya []*model.BiayaLainnya) error {
	if len(biaya) == 0 {
		return nil
	}
	valueStrings := make([]string, 0, len(biaya))
	args := make([]interface{}, 0, len(biaya)*4)

	for i, b := range biaya {
		b.FkBaID = baID
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		args = append(args, b.FkBaID, b.JenisBiaya, b.Jumlah, b.Lampiran)
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (ba_id, jenis_biaya, jumlah, lampiran)
		VALUES %s`,
		constants.TABLE_BIAYA_LAINNYA,
		strings.Join(valueStrings, ", "),
	)

	query = r.db.RebindTxx(query)

	if _, err := r.db.ExecTxx(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *TicketRepository) GetBaDetailByTicketID(ctx context.Context, ticketID int64) (*dto.BaDetailResponse, error) {
	query := fmt.Sprintf(`
		SELECT 
			ba.id,
			ba.ticket_id,
			t.nomor_tiket,
			ba.gambar_perangkat,
			ba.gambar_speedtest,
			ba.detail_ba
		FROM %s ba
		JOIN %s t ON ba.ticket_id = t.id
		WHERE ba.ticket_id = $1
	`, constants.TABLE_BA, constants.TABLE_TICKET)

	var ba dto.BaDetailResponse
	if err := r.db.Get(ctx, &ba, query, ticketID); err != nil {
		return nil, err
	}

	// fetch biaya lainnya
	biayaQuery := fmt.Sprintf(`
		SELECT 
			jenis_biaya,
			jumlah,
			lampiran
		FROM %s
		WHERE ba_id = $1
	`, constants.TABLE_BIAYA_LAINNYA)

	var biaya []*dto.BiayaLainnya
	if err := r.db.Select(ctx, &biaya, biayaQuery, ba.ID); err != nil {
		return nil, err
	}

	ba.BiayaLainnya = biaya
	return &ba, nil
}
