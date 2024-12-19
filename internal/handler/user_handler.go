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
		utils.SendError(ctx, err)
		return
	}

	users, err := h.userService.GetUsers(&req)

	if err != nil {
		utils.SendError(ctx, err)
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
		utils.SendError(ctx, err)
		return
	}

	users, err := h.userService.GetAdmins(&req)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUser @Summary Get a user
// @Description Get a user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := h.userService.GetUser(userId)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// PromoteUser @Summary Promote a user
// @Description Promote a user to admin
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id}/promote [put]
func (h *UserHandler) PromoteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := h.userService.ChangeRole(userId, "admin")

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DemoteUser @Summary Demote a user
// @Description Demote a user to user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id}/demote [put]
func (h *UserHandler) DemoteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := h.userService.ChangeRole(userId, "user")

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
