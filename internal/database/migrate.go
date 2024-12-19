package database

import (
	"gorm.io/gorm"
	"ordent-test/internal/model"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Comment{},
		&model.CommentLike{},
		&model.Article{},
		&model.User{},
	)

	return err
}
