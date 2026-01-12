package application

import (
	"errors"
	"prueba-tecnica-back/internal/customer/domain"
)

type CreateCustomerUseCase struct {
	customerRepo domain.CustomerRepository
}

func NewCreateCustomerUseCase(customerRepo domain.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		customerRepo: customerRepo,
	}
}

func (u *CreateCustomerUseCase) Execute(name, email string) (*domain.Customer, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	customer := &domain.Customer{
		Name:  name,
		Email: email,
	}

	err := u.customerRepo.Save(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
