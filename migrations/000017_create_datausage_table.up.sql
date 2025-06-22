CREATE TABLE tbl_datausage (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    service_id integer,
    data_usage float,
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_id) REFERENCES tbl_service(id)
);

CREATE INDEX datausage_service_id_fk ON tbl_datausage(service_id);
