package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	SecureID  string    `json:"secure_id" gorm:"type:char(36);uniqueIndex;not null"`
	ArticleID *uint     `json:"-" gorm:"index;default:null"`
	ParentID  *uint     `json:"-" gorm:"index;default:null"`
	UserID    uint      `json:"-" gorm:"index;not null"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	Article       *Article   `json:"-" gorm:"foreignKey:ArticleID"`
	ParentComment *Comment   `json:"-" gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReplyComments []*Comment `json:"reply_comments,omitempty" gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User          User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (a *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	a.SecureID = uuid.New().String()
	return
}
