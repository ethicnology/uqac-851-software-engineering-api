package model

import (
	"time"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&Bank{})
}

// Bank represents a bank account
type Bank struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	ID        uint64         `json:"id" gorm:"primarykey"`
	Balance   float64        `json:"balance"`
	UserID    uint64         `json:"user_id" gorm:"not null"`
}
