package bank

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

// Index banks account for a user
func Index(response *goyave.Response, request *goyave.Request) {
	banks := []model.Bank{}
	result := database.Conn().Where(&model.Bank{UserID: request.Extra["UserID"].(uint64)}).Find(&banks)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, banks)
	}
}

// Show a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["bank_id"], 10, 64)
	bank := model.Bank{}
	result := database.Conn().Where(&model.Bank{ID: id, UserID: request.Extra["UserID"].(uint64)}).First(&bank)
	if response.HandleDatabaseError(result) {
		computedBalance := ComputeBalance(bank.ID)
		if err := database.Conn().Model(&bank).Updates(model.Bank{
			Balance: computedBalance,
		}).Error; err != nil {
			response.Error(err)
		}
		response.JSON(http.StatusOK, map[string]interface{}{
			"id ":     bank.ID,
			"balance": computedBalance,
		})
	}
}

// Store a bank account
func Store(response *goyave.Response, request *goyave.Request) {
	// Email verificatoin
	user := model.User{}
	database.Conn().Where("email = ?", request.Params["email"]).First(&user)
	if !user.Verified { // if verified is false
		response.Status(http.StatusForbidden)
	}
	// Up to 10,000 first customers are credited with 1000.00$ on their account
	var count int64
	var balance float64
	database.Conn().Model(&model.Bank{}).Count(&count)
	if count <= 10000 {
		balance = 1000.00
	} else {
		balance = 0.0
	}
	// Create bank account
	bank := model.Bank{
		Balance: balance,
		UserID:  request.Extra["UserID"].(uint64),
	}
	if err := database.GetConnection().Create(&bank).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id ": bank.ID,
		})
	}
}

// Destroy a bank account
func Destroy(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["bank_id"], 10, 64)
	bank := model.Bank{}
	db := database.Conn()
	result := db.Select("id").Where(&model.Bank{ID: id, UserID: request.Extra["UserID"].(uint64)}).First(&bank)
	if response.HandleDatabaseError(result) {
		if err := db.Delete(&bank).Error; err != nil {
			response.Error(err)
		}
	}
}

func ComputeBalance(id uint64) float64 {
	balance := 0.0
	if id <= 10000 {
		balance += 1000.0
	}
	spent := []model.Operation{}
	database.Conn().Where(&model.Operation{SenderID: id}).Find(&spent)
	acquittedInvoices := []model.Operation{}
	database.Conn().Where(&model.Operation{BankID: id, Acquitted: true}).Find(&acquittedInvoices)
	receivedTransfers := []model.Operation{}
	database.Conn().Where(&model.Operation{BankID: id, Instant: true, Scheduled: false, Verified: true}).Or(&model.Operation{BankID: id, Scheduled: true, Instant: false, Verified: true}).Find(&receivedTransfers)
	for _, outcome := range spent {
		balance -= outcome.Amount
	}
	for _, invoice := range acquittedInvoices {
		balance += invoice.Amount
	}
	today := time.Now()

	for _, transfer := range receivedTransfers {
		if transfer.Scheduled {
			if transfer.Date.Before(today) || transfer.Date == today {
				balance += transfer.Amount
			}
		}
		if transfer.Instant {
			balance += transfer.Amount
		}
	}
	return balance
}
