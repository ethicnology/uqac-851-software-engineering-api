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
	user := model.User{}
	database.Conn().Where("email = ?", request.Params["email"]).First(&user)
	if !user.Verified { // if verified is false
		response.Status(http.StatusForbidden)
	}
	bank := model.Bank{
		Balance: request.Numeric("balance"),
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
