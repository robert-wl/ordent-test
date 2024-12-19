package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);uniqueIndex;not null"`
	ArticleID *uint     `json:"article_id" gorm:"index;default:null"`
	ParentID  *uint     `json:"reply_id" gorm:"index;default:null"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	Body      string    `json:"body" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	Article *Article `json:"article" gorm:"foreignKey:ArticleID"`
	Comment *Comment `json:"comment" gorm:"foreignKey:ParentID"`
	User    User     `json:"user" gorm:"foreignKey:UserID"`
}

func (a *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	a.SecureID = uuid.New().String()
	return
}
