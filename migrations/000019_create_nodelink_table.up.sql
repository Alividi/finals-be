CREATE TABLE tbl_nodelink (
    id integer PRIMARY KEY,
    custumer_id integer,
    service_id integer,
    service_line_number VARCHAR(50),
    is_problem BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (custumer_id) REFERENCES tbl_produk(id),
    FOREIGN KEY (service_id) REFERENCES tbl_service(id),
    FOREIGN KEY (service_line_number) REFERENCES tbl_service_line(service_line_number)
);

CREATE INDEX nodelink_custumer_id_fk ON tbl_nodelink(custumer_id);
CREATE INDEX nodelink_service_id_fk ON tbl_nodelink(service_id);
CREATE INDEX nodelink_service_line_number_fk ON tbl_nodelink(service_line_number);

