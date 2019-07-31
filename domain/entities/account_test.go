package entities

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountSuite struct {
	suite.Suite
}

func (suite *AccountSuite) TestNewAccount() {
	account := NewAccount()

	suite.NotNil(account)
	suite.Equal("", account.Username)
	suite.Equal("", account.SecurePassword)
	suite.Equal("", account.Email)
	suite.Empty(account.Characters)
}

func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}
