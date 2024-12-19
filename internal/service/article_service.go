package service

import (
	"fmt"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
)

type ArticleService interface {
	GetArticles() ([]*model.Article, error)
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

func (s *articleService) GetArticles() ([]*model.Article, error) {
	articles, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *articleService) GetArticle(articleId string) (*model.Article, error) {
	article, err := s.repo.FindBySecureID(articleId)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	return createdArticle, nil
}

func (s *articleService) UpdateArticle(user *model.User, articleId string, dto *dto.UpdateArticleRequest) (*model.Article, error) {
	article, err := s.repo.FindBySecureID(articleId)

	if err != nil {
		return nil, err
	}

	if article.UserID != user.ID {
		return nil, fmt.Errorf("unauthorized")
	}

	article.Title = dto.Title
	article.Body = dto.Body

	updatedArticle, err := s.repo.Update(article)

	if err != nil {
		return nil, err
	}

	return updatedArticle, nil
}

func (s *articleService) DeleteArticle(user *model.User, articleId string) error {
	article, err := s.repo.FindBySecureID(articleId)

	if err != nil {
		return err
	}

	if article.UserID != user.ID {
		return fmt.Errorf("unauthorized")
	}

	if err := s.repo.Delete(article); err != nil {
		return err
	}

	return nil
}
