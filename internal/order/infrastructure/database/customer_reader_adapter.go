package database

import (
	"gorm.io/gorm"
)

type CustomerReaderAdapter struct {
	db *gorm.DB
}

func NewCustomerReaderAdapter(db *gorm.DB) *CustomerReaderAdapter {
	return &CustomerReaderAdapter{db: db}
}

func (a *CustomerReaderAdapter) ExistsCustomer(id uint) (bool, error) {
	var exists bool
	err := a.db.Raw("SELECT EXISTS(SELECT 1 FROM customers WHERE id = ?)", id).Scan(&exists).Error
	return exists, err
}
