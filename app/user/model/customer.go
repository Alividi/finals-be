package model

type Customer struct {
	CustomerID       int64  `db:"id"`
	FkUserId         int64  `db:"user_id"`
	NamaPerusahaan   string `db:"nama_perusahaan"`
	EmailPerusahaan  string `db:"email_perusahaan"`
	NoTelpPerusahaan string `db:"no_telp_perusahaan"`
	NoNpwpPerusahaan string `db:"no_npwp_perusahaan"`
}

type Alamat struct {
	AlamatID     int64   `db:"id"`
	FkCustomerID int64   `db:"customer_id"`
	Provinsi     string  `db:"provinsi"`
	Kabupaten    string  `db:"kabupaten"`
	Kecamatan    string  `db:"kecamatan"`
	Kelurahan    string  `db:"kelurahan"`
	RT           string  `db:"rt"`
	RW           string  `db:"rw"`
	Alamat       string  `db:"alamat"`
	Latitude     float64 `db:"latitude"`
	Longitude    float64 `db:"longitude"`
}

type CustomerWithAlamat struct {
	Customer
	Alamat Alamat
}
