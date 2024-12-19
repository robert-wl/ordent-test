package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);uniqueIndex;not null"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}

func (a *User) BeforeCreate(tx *gorm.DB) (err error) {
	a.SecureID = uuid.New().String()
	return
}
