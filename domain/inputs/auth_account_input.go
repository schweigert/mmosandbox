package inputs

import (
	"github.com/schweigert/mmosandbox/domain/entities"
)

// AuthAccountInput form
type AuthAccountInput struct {
	Username       string `json:"username"`
	SecurePassword string `json:"secure_password"`
}

// NewAuthAccountInput constructor
func NewAuthAccountInput() *AuthAccountInput {
	return &AuthAccountInput{}
}

// BuildAccount based in input request
func (in *AuthAccountInput) BuildAccount() (account *entities.Account) {
	account = entities.NewAccount()

	account.Username = in.Username
	account.SecurePassword = in.SecurePassword

	return account
}
