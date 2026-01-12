package application

import (
	"prueba-tecnica-back/internal/order/domain"
)

type CalculateTotalSpentUseCase struct {
	orderRepo domain.OrderRepository
}

func NewCalculateTotalSpentUseCase(orderRepo domain.OrderRepository) *CalculateTotalSpentUseCase {
	return &CalculateTotalSpentUseCase{
		orderRepo: orderRepo,
	}
}

func (u *CalculateTotalSpentUseCase) Execute(customerID uint) (float64, error) {
	return u.orderRepo.CalculateTotalSpentByCustomer(customerID)
}
