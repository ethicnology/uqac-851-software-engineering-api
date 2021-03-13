package user

import (
	"net/http"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

// Index users
func Index(response *goyave.Response, request *goyave.Request) {
	users := []model.User{}
	result := database.Conn().Find(&users)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, users)
	}
}

// Show a user
func Show(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	result := database.Conn().First(&user, request.Params["id"])
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, user)
	}
}

// Store a user
func Store(response *goyave.Response, request *goyave.Request) {
	user := model.User{
		Email:    request.String("email"),
		Password: request.String("password"),
	}
	if err := database.GetConnection().Create(&user).Error; err != nil {
		response.Error(err)
	} else {
		response.JSON(http.StatusCreated, map[string]interface{}{
			"id": user.ID,
		})
	}
}

// Update a user
func Update(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	db := database.GetConnection()
	result := db.Select("id").First(&user, request.Params["id"])
	if response.HandleDatabaseError(result) {
		if err := db.Model(&user).Updates(model.User{
			Email:     request.String("email"),
			Password:  request.String("password"),
			FirstName: request.String("first_name"),
			LastName:  request.String("last_name"),
		}).Error; err != nil {
			response.Error(err)
		}
	}
}

// Destroy an user
func Destroy(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	db := database.Conn()
	result := db.Select("id").First(&user, request.Params["id"])
	if response.HandleDatabaseError(result) {
		if err := db.Delete(&user).Error; err != nil {
			response.Error(err)
		}
	}
}
