package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
)

type CommentRepository interface {
	FindBySecureID(secureID string) (*model.Comment, error)
	FindByArticleID(articleID uint) ([]*model.Comment, error)
	Create(comment *model.Comment) (*model.Comment, error)
	Delete(comment *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) FindBySecureID(secureID string) (*model.Comment, error) {
	var comment model.Comment

	err := r.db.Where("secure_id = ?", secureID).
		Preload("User").
		Preload("ReplyComments").
		First(&comment).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) FindByArticleID(articleID uint) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.db.Where("article_id = ?", articleID).
		Preload("User").
		Preload("ReplyComments").
		Preload("ReplyComments.User").
		Find(&comments).Error

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) Create(comment *model.Comment) (*model.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return r.FindBySecureID(comment.SecureID)
}

func (r *commentRepository) Delete(comment *model.Comment) error {
	return r.db.Delete(comment).Error
}
