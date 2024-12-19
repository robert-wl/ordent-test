package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	Username  string    `json:"username" gorm:"unique_index;not null"`
	Email     string    `json:"email" gorm:"unique_index;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}
