package domain

import "time"

type Order struct {
	ID              uint        `json:"id"`
	Status          string      `json:"status"`
	Total           *float64    `json:"total,omitempty"`
	ShippingAddress *string     `json:"shipping_address,omitempty"`
	ShippedAt       *time.Time  `json:"shipped_at,omitempty"`
	CustomerID      uint        `json:"customer_id"`
	Customer        *Customer   `json:"customer,omitempty"`
	Items           []OrderItem `json:"items"`
	CreatedAt       time.Time   `json:"created_at"`
}

type Customer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
