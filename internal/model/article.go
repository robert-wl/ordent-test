package model

type Article struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	SecureID  string `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	UserID    uint   `json:"user_id" gorm:"index;not null"`
	Title     string `json:"title" gorm:"not null"`
	Body      string `json:"body" gorm:"not null"`
	Views     uint   `json:"views" gorm:"default:0"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
