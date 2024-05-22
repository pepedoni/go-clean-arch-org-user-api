package organization

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/uuid"
)

type OrganizationService struct {
	repo          OrganizationRepositoryInterface
	uuidGenerator uuid.UUIDGeneratorInterface
}

func NewOrganizationService(repo OrganizationRepositoryInterface, uuidGenerator uuid.UUIDGeneratorInterface) OrganizationServiceInterface {
	return &OrganizationService{repo: repo, uuidGenerator: uuidGenerator}
}

func (s *OrganizationService) Create(orgCreate *Organization) (*Organization, error) {
	if err := orgCreate.Validate(); err != nil {
		return nil, err
	}

	id, errGenerateId := s.uuidGenerator.Generate()
	if errGenerateId != nil {
		return nil, errGenerateId
	}
	orgCreate.Id = id.String()

	return s.repo.Create(orgCreate)
}

func (s *OrganizationService) GetById(id string) (*Organization, error) {
	return s.repo.GetById(id)
}

func (s *OrganizationService) Update(orgUpdate *Organization) (*Organization, error) {
	oldOrganization, errGet := s.repo.GetById(orgUpdate.Id)
	if errGet != nil {
		return nil, errGet
	}

	orgUpdate.FormatDocument()

	if err := orgUpdate.Validate(); err != nil {
		return nil, err
	}
	oldOrganization.Name = orgUpdate.Name
	oldOrganization.Document = orgUpdate.Document

	return s.repo.Update(oldOrganization)
}

func (s *OrganizationService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *OrganizationService) Get(filter dto.FilterGetOrganizationDTO) (*response.PaginationReponse[[]Organization], error) {
	return s.repo.Get(filter)
}
