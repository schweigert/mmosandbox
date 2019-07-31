package inputs

import "github.com/schweigert/mmosandbox/domain/entities"

// CreateCharacterInput form
type CreateCharacterInput struct {
	Auth AuthAccountInput

	Name string `json:"name"`
}

// NewCreateCharacterInput constructor
func NewCreateCharacterInput() *CreateCharacterInput {
	return &CreateCharacterInput{}
}

// BuildAccount from this form
func (in *CreateCharacterInput) BuildAccount() *entities.Account {
	return in.Auth.BuildAccount()
}

// BuildCharacter from this form
func (in *CreateCharacterInput) BuildCharacter() *entities.Character {
	character := entities.NewCharacter()
	character.Name = in.Name
	return character
}
