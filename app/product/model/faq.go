package model

type FAQ struct {
	ID           int64  `db:"id"`
	FkKategoriId int64  `db:"kategori_produk_id"`
	Pertanyaan   string `db:"pertanyaan"`
	Jawaban      string `db:"jawaban"`
}
