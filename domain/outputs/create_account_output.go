package outputs

import "github.com/schweigert/mmosandbox/domain/entities"

// CreateAccountOutput put some information from domain
type CreateAccountOutput struct {
	Account entities.Account `json:"account"`
	Success bool             `json:"success"`
}
