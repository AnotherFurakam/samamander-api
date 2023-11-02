package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	IdPost   uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title    string         `gorm:"type:varchar(200);not null;index"`
	Body     string         `gorm:"type:varchar(700);not null"`
	IsActive bool           `gorm:"type:bool;default:true"`
	CreateAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

type GetPostDto struct {
	IdPost   uuid.UUID `json:"idPost"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	IsActive bool      `json:"isActive"`
	CreateAt time.Time `json:"createAt"`
}

type PostDto struct {
	Title string `json:"title" validate:"required,min=5,max=200"`
	Body  string `json:"body" validate:"required,min=5,max=700"`
}
