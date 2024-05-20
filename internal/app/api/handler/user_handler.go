package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pepedoni/go-clean-arch-org-user-api/internal/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/utils/errors/rest_errors"
)

type UserHandler struct {
	UserService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (uh *UserHandler) Create(c *gin.Context) {
	var userDto dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("invalid json body"))
		return
	}

	user := user.NewUser(userDto.Name, userDto.Email, userDto.Phone, userDto.Document)

	userCreated, errCreate := uh.UserService.Create(user)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError(errCreate.Error()))
		return
	}
	c.JSON(http.StatusCreated, userCreated)
}

func (uh *UserHandler) Get(c *gin.Context) {

	users, err := uh.UserService.Get(1, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userUpdate, errUpdate := uh.UserService.Create(&user)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, errUpdate)
		return
	}
	c.JSON(http.StatusOK, userUpdate)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uh.UserService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (uh *UserHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	if err := uh.UserService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	user, errGetUser := uh.UserService.GetById(id)
	if errGetUser != nil {
		c.JSON(http.StatusInternalServerError, errGetUser)
		return
	}
	c.JSON(http.StatusOK, user)
}
