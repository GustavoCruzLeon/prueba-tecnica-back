package application

import (
	"prueba-tecnica-back/internal/product/domain"
)

type ListProductsUseCase struct {
	productRepo domain.ProductRepository
}

func NewListProductsUseCase(productRepo domain.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{
		productRepo: productRepo,
	}
}

func (u *ListProductsUseCase) Execute(categoryFilter, sortOrder string, page, limit int) ([]domain.Product, int64, error) {
	return u.productRepo.FindAll(categoryFilter, sortOrder, page, limit)
}
