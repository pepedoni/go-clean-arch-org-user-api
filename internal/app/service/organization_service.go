package service

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"
)

type OrganizationService struct {
	repo organization.OrganizationRepositoryInterface
}

func NewOrganizationService(repo organization.OrganizationRepositoryInterface) organization.OrganizationServiceInterface {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) Create(organization *organization.Organization) (*organization.Organization, error) {
	return s.repo.Create(organization)
}

func (s *OrganizationService) GetById(id string) (*organization.Organization, error) {
	return s.repo.GetById(id)
}

func (s *OrganizationService) Update(organization *organization.Organization) (*organization.Organization, error) {
	return s.repo.Update(organization)
}

func (s *OrganizationService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *OrganizationService) Get(page, limit int) (*response.PaginationReponse[[]organization.Organization], error) {
	return s.repo.Get(page, limit)
}
