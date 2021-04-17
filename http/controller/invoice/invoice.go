package invoice

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type Invoice struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint64    `json:"id"`
	Amount    float64   `json:"amount"`
	SenderID  uint64    `json:"sender_id"`
	BankID    uint64    `json:"receiver_id"`
	Invoice   bool      `json:"invoice"`
	Acquitted bool      `json:"acquitted"`
	DueDate   time.Time `json:"due_date"`
}

// Index invoices operations for a bank account
func Index(response *goyave.Response, request *goyave.Request) {
	invoices := []Invoice{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{Invoice: true, Transfer: false, BankID: request.Extra["BankID"].(uint64)}).Find(&invoices)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, invoices)
	}
}

// Show invoice operation for a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["invoice_id"], 10, 64)
	invoice := Invoice{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: true, Transfer: false, BankID: request.Extra["BankID"].(uint64)}).First(&invoice, id)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, invoice)
	}
}

// Store invoice operation for a bank account
func Store(response *goyave.Response, request *goyave.Request) {
	operation := model.Operation{
		Transfer:  false,
		Invoice:   true,
		Amount:    request.Numeric("amount"),
		SenderID:  request.Extra["BankID"].(uint64),
		BankID:    uint64(request.Numeric("receiver_id")),
		Acquitted: request.Bool("acquitted"),
		DueDate:   request.Date("due_date"),
	}
	if err := database.GetConnection().Create(&operation).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id": operation.ID,
		})
	}
}

// Update invoice operation
func Update(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["invoice_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: true, Transfer: false, BankID: request.Extra["BankID"].(uint64)}).First(&operation)
	if response.HandleDatabaseError(result) {
		if err := database.Conn().Model(&operation).Updates(model.Operation{
			Amount:    request.Numeric("amount"),
			Acquitted: request.Bool("acquitted"),
			DueDate:   request.Date("due_date"),
		}).Error; err != nil {
			response.Error(err)
		}
	}
}

// Destroy invoice operation
func Destroy(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["invoice_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Where(&model.Operation{ID: id, Invoice: true, Transfer: false, BankID: request.Extra["BankID"].(uint64)}).First(&operation)
	if response.HandleDatabaseError(result) {
		if err := database.Conn().Delete(&operation).Error; err != nil {
			response.Error(err)
		}
	}
}
