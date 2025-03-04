package request

type BlogRequest struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Category    string `json:"category" form:"category" binding:"required"`
	Content     string `json:"content" form:"content" binding:"required"`
}

type UserRequest struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Age   int    `json:"age" form:"age" binding:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}