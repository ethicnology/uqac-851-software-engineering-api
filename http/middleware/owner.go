package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/config"
	"goyave.dev/goyave/v3/database"
)

func Owner(next goyave.Handler) goyave.Handler {
	return func(response *goyave.Response, request *goyave.Request) {
		email := request.Params["email"]

		// Get Bearer Token
		tokenString, _ := request.BearerToken()

		// Verify Bearer Token Signature
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.GetString("auth.jwt.secret")), nil
		})

		// Unpack claims from JWT
		claims, _ := token.Claims.(jwt.MapClaims)

		// Check if the User is the owner of the resource.
		if email != claims["userid"] {
			response.Status(http.StatusForbidden)
			return
		} else {
			user := model.User{}
			database.Conn().Where("email = ?", claims["userid"]).First(&user)
			request.Extra["UserID"] = user.ID
			request.Extra["UserEmail"] = user.Email
		}
		next(response, request) // Pass to the next handler
	}
}

func BankOwner(next goyave.Handler) goyave.Handler {
	return func(response *goyave.Response, request *goyave.Request) {
		id, _ := strconv.ParseUint(request.Params["bank_id"], 10, 64)
		bank := model.Bank{}
		if err := database.Conn().Where(&model.Bank{ID: id, UserID: request.Extra["UserID"].(uint64)}).First(&bank).Error; err != nil {
			response.Status(http.StatusForbidden)
		} else {
			request.Extra["BankID"] = bank.ID
		}
		next(response, request) // Pass to the next handler
	}
}
