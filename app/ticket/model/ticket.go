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
