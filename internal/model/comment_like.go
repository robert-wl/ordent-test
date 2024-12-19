package model

type CommentLike struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	SecureID  string `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	CommentID uint   `json:"comment_id" gorm:"index;not null"`
	UserID    uint   `json:"user_id" gorm:"index;not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
