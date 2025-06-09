CREATE TABLE tbl_service_line (
    service_line_number VARCHAR(50) PRIMARY KEY,
    address_reference_id integer,
    nickname VARCHAR(255),
    active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (address_reference_id) REFERENCES tbl_address_reference(id)
);

CREATE INDEX service_line_address_reference_id_fk ON tbl_service_line(address_reference_id);
