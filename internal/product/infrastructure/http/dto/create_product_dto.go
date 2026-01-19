package dto

type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required,oneof=Electronics Clothing Books"`
	Price    float64 `json:"price" binding:"required,gt=0"`
}
