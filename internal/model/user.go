package model

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	SecureID  string `json:"secure_id" gorm:"type:char(36);unique_index;not null"`
	Username  string `json:"username" gorm:"unique_index;not null"`
	Email     string `json:"email" gorm:"unique_index;not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
