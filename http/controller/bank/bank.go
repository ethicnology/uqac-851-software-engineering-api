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
		response.Status(http.StatusBadRequest)
	} else {
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
	outcomeInvoices := []model.Operation{}
	database.Conn().Where(&model.Operation{SenderID: id, Invoice: true, Acquitted: true}).Find(&outcomeInvoices)
	outcomeTransfers := []model.Operation{}
	database.Conn().Where(&model.Operation{SenderID: id, Transfer: true}).Find(&outcomeTransfers)
	incomeInvoices := []model.Operation{}
	database.Conn().Where(&model.Operation{BankID: id, Invoice: true, Acquitted: true}).Find(&incomeInvoices)
	incomeTransfers := []model.Operation{}
	database.Conn().Where(&model.Operation{BankID: id, Transfer: true, Instant: true, Scheduled: false, Verified: true}).Or(&model.Operation{BankID: id, Transfer: true, Scheduled: true, Instant: false, Verified: true}).Find(&incomeTransfers)
	for _, outcomeInvoice := range outcomeInvoices {
		balance -= outcomeInvoice.Amount
	}
	for _, outcomeTransfer := range outcomeTransfers {
		balance -= outcomeTransfer.Amount
	}
	for _, incomeInvoice := range incomeInvoices {
		balance += incomeInvoice.Amount
	}

	today := time.Now()
	for _, incomeTransfer := range incomeTransfers {
		if incomeTransfer.Scheduled {
			if incomeTransfer.Date.Before(today) || incomeTransfer.Date == today {
				balance += incomeTransfer.Amount
			}
		}
		if incomeTransfer.Instant {
			balance += incomeTransfer.Amount
		}
	}
	return balance
}
