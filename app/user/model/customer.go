package model

type Customer struct {
	ID               string  `db:"id"`
	NamaPerusahaan   string  `db:"nama_perusahaan"`
	EmailPerusahaan  string  `db:"email_perusahaan"`
	NoTelpPerusahaan string  `db:"no_telp_perusahaan"`
	NoNpwpPerusahaan string  `db:"no_npwp_perusahaan"`
	Provinsi         string  `db:"provinsi"`
	Kabupaten        string  `db:"kabupaten"`
	Kecamatan        string  `db:"kecamatan"`
	Kelurahan        string  `db:"kelurahan"`
	Alamat           string  `db:"alamat"`
	Latitude         float64 `db:"latitude"`
	Longitude        float64 `db:"longitude"`
}

type Alamat struct {
	ID         string  `db:"id"`
	CustomerID string  `db:"customer_id"`
	Provinsi   string  `db:"provinsi"`
	Kabupaten  string  `db:"kabupaten"`
	Kecamatan  string  `db:"kecamatan"`
	Kelurahan  string  `db:"kelurahan"`
	RT         string  `db:"rt"`
	RW         string  `db:"rw"`
	Alamat     string  `db:"alamat"`
	Latitude   float64 `db:"latitude"`
	Longitude  float64 `db:"longitude"`
}
