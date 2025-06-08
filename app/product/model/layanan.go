package model

type Layanan struct {
	ID           int64  `db:"id"`
	FkProdukId   int64  `db:"produk_id"`
	NamaLayanan  string `db:"nama_layanan"`
	HargaLayanan int    `db:"harga_layanan"`
}
