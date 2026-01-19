package database

import (
	"prueba-tecnica-back/internal/product/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Save(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepositoryImpl) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}

func (r *ProductRepositoryImpl) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepositoryImpl) FindAll(categoryFilter string, sortOrder string, page, limit int) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	query := r.db.Model(&domain.Product{})

	if categoryFilter != "" {
		query = query.Where("category = ?", categoryFilter)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	order := "created_at DESC"
	if sortOrder == "asc" {
		order = "price ASC"
	} else if sortOrder == "desc" {
		order = "price DESC"
	}

	err = query.Offset(offset).Limit(limit).Order(order).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
