package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/service"
	"ordent-test/pkg/utils"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(s service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: s,
	}
}

// CreateArticle @Summary Create an article
// @Description Create an article with the provided data
// @Tags articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateArticleRequest true "Article data"
// @Success 201 {object} model.Article
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles [post]
func (h *ArticleHandler) CreateArticle(ctx *gin.Context) {
	var req dto.CreateArticleRequest

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

	user := ctx.MustGet("user").(*model.User)

	article, err := h.articleService.CreateArticle(user, &req)

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

	ctx.JSON(http.StatusCreated, article)
}
