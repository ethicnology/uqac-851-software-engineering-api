package transfer

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type Transfer struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint64    `json:"id"`
	Amount    float64   `json:"amount"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	SenderID  uint64    `json:"sender_id"`
	BankID    uint64    `json:"receiver_id"`
	Transfer  bool      `json:"transfer"`
	Instant   bool      `json:"instant"`
	Scheduled bool      `json:"scheduled"`
	Date      time.Time `json:"date"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	Verified  bool      `json:"verified"`
	Try       uint      `json:"try"`
}

// Index transfers operations for a bank account
func Index(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	transfers := []Transfer{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{Invoice: false, Transfer: true, BankID: userBankId}).Or(&model.Operation{Invoice: false, Transfer: true, SenderID: userBankId}).Find(&transfers)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, transfers)
	}
}

// Show transfer operation for a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	id, _ := strconv.ParseUint(request.Params["transfer_id"], 10, 64)
	transfer := Transfer{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: false, Transfer: true, BankID: userBankId}).Or(&model.Operation{ID: id, Invoice: false, Transfer: true, BankID: userBankId}).First(&transfer, id)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, transfer)
	}
}

// Store transfers operation for a bank account
func Store(response *goyave.Response, request *goyave.Request) {
	// Verify if receiver exist
	receiver := model.User{}
	result := database.Conn().Where("email = ?", request.String("to")).First(&receiver)
	if response.HandleDatabaseError(result) {
		// Get the Bank account ID of receiver
		receiverBank := model.Bank{}
		database.Conn().Where(&model.Bank{UserID: receiver.ID}).First(&receiverBank)
		// Construct operation object
		operation := model.Operation{
			Invoice:   false,
			Transfer:  true,
			From:      request.Extra["UserEmail"].(string),
			To:        receiver.Email,
			SenderID:  request.Extra["BankID"].(uint64),
			BankID:    receiverBank.ID,
			Amount:    request.Numeric("amount"),
			Instant:   request.Bool("instant"),
			Scheduled: request.Bool("scheduled"),
			Date:      request.Date("date"),
			Question:  request.String("question"),
			Answer:    request.String("answer"),
			Verified:  false,
			Try:       0,
		}
		if err := database.GetConnection().Create(&operation).Error; err != nil {
			response.Error(err)
		} else {
			response.JSON(http.StatusCreated, map[string]interface{}{
				"id": operation.ID,
			})
		}
	}
}

// Destroy transfer operation
func Destroy(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["transfer_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Where(&model.Operation{ID: id, Invoice: false, Transfer: true, BankID: request.Extra["BankID"].(uint64)}).First(&operation)
	if response.HandleDatabaseError(result) {
		if err := database.Conn().Delete(&operation).Error; err != nil {
			response.Error(err)
		}
	}
}
