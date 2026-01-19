package dto

type ListProductsQuery struct {
	Category  string `form:"category"`
	SortOrder string `form:"sort_order,default=desc"`
	Page      int    `form:"page,default=1"`
	Limit     int    `form:"limit,default=10"`
}
