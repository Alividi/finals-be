package model

type Product struct {
	ID                 string `db:"id"`
	FkKategoriProdukId string `db:"kategori_produk_id"`
	NamaProduk         string `db:"nama_produk"`
	DeskripsiProduk    string `db:"deskripsi_produk"`
	GambarProduk       string `db:"gambar_produk"`
	SpesifikasiProduk  string `db:"spesifikasi_produk"`
}

type KategoriProduk struct {
	ID                 string `db:"id"`
	NamaKetegoriProduk string `db:"nama_kategori_produk"`
}
