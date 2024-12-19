package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
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
