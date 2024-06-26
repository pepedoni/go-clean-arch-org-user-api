package dto

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
}

type FilterGetUserDTO struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
