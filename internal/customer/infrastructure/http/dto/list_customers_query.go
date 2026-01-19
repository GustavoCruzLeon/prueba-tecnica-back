package dto

type ListCustomersQuery struct {
	Name  string `form:"name"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}
