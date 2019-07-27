package dbrepositories

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/entities"

	"github.com/schweigert/mmosandbox/infra/db"

	"github.com/stretchr/testify/suite"
)

type AccountRepositorySuite struct {
	suite.Suite
}

func (suite *AccountRepositorySuite) SetupTest() {
	db.Clear()
	db.Migrate()
}

func (suite *AccountRepositorySuite) TestNewAccountRepository() {
	suite.NotPanics(func() {
		suite.NotNil(NewAccountRepository())
	})
}

func (suite *AccountRepositorySuite) TestUsernameHasTaken() {
	repository := NewAccountRepository()
	account := &entities.Account{Username: "MeninoNeymar"}

	suite.False(repository.UsernameHasTaken(account.Username))
	suite.True(repository.Create(account))
	suite.True(repository.UsernameHasTaken(account.Username))
}

func (suite *AccountRepositorySuite) TestCreate() {
	repository := NewAccountRepository()
	account := &entities.Account{Username: "MeninoNeymar"}

	suite.True(repository.Create(account))
}

func TestAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(AccountRepositorySuite))
}
