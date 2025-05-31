CREATE TABLE tbl_faq_kategori_produk (
    id VARCHAR(50) PRIMARY KEY,
    kategori_produk_id VARCHAR(50),
    pertanyaan TEXT,
    jawaban TEXT,
    FOREIGN KEY (kategori_produk_id) REFERENCES tbl_kategori_produk(id)
);

CREATE INDEX faq_kategori_produk_id_fk ON tbl_faq_kategori_produk(kategori_produk_id);