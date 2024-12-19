package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ordent-test/internal/handler"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/internal/service"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.LogIn)
		}
	}
	return r
}
