package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	IdProduct   uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string         `gorm:"type:varchar(300);not null;index;unique"`
	Description string         `gorm:"type:varchar(500);not null"`
	IsActive    bool           `gorm:"type:bool;default:true"`
	CreateAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
}

type GetProductDto struct {
	IdProduct   uuid.UUID `json:"idProduct"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"isActive"`
	CreateAt    time.Time `json:"createAt"`
}

type ProductDto struct {
	Name        string `json:"name" validate:"required,min=5,max=300"`
	Description string `json:"description" validate:"required,min=10,max=500"`
}
