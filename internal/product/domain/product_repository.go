package domain

type ProductRepository interface {
	Save(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
	FindByID(id uint) (*Product, error)
	FindAll(categoryFilter string, sortOrder string, page, limit int) ([]Product, int64, error)
}
