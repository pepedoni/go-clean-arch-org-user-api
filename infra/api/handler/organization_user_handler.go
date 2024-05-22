package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/organization_user"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/errors/rest_errors"
)

type OrganizationUserHandler struct {
	OrganizationUserService organization_user.OrganizationUserServiceInterface
}

func NewOrganizationUserHandler(organizationUserService organization_user.OrganizationUserServiceInterface) *OrganizationUserHandler {
	return &OrganizationUserHandler{OrganizationUserService: organizationUserService}
}

func (uh *OrganizationUserHandler) Create(c *gin.Context) {
	var organizationDto dto.CreateOrganizationUserDTO
	userId := c.Param("userId")
	orgId := c.Param("orgId")

	if userId == "" || orgId == "" {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("userId and orgId are required"))
		return
	}
	organizationDto.UserId = userId
	organizationDto.OrganizationId = orgId

	organization := organization_user.NewOrganizationUser(organizationDto.UserId, organizationDto.OrganizationId)

	organizationCreated, errCreate := uh.OrganizationUserService.Create(organization)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError(errCreate.Error()))
		return
	}
	c.JSON(http.StatusCreated, organizationCreated)
}

func (uh *OrganizationUserHandler) DeleteOrganization(c *gin.Context) {
	userId := c.Param("userId")
	orgId := c.Param("orgId")
	if err := uh.OrganizationUserService.Delete(userId, orgId); err != nil {
		if err.Error() == constants.NOT_FOUND {
			c.JSON(http.StatusNotFound, rest_errors.NewNotFoundError("organization user not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
