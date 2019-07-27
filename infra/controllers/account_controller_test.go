package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/infra/db"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/routes"

	"github.com/stretchr/testify/suite"
)

type AccountControllerSuite struct {
	suite.Suite
}

func (suite *AccountControllerSuite) SetupTest() {
	domain.AccountRepository = dbrepositories.NewAccountRepository()
}

func (suite *AccountControllerSuite) TestNewAccountController() {
	suite.NotNil(NewAccountController())
}

func (suite *AccountControllerSuite) TestRoutes() {
	engine := gin.Default()
	NewAccountController().Routes(engine)

	suite.Equal(1, len(engine.Routes()))
}

func (suite *AccountControllerSuite) TestCreate() {
	db.Clear()

	input := inputs.NewCreateAccountInput()

	input.Email = "testing@example.onion"
	input.Username = "a_very_secure_user"
	input.Password = "a very secure pass"

	inputBody, err := json.Marshal(input)

	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPut, routes.Account, bytes.NewReader(inputBody))
	req.Header.Set("Content-Type", "application/json")

	suite.NoError(err)

	rec := httptest.NewRecorder()
	engine := gin.Default()
	NewAccountController().Routes(engine)
	engine.ServeHTTP(rec, req)

	suite.Equal(http.StatusCreated, rec.Code)
}

func TestAccountControllerSuite(t *testing.T) {
	suite.Run(t, new(AccountControllerSuite))
}
