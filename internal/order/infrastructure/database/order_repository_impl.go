package database

import (
	"prueba-tecnica-back/internal/order/domain"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Save(order *domain.Order) error {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range order.Items {
		order.Items[i].OrderID = order.ID
		if err := tx.Create(&order.Items[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *OrderRepositoryImpl) FindByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *OrderRepositoryImpl) FindByCustomerID(customerID uint, page, limit int) ([]domain.Order, int64, error) {
	var orders []domain.Order
	var total int64

	query := r.db.Where("customer_id = ?", customerID)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (r *OrderRepositoryImpl) CalculateTotalSpentByCustomer(customerID uint) (float64, error) {
	var total float64
	err := r.db.Model(&domain.Order{}).
		Where("customer_id = ? AND status = ?", customerID, "completed").
		Select("COALESCE(SUM(total), 0)").
		Scan(&total).Error
	return total, err
}
