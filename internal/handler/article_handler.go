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

// GetArticles @Summary Get all articles
// @Description Get all articles
// @Tags articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Article
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles [get]
func (h *ArticleHandler) GetArticles(ctx *gin.Context) {
	articles, err := h.articleService.GetArticles()

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

	ctx.JSON(http.StatusOK, articles)
}

// GetArticle @Summary Get an article
// @Description Get an article by its ID
// @Tags articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Article ID"
// @Success 200 {object} model.Article
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles/{id} [get]
func (h *ArticleHandler) GetArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	article, err := h.articleService.GetArticle(articleId)

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

	ctx.JSON(http.StatusOK, article)
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

// UpdateArticle @Summary Update an article
// @Description Update an article with the provided data
// @Tags articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Article ID"
// @Param request body dto.UpdateArticleRequest true "Article data"
// @Success 200 {object} model.Article
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles/{id} [put]
func (h *ArticleHandler) UpdateArticle(ctx *gin.Context) {
	var req dto.UpdateArticleRequest

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

	articleId := ctx.Param("id")
	user := ctx.MustGet("user").(*model.User)

	article, err := h.articleService.UpdateArticle(user, articleId, &req)

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

	ctx.JSON(http.StatusOK, article)
}
