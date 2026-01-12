package application

import (
	"prueba-tecnica-back/internal/order/domain"
)

type ListOrdersByCustomerUseCase struct {
	orderRepo domain.OrderRepository
}

func NewListOrdersByCustomerUseCase(orderRepo domain.OrderRepository) *ListOrdersByCustomerUseCase {
	return &ListOrdersByCustomerUseCase{
		orderRepo: orderRepo,
	}
}

func (u *ListOrdersByCustomerUseCase) Execute(customerID uint, page, limit int) ([]domain.Order, int64, error) {
	return u.orderRepo.FindByCustomerID(customerID, page, limit)
}
