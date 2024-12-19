package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-test/pkg/auth"
	"ordent-test/pkg/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			ctx.Next()
			return
		}

		header = strings.Replace(header, "Bearer ", "", 1)

		token := strings.Replace(header, "Bearer ", "", 1)

		user, err := auth.ParseJWT(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewErrorResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"Invalid token",
			))
			return
		}

		fmt.Println(user)
		ctx.Set("user", user)
	}
}
