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
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Success 200 {array} model.Article
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles [get]
func (h *ArticleHandler) GetArticles(ctx *gin.Context) {
	var req dto.GetArticleRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.SendError(ctx, err)
		return
	}

	articles, err := h.articleService.GetArticles(&req)

	if err != nil {
		utils.SendError(ctx, err)
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
		utils.SendError(ctx, err)
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
		utils.SendError(ctx, err)
		return
	}

	user := ctx.MustGet("user").(*model.User)

	article, err := h.articleService.CreateArticle(user, &req)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

// UpdateArticle @Summary Update an article
// @Description Update an article with the provided data, only the owner or admin can update
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
		utils.SendError(ctx, err)
		return
	}

	articleId := ctx.Param("id")
	user := ctx.MustGet("user").(*model.User)

	article, err := h.articleService.UpdateArticle(user, articleId, &req)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// DeleteArticle @Summary Delete an article
// @Description Delete an article by its ID, only the owner or admin can delete
// @Tags articles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Article ID"
// @Success 204
// @Failure 400 {object} utils.ErrorResponse
// @Router /articles/{id} [delete]
func (h *ArticleHandler) DeleteArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")
	user := ctx.MustGet("user").(*model.User)

	err := h.articleService.DeleteArticle(user, articleId)

	if err != nil {
		utils.SendError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
