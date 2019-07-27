package domain

// SessionRules manage some rules for session with accounts
type SessionRules struct{}

// NewSessionRules constructor
func NewSessionRules() *SessionRules {
	return &SessionRules{}
}

// CreateAccount rule
func (module *SessionRules) CreateAccount(in *CreateAccountInput) *CreateAccountOutput {
	out := NewCreateAccountOutput()

	out.Account = in.BuildAccount()

	out.Success = in.Valid() &&
		!AccountRepository.UsernameHasTaken(in.Username) &&
		AccountRepository.Create(in.BuildAccount())

	return out
}
