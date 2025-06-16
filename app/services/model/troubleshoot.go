package model

type Problem struct {
	ID                int64  `db:"id"`
	NamaGangguan      string `db:"nama_gangguan"`
	DeskripsiGangguan string `db:"deskripsi_gangguan"`
}

type Step struct {
	ID           int64  `db:"id"`
	FkGangguanId int64  `db:"gangguan_id"`
	Step         string `db:"step"`
	StepNumber   int64  `db:"step_number"`
}

type Substep struct {
	ID        int64  `db:"id"`
	FkStepId  int64  `db:"step_id"`
	Substep   string `db:"substep"`
	Gambar    string `db:"gambar"`
	Deskripsi string `db:"deskripsi"`
}
