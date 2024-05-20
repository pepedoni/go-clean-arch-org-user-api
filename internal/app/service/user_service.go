package service

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"
)

type UserService struct {
	repo user.UserRepositoryInterface
}

func NewUserService(repo user.UserRepositoryInterface) user.UserServiceInterface {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *user.User) (*user.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return s.repo.Create(user)
}

func (s *UserService) GetById(id string) (*user.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Update(user *user.User) (*user.User, error) {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) Get(page int, limit int) (*response.PaginationReponse[[]user.User], error) {
	return s.repo.Get(page, limit)
}
