CREATE TABLE tbl_layanan (
    id VARCHAR(50) PRIMARY KEY,
    produk_id VARCHAR(50),
    nama_layanan VARCHAR(100),
    harga_layanan INTEGER,
    FOREIGN KEY (produk_id) REFERENCES tbl_produk(id)
);

CREATE INDEX layanan_produk_id_fk ON tbl_layanan(produk_id);
