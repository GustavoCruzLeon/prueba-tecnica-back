package application

import (
	"errors"
	"prueba-tecnica-back/internal/order/domain"
)

type ProductReader interface {
	GetProductPrice(id uint) (float64, error)
}

type CustomerReader interface {
	ExistsCustomer(id uint) (bool, error)
}

type CreateOrderUseCase struct {
	orderRepo      domain.OrderRepository
	productReader  ProductReader
	customerReader CustomerReader
}

func NewCreateOrderUseCase(
	orderRepo domain.OrderRepository,
	productReader ProductReader,
	customerReader CustomerReader,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:      orderRepo,
		productReader:  productReader,
		customerReader: customerReader,
	}
}

func (u *CreateOrderUseCase) Execute(customerID uint, items []domain.OrderItem) (*domain.Order, error) {
	if customerID == 0 {
		return nil, errors.New("customer ID is required")
	}
	if len(items) == 0 {
		return nil, errors.New("at least one item is required")
	}

	exists, err := u.customerReader.ExistsCustomer(customerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("customer not found")
	}

	var total float64
	for i := range items {
		if items[i].ProductID == 0 || items[i].Quantity <= 0 {
			return nil, errors.New("invalid product ID or quantity")
		}

		price, err := u.productReader.GetProductPrice(items[i].ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		items[i].Price = price
		total += price * float64(items[i].Quantity)
	}

	order := &domain.Order{
		Status:     "pending",
		Total:      &total,
		CustomerID: customerID,
		Items:      items,
	}

	err = u.orderRepo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
