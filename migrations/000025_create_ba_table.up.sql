CREATE TABLE tbl_ba (
    id integer PRIMARY KEY,
    ticket_id integer,
    gambar_perangkat varchar(100),
    gambar_speedtest varchar(100),
    detail_ba text,
    FOREIGN KEY (ticket_id) REFERENCES tbl_ticket(id)
);

CREATE INDEX ba_ticket_id_fk ON tbl_ba(ticket_id);