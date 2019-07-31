package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateCharacterInputSuite struct {
	suite.Suite
}

func (suite *CreateCharacterInputSuite) TestNewCreateCharacterInput() {
	input := NewCreateCharacterInput()
	suite.NotNil(input)
	suite.Empty(input.Name)
}

func (suite *CreateCharacterInputSuite) TestBuildAccount() {
	input := NewCreateCharacterInput()

	input.Auth.Username = "testing_fake_onion"
	input.Auth.SecurePassword = "a super secure hash"

	account := input.BuildAccount()

	suite.Equal("testing_fake_onion", account.Username)
	suite.Equal("a super secure hash", account.SecurePassword)
}

func (suite *CreateCharacterInputSuite) TestBuildCharacter() {
	input := NewCreateCharacterInput()

	input.Name = "Fake Lord Valdemorth"

	character := input.BuildCharacter()

	suite.Equal("Fake Lord Valdemorth", character.Name)
}

func TestCreateCharacterInputSuite(t *testing.T) {
	suite.Run(t, new(CreateCharacterInputSuite))
}
