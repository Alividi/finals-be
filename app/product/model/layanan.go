package model

type Layanan struct {
	ID           string `db:"id"`
	FkProdukId   string `db:"produk_id"`
	NamaLayanan  string `db:"nama_layanan"`
	HargaLayanan int    `db:"harga_layanan"`
}
