package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3/database"
	"syreclabs.com/go/faker"
)

func init() {
	database.RegisterModel(&Transfer{})
}

// Transfer represents a transfer account
type Transfer struct {
	TransferId        uint64       		   `json:"id" gorm:"primarykey"`
	Operation 		  model.Operation      `json:"operation" gorm:"not null"`
	Instant			  bool                 `json:"instant"`
	Scheduled		  bool				   `json:"scheduled"`
	Question	 	  string			   `json:"question"`
	Answer  	 	  string			   `json:"answer"`
}

func TransferGenerator() interface{} {
	transfer := &Transfer{}
	Operation := database.NewFactory(OperationGenerator).Save(1).([]*Operation)[0]				// Todo : Only ID of Operation ?
	typeTransfer := faker.Bool()
	transfer.Operation = Operation
	transfer.Instant = typeTransfer
	transfer.Scheduled = !typeTransfer
	return transfer
}
