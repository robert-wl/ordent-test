package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
	"ordent-test/pkg/pagination"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Find(search *string, pagination *pagination.Pagination) ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindBySecureID(secureID string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Find(search *string, pagination *pagination.Pagination) ([]*model.User, error) {
	var users []*model.User

	err := r.db.
		Scopes(pagination.Paginate()).
		Where("username LIKE ?", "%"+*search+"%").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindBySecureID(secureID string) (*model.User, error) {
	var user model.User

	err := r.db.Where("secure_id = ?", secureID).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
