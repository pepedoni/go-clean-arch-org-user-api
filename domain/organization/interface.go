package organization

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
)

type OrganizationServiceInterface interface {
	Create(organization *Organization) (*Organization, error)
	Get(filter dto.FilterGetOrganizationDTO) (*response.PaginationReponse[[]Organization], error)
	GetById(id string) (*Organization, error)
	Delete(id string) error
	Update(organization *Organization) (*Organization, error)
}

type OrganizationRepositoryInterface interface {
	Create(organization *Organization) (*Organization, error)
	Get(filter dto.FilterGetOrganizationDTO) (*response.PaginationReponse[[]Organization], error)
	GetById(id string) (*Organization, error)
	Delete(id string) error
	Update(organization *Organization) (*Organization, error)
}
