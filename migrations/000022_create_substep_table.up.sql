CREATE TABLE tbl_substep (
    id integer PRIMARY KEY,
    step_id integer,
    substep varchar(100),
    gambar varchar(100),
    deskripsi text,
    FOREIGN KEY (step_id) REFERENCES tbl_step(id)
);

CREATE INDEX substep_step_id_fk ON tbl_substep(step_id);