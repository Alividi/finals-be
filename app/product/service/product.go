package service

import (
	"context"
	"finals-be/app/product/dto"
	productrepo "finals-be/app/product/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"
)

type ProductService struct {
	cfg               *config.Config
	db                *connection.MultiInstruction
	productRepository productrepo.IProductRepository
}

func NewProductService(cfg *config.Config, conn *connection.SQLServerConnectionManager) *ProductService {
	db := conn.GetTransaction()
	return &ProductService{
		cfg:               cfg,
		db:                db,
		productRepository: productrepo.NewProductRepository(db),
	}
}

func (p *ProductService) GetProducts(ctx context.Context) (response []dto.GetProductsResponse, err error) {
	products, err := p.productRepository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	response = make([]dto.GetProductsResponse, len(products))
	for i, product := range products {
		response[i] = dto.GetProductsResponse{
			ID:        product.ID,
			Nama:      product.NamaProduk,
			Deskripsi: product.DeskripsiProduk,
			Image:     product.GambarProduk,
		}
	}

	return response, nil
}

func (p *ProductService) GetProductById(ctx context.Context, productId string) (response dto.GetProductDetailResponse, err error) {
	product, err := p.productRepository.GetProductById(ctx, productId)
	if err != nil {
		return dto.GetProductDetailResponse{}, err
	}

	modelLayanan, err := p.productRepository.GetLayananByProductId(ctx, productId)
	if err != nil {
		return dto.GetProductDetailResponse{}, err
	}

	modelPerangkat, err := p.productRepository.GetPerangkatByProductId(ctx, productId)
	if err != nil {
		return dto.GetProductDetailResponse{}, err
	}

	var layanan []dto.Layanan
	for _, l := range modelLayanan {
		layanan = append(layanan, dto.Layanan{
			NamaLayanan:  l.NamaLayanan,
			HargaLayanan: l.HargaLayanan,
		})
	}

	var perangkat []dto.Perangkat
	for _, p := range modelPerangkat {
		perangkat = append(perangkat, dto.Perangkat{
			ID:              p.ID,
			FkKategoriId:    p.FkKategoriId,
			NamaProduk:      p.NamaProduk,
			DeskripsiProduk: p.DeskripsiProduk,
			HargaProduk:     p.HargaProduk,
			GambarProduk:    p.GambarProduk,
		})
	}

	response = dto.GetProductDetailResponse{
		ID:          product.ID,
		Nama:        product.NamaProduk,
		Spesifikasi: product.SpesifikasiProduk,
		Image:       product.GambarProduk,
		Perangkat:   perangkat,
		Layanan:     layanan,
	}

	return response, nil
}

func (p *ProductService) GetFaqByKategoriId(ctx context.Context, kategoriId string) (response []dto.GetFaqResponse, err error) {
	faqs, err := p.productRepository.GetFaqByKategoriId(ctx, kategoriId)
	if err != nil {
		return nil, err
	}

	response = make([]dto.GetFaqResponse, len(faqs))
	for i, faq := range faqs {
		response[i] = dto.GetFaqResponse{
			ID:                faq.ID,
			KategoriProductId: faq.FkKategoriId,
			Pertanyaan:        faq.Pertanyaan,
			Jawaban:           faq.Jawaban,
		}
	}

	return response, nil
}
