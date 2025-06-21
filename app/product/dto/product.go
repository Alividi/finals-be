package dto

type Perangkat struct {
	ID              int64  `json:"id" db:"id"`
	FkKategoriId    int64  `json:"kategori_produk_id" db:"kategori_produk_id"`
	NamaProduk      string `json:"nama_produk" db:"nama_produk"`
	DeskripsiProduk string `json:"deskripsi_produk" db:"deskripsi_produk"`
	HargaProduk     int    `json:"harga_produk" db:"harga_produk"`
	GambarProduk    string `json:"gambar_produk" db:"gambar_produk"`
}

type Layanan struct {
	NamaLayanan  string `json:"nama_layanan" db:"nama_layanan"`
	HargaLayanan int    `json:"harga_layanan" db:"harga_layanan"`
}

type GetProductDetailResponse struct {
	ID          int64       `json:"id" db:"id"`
	Nama        string      `json:"nama" db:"nama_produk"`
	Spesifikasi string      `json:"spesifikasi" db:"spesifikasi_produk"`
	Image       string      `json:"image" db:"gambar_produk"`
	Perangkat   []Perangkat `json:"perangkat"`
	Layanan     []Layanan   `json:"layanan"`
}

type GetProductsResponse struct {
	ID        int64  `json:"id" db:"id"`
	Nama      string `json:"nama" db:"nama_produk"`
	Deskripsi string `json:"deskripsi" db:"deskripsi_produk"`
	Image     string `json:"image" db:"gambar_produk"`
}

type GetFaqResponse struct {
	ID                int64  `json:"id" db:"id"`
	KategoriProductId int64  `json:"kategori_produk_id" db:"kategori_produk_id"`
	Pertanyaan        string `json:"pertanyaan" db:"pertanyaan"`
	Jawaban           string `json:"jawaban" db:"jawaban"`
}
