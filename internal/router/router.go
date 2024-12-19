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
	articleRepo := repository.NewArticleRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	authService := service.NewAuthService(userRepo)
	articleService := service.NewArticleService(articleRepo)
	commentService := service.NewCommentService(commentRepo, articleRepo)

	authHandler := handler.NewAuthHandler(authService)
	articleHandler := handler.NewArticleHandler(articleService)
	commentHandler := handler.NewCommentHandler(commentService)

	docs.SwaggerInfo.Title = "Ordent Test API"

	authMiddleware := middleware.AuthMiddleware(userRepo)

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

			auth.GET("/me", authMiddleware, authHandler.Me)
		}

		articles := v1.Group("/articles")
		articles.Use(authMiddleware)
		{
			articles.GET("", articleHandler.GetArticles)
			articles.GET("/:id", articleHandler.GetArticle)
			articles.GET("/:id/comments", commentHandler.GetCommentsByArticle)
			articles.POST("", articleHandler.CreateArticle)
			articles.PUT("/:id", articleHandler.UpdateArticle)
			articles.DELETE("/:id", articleHandler.DeleteArticle)
		}

		comments := v1.Group("/comments")
		comments.Use(authMiddleware)
		{
			comments.GET("/:id", commentHandler.GetComment)
			comments.POST("", commentHandler.CreateComment)
			comments.PUT("/:id", commentHandler.UpdateComment)
			comments.DELETE("/:id", commentHandler.DeleteComment)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
