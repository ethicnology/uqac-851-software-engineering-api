package bank

import (
	"net/http"
	"strconv"

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
		response.JSON(http.StatusOK, bank)
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

// Update a bank account
func Update(response *goyave.Response, request *goyave.Request) {
	id, _ := strconv.ParseUint(request.Params["bank_id"], 10, 64)
	bank := model.Bank{}
	db := database.GetConnection()
	result := db.Select("id").Where(&model.Bank{ID: id, UserID: request.Extra["UserID"].(uint64)}).First(&bank, id)
	if response.HandleDatabaseError(result) {
		if err := db.Model(&bank).Updates(model.Bank{
			Balance: request.Numeric("balance"),
		}).Error; err != nil {
			response.Error(err)
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
