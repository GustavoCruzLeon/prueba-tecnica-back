package domain

type CustomerRepository interface {
	Save(customer *Customer) error
	Update(customer *Customer) error
	Delete(id uint) error
	FindByID(id uint) (*Customer, error)
	FindAll(nameFilter string, page, limit int) ([]Customer, int64, error)
}
