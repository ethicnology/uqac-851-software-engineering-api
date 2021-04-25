package invoice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/bank"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/user"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type Invoice struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint64    `json:"id"`
	Amount    float64   `json:"amount"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	SenderID  uint64    `json:"sender_id"`
	BankID    uint64    `json:"receiver_id"`
	Invoice   bool      `json:"invoice"`
	Acquitted bool      `json:"acquitted"`
	DueDate   time.Time `json:"due_date"`
}

// Index invoices operations for a bank account
func Index(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	invoices := []Invoice{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{Invoice: true, Transfer: false, BankID: userBankId}).Or(&model.Operation{Invoice: true, Transfer: false, SenderID: userBankId}).Find(&invoices)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, invoices)
	}
}

// Show invoice operation for a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	id, _ := strconv.ParseUint(request.Params["invoice_id"], 10, 64)
	invoice := Invoice{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: true, Transfer: false, BankID: userBankId}).Or(&model.Operation{ID: id, Invoice: true, Transfer: false, SenderID: userBankId}).First(&invoice, id)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, invoice)
	}
}

// Store invoice operation for a bank account
func Store(response *goyave.Response, request *goyave.Request) {
	// Verify if receiver exist
	owing := model.User{}
	result := database.Conn().Where("email = ?", request.String("to")).First(&owing)
	if response.HandleDatabaseError(result) {
		// Get the Bank account ID of receiver
		owingBank := model.Bank{}
		database.Conn().Where(&model.Bank{UserID: owing.ID}).First(&owingBank)
		// Construct operation object
		operation := model.Operation{
			Transfer:  false,
			Invoice:   true,
			Amount:    request.Numeric("amount"),
			From:      owing.Email,
			To:        request.Extra["UserEmail"].(string),
			SenderID:  owingBank.ID,
			BankID:    request.Extra["BankID"].(uint64),
			Acquitted: false,
			DueDate:   request.Date("due_date"),
		}
		if err := database.GetConnection().Create(&operation).Error; err != nil {
			response.Error(err)
		} else {
			operationId := strconv.FormatUint(operation.ID, 10)
			amount := fmt.Sprintf("%f", operation.Amount)
			response.JSON(http.StatusCreated, map[string]interface{}{
				"id": operation.ID,
			})
			user.SendEmail(operation.From, "Prix-Banque Invoice n"+operationId+" To Fullfill", "From "+operation.From+" to "+operation.To+" amount "+amount+"$")
			user.SendEmail(operation.To, "Prix-Banque Invoice n"+operationId+" Created", "From "+operation.From+" to "+operation.To+" amount "+amount+"$")
		}
	}
}

func Update(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	id, _ := strconv.ParseUint(request.Params["invoice_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: true, Transfer: false, SenderID: userBankId}).First(&operation)
	if response.HandleDatabaseError(result) {
		operationId := strconv.FormatUint(operation.ID, 10)
		amount := fmt.Sprintf("%f", operation.Amount)
		balance := bank.ComputeBalance(userBankId)
		if operation.Amount > balance {
			response.Status(http.StatusBadRequest)
		} else {
			if err := database.Conn().Model(&operation).Updates(model.Operation{
				Acquitted: request.Bool("acquitted"),
			}).Error; err != nil {
				response.Error(err)
			}
			user.SendEmail(operation.From, "Prix-Banque Invoice n"+operationId+" Completed", "From "+operation.From+" to "+operation.To+" amount "+amount+"$")
			user.SendEmail(operation.To, "Prix-Banque Invoice n"+operationId+" Completed", "From "+operation.From+" to "+operation.To+" amount "+amount+"$")
		}
	}
}

// Destroy invoice operation
func Destroy(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["invoice_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Where(&model.Operation{ID: id, Invoice: true, Transfer: false, Acquitted: false, BankID: request.Extra["BankID"].(uint64)}).First(&operation)
	if response.HandleDatabaseError(result) {
		if err := database.Conn().Delete(&operation).Error; err != nil {
			response.Error(err)
		}
	}
}
