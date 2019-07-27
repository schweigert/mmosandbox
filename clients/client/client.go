package client

import (
	"github.com/dmgk/faker"
	"github.com/schweigert/mmosandbox/domain/inputs"
)

// Client flow vars
var (
	CreateAccountInput    *inputs.CreateAccountInput
	UsedCreateAccountFlow CreateAccountFlow
)

// CreateAccountFlow used in client
type CreateAccountFlow interface {
	CreateAccountOperation(in *inputs.CreateAccountInput) bool
}

// FakeCreateAccountInput instance a new CreateAccountInput
func FakeCreateAccountInput() {
	CreateAccountInput = inputs.NewCreateAccountInput()
	CreateAccountInput.Username = faker.Internet().UserName()
	CreateAccountInput.Password = faker.Internet().Password(8, 64)
	CreateAccountInput.Email = faker.Internet().SafeEmail()
}

// BotFlow loop
func BotFlow() {
	for {
		FakeCreateAccountInput()
		if UsedCreateAccountFlow.CreateAccountOperation(CreateAccountInput) {
			break
		}
	}
}
