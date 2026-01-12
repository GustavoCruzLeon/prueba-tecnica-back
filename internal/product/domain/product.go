package domain

import "time"

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	QRCode    string    `json:"qr_code,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
