CREATE TABLE tbl_produk (
    id integer PRIMARY KEY,
    kategori_produk_id integer,
    nama_produk VARCHAR(100),
    deskripsi_produk TEXT,
    gambar_produk VARCHAR(255),
    spesifikasi_produk TEXT
);

CREATE INDEX produk_kategori_produk_id_fk ON tbl_produk(kategori_produk_id);