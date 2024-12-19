package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-test/internal/dto"
	"ordent-test/internal/service"
	"ordent-test/pkg/utils"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) LogIn(ctx *gin.Context) {
	var req dto.LogInRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(err.Error(), http.StatusBadRequest, err.Error()),
		)
		return
	}

	token, err := h.authService.LogIn(&req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(err.Error(), http.StatusBadRequest, err.Error()),
		)
		return
	}

	ctx.JSON(http.StatusOK, dto.LogInResponse{
		AccessToken: *token,
	})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(err.Error(), http.StatusBadRequest, err.Error()),
		)
		return
	}

	err := h.authService.Register(&req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(err.Error(), http.StatusBadRequest, err.Error()),
		)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
