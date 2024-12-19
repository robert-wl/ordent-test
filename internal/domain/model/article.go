package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);uniqueIndex;not null"`
	UserID    uint      `json:"-" gorm:"index;not null"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	Views     uint      `json:"views" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	a.SecureID = uuid.New().String()
	return
}
