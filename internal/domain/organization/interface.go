package organization

import "github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"

type OrganizationServiceInterface interface {
	Create(organization *Organization) (*Organization, error)
	Get(page, limit int) (*response.PaginationReponse[[]Organization], error)
	GetById(id string) (*Organization, error)
	Delete(id string) error
	Update(organization *Organization) (*Organization, error)
}

type OrganizationRepositoryInterface interface {
	Create(organization *Organization) (*Organization, error)
	Get(page, limit int) (*response.PaginationReponse[[]Organization], error)
	GetById(id string) (*Organization, error)
	Delete(id string) error
	Update(organization *Organization) (*Organization, error)
}
