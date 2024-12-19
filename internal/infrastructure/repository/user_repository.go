package repository

import (
	"gorm.io/gorm"
	"ordent-test/internal/domain/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmailAndPassword(email, password string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmailAndPassword(email, password string) (*model.User, error) {
	var user model.User

	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
