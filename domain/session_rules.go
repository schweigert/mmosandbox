package domain

import (
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
)

// SessionRules manage some rules for session with accounts
type SessionRules struct{}

// NewSessionRules constructor
func NewSessionRules() *SessionRules {
	return &SessionRules{}
}

// CreateAccount rule
func (module *SessionRules) CreateAccount(in *inputs.CreateAccountInput) *outputs.CreateAccountOutput {
	out := outputs.NewCreateAccountOutput()

	out.Account = in.BuildAccount()

	out.Success = in.Valid() &&
		!AccountRepository.UsernameHasTaken(in.Username) &&
		AccountRepository.Create(in.BuildAccount())

	return out
}
