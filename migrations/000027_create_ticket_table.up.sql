CREATE TABLE tbl_ticket (
    id integer PRIMARY KEY,
    nomor_ticket varchar(100),
    service_id integer,
    status varchar(50),
    customer_id integer,
    gangguan_id integer,
    teknisi_id integer,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_id) REFERENCES tbl_service(id),
    FOREIGN KEY (customer_id) REFERENCES tbl_customer(id),
    FOREIGN KEY (gangguan_id) REFERENCES tbl_gangguan(id),
    FOREIGN KEY (teknisi_id) REFERENCES tbl_teknisi(id)
);

CREATE INDEX ticket_service_id_fk ON tbl_ticket(service_id);
CREATE INDEX ticket_customer_id_fk ON tbl_ticket(customer_id);
CREATE INDEX ticket_gangguan_id_fk ON tbl_ticket(gangguan_id);  
CREATE INDEX ticket_teknisi_id_fk ON tbl_ticket(teknisi_id);
