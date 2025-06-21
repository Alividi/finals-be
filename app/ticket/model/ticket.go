package model

type Ticket struct {
	ID         int64  `json:"id" db:"id"`
	NomorTiket string `json:"nomor_tiket" db:"nomor_tiket"`
	ServiceId  int64  `json:"service_id" db:"service_id"`
	CustomerId int64  `json:"customer_id" db:"customer_id"`
	Status     string `json:"status" db:"status"`
	GangguanId int64  `json:"gangguan_id" db:"gangguan_id"`
	TeknisiId  *int64 `json:"teknisi_id" db:"teknisi_id"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}

type Ba struct {
	ID              int64  `db:"id" json:"id"`
	FkTicketID      int64  `db:"ticket_id" json:"ticket_id"`
	GambarPerangkat string `db:"gambar_perangkat" json:"gambar_perangkat"`
	GambarSpeedtest string `db:"gambar_speedtest" json:"gambar_speedtest"`
	DetailBa        string `db:"detail_ba" json:"detail_ba"`
}

type BiayaLainnya struct {
	ID         int64  `db:"id" json:"id"`
	FkBaID     int64  `db:"ba_id" json:"ba_id"`
	JenisBiaya string `db:"jenis_biaya" json:"jenis_biaya"`
	Jumlah     int64  `db:"jumlah" json:"jumlah"`
	Lampiran   string `db:"lampiran" json:"lampiran"`
}
