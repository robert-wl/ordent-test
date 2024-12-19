package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"ordent-test/docs"
	"ordent-test/internal/handler"
	"ordent-test/internal/infrastructure/middleware"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/internal/service"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	docs.SwaggerInfo.Title = "Ordent Test API"

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.LogIn)
			auth.POST("/register", authHandler.Register)

			auth.GET("/me", middleware.AuthMiddleware(), authHandler.Me)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
