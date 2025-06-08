CREATE TABLE tbl_faq_kategori_produk (
    id integer PRIMARY KEY,
    kategori_produk_id integer,
    pertanyaan TEXT,
    jawaban TEXT,
    FOREIGN KEY (kategori_produk_id) REFERENCES tbl_kategori_produk(id)
);

CREATE INDEX faq_kategori_produk_id_fk ON tbl_faq_kategori_produk(kategori_produk_id);