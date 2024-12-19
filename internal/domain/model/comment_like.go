package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CommentLike struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);uniqueIndex;not null"`
	CommentID uint      `json:"-" gorm:"index;not null"`
	UserID    uint      `json:"-" gorm:"index;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	Comment Comment `json:"comment" gorm:"foreignKey:CommentID"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
}

func (a *CommentLike) BeforeCreate(tx *gorm.DB) (err error) {
	a.SecureID = uuid.New().String()
	return
}
