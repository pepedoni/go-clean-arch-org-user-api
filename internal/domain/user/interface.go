package user

import "github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/errors/rest_errors"

type UserServiceInterface interface {
	CreateUser(user User) (*User, *rest_errors.RestErr)
	GetUserById(id string) (*User, *rest_errors.RestErr)
	DeleteUser(id string) *rest_errors.RestErr
	UpdateUser(user User) (*User, *rest_errors.RestErr)
}
