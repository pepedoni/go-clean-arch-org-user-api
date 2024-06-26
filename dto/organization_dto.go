package dto

type CreateOrganizationDTO struct {
	Name     string `json:"name" binding:"required"`
	Document string `json:"document" binding:"required"`
}

type FilterGetOrganizationDTO struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
