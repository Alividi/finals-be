package model

import "time"

type Admin struct {
	ID       int64     `db:"id"`
	NIK      string    `db:"nik"`
	Nama     string    `db:"nama"`
	NPWP     string    `db:"npwp"`
	TglLahir time.Time `db:"tgl_lahir"`
}
