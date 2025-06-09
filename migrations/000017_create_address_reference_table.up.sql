CREATE TABLE tbl_address_reference (
    id integer PRIMARY KEY,
    address_line VARCHAR(255),
    locality VARCHAR(100),
    latitude DOUBLE PRECISION,
	longitude DOUBLE PRECISION
);