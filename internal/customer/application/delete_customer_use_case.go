package application

import (
	"errors"
	"prueba-tecnica-back/internal/customer/domain"
)

type DeleteCustomerUseCase struct {
	customerRepo domain.CustomerRepository
}

func NewDeleteCustomerUseCase(customerRepo domain.CustomerRepository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{
		customerRepo: customerRepo,
	}
}

func (u *DeleteCustomerUseCase) Execute(id uint) error {
	if id == 0 {
		return errors.New("customer ID is required")
	}

	return u.customerRepo.Delete(id)
}
