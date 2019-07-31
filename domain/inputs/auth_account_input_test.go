package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthAccountInputSuite struct {
	suite.Suite
}

func (suite *AuthAccountInputSuite) TestNewAuthAccountInput() {
	input := NewAuthAccountInput()
	suite.NotNil(input)
	suite.Empty(input.Username)
	suite.Empty(input.SecurePassword)
}

func (suite *AuthAccountInputSuite) TestBuildAccount() {
	input := NewAuthAccountInput()

	input.Username = "testing_fake_onion"
	input.SecurePassword = "a super secure hash"

	account := input.BuildAccount()

	suite.Equal("testing_fake_onion", account.Username)
	suite.Equal("a super secure hash", account.SecurePassword)
}

func TestAuthAccountInputSuite(t *testing.T) {
	suite.Run(t, new(AuthAccountInputSuite))
}
