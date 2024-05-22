package organization_user

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/uuid"
)

type OrganizationUserService struct {
	repo          OrganizationUserRepositoryInterface
	uuidGenerator uuid.UUIDGeneratorInterface
}

func NewOrganizationService(repo OrganizationUserRepositoryInterface, uuidGenerator uuid.UUIDGeneratorInterface) OrganizationUserServiceInterface {
	return &OrganizationUserService{repo: repo, uuidGenerator: uuidGenerator}
}

func (s *OrganizationUserService) Create(orgCreate *OrganizationUser) (*OrganizationUser, error) {
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

func (s *OrganizationUserService) Delete(userId, organizationId string) error {
	return s.repo.Delete(userId, organizationId)
}
