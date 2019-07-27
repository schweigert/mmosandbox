package domain

import (
	"github.com/schweigert/mmosandbox/domain/entities"
)

// CreateAccountOutput form
type CreateAccountOutput struct {
	Account *entities.Account `json:"account"`
	Success bool              `json:"success"`
}

// NewCreateAccountOutput constructor
func NewCreateAccountOutput() *CreateAccountOutput {
	return &CreateAccountOutput{}
}
