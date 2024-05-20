package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/errors/rest_errors"
)

type OrganizationHandler struct {
	OrganizationService organization.OrganizationServiceInterface
}

func NewOrganizationHandler(organizationService organization.OrganizationServiceInterface) *OrganizationHandler {
	return &OrganizationHandler{OrganizationService: organizationService}
}

func (uh *OrganizationHandler) Create(c *gin.Context) {
	var organizationDto dto.CreateOrganizationDTO
	if err := c.ShouldBindJSON(&organizationDto); err != nil {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("invalid json body"))
		return
	}

	organization := organization.NewOrganization(organizationDto.Name, organizationDto.Document)

	organizationCreated, errCreate := uh.OrganizationService.Create(organization)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError(errCreate.Error()))
		return
	}
	c.JSON(http.StatusCreated, organizationCreated)
}

func (uh *OrganizationHandler) Get(c *gin.Context) {

	organizations, err := uh.OrganizationService.Get(1, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, organizations)
}

func (uh *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	var organization organization.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	organizationUpdate, errUpdate := uh.OrganizationService.Create(&organization)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, errUpdate)
		return
	}
	c.JSON(http.StatusOK, organizationUpdate)
}

func (uh *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	id := c.Param("id")
	if err := uh.OrganizationService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (uh *OrganizationHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	if err := uh.OrganizationService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	organization, errGetOrganization := uh.OrganizationService.GetById(id)
	if errGetOrganization != nil {
		c.JSON(http.StatusInternalServerError, errGetOrganization)
		return
	}
	c.JSON(http.StatusOK, organization)
}
