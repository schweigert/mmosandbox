package outputs

import (
	"net/http"

	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmosandbox/domain/entities"
)

// CreateCharacterOutput form
type CreateCharacterOutput struct {
	Account   *entities.Account   `json:"account"`
	Character *entities.Character `json:"character"`
	Auth      bool                `json:"auth"`
	Success   bool                `json:"success"`
}

// NewCreateCharacterOutput constructor
func NewCreateCharacterOutput() *CreateCharacterOutput {
	return &CreateCharacterOutput{}
}

// HTTPCode to output over web
func (out *CreateCharacterOutput) HTTPCode() int {
	return ternary.Int(out.Success && out.Auth, http.StatusCreated, http.StatusBadRequest)
}
