package domain

type OrderRepository interface {
	Save(order *Order) error
	FindByID(id uint) (*Order, error)
	FindByCustomerID(customerID uint, page, limit int) ([]Order, int64, error)
	CalculateTotalSpentByCustomer(customerID uint) (float64, error)
}
