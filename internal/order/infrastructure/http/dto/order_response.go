package dto

type OrderResponse struct {
	ID              uint        `json:"id"`
	Status          string      `json:"status"`
	Total           *float64    `json:"total,omitempty"`
	ShippingAddress *string     `json:"shipping_address,omitempty"`
	ShippedAt       *string     `json:"shipped_at,omitempty"`
	CustomerID      uint        `json:"customer_id"`
	Customer        *Customer   `json:"customer,omitempty"`
	Items           []OrderItem `json:"items"`
	CreatedAt       string      `json:"created_at"`
}

type Customer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OrderItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
