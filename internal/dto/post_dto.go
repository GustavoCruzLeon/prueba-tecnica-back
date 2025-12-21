package dto

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
}

type ListPostsQuery struct {
	Page      int    `form:"page,default=1"`
	Limit     int    `form:"limit,default=10"`
	UserID    *uint  `form:"user_id"`
	SortOrder string `form:"sort_order,default=desc"`
}
