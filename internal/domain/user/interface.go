package user

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/response"
)

type UserServiceInterface interface {
	Create(user *User) (*User, error)
	Get(page int, limit int) (*response.PaginationReponse[[]User], error)
	GetById(id string) (*User, error)
	Delete(id string) error
	Update(user *User) (*User, error)
}

type UserRepositoryInterface interface {
	Create(user *User) (*User, error)
	Get(page int, limit int) (*response.PaginationReponse[[]User], error)
	GetById(id string) (*User, error)
	Delete(id string) error
	Update(user *User) (*User, error)
}
