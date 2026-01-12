package application

import (
	"errors"
	"prueba-tecnica-back/internal/product/domain"
)

type UpdateProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewUpdateProductUseCase(productRepo domain.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepo: productRepo,
	}
}

func (u *UpdateProductUseCase) Execute(id uint, name, category string, price float64) (*domain.Product, error) {
	if id == 0 {
		return nil, errors.New("product ID is required")
	}
	if name == "" || category == "" || price <= 0 {
		return nil, errors.New("name, category and positive price are required")
	}

	product, err := u.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Category = category
	product.Price = price

	err = u.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
