package model

import "time"

type Article struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	Views     uint      `json:"views" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}
