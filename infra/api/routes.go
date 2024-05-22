package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/login"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/organization"
	"github.com/pepedoni/go-clean-arch-org-user-api/domain/user"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/api/handler"
	redisCache "github.com/pepedoni/go-clean-arch-org-user-api/infra/cache/redis"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/database/postgres"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/middlewares"
	"github.com/pepedoni/go-clean-arch-org-user-api/infra/repository/postgres_repository"
	"github.com/pepedoni/go-clean-arch-org-user-api/utils/uuid"
)

var (
	router = gin.Default()
)

func mapUserRoutes(api *gin.RouterGroup, postgresConnection postgres.PoolInterface, redisConnection *redis.Client) {
	usersGroup := api.Group("/users")
	usersGroup.Use(middlewares.OAuthMiddleware())

	// repository := memory_repository.NewUserMemoryRepository()
	repository := postgres_repository.NewUserPostgresRepository(postgresConnection)
	service := user.NewUserService(repository, uuid.New())
	userHandler := handler.NewUserHandler(service)

	usersGroup.GET("", userHandler.Get)
	usersGroup.GET("/:id", middlewares.RedisCacheByRequestUri(redisConnection, 600), userHandler.GetById)
	usersGroup.POST("", userHandler.Create)
	usersGroup.PUT("/:id", middlewares.InvalidateCacheByRequestUri(redisConnection), userHandler.UpdateUser)
	usersGroup.DELETE("/:id", middlewares.InvalidateCacheByRequestUri(redisConnection), userHandler.DeleteUser)
}

func mapOrganizationRoutes(api *gin.RouterGroup, postgresConnection postgres.PoolInterface, redisConnection *redis.Client) {
	organizationsGroup := api.Group("/organizations")
	organizationsGroup.Use(middlewares.OAuthMiddleware())

	// repository := memory_repository.NewOrganizationMemoryRepository()
	repository := postgres_repository.NewOrganizationPostgresRepository(postgresConnection)
	service := organization.NewOrganizationService(repository, uuid.New())
	organizationHandler := handler.NewOrganizationHandler(service)

	organizationsGroup.GET("", middlewares.RedisCacheByRequestUri(redisConnection, 600), organizationHandler.Get)
	organizationsGroup.GET("/:id", middlewares.RedisCacheByRequestUri(redisConnection, 600), organizationHandler.GetById)
	organizationsGroup.POST("", middlewares.InvalidateCacheLikeRequestUri(redisConnection), organizationHandler.Create)
	organizationsGroup.PUT("/:id", middlewares.InvalidateCacheByRequestUri(redisConnection), organizationHandler.UpdateOrganization)
	organizationsGroup.DELETE("/:id", middlewares.InvalidateCacheByRequestUri(redisConnection), organizationHandler.DeleteOrganization)
}

func mapLoginRoutes(api *gin.RouterGroup) {
	loginGroup := api.Group("/login")

	service := login.NewLoginService()
	loginHandler := handler.NewLoginHandler(service)

	loginGroup.POST("", loginHandler.Login)
}

func StartApplication() {
	api := router.Group("/api")

	ctx := context.Background()
	postgresConnection := postgres.GetConnection(ctx)
	defer postgresConnection.Close()

	redisConnection := redisCache.GetConnection()
	defer redisConnection.Close()

	postgres.RunMigrations()

	mapUserRoutes(api, postgresConnection, redisConnection)
	mapOrganizationRoutes(api, postgresConnection, redisConnection)
	mapLoginRoutes(api)

	router.Run(":8080")
}
