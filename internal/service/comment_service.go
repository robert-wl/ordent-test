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

type CommentService interface {
	GetComment(commentID string) (*model.Comment, error)
	GetArticleComments(articleID string, dto *dto.GetArticleCommentRequest) ([]*model.Comment, error)
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

func (s *commentService) GetArticleComments(articleID string, dto *dto.GetArticleCommentRequest) ([]*model.Comment, error) {
	article, err := s.articleRepo.FindBySecureID(articleID)

	if err != nil {
		return nil, utils.NewAppError(
			fmt.Errorf("article with id %s not found", articleID),
			http.StatusNotFound,
			"Article not found",
		)
	}

	if dto.Search == nil {
		dto.Search = new(string)
	}

	if dto.Pagination == nil {
		dto.Pagination = new(pagination.Pagination)
	}

	return s.commentRepo.FindByArticleID(article.ID, dto.Search, dto.Pagination)
}

func (s *commentService) CreateComment(user *model.User, dto *dto.CreateCommentRequest) (*model.Comment, error) {
	if dto.ArticleID == nil && dto.ParentID == nil {
		return nil, utils.NewAppError(
			fmt.Errorf("article_id or parent_id is required"),
			http.StatusBadRequest,
			"Bad Request",
		)
	}

	var article *uint
	var parent *uint

	if dto.ArticleID != nil {
		fArticle, err := s.articleRepo.FindBySecureID(*dto.ArticleID)

		if err != nil {
			return nil, utils.NewAppError(
				fmt.Errorf("article with id %s not found", *dto.ArticleID),
				http.StatusNotFound,
				"Article not found",
			)
		}

		article = &fArticle.ID
	}

	if dto.ParentID != nil {
		fParent, err := s.commentRepo.FindBySecureID(*dto.ParentID)

		if err != nil {
			return nil, utils.NewAppError(
				fmt.Errorf("parent comment with id %s not found", *dto.ParentID),
				http.StatusNotFound,
				"Parent comment not found",
			)
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
		return nil, utils.NewAppError(
			fmt.Errorf("comment with id %s not found", commentID),
			http.StatusNotFound,
			"Comment not found",
		)
	}

	if !user.IsAdmin() && comment.UserID != user.ID {
		return nil, utils.NewAppError(
			fmt.Errorf("unauthorized to update comment"),
			http.StatusForbidden,
			"Unauthorized",
		)
	}

	comment.Title = dto.Title
	comment.Body = dto.Body

	return s.commentRepo.Update(comment)
}

func (s *commentService) DeleteComment(user *model.User, commentID string) error {
	comment, err := s.commentRepo.FindBySecureID(commentID)

	if err != nil {
		return utils.NewAppError(
			fmt.Errorf("comment with id %s not found", commentID),
			http.StatusNotFound,
			"Comment not found",
		)
	}

	if !user.IsAdmin() && comment.UserID != user.ID {
		return utils.NewAppError(
			fmt.Errorf("unauthorized to delete comment"),
			http.StatusForbidden,
			"Unauthorized",
		)
	}

	return s.commentRepo.Delete(comment)
}
