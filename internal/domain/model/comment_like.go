package model

import "time"

type CommentLike struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	CommentID uint      `json:"comment_id" gorm:"index;not null"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	Comment Comment `json:"comment" gorm:"foreignKey:CommentID"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
}
