package domain

import (
	"net/http"
	"testing"

	"github.com/schweigert/mmosandbox/domain/outputs"

	"github.com/schweigert/mmosandbox/domain/inputs"
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

	input := inputs.NewCreateAccountInput()
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

func (suite *SessionRulesSuite) TestAllowAuthentication() {
	AccountRepository = &AccountRepositoryMock{
		UsernameAndPasswordAreEqualResult: true,
	}

	input := inputs.NewAuthAccountInput()
	input.Username = "testing_username"
	input.SecurePassword = "ABCDEFGHIJKLMNOPQRSTUWXZ"

	suite.True(NewSessionRules().AllowAuthentication(input))

	AccountRepository = &AccountRepositoryMock{
		UsernameAndPasswordAreEqualResult: false,
	}

	suite.False(NewSessionRules().AllowAuthentication(input))
}

func (suite *SessionRulesSuite) TestStartSession() {
	AccountRepository = &AccountRepositoryMock{
		UsernameAndPasswordAreEqualResult: true,
	}

	TokenRepository = &TokenRepositoryMock{
		GenerateTokenResult: "test ok",
	}

	input := inputs.NewAuthAccountInput()
	input.Username = "testing_username"
	input.SecurePassword = "ABCDEFGHIJKLMNOPQRSTUWXZ"
	out := outputs.NewStartSessionOutput()

	NewSessionRules().StartSession(input, out)
	suite.True(out.Success)
	suite.Equal("test ok", out.Token)
}

func (suite *SessionRulesSuite) TestCheckSession() {
	AccountRepository = &AccountRepositoryMock{
		UsernameHasTakenResult: true,
	}

	TokenRepository = &TokenRepositoryMock{
		CheckUsernameResult: true,
	}

	in := inputs.NewCheckSessionInput()
	suite.True(NewSessionRules().CheckSession(in))

	AccountRepository = &AccountRepositoryMock{
		UsernameHasTakenResult: false,
	}

	TokenRepository = &TokenRepositoryMock{
		CheckUsernameResult: true,
	}

	in = inputs.NewCheckSessionInput()
	suite.False(NewSessionRules().CheckSession(in))

	AccountRepository = &AccountRepositoryMock{
		UsernameHasTakenResult: true,
	}

	TokenRepository = &TokenRepositoryMock{
		CheckUsernameResult: false,
	}

	in = inputs.NewCheckSessionInput()
	suite.False(NewSessionRules().CheckSession(in))

	AccountRepository = &AccountRepositoryMock{
		UsernameHasTakenResult: false,
	}

	TokenRepository = &TokenRepositoryMock{
		CheckUsernameResult: false,
	}

	in = inputs.NewCheckSessionInput()
	suite.False(NewSessionRules().CheckSession(in))
}

func TestSessionRulesSuite(t *testing.T) {
	suite.Run(t, new(SessionRulesSuite))
}
