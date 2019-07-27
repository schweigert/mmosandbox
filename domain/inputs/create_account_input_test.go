package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateAccountInputSuite struct {
	suite.Suite
}

func (suite *CreateAccountInputSuite) TestNewCreateAccountInput() {
	input := NewCreateAccountInput()
	suite.NotNil(input)
	suite.Empty(input.Username)
	suite.Empty(input.Password)
	suite.Empty(input.Email)
}

func (suite *CreateAccountInputSuite) TestValid() {
	input := NewCreateAccountInput()

	input.Email = "testing@fake.onion"
	input.Username = "testing_fake_onion"
	input.Password = "testing_fake_pass"

	suite.True(input.Valid())
}

func (suite *CreateAccountInputSuite) TestBuildAccount() {
	input := NewCreateAccountInput()

	input.Email = "testing@fake.onion"
	input.Username = "testing_fake_onion"
	input.Password = "testing_fake_pass"

	account := input.BuildAccount()

	suite.Equal(input.Username, account.Username)
	suite.Equal(input.Email, account.Email)
	suite.Equal("be9b48185c3be4fe4098f00d4146e51b099eecb79e62746b9ed5cf862bc4d32f", account.SecurePassword)
}

func TestCreateAccountInputSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountInputSuite))
}
