package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/app/api/handler"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/app/repository/memory_repository"
	"github.com/pepedoni/go-clean-arch-org-user-api/internal/app/service"
)

var (
	router = gin.Default()
)

func mapUserRoutes(api *gin.RouterGroup) {
	usersGroup := api.Group("/users")

	repository := memory_repository.NewUserMemoryRepository()
	service := service.NewUserService(repository)
	userHandler := handler.NewUserHandler(service)

	usersGroup.GET("", userHandler.Get)
	usersGroup.GET("/:id", userHandler.GetById)
	usersGroup.POST("", userHandler.Create)
	usersGroup.PUT("/:id", userHandler.UpdateUser)
	usersGroup.DELETE("/:id", userHandler.DeleteUser)

}

func StartApplication() {
	// router.use(OAuthMiddleware())
	api := router.Group("/api")

	mapUserRoutes(api)
	fmt.Println("Starting (111) the application...")
	router.Run(":8080")

}
