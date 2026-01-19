package application

import (
	"encoding/base64"
	"errors"
	"prueba-tecnica-back/internal/product/domain"

	"github.com/skip2/go-qrcode"
)

type GenerateQRUseCase struct {
	productRepo domain.ProductRepository
}

func NewGenerateQRUseCase(productRepo domain.ProductRepository) *GenerateQRUseCase {
	return &GenerateQRUseCase{
		productRepo: productRepo,
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

	qrData := "Product ID: " + string(rune(product.ID)) + ", Name: " + product.Name + ", Price: $" + string(rune(int(product.Price)))

	// Generar el QR en formato PNG
	png, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	// Convertir a base64
	qrBase64 := base64.StdEncoding.EncodeToString(png)
	product.QRCode = "data:image/png;base64," + qrBase64

	err = u.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
