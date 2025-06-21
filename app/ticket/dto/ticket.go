package dto

import "mime/multipart"

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

type BiayaLainnyaRequest struct {
	JenisBiaya     string                `json:"jenis_biaya" db:"jenis_biaya"`
	Jumlah         int64                 `json:"jumlah" db:"jumlah"`
	Lampiran       multipart.File        `json:"lampiran"`
	LampiranHeader *multipart.FileHeader `json:"lampiran_headers"`
}

type CreateBaRequest struct {
	TicketID              int64                  `json:"ticket_id" validate:"required"`
	GambarPerangkat       multipart.File         `json:"gambar_perangkat"`
	GambarPerangkatHeader *multipart.FileHeader  `json:"gambar_perangkat_headers"`
	GambarSpeedtest       multipart.File         `json:"gambar_speedtest"`
	GambarSpeedtestHeader *multipart.FileHeader  `json:"gambar_speedtest_headers"`
	DetailBa              string                 `json:"detail_ba" validate:"required"`
	BiayaLainnya          []*BiayaLainnyaRequest `json:"biaya_lainnya"`
}

type BiayaLainnya struct {
	JenisBiaya string `json:"jenis_biaya" db:"jenis_biaya"`
	Jumlah     int64  `json:"jumlah" db:"jumlah"`
	Lampiran   string `json:"lampiran" db:"lampiran"`
}

type BaDetailResponse struct {
	ID              int64           `json:"id" db:"id"`
	TicketID        int64           `json:"ticket_id" db:"ticket_id"`
	NomorTiket      string          `json:"nomor_tiket" db:"nomor_tiket"`
	GambarPerangkat string          `json:"gambar_perangkat" db:"gambar_perangkat"`
	GambarSpeedtest string          `json:"gambar_speedtest" db:"gambar_speedtest"`
	DetailBa        string          `json:"detail_ba" db:"detail_ba"`
	BiayaLainnya    []*BiayaLainnya `json:"biaya_lainnya" db:"biaya_lainnya"`
}
