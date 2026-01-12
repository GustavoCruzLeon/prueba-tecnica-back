package application

import (
	"errors"
	"prueba-tecnica-back/internal/product/domain"
)

type CreateProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewCreateProductUseCase(productRepo domain.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepo: productRepo,
	}
}

func (u *CreateProductUseCase) Execute(name, category string, price float64) (*domain.Product, error) {
	if name == "" || category == "" || price <= 0 {
		return nil, errors.New("name, category and positive price are required")
	}

	product := &domain.Product{
		Name:     name,
		Category: category,
		Price:    price,
	}

	err := u.productRepo.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
