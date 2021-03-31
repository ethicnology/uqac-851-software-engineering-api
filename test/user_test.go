package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/route"
	"goyave.dev/goyave/v3"
)

type UserTestSuite struct {
	goyave.TestSuite
}

func (suite *UserTestSuite) CreateUser() {
	suite.ClearDatabase()
	user := &model.User{
		Email:    "a@a.fr",
		Password: "a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11",
	}
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(user)
		resp, err := suite.Post("/auth/register", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(201, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) LoginUser() {
	user := &model.User{
		Email:    "a@a.fr",
		Password: "a441b15fe9a3cf56661190a0b93b9dec7d04127288cc87250967cf3b52894d11",
	}
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(user)
		resp, err := suite.Post("/auth/login", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) GetUser() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTMxNjMxODQsIm5iZiI6MTYxNzE2MzE4NCwidXNlcmlkIjoiYUBhLmZyIn0.z-wZpGB0B_vUjDCMLEBiB5HzST5imglHwEbHh5YoLgY"
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		}
		resp, err := suite.Get("/users/a@a.fr", headers)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) PatchUser() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTMxNjMxODQsIm5iZiI6MTYxNzE2MzE4NCwidXNlcmlkIjoiYUBhLmZyIn0.z-wZpGB0B_vUjDCMLEBiB5HzST5imglHwEbHh5YoLgY"
	user := &model.User{
		Email:     "a@a.fr",
		FirstName: "Léon",
		LastName:  "Sokolov",
	}
	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		}
		body, _ := json.Marshal(user)
		resp, err := suite.Patch("/users/a@a.fr", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(204, resp.StatusCode)
		}
	})
}

func (suite *UserTestSuite) DeleteUser() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTMxNjMxODQsIm5iZiI6MTYxNzE2MzE4NCwidXNlcmlkIjoiYUBhLmZyIn0.z-wZpGB0B_vUjDCMLEBiB5HzST5imglHwEbHh5YoLgY"

	suite.RunServer(route.Register, func() {
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		}
		resp, err := suite.Delete("/users/a@a.fr", headers, nil)
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
