CREATE TABLE tbl_service (
    id integer PRIMARY KEY,
    produk_id integer,
    nama_service VARCHAR(100),
    deskripsi_service VARCHAR(100),
    gambar_service VARCHAR(100),
    harga_service INTEGER,
    FOREIGN KEY (produk_id) REFERENCES tbl_produk(id)
);

CREATE INDEX service_produk_id_fk ON tbl_service(produk_id);