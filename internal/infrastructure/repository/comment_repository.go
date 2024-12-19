package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
)

type CommentRepository interface {
	FindBySecureID(secureID string) (*model.Comment, error)
	Create(comment *model.Comment) (*model.Comment, error)
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
		First(&comment).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) Create(comment *model.Comment) (*model.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return r.FindBySecureID(comment.SecureID)
}
