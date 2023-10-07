package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	IdUser   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserName string    `gorm:"type:varchar(40);index"`
	Email    string    `gorm:"type:varchar(100);unique;index"`
	Password string    `gorm:"type:varchar(50)"`
	IsActive bool      `gorm:"type:bool;default:true"`
	CreateAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeleteAt time.Time `gorm:"index"`
}

type GetUserDto struct {
	IdUser   uuid.UUID `json:"idUser"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	IsActive bool      `json:"isActive"`
	CreateAt time.Time `json:"createAt"`
}

type CreateUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
