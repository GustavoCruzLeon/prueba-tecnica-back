package database

import (
	"prueba-tecnica-back/internal/customer/domain"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{db: db}
}

func (r *CustomerRepositoryImpl) Save(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepositoryImpl) Update(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}

func (r *CustomerRepositoryImpl) FindByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *CustomerRepositoryImpl) FindAll(nameFilter string, page, limit int) ([]domain.Customer, int64, error) {
	var customers []domain.Customer
	var total int64

	query := r.db.Model(&domain.Customer{})

	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&customers).Error
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}
