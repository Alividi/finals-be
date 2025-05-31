package handler

import (
	ProductService "finals-be/app/product/service"

	"finals-be/internal/lib/helper"
	"net/http"
)

type ProductHandler struct {
	productService *ProductService.ProductService
}

func NewProductHandler(productService *ProductService.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}
func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.productService.GetProducts(r.Context())
	if err != nil {
		helper.WriteResponse(r.Context(), w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, nil, products)
}

func (p *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := helper.GetURLParamString(r, "id")
	if productId == "" {
		helper.WriteResponse(r.Context(), w, helper.NewErrBadRequest("Product ID is required"), nil)
		return
	}

	product, err := p.productService.GetProductById(r.Context(), productId)
	if err != nil {
		helper.WriteResponse(r.Context(), w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, nil, product)
}

func (p *ProductHandler) GetFAQById(w http.ResponseWriter, r *http.Request) {
	faqId := helper.GetURLParamString(r, "id")
	if faqId == "" {
		helper.WriteResponse(r.Context(), w, helper.NewErrBadRequest("FAQ ID is required"), nil)
		return
	}

	faq, err := p.productService.GetFaqByKategoriId(r.Context(), faqId)
	if err != nil {
		helper.WriteResponse(r.Context(), w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, nil, faq)
}
