package database

import (
	"gorm.io/gorm"
)

type ProductReaderAdapter struct {
	db *gorm.DB
}

func NewProductReaderAdapter(db *gorm.DB) *ProductReaderAdapter {
	return &ProductReaderAdapter{db: db}
}

func (a *ProductReaderAdapter) GetProductPrice(id uint) (float64, error) {
	var price float64
	err := a.db.Raw("SELECT price FROM products WHERE id = ?", id).Scan(&price).Error
	if err != nil {
		return 0, err
	}
	return price, nil
}
