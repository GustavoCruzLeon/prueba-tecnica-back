package application

import (
	"errors"
	"prueba-tecnica-back/internal/customer/domain"
)

type UpdateCustomerUseCase struct {
	customerRepo domain.CustomerRepository
}

func NewUpdateCustomerUseCase(customerRepo domain.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		customerRepo: customerRepo,
	}
}

func (u *UpdateCustomerUseCase) Execute(id uint, name, email string) (*domain.Customer, error) {
	if id == 0 {
		return nil, errors.New("customer ID is required")
	}
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	customer, err := u.customerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	customer.Name = name
	customer.Email = email

	err = u.customerRepo.Update(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
