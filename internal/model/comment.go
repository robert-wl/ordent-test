package model

type Comment struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	SecureID  string `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	ArticleID uint   `json:"article_id" gorm:"index;not null"`
	ReplyID   uint   `json:"reply_id" gorm:"index"`
	UserID    uint   `json:"user_id" gorm:"index;not null"`
	Body      string `json:"body" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
