package db

import (
	"gorm.io/gorm"
	model2 "ordent-test/internal/domain/model"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model2.Comment{},
		&model2.CommentLike{},
		&model2.Article{},
		&model2.User{},
	)

	return err
}
