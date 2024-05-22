package user

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/uuid"
)

type UserService struct {
	repo          UserRepositoryInterface
	uuidGenerator uuid.UUIDGeneratorInterface
}

func NewUserService(repo UserRepositoryInterface, uuidGenerator uuid.UUIDGeneratorInterface) UserServiceInterface {
	return &UserService{repo: repo, uuidGenerator: uuidGenerator}
}

func (s *UserService) Create(userCreate *User) (*User, error) {
	if err := userCreate.Validate(); err != nil {
		return nil, err
	}

	id, errGenerateId := s.uuidGenerator.Generate()
	if errGenerateId != nil {
		return nil, errGenerateId
	}
	userCreate.Id = id.String()

	return s.repo.Create(userCreate)
}

func (s *UserService) GetById(id string) (*User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Update(userUpdate *User) (*User, error) {
	oldUser, errGet := s.repo.GetById(userUpdate.Id)
	if errGet != nil {
		return nil, errGet
	}

	userUpdate.FormatDocument()
	userUpdate.FormatPhone()

	if err := userUpdate.Validate(); err != nil {
		return nil, err
	}
	oldUser.Name = userUpdate.Name
	oldUser.Email = userUpdate.Email
	oldUser.Phone = userUpdate.Phone
	oldUser.Document = userUpdate.Document

	return s.repo.Update(oldUser)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) Get(filter dto.FilterGetUserDTO) (*response.PaginationReponse[[]User], error) {
	return s.repo.Get(filter)
}
