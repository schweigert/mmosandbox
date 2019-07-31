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
	return nil
}
