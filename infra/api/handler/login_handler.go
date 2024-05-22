package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/login"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/errors/rest_errors"
)

type LoginHandler struct {
	LoginService login.LoginServiceInterface
}

func NewLoginHandler(service login.LoginServiceInterface) *LoginHandler {
	return &LoginHandler{
		LoginService: service,
	}
}

func (lh *LoginHandler) Login(c *gin.Context) {
	var loginDto dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("invalid json body"))
		return
	}

	token, err := lh.LoginService.Login(&loginDto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": *token})
}
