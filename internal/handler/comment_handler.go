package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/service"
	"ordent-test/pkg/utils"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(s service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: s,
	}
}

// CreateComment @Summary Create a comment
// @Description Create a comment on an article or another comment
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCommentRequest true "Comment details"
// @Success 200 {object} model.Comment
// @Failure 400 {object} utils.ErrorResponse
// @Router /comments [post]
func (h *CommentHandler) CreateComment(ctx *gin.Context) {
	var req dto.CreateCommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	user := ctx.MustGet("user").(*model.User)

	comment, err := h.commentService.CreateComment(user, &req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// GetComment @Summary Get a comment
// @Description Get a comment by its ID
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Success 200 {object} model.Comment
// @Failure 400 {object} utils.ErrorResponse
// @Router /comments/{id} [get]
func (h *CommentHandler) GetComment(ctx *gin.Context) {
	commentId := ctx.Param("id")

	comment, err := h.commentService.GetComment(commentId)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// GetCommentsByArticle @Summary Get comments by article
// @Description Get all comments on an article
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Param id path string true "Article ID"
// @Success 200 {array} model.Comment
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles/{id}/comments [get]
func (h *CommentHandler) GetCommentsByArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	var req dto.GetArticleCommentRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	comments, err := h.commentService.GetArticleComments(articleId, &req)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				"Bad Request",
				http.StatusBadRequest,
				err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// DeleteComment @Summary Delete a comment
// @Description Delete a comment by its ID, only the owner or admin can delete
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Success 200
// @Failure 400 {object} utils.ErrorResponse
// @Router /comments/{id} [delete]
func (h *CommentHandler) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("id")
	user := ctx.MustGet("user").(*model.User)

	err := h.commentService.DeleteComment(user, commentId)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// UpdateComment @Summary Update a comment
// @Description Update a comment by its ID, only the owner or admin can update
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Comment ID"
// @Param request body dto.UpdateCommentRequest true "Comment details"
// @Success 200 {object} model.Comment
// @Failure 400 {object} utils.ErrorResponse
// @Router /comments/{id} [put]
func (h *CommentHandler) UpdateComment(ctx *gin.Context) {
	var req dto.UpdateCommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendError(ctx, err)
		return
	}

	commentId := ctx.Param("id")
	user := ctx.MustGet("user").(*model.User)

	comment, err := h.commentService.UpdateComment(user, commentId, &req)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
