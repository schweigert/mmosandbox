package domain

import (
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
)

// AccountRules operations over this account
type AccountRules struct{}

// NewAccountRules constructor
func NewAccountRules() *AccountRules {
	return &AccountRules{}
}

// CreateCharacter based in this input
func (rules *AccountRules) CreateCharacter(in *inputs.CreateCharacterInput) *outputs.CreateCharacterOutput {
	account := in.BuildAccount()
	character := in.BuildCharacter()

	output := outputs.NewCreateCharacterOutput()

	output.Account = account
	output.Character = character
	output.Auth = NewSessionRules().AllowAuthentication(&in.Auth)
	output.Success = output.Auth && CharacterRepository.CreateInAccount(character, account)

	return output
}
