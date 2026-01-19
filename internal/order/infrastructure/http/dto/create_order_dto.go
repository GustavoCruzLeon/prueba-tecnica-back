package dto

type CreateOrderRequest struct {
	CustomerID uint        `json:"customer_id" binding:"required"`
	Items      []OrderItem `json:"items" binding:"required,min=1"`
}
