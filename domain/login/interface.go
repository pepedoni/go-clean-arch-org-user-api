package login

import (
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
)

type LoginServiceInterface interface {
	Login(loginRequest *dto.LoginRequestDTO) (*string, error)
}
