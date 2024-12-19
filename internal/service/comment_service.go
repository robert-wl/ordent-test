package service

import (
	"fmt"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
)

type CommentService interface {
	CreateComment(user *model.User, dto *dto.CreateCommentRequest) (*model.Comment, error)
}

type commentService struct {
	commentRepo repository.CommentRepository
	articleRepo repository.ArticleRepository
}

func NewCommentService(cr repository.CommentRepository, ar repository.ArticleRepository) CommentService {
	return &commentService{
		commentRepo: cr,
		articleRepo: ar,
	}
}

func (s *commentService) CreateComment(user *model.User, dto *dto.CreateCommentRequest) (*model.Comment, error) {
	if dto.ArticleID == nil && dto.ParentID == nil {
		return nil, fmt.Errorf("either article or parent comment is required")
	}

	var article *uint
	var parent *uint

	if dto.ArticleID != nil {
		fArticle, err := s.articleRepo.FindBySecureID(*dto.ArticleID)

		if err != nil {
			return nil, fmt.Errorf("article with id %s not found", *dto.ArticleID)
		}

		article = &fArticle.ID
	}

	if dto.ParentID != nil {
		fParent, err := s.commentRepo.FindBySecureID(*dto.ParentID)

		if err != nil {
			return nil, fmt.Errorf("parent comment with id %s not found", *dto.ParentID)
		}

		parent = &fParent.ID
	}

	comment := &model.Comment{
		ArticleID: article,
		ParentID:  parent,
		UserID:    user.ID,
		Title:     dto.Title,
		Body:      dto.Body,
	}

	return s.commentRepo.Create(comment)
}
