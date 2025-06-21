package model

type Teknisi struct {
	ID     int64  `db:"id"`
	Nama   string `db:"nama"`
	Email  string `db:"email"`
	NoTelp string `db:"no_telp"`
	Status string `db:"status"`
	Base   string `db:"base"`
}
