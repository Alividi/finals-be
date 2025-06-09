CREATE TABLE tbl_datausage (
    id integer PRIMARY KEY,
    service_line_number VARCHAR(50),
    data_usage float,
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_line_number) REFERENCES tbl_service_line(service_line_number)
);

CREATE INDEX datausage_service_line_number_fk ON tbl_datausage(service_line_number);