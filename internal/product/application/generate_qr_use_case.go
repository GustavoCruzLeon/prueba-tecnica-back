package application

import (
	"errors"
	"prueba-tecnica-back/internal/product/domain"
)

type GenerateQRUseCase struct {
	productRepo domain.ProductRepository
	qrGenerator func(data string) (string, error)
}

func NewGenerateQRUseCase(productRepo domain.ProductRepository, qrGenerator func(string) (string, error)) *GenerateQRUseCase {
	return &GenerateQRUseCase{
		productRepo: productRepo,
		qrGenerator: qrGenerator,
	}
}

func (u *GenerateQRUseCase) Execute(productID uint) (*domain.Product, error) {
	if productID == 0 {
		return nil, errors.New("product ID is required")
	}

	product, err := u.productRepo.FindByID(productID)
	if err != nil {
		return nil, err
	}

	qrData := "Product ID: " + string(rune(product.ID)) + ", Name: " + product.Name
	qrCode, err := u.qrGenerator(qrData)
	if err != nil {
		return nil, err
	}

	product.QRCode = qrCode
	err = u.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
