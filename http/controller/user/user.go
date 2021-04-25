package user

import (
	"net/http"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

// Show a user
func Show(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	result := database.Conn().Where("email = ?", request.Params["email"]).First(&user)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, user)
	}
}

// Store a user
func Store(response *goyave.Response, request *goyave.Request) {
	code, err := GenerateRandomString(8)
	if err != nil {
		panic(err)
	}
	user := model.User{
		Email:            request.String("email"),
		Password:         request.String("password"),
		FirstName:        request.String("first_name"),
		LastName:         request.String("last_name"),
		Verified:         false,
		VerificationCode: code,
	}
	if err := database.GetConnection().Create(&user).Error; err != nil {
		response.Error(err)
	} else {
		SendEmail(request.String("email"), "Prix-Banque Account Verification", "Please proceed to email verification using the code : "+code)
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id": user.ID,
		})
	}
}

// Email verification for a user
func Verify(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	result := database.Conn().Where("email = ?", request.Params["email"]).First(&user)
	if response.HandleDatabaseError(result) {
		if user.VerificationCode == request.Params["verification_code"] {
			if err := database.Conn().Model(&user).Updates(model.User{
				Verified: true,
			}).Error; err != nil {
				response.Error(err)
			}
			SendEmail(request.Params["email"], "Prix-Banque Account Verified", "Your account is now verified, you can create a bank account")
		} else {
			response.Status(http.StatusBadRequest)
		}
	}
}

// Update a user
func Update(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	db := database.GetConnection()
	result := db.Where("email = ?", request.Params["email"]).First(&user)
	if response.HandleDatabaseError(result) {
		if err := db.Model(&user).Updates(model.User{
			Email:     request.String("email"),
			FirstName: request.String("first_name"),
			LastName:  request.String("last_name"),
		}).Error; err != nil {
			response.Error(err)
		}
	}
}

// Destroy a user
func Destroy(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	db := database.Conn()
	result := db.Where("email = ?", request.Params["email"]).First(&user)
	if response.HandleDatabaseError(result) {
		if err := db.Delete(&user).Error; err != nil {
			response.Error(err)
		}
	}
}
