package model

import (
	"time"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&Operation{})
}

// Operation represents a operation account
type Operation struct {
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	ID         uint64         `json:"id" gorm:"primarykey"`
	Amount     float64        `json:"amount"`
	BankID     uint64         `json:"bank_id" gorm:"not null"` // Foreign key of table bank
	ReceiverID uint64         `json:"receiver_id" gorm:"not null"`
}

// func OperationGenerator() interface{} {
// 	operation := &Operation{}
// 	BankID := database.NewFactory(UserGenerator).Save(1).([]*User)[0].ID
// 	ReceiverID := database.NewFactory(UserGenerator).Save(1).([]*User)[0].ID
// 	FakeFloat, _ := strconv.ParseFloat(faker.Number().Decimal(8, 2), 64)
// 	operation.Amount = FakeFloat
// 	operation.ReceiverId = ReceiverID
// 	return operation
// }
