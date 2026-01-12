package application

import (
	"prueba-tecnica-back/internal/customer/domain"
)

type ListCustomersUseCase struct {
	customerRepo domain.CustomerRepository
}

func NewListCustomersUseCase(customerRepo domain.CustomerRepository) *ListCustomersUseCase {
	return &ListCustomersUseCase{
		customerRepo: customerRepo,
	}
}

func (u *ListCustomersUseCase) Execute(nameFilter string, page, limit int) ([]domain.Customer, int64, error) {
	return u.customerRepo.FindAll(nameFilter, page, limit)
}
