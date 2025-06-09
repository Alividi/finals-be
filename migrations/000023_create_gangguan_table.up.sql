CREATE TABLE tbl_gangguan (
    id integer primary key,
    nama_gangguan varchar(100),
    deskripsi_gangguan text
);

ALTER TABLE tbl_nodelink
ADD COLUMN id_gangguan INTEGER REFERENCES tbl_gangguan(id);

CREATE INDEX gangguan_id_fk ON tbl_nodelink(id_gangguan);
