package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/config"
	"goyave.dev/goyave/v3/database"
)

func Owner(next goyave.Handler) goyave.Handler {
	return func(response *goyave.Response, request *goyave.Request) {
		email, _ := request.Params["email"]

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
		}
		next(response, request) // Pass to the next handler
	}
}
