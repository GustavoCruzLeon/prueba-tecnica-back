package dto

type ProductResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Price     float64 `json:"price"`
	QRCode    string  `json:"qr_code,omitempty"`
	CreatedAt string  `json:"created_at"`
}
