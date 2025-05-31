package model

type FAQ struct {
	ID           string `db:"id"`
	FkKategoriId string `db:"kategori_produk_id"`
	Pertanyaan   string `db:"pertanyaan"`
	Jawaban      string `db:"jawaban"`
}
