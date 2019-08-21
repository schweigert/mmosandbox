package dbweb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/routes"
	"github.com/stretchr/testify/suite"
)

type AccountRepositorySuite struct {
	suite.Suite
}

func (suite *AccountRepositorySuite) SetupTest() {
	dbrepositories.UseAccountRepository()
	dbrepositories.UseCharacterRepository()
}

func (suite *AccountRepositorySuite) TestNewAccountController() {
	suite.NotNil(NewAccountRepositoryHandler())
}

func (suite *AccountRepositorySuite) TestRoutes() {
	engine := gin.Default()
	NewAccountRepositoryHandler().Routes(engine)

	suite.Equal(1, len(engine.Routes()))
}

func (suite *AccountRepositorySuite) TestCreate() {
	db.Clear()

	account := entities.NewAccount()
	account.Username = "testing_now"

	suite.True(dbrepositories.NewAccountRepository().Create(account))

	inputBody, err := json.Marshal(account.Username)
	suite.NoError(err)
	req, err := http.NewRequest(http.MethodPost, routes.AccountRepositoryUsernameHasTaken, bytes.NewReader(inputBody))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	engine := gin.Default()
	NewAccountRepositoryHandler().Routes(engine)
	engine.ServeHTTP(rec, req)

	suite.Equal(http.StatusOK, rec.Code)
	hasTaken := false
	err = json.Unmarshal(rec.Body.Bytes(), &hasTaken)
	suite.NoError(err)
	suite.True(hasTaken)
}

func TestAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(AccountRepositorySuite))
}
