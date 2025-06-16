CREATE TABLE tbl_service (
    id integer PRIMARY KEY,
    product_id integer,
    customer_id integer,
    nama_service VARCHAR(100),
    address_line VARCHAR(255),
    locality VARCHAR(100),
    latitude DOUBLE PRECISION,
	longitude DOUBLE PRECISION,
    service_line_number VARCHAR(50),
    nickname VARCHAR(255),
    active integer NOT NULL DEFAULT 0,
    ip_kit integer,
    kit_sn VARCHAR(50),
    ssid VARCHAR(50),
    activation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_problem BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (customer_id) REFERENCES tbl_produk(id),
    FOREIGN KEY (product_id) REFERENCES tbl_produk(id)
);

CREATE INDEX service_product_id_fk ON tbl_service(product_id);
CREATE INDEX service_customer_id_fk ON tbl_service(customer_id);