package dto

type InsertTicketRequest struct {
	UserId     int64 `json:"user_id" validate:"required"`
	CustomerID int64 `json:"-"`
	ServiceId  int64 `json:"service_id" validate:"required"`
	GangguanId int64 `json:"gangguan_id" validate:"required"`
}

type TicketsResponse struct {
	ID             int64   `json:"id" db:"id"`
	Status         string  `json:"status" db:"status"`
	NomorTiket     string  `json:"nomor_tiket" db:"nomor_tiket"`
	NamaService    string  `json:"nama_service" db:"nama_service"`
	NamaPerusahaan string  `json:"nama_perusahaan" db:"nama_perusahaan"`
	NamaTeknisi    *string `json:"nama_teknisi" db:"nama"`
	AddressLine    string  `json:"address_line" db:"address_line"`
	NamaGangguan   string  `json:"nama_gangguan" db:"nama_gangguan"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
}

type TicketsSummaryResponse struct {
	OpenCount       int64 `json:"open_count"`
	InProgressCount int64 `json:"in_progress_count"`
	ClosedCount     int64 `json:"closed_count"`
	TotalCount      int64 `json:"total_count"`
}

type TicketDetailResponse struct {
	ID             int64   `json:"id" db:"id"`
	Status         string  `json:"status" db:"status"`
	NomorTiket     string  `json:"nomor_tiket" db:"nomor_tiket"`
	NamaService    string  `json:"nama_service" db:"nama_service"`
	Ipkit          string  `json:"ip_kit" db:"ip_kit"`
	KitSn          string  `json:"kit_sn" db:"kit_sn"`
	SSID           string  `json:"ssid" db:"ssid"`
	NamaPerusahaan string  `json:"nama_perusahaan" db:"nama_perusahaan"`
	NamaTeknisi    *string `json:"nama_teknisi" db:"nama"`
	AddressLine    string  `json:"address_line" db:"address_line"`
	NamaGangguan   string  `json:"nama_gangguan" db:"nama_gangguan"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
}

type AsignTicketRequest struct {
	TicketId  int64 `json:"ticket_id" validate:"required"`
	TeknisiId int64 `json:"teknisi_id" validate:"required"`
}
