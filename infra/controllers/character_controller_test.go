package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schweigert/mmosandbox/domain/entities"

	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/infra/db"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/routes"

	"github.com/stretchr/testify/suite"
)

type CharacterControllerSuite struct {
	suite.Suite
}

func (suite *CharacterControllerSuite) SetupTest() {
	dbrepositories.UseAccountRepository()
	dbrepositories.UseCharacterRepository()
}

func (suite *CharacterControllerSuite) TestNewAccountController() {
	suite.NotNil(NewAccountController())
}

func (suite *CharacterControllerSuite) TestRoutes() {
	engine := gin.Default()
	NewAccountController().Routes(engine)

	suite.Equal(1, len(engine.Routes()))
}

func (suite *CharacterControllerSuite) TestCreate() {
	db.Clear()
	account := entities.NewAccount()
	account.SecurePassword = "a super hash"
	account.Username = "a super user"
	suite.True(domain.AccountRepository.Create(account))

	input := inputs.NewCreateCharacterInput()
	input.Auth.SecurePassword = account.SecurePassword
	input.Auth.Username = account.Username
	input.Name = "YAHO! Character Name!"

	inputBody, err := json.Marshal(input)

	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPut, routes.Character, bytes.NewReader(inputBody))
	req.Header.Set("Content-Type", "application/json")

	suite.NoError(err)

	rec := httptest.NewRecorder()
	engine := gin.Default()
	NewCharacterController().Routes(engine)
	engine.ServeHTTP(rec, req)

	suite.Equal(http.StatusCreated, rec.Code)
}

func TestCharacterControllerSuite(t *testing.T) {
	suite.Run(t, new(CharacterControllerSuite))
}
