package dto

type CreateOrganizationUserDTO struct {
	UserId         string `json:"user_id" binding:"required"`
	OrganizationId string `json:"organization_id" binding:"required"`
}
