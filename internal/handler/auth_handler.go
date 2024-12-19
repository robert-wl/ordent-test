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

// LogIn @Summary Log in a user
// @Description Authenticate a user with their credentials and return an access token
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dto.LogInRequest true "User credentials"
// @Success 200 {object} dto.LogInResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) LogIn(ctx *gin.Context) {
	var req dto.LogInRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()),
		)
		return
	}

	token, err := h.authService.LogIn(&req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				err.Error()),
		)
		return
	}

	ctx.JSON(http.StatusOK, dto.LogInResponse{
		AccessToken: *token,
	})
}

// Register @Summary Register a user
// @Description Register a user with their credentials
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User credentials"
// @Success 200
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()),
		)
		return
	}

	user, err := h.authService.Register(&req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()),
		)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
