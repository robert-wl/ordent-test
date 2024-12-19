package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-test/internal/dto"
	"ordent-test/internal/service"
	"ordent-test/pkg/utils"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

// GetUsers @Summary Get users
// @Description Get users
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Success 200 {array} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(ctx *gin.Context) {
	var req dto.GetUserRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	users, err := h.userService.GetUsers(&req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetAdmins @Summary Get admins
// @Description Get admins
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Success 200 {array} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/admins [get]
func (h *UserHandler) GetAdmins(ctx *gin.Context) {
	var req dto.GetAdminRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	users, err := h.userService.GetAdmins(&req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
