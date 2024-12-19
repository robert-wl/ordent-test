package service

import (
	"fmt"
	"ordent-test/internal/domain/model"
	"ordent-test/internal/dto"
	"ordent-test/internal/infrastructure/repository"
)

type CommentService interface {
	GetComment(commentID string) (*model.Comment, error)
	GetCommentsByArticleSecureID(articleID string) ([]*model.Comment, error)
	CreateComment(user *model.User, dto *dto.CreateCommentRequest) (*model.Comment, error)
	UpdateComment(user *model.User, commentID string, dto *dto.UpdateCommentRequest) (*model.Comment, error)
	DeleteComment(user *model.User, commentID string) error
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

func (s *commentService) GetComment(commentID string) (*model.Comment, error) {
	return s.commentRepo.FindBySecureID(commentID)
}

func (s *commentService) GetCommentsByArticleSecureID(articleID string) ([]*model.Comment, error) {
	article, err := s.articleRepo.FindBySecureID(articleID)

	if err != nil {
		return nil, fmt.Errorf("article with id %s not found", articleID)
	}

	return s.commentRepo.FindByArticleID(article.ID)
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

func (s *commentService) UpdateComment(user *model.User, commentID string, dto *dto.UpdateCommentRequest) (*model.Comment, error) {
	comment, err := s.commentRepo.FindBySecureID(commentID)

	if err != nil {
		return nil, fmt.Errorf("comment with id %s not found", commentID)
	}

	if comment.UserID != user.ID {
		return nil, fmt.Errorf("unauthorized to update comment")
	}

	comment.Title = dto.Title
	comment.Body = dto.Body

	return s.commentRepo.Update(comment)
}

func (s *commentService) DeleteComment(user *model.User, commentID string) error {
	comment, err := s.commentRepo.FindBySecureID(commentID)

	if err != nil {
		return fmt.Errorf("comment with id %s not found", commentID)
	}

	if comment.UserID != user.ID {
		return fmt.Errorf("unauthorized to delete comment")
	}

	return s.commentRepo.Delete(comment)
}
