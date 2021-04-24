package operation

import (
	"net/http"
	"strconv"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

// Index operations for a bank account
func Index(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	operations := []model.Operation{}
	result := database.Conn().Where(&model.Operation{BankID: userBankId}).Or(&model.Operation{SenderID: userBankId}).Find(&operations)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, operations)
	}
}

// Show operation for a bank account
func Show(response *goyave.Response, request *goyave.Request) {
	userBankId := request.Extra["BankID"].(uint64)
	id, _ := strconv.ParseUint(request.Params["operation_id"], 10, 64)
	operation := model.Operation{}
	result := database.Conn().Where(&model.Operation{ID: id, BankID: userBankId}).Or(&model.Operation{ID: id, SenderID: userBankId}).First(&operation, id)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, operation)
	}
}
