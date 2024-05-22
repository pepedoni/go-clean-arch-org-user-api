package login

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
)

const (
	SECRET_KEY_JWT = "SECRET_KEY_JWT"
)

var (
	secretKey  = []byte(os.Getenv(SECRET_KEY_JWT))
	validUsers = map[string]string{
		"teste@teste.com": "123456",
	}
)

type LoginService struct {
}

func NewLoginService() LoginServiceInterface {
	return &LoginService{}
}

func (s *LoginService) Login(loginRequest *dto.LoginRequestDTO) (*string, error) {
	if validUsers[loginRequest.Email] != loginRequest.Password {
		return nil, errors.New("invalid credentials")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loginRequest.Email,
		"iss": "go-clean-arch-org-user-api",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, errSign := claims.SignedString(secretKey)
	if errSign != nil {
		return nil, errSign
	}

	return &tokenString, nil
}
