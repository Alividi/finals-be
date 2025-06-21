CREATE TABLE tbl_biaya_lainnya (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ba_id integer,
    jenis_biaya varchar(100),
    jumlah integer,
    lampiran text,
    FOREIGN KEY (ba_id) REFERENCES tbl_ba(id)
);

CREATE INDEX biaya_lainnya_ba_id_fk ON tbl_biaya_lainnya(ba_id);