CREATE TABLE tbl_perangkat (
    id integer PRIMARY KEY,
    produk_id integer,
    kategori_produk_id integer,
    nama_produk VARCHAR(100),
    deskripsi_produk TEXT,
    harga_produk INTEGER,
    gambar_produk VARCHAR(255),
    FOREIGN KEY (produk_id) REFERENCES tbl_produk(id),
    FOREIGN KEY (kategori_produk_id) REFERENCES tbl_kategori_produk(id)
);

CREATE INDEX perangkat_produk_id_fk ON tbl_perangkat(produk_id);
CREATE INDEX perangkat_kategori_produk_id_fk ON tbl_perangkat(kategori_produk_id);