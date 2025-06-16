CREATE TABLE tbl_step (
    id integer PRIMARY KEY,
    gangguan_id integer,
    step varchar(100),
    step_number integer,
    FOREIGN KEY (gangguan_id) REFERENCES tbl_gangguan(id)
);

CREATE INDEX step_gangguan_id_fk ON tbl_step(gangguan_id);