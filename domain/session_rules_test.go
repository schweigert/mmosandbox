package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SessionRulesSuite struct {
	suite.Suite
}

func (suite *SessionRulesSuite) TestCreateAccount() {
	AccountRepository = &AccountRepositoryMock{
		UsernameHasTakenResult: false,
		CreateResult:           true,
	}

	input := NewCreateAccountInput()
	input.Username = "testingUser"
	input.Password = "testingPass"
	input.Email = "testing@test.onion"

	out := NewSessionRules().CreateAccount(input)

	suite.NotNil(out)
	suite.NotNil(out.Account)
	suite.Equal(input.Username, out.Account.Username)
	suite.Equal(input.Email, out.Account.Email)

	suite.NotEmpty(out.Account.SecurePassword)

	suite.True(out.Success)
	suite.Equal(http.StatusCreated, out.HTTPCode())
}

func TestSessionRulesSuite(t *testing.T) {
	suite.Run(t, new(SessionRulesSuite))
}
