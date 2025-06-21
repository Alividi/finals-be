package repository

import (
	"context"
	"database/sql"
	"finals-be/app/product/model"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"finals-be/internal/lib/helper"
	"fmt"
)

type IProductRepository interface {
	GetProducts(ctx context.Context) (products []model.Product, err error)
	GetLayananByProductId(ctx context.Context, productId int64) (layanan []model.Layanan, err error)
	GetPerangkatByProductId(ctx context.Context, productId int64) (perangkat []model.Perangkat, err error)
	GetProductById(ctx context.Context, id int64) (product model.Product, err error)
	GetFaqByKategoriId(ctx context.Context, id int64) (faq []model.FAQ, err error)
}

type ProductRepository struct {
	db connection.Connection
}

func NewProductRepository(db connection.Connection) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts(ctx context.Context) (products []model.Product, err error) {
	query := fmt.Sprintf(
		`SELECT 
			id, 
			nama_produk, 
			deskripsi_produk, 
			gambar_produk 
		FROM %s`,
		constants.TABLE_PRODUK)

	err = r.db.Select(ctx, &products, query)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("No products found")
		}
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetPerangkatByProductId(ctx context.Context, productId int64) (perangkat []model.Perangkat, err error) {
	query := fmt.Sprintf(
		`SELECT 
			id,
			kategori_produk_id, 
			nama_produk, 
			deskripsi_produk,
			harga_produk, 
			gambar_produk 
		FROM %s 
		WHERE produk_id = $1`,
		constants.TABLE_PERANGKAT)

	err = r.db.Select(ctx, &perangkat, query, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("No perangkat found for the given product ID")
		}
		return nil, err
	}

	return perangkat, nil
}

func (r *ProductRepository) GetLayananByProductId(ctx context.Context, productId int64) (layanan []model.Layanan, err error) {
	query := fmt.Sprintf(
		`SELECT 
			id, 
			nama_layanan, 
			harga_layanan 
		FROM %s 
		WHERE produk_id = $1`,
		constants.TABLE_LAYANAN)

	err = r.db.Select(ctx, &layanan, query, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("No layanan found for the given product ID")
		}
		return nil, err
	}

	return layanan, nil
}

func (r *ProductRepository) GetProductById(ctx context.Context, productId int64) (product model.Product, err error) {
	query := fmt.Sprintf(
		`SELECT 
			id, 
			nama_produk, 
			spesifikasi_produk, 
			gambar_produk 
		FROM %s 
		WHERE id = $1`,
		constants.TABLE_PRODUK)

	err = r.db.Get(ctx, &product, query, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, helper.NewErrNotFound("Product not found")
		}
		return model.Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) GetFaqByKategoriId(ctx context.Context, id int64) (faq []model.FAQ, err error) {
	query := fmt.Sprintf(
		`SELECT 
			id, 
			kategori_produk_id, 
			pertanyaan, 
			jawaban 
		FROM %s 
		WHERE kategori_produk_id = $1`,
		constants.TABLE_FAQ)

	err = r.db.Select(ctx, &faq, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("No FAQs found for the given category ID")
		}
		return nil, err
	}

	return faq, nil
}
