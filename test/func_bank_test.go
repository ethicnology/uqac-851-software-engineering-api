package test

import (
	"fmt"
	"testing"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/route"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/database"
)

type BankTestSuite struct {
	goyave.TestSuite
	UserID   uint64
	Email    string
	Password string
	Token    string
	BankID   uint64
	Balance  float64
}

func (suite *BankTestSuite) SetupTest() {
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

	bankFactory := database.NewFactory(model.BankGenerator)
	bankOverride := &model.Bank{
		UserID: suite.UserID,
	}
	bank := bankFactory.Override(bankOverride).Save(1).([]*model.Bank)[0]
	suite.BankID = bank.ID
	suite.Balance = bank.Balance
}

func (suite *BankTestSuite) TestStore() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		resp, err := suite.Post("/users/"+suite.Email+"/banks", headers, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			// 400 Because account should be email verified
			suite.Equal(400, resp.StatusCode)
		}
	})
}

func (suite *BankTestSuite) TestIndex() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		resp, err := suite.Get("/users/"+suite.Email+"/banks", headers)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *BankTestSuite) TestShow() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		resp, err := suite.Get("/users/"+suite.Email+"/banks/"+fmt.Sprint(suite.BankID), headers)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *BankTestSuite) TestDestroy() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + suite.Token,
		}
		resp, err := suite.Delete("/users/"+suite.Email+"/banks/"+fmt.Sprint(suite.BankID), headers, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(204, resp.StatusCode)
		}
	})
}

func TestBankSuite(t *testing.T) {
	goyave.RunTest(t, new(BankTestSuite))
}
