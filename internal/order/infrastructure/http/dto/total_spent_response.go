package dto

type TotalSpentResponse struct {
	CustomerID uint    `json:"customer_id"`
	Total      float64 `json:"total"`
}
