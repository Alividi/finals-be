CREATE TABLE tbl_gangguan (
    id integer primary key,
    nama_gangguan varchar(100),
    deskripsi_gangguan text
);

ALTER TABLE tbl_service
ADD COLUMN gangguan_id INTEGER REFERENCES tbl_gangguan(id);

CREATE INDEX gangguan_id_fk ON tbl_service(gangguan_id);
