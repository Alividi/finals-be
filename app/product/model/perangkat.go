package model

type Perangkat struct {
	ID              int64  `db:"id"`
	FkKategoriId    int64  `db:"kategori_produk_id"`
	FkProdukId      int64  `db:"produk_id"`
	NamaProduk      string `db:"nama_produk"`
	DeskripsiProduk string `db:"deskripsi_produk"`
	HargaProduk     int    `db:"harga_produk"`
	GambarProduk    string `db:"gambar_produk"`
}
