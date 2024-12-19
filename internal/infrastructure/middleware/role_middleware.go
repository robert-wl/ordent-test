package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-test/internal/domain/model"
	"ordent-test/pkg/utils"
)

func Role(requiredRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*model.User)

		for _, role := range requiredRoles {
			if user.Role == role {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, utils.NewErrorResponse(
			"Forbidden",
			http.StatusForbidden,
			"User does not have the required role",
		))
	}
}
