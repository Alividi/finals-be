CREATE TABLE tbl_ba (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ticket_id integer,
    gambar_perangkat text,
    gambar_speedtest text,
    detail_ba text,
    FOREIGN KEY (ticket_id) REFERENCES tbl_ticket(id)
);

CREATE INDEX ba_ticket_id_fk ON tbl_ba(ticket_id);