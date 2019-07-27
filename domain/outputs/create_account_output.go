package outputs

import (
	"net/http"

	"github.com/krakenlab/ternary"
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

// HTTPCode to output over web
func (out *CreateAccountOutput) HTTPCode() int {
	return ternary.Int(out.Success, http.StatusCreated, http.StatusBadRequest)
}
