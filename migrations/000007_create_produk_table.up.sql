CREATE TABLE tbl_produk (
    id VARCHAR(50) PRIMARY KEY,
    kategori_produk_id VARCHAR(50),
    nama_produk VARCHAR(100),
    deskripsi_produk TEXT,
    gambar_produk VARCHAR(255),
    spesifikasi_produk TEXT
);

CREATE INDEX produk_kategori_produk_id_fk ON tbl_produk(kategori_produk_id);