package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/route"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/database"
	_ "goyave.dev/goyave/v3/database/dialect/mysql"
)

type UserTestSuite struct {
	goyave.TestSuite
	UserID   uint64
	Email    string
	Password string
	Token    string
}

func (suite *UserTestSuite) SetupTest() {
	suite.ClearDatabase()
	factory := database.NewFactory(model.UserGenerator)
	override := &model.User{
		Email:    "murray@bookchin.org",
		Password: "a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11",
	}
	suite.Password = override.Password
	user := factory.Override(override).Save(1).([]*model.User)[0]
	suite.UserID = user.ID
	suite.Email = user.Email
	suite.Token, _ = auth.GenerateToken(user.Email)
}

func (suite *UserTestSuite) TestStore() {
	suite.RunServer(route.Register, func() {
		request := map[string]interface{}{
			"email":      "user1prixbanque@gmail.com",
			"password":   suite.Password,
			"first_name": "user1",
			"last_name":  "lambda",
		}
		body, _ := json.Marshal(request)
		headers := map[string]string{"Content-Type": "application/json"}
		resp, err := suite.Post("/auth/register", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(201, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) TestLogin() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		request := map[string]interface{}{
			"email":    suite.Email,
			"password": suite.Password,
		}
		body, _ := json.Marshal(request)
		resp, err := suite.Post("/auth/login", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) TestShow() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		resp, err := suite.Get("/users/"+suite.Email, headers)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) TestUpdate() {
	user := &model.User{
		Email:     suite.Email,
		FirstName: "Léon",
		LastName:  "Sokolov",
	}
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		body, _ := json.Marshal(user)
		resp, err := suite.Patch("/users/"+suite.Email, headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(204, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) TestDestroy() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		resp, err := suite.Delete("/users/"+suite.Email, headers, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(204, resp.StatusCode)
		}
	})
}

func TestUserSuite(t *testing.T) {
	goyave.RunTest(t, new(UserTestSuite))
}
