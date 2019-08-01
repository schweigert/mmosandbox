package domain

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/inputs"

	"github.com/stretchr/testify/suite"
)

type AccountRulesSuite struct {
	suite.Suite
}

func (suite *AccountRulesSuite) TestNewAccountRules() {
	suite.NotNil(NewAccountRules())
}

func (suite *AccountRulesSuite) TestCreateCharacter() {
	CharacterRepository = &CharacterRepositoryMock{
		CreateInAccountResult: true,
		NameHasTakenResult:    false,
	}

	AccountRepository = &AccountRepositoryMock{
		UsernameAndPasswordAreEqualResult: true,
	}

	rule := NewAccountRules()

	in := inputs.NewCreateCharacterInput()
	in.Auth.Username = "username"
	in.Auth.SecurePassword = "secure hash"

	in.Name = "A Super Character Name"

	out := rule.CreateCharacter(in)

	suite.Equal(in.Auth.Username, out.Account.Username)
	suite.Equal(in.Auth.SecurePassword, out.Account.SecurePassword)
	suite.Equal(in.Name, out.Character.Name)
	suite.Equal(uint(1), out.Character.Level)
	suite.Equal(uint(0), out.Character.MapIndex)
	suite.Equal(int(0), out.Character.MapXPosition)
	suite.Equal(int(0), out.Character.MapYPosition)

	suite.True(out.Success)
}

func TestAccountRulesSuite(t *testing.T) {
	suite.Run(t, new(AccountRulesSuite))
}
