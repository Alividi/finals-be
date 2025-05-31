package model

type Perangkat struct {
	ID              string `db:"id"`
	FkKategoriId    string `db:"kategori_produk_id"`
	FkProdukId      string `db:"produk_id"`
	NamaProduk      string `db:"nama_produk"`
	DeskripsiProduk string `db:"deskripsi_produk"`
	HargaProduk     int    `db:"harga_produk"`
	GambarProduk    string `db:"gambar_produk"`
}
