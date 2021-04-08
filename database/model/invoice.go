package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3/database"
	"syreclabs.com/go/faker"
)

func init() {
	database.RegisterModel(&Invoice{})
}

// Invoice represents a invoice account
type Invoice struct {
	InvoiceId        uint64       		   `json:"id" gorm:"primarykey"`
	Operation 		 model.Operation       `json:"operation" gorm:"not null"`
	Acquitted		 bool             	   `json:"acquitted"`
	DueDate		  	 time.Time			   `json:"due_date"`
}

func InvoiceGenerator() interface{} {
	invoice := &Invoice{}
	Operation := database.NewFactory(OperationGenerator).Save(1).([]*Operation)[0]				// Todo : Only ID of Operation ?
	typeInvoice := faker.Bool()
	invoice.Operation = Operation
	invoice.Acquitted = typeInvoice
	return invoice
}
