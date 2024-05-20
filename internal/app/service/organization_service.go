package service

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/uuid"
)

type OrganizationService struct {
	repo          organization.OrganizationRepositoryInterface
	uuidGenerator uuid.UUIDGeneratorInterface
}

func NewOrganizationService(repo organization.OrganizationRepositoryInterface, uuidGenerator uuid.UUIDGeneratorInterface) organization.OrganizationServiceInterface {
	return &OrganizationService{repo: repo, uuidGenerator: uuidGenerator}
}

func (s *OrganizationService) Create(organization *organization.Organization) (*organization.Organization, error) {
	id, errGenerateId := s.uuidGenerator.Generate()
	if errGenerateId != nil {
		return nil, errGenerateId
	}
	organization.Id = id.String()

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
