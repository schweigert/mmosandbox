package client

import (
	"fmt"

	"github.com/dmgk/faker"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
)

// Client flow vars
var (
	CreateAccountInput    *inputs.CreateAccountInput
	UsedCreateAccountFlow CreateAccountFlow
	Account               *entities.Account

	CreateCharacterInput    *inputs.CreateCharacterInput
	UsedCreateCharacterFlow CreateCharacterFlow
	Character               *entities.Character
)

// CreateAccountFlow used in client
type CreateAccountFlow interface {
	CreateAccountOperation(in *inputs.CreateAccountInput) (*entities.Account, bool)
}

// CreateCharacterFlow used in client
type CreateCharacterFlow interface {
	CreateCharacterOperation(in *inputs.CreateCharacterInput) (*entities.Character, bool)
}

// FakeName generator
func FakeName() string {
	return fmt.Sprintf("%s_%d", faker.Internet().UserName(), faker.RandomInt(1950, 2019))
}

// FakeCreateAccountInput instance a new CreateAccountInput
func FakeCreateAccountInput() {
	CreateAccountInput = inputs.NewCreateAccountInput()
	CreateAccountInput.Username = FakeName()
	CreateAccountInput.Password = faker.Internet().Password(8, 64)
	CreateAccountInput.Email = faker.Internet().SafeEmail()
}

// FakeCreateCharacterInput instance a new CreateCharacterInput
func FakeCreateCharacterInput() {
	CreateCharacterInput = inputs.NewCreateCharacterInput()
	CreateCharacterInput.Auth.Username = Account.Username
	CreateCharacterInput.Auth.SecurePassword = Account.SecurePassword
	CreateCharacterInput.Name = FakeName()
}

// BotFlow loop
func BotFlow() {
	for {
		FakeCreateAccountInput()
		account, ok := UsedCreateAccountFlow.CreateAccountOperation(CreateAccountInput)
		if ok {
			Account = account
			break
		}
	}

	for {
		FakeCreateCharacterInput()
		character, ok := UsedCreateCharacterFlow.CreateCharacterOperation(CreateCharacterInput)
		if ok {
			Character = character
			break
		}
	}

	Account.Characters = []entities.Character{*Character}
	Character.Account = *Account

	fmt.Println("BootFlow ended with:", Account)
}
