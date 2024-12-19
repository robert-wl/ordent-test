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

// GetCommentsByArticle @Summary Get comments by article
// @Description Get all comments on an article
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Article ID"
// @Success 200 {array} model.Comment
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles/{id}/comments [get]
func (h *CommentHandler) GetCommentsByArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	comments, err := h.commentService.GetCommentsByArticleSecureID(articleId)

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
