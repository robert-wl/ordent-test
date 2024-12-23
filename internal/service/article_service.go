package service

import (
	"fmt"
	"net/http"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
	"ordent-test/pkg/pagination"
	"ordent-test/pkg/utils"
)

type ArticleService interface {
	GetArticles(dto *dto.GetArticleRequest) ([]*model.Article, error)
	GetArticle(articleId string) (*model.Article, error)
	CreateArticle(user *model.User, dto *dto.CreateArticleRequest) (*model.Article, error)
	UpdateArticle(user *model.User, articleId string, dto *dto.UpdateArticleRequest) (*model.Article, error)
	DeleteArticle(user *model.User, articleId string) error
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(r repository.ArticleRepository) ArticleService {
	return &articleService{
		repo: r,
	}
}

func (s *articleService) GetArticles(dto *dto.GetArticleRequest) ([]*model.Article, error) {
	if dto.Search == nil {
		dto.Search = new(string)
	}

	if dto.Pagination == nil {
		dto.Pagination = new(pagination.Pagination)
	}

	articles, err := s.repo.Find(dto.Search, dto.Pagination)

	if err != nil {
		return nil, utils.NewAppError(
			err,
			http.StatusInternalServerError,
			"Failed to get articles",
		)
	}

	return articles, nil
}

func (s *articleService) GetArticle(articleId string) (*model.Article, error) {
	article, err := s.repo.FindBySecureID(articleId)

	if err != nil {
		return nil, utils.NewAppError(
			err,
			http.StatusNotFound,
			"Article not found",
		)
	}

	return article, nil
}

func (s *articleService) CreateArticle(user *model.User, dto *dto.CreateArticleRequest) (*model.Article, error) {
	article := &model.Article{
		UserID: user.ID,
		Title:  dto.Title,
		Body:   dto.Body,
	}

	createdArticle, err := s.repo.Create(article)

	if err != nil {
		return nil, utils.NewAppError(
			err,
			http.StatusInternalServerError,
			"Failed to create article",
		)
	}

	return createdArticle, nil
}

func (s *articleService) UpdateArticle(user *model.User, articleId string, dto *dto.UpdateArticleRequest) (*model.Article, error) {
	article, err := s.repo.FindBySecureID(articleId)

	if err != nil {
		return nil, utils.NewAppError(
			err,
			http.StatusNotFound,
			"Article not found",
		)
	}

	if !user.IsAdmin() && article.UserID != user.ID {
		return nil, utils.NewAppError(
			fmt.Errorf("unauthorized"),
			http.StatusForbidden,
			"Unauthorized",
		)
	}

	article.Title = dto.Title
	article.Body = dto.Body

	updatedArticle, err := s.repo.Update(article)

	if err != nil {
		return nil, utils.NewAppError(
			err,
			http.StatusInternalServerError,
			"Failed to update article",
		)
	}

	return updatedArticle, nil
}

func (s *articleService) DeleteArticle(user *model.User, articleId string) error {
	article, err := s.repo.FindBySecureID(articleId)

	if err != nil {
		return utils.NewAppError(
			err,
			http.StatusNotFound,
			"Article not found",
		)
	}

	if !user.IsAdmin() && article.UserID != user.ID {
		return utils.NewAppError(
			fmt.Errorf("unauthorized"),
			http.StatusForbidden,
			"Unauthorized",
		)
	}

	if err := s.repo.Delete(article); err != nil {
		return utils.NewAppError(
			err,
			http.StatusInternalServerError,
			"Failed to delete article",
		)
	}

	return nil
}
