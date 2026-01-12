package application

import (
	"errors"
	"prueba-tecnica-back/internal/product/domain"
)

type DeleteProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewDeleteProductUseCase(productRepo domain.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepo: productRepo,
	}
}

func (u *DeleteProductUseCase) Execute(id uint) error {
	if id == 0 {
		return errors.New("product ID is required")
	}

	return u.productRepo.Delete(id)
}
