package user

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/response"
)

type UserServiceInterface interface {
	Create(u *User) (*User, error)
	Get(filter dto.FilterGetUserDTO) (*response.PaginationReponse[[]User], error)
	GetById(id string) (*User, error)
	Delete(id string) error
	Update(u *User) (*User, error)
}

type UserRepositoryInterface interface {
	Create(user *User) (*User, error)
	Get(filter dto.FilterGetUserDTO) (*response.PaginationReponse[[]User], error)
	GetById(id string) (*User, error)
	Delete(id string) error
	Update(user *User) (*User, error)
}
