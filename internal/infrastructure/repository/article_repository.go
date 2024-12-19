package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
)

type ArticleRepository interface {
	Create(article *model.Article) (*model.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) Create(article *model.Article) (*model.Article, error) {
	if err := r.db.Create(article).Error; err != nil {
		return nil, err
	}

	return article, nil
}
