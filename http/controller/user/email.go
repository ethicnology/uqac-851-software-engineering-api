package user

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
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

func SendEmail(address string, subject string, body string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fromAddress := os.Getenv("MAIL_USER")
	toAddresses := []string{address}

	msg := "From: " + fromAddress + "\n" +
		"To: " + address + "\n" +
		"Subject:" + subject + "\n\n" +
		body

	auth := LoginAuth(os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
	err = smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, fromAddress, toAddresses, []byte(msg))
	if err != nil {
		fmt.Print(err)
	}
}

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
