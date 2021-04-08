package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3/database"
	"syreclabs.com/go/faker"
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

func BankGenerator() interface{} {
	bank := &Bank{}
	UserID := database.NewFactory(UserGenerator).Save(1).([]*User)[0].ID
	FakeFloat, _ := strconv.ParseFloat(faker.Number().Decimal(8, 2), 64)
	bank.Balance = FakeFloat
	bank.UserID = UserID
	return bank
}
