package service

import (
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
)

type ArticleService interface {
	CreateArticle(user *model.User, dto *dto.CreateArticleRequest) (*model.Article, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(r repository.ArticleRepository) ArticleService {
	return &articleService{
		repo: r,
	}
}

func (s *articleService) CreateArticle(user *model.User, dto *dto.CreateArticleRequest) (*model.Article, error) {
	article := &model.Article{
		UserID: user.ID,
		Title:  dto.Title,
		Body:   dto.Body,
	}

	createdArticle, err := s.repo.Create(article)

	if err != nil {
		return nil, err
	}

	return createdArticle, nil
}
