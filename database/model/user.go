package model

import (
	"reflect"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/config"
	"goyave.dev/goyave/v3/database"
	"goyave.dev/goyave/v3/middleware/ratelimiter"
)

func init() {
	database.RegisterModel(&User{})

	config.Register("app.bcryptCost", config.Entry{
		Value:            10,
		Type:             reflect.Int,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})
}

// User represents a user
type User struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	ID        uint           `json:"id" gorm:"primarykey"`
	Email     string         `json:"email" gorm:"type:char(100);uniqueIndex;not null" auth:"username"`
	Password  string         `json:"-" gorm:"size:64;not null" auth:"password"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
}

// BeforeCreate hook executed before a User record is inserted in the database.
// Ensures the password is encrypted using bcrypt, with the cost defined by the
// config entry "app.bcryptCost".
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.bcryptPassword(tx)
}

// BeforeUpdate hook executed before a User record is updated in the database.
// Ensures the password is encrypted using bcrypt, with the cost defined by the
// config entry "app.bcryptCost".
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") {
		return u.bcryptPassword(tx)
	}
	return nil
}

func (u *User) bcryptPassword(tx *gorm.DB) error {
	var newPass string
	switch u := tx.Statement.Dest.(type) {
	case map[string]interface{}:
		newPass = u["password"].(string)
	case *User:
		newPass = u.Password
	case []*User:
		newPass = u[tx.Statement.CurDestIndex].Password
	}

	b, err := bcrypt.GenerateFromPassword([]byte(newPass), config.GetInt("app.bcryptCost"))
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("password", b)
	return nil
}

// RateLimiterFunc returns rate limiting configuration
// Anonymous users have a quota of 50 requests per minute while
// Authenticated users are limited to 500 requests per minute.
func RateLimiterFunc(request *goyave.Request) ratelimiter.Config {
	var id interface{} = nil
	quota := 50
	if request.User != nil {
		id = request.User.(*User).ID
		quota = 500
	}
	return ratelimiter.Config{
		ClientID:      id,
		RequestQuota:  quota,
		QuotaDuration: time.Minute,
	}
}
