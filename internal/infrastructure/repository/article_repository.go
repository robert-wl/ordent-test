package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
	"ordent-test/pkg/pagination"
)

type ArticleRepository interface {
	Find(search *string, pagination *pagination.Pagination) ([]*model.Article, error)
	FindBySecureID(secureID string) (*model.Article, error)
	Create(article *model.Article) (*model.Article, error)
	Update(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) Find(search *string, pagination *pagination.Pagination) ([]*model.Article, error) {
	var articles []*model.Article

	err := r.db.
		Scopes(pagination.Paginate()).
		Where("title LIKE ?", "%"+*search+"%").
		Where("body LIKE ?", "%"+*search+"%").
		Preload("User").
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepository) FindBySecureID(secureID string) (*model.Article, error) {
	var article model.Article

	err := r.db.Where("secure_id = ?", secureID).
		Preload("User").
		First(&article).Error

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *articleRepository) Create(article *model.Article) (*model.Article, error) {
	if err := r.db.Create(article).Error; err != nil {
		return nil, err
	}

	return r.FindBySecureID(article.SecureID)
}

func (r *articleRepository) Update(article *model.Article) (*model.Article, error) {
	if err := r.db.Save(article).Error; err != nil {
		return nil, err
	}

	return r.FindBySecureID(article.SecureID)
}

func (r *articleRepository) Delete(article *model.Article) error {
	return r.db.Delete(article).Error
}
