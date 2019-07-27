package domain

import (
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/validators"
	"github.com/schweigert/mmosandbox/lib/security"
)

// CreateAccountInput form
type CreateAccountInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// NewCreateAccountInput constructor
func NewCreateAccountInput() *CreateAccountInput {
	return &CreateAccountInput{}
}

// Valid input form
func (in *CreateAccountInput) Valid() bool {
	return validators.Email(in.Email) &&
		validators.StringSize(in.Username, 8, 128) &&
		validators.StringSize(in.Password, 8, 128) &&
		validators.StringSize(in.Email, 8, 128)
}

// BuildAccount based in input request
func (in *CreateAccountInput) BuildAccount() (account *entities.Account) {
	account = entities.NewAccount()

	account.Username = in.Username
	account.Email = in.Email
	account.SecurePassword = security.Sum(in.Password)

	return account
}
