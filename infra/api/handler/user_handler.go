package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pepedoni/go-clean-arch-org-user-api/constants"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/dto"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/errors/rest_errors"
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
	filter := dto.FilterGetUserDTO{
		Page:  1,
		Limit: 10,
	}
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("invalid query params"))
		return
	}

	users, err := uh.UserService.Get(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user user.User

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("id is required"))
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user.Id = id
	userUpdate, errUpdate := uh.UserService.Update(&user)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError(errUpdate.Error()))
		return
	}
	c.JSON(http.StatusOK, userUpdate)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uh.UserService.Delete(id); err != nil {
		if err.Error() == constants.NOT_FOUND {
			c.JSON(http.StatusNotFound, rest_errors.NewNotFoundError("user not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (uh *UserHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	user, errGetUser := uh.UserService.GetById(id)
	if errGetUser != nil {
		c.JSON(http.StatusInternalServerError, errGetUser)
		return
	}
	fmt.Println("vai terminar de executar o GetById")
	c.JSON(http.StatusOK, user)
}
