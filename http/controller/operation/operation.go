package operation

import (
	"net/http"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

// Index operations for a bank account
func Index(response *goyave.Response, request *goyave.Request) {
	operations := []model.Operation{}
	result := database.Conn().Where(&model.Operation{BankID: request.Extra["BankID"].(uint64)}).Find(&operations)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, operations)
	}
}

// Show operation for a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	operation := model.Operation{}
	result := database.Conn().Where(&model.Operation{BankID: request.Extra["BankID"].(uint64)}).First(&operation)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, operation)
	}
}

// Store operation for a bank account
func Store(response *goyave.Response, request *goyave.Request) {
	operation := model.Operation{
		Amount:   request.Numeric("amount"),
		SenderID: request.Extra["BankID"].(uint64),
		BankID:   uint64(request.Numeric("receiver_id")),
	}
	if err := database.GetConnection().Create(&operation).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id ": operation.ID,
		})
	}
}
