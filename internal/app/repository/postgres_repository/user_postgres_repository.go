package postgres_repository

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"
)

type UserPostgresRepository struct{}

func NewUserPostgresRepository() user.UserRepositoryInterface {
	return &UserPostgresRepository{}
}

func (r *UserPostgresRepository) Create(u *user.User) (*user.User, error) {
	return nil, nil
}

func (r *UserPostgresRepository) Get(page int, limit int) (*response.PaginationReponse[[]user.User], error) {
	return nil, nil
}

func (r *UserPostgresRepository) GetById(id string) (*user.User, error) {
	return nil, nil
}

func (r *UserPostgresRepository) Delete(id string) error {
	return nil
}

func (r *UserPostgresRepository) Update(u *user.User) (*user.User, error) {
	return nil, nil
}
