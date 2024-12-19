package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ordent-test/pkg/auth"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Contains(ctx.Request.URL.Path, "auth") {
			ctx.Next()
			return
		}

		header := ctx.GetHeader("Authorization")
		if !strings.Contains(header, "Bearer") {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		token := strings.Replace(header, "Bearer ", "", 1)

		user, err := auth.ParseJWT(token)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		fmt.Println("hahahaH", user)
		ctx.Set("user", user)
	}
}
