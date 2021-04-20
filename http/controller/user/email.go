package user

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown from server")
		}
	}
	return nil, nil
}

func SendEmail(address string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fromAddress := os.Getenv("MAIL_USER")
	toAddresses := []string{address}

	msg := "From: " + fromAddress + "\n" +
		"To: " + toAddresses[0] + "\n" +
		"Subject: Welcome to Prix-Banque\n\n" +
		"You've just created a Prix-Banque account !"

	auth := LoginAuth(os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
	err = smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, fromAddress, toAddresses, []byte(msg))
	if err != nil {
		fmt.Print(err)
	}
}
