package test

import (
	"testing"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

type FactoryTestSuite struct {
	goyave.TestSuite
	UserID uint64
}

// TestUserGenerator unit test model.UserGenerator
func (suite *FactoryTestSuite) TestUserGenerator() {
	factory := database.NewFactory(model.UserGenerator)
	user := factory.Generate(1).([]*model.User)[0]
	override := &model.User{
		Email: "arsene@lupin.fr",
	}
	user = factory.Override(override).Generate(1).([]*model.User)[0]
	suite.Equal("arsene@lupin.fr", user.Email)
	suite.UserID = user.ID
}

// TestBankGenerator unit test model.BankGenerator
func (suite *FactoryTestSuite) TestBankGenerator() {
	factory := database.NewFactory(model.BankGenerator)
	bank := factory.Generate(1).([]*model.Bank)[0]
	override := &model.Bank{
		Balance: 100.5,
	}
	bank = factory.Override(override).Generate(1).([]*model.Bank)[0]
	suite.Equal(100.5, bank.Balance)
}

func TestFactorySuite(t *testing.T) {
	goyave.RunTest(t, new(FactoryTestSuite))
}
