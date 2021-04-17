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
	SenderID  uint64    `json:"sender_id"`
	BankID    uint64    `json:"receiver_id"`
	Transfer  bool      `json:"transfer"`
	Instant   bool      `json:"instant"`
	Scheduled bool      `json:"scheduled"`
	Date      time.Time `json:"date"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
}

// Index transfers operations for a bank account
func Index(response *goyave.Response, request *goyave.Request) {
	transfers := []Transfer{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{Invoice: false, Transfer: true, BankID: request.Extra["BankID"].(uint64)}).Find(&transfers)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, transfers)
	}
}

// Show transfer operation for a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["transfer_id"], 10, 64)
	transfer := Transfer{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: false, Transfer: true, BankID: request.Extra["BankID"].(uint64)}).First(&transfer, id)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, transfer)
	}
}

// Store transfers operation for a bank account
func Store(response *goyave.Response, request *goyave.Request) {
	operation := model.Operation{
		Invoice:   false,
		Transfer:  true,
		SenderID:  request.Extra["BankID"].(uint64),
		BankID:    uint64(request.Numeric("receiver_id")),
		Amount:    request.Numeric("amount"),
		Scheduled: request.Bool("scheduled"),
		Date:      request.Date("date"),
		Question:  request.String("question"),
		Answer:    request.String("answer"),
	}
	if err := database.GetConnection().Create(&operation).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id": operation.ID,
		})
	}
}

// Update transfer operation
func Update(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["transfer_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Model(model.Operation{}).Where(&model.Operation{ID: id, Invoice: false, Transfer: true, BankID: request.Extra["BankID"].(uint64)}).First(&operation)
	if response.HandleDatabaseError(result) {
		if err := database.Conn().Model(&operation).Updates(model.Operation{
			Amount:    request.Numeric("amount"),
			Scheduled: request.Bool("scheduled"),
			Date:      request.Date("date"),
			Question:  request.String("question"),
			Answer:    request.String("answer"),
		}).Error; err != nil {
			response.Error(err)
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
