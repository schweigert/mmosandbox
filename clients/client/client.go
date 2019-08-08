package client

import (
	"fmt"

	"github.com/schweigert/mmosandbox/domain/outputs"

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

	AuthAccountInput inputs.AuthAccountInput
	UsedSessionFlow  SessionFlow
	Session          *outputs.StartSessionOutput

	SpawnCharacterInput inputs.SpawnCharacterInput
)

// CreateAccountFlow used in client
type CreateAccountFlow interface {
	CreateAccountOperation(in *inputs.CreateAccountInput) (*entities.Account, bool)
}

// CreateCharacterFlow used in client
type CreateCharacterFlow interface {
	CreateCharacterOperation(in *inputs.CreateCharacterInput) (*entities.Character, bool)
}

// SessionFlow used in client
type SessionFlow interface {
	StartSession(in inputs.AuthAccountInput) (*outputs.StartSessionOutput, bool)

	SpawnCharacter(in inputs.SpawnCharacterInput) (*outputs.CheckSessionOutput, bool)
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

func createAccountFlowExec() {
	for {
		FakeCreateAccountInput()
		account, ok := UsedCreateAccountFlow.CreateAccountOperation(CreateAccountInput)
		if ok {
			Account = account
			break
		}
	}
}

func createCharacterFlowExec() {
	for {
		FakeCreateCharacterInput()
		character, ok := UsedCreateCharacterFlow.CreateCharacterOperation(CreateCharacterInput)
		if ok {
			Character = character
			break
		}
	}
}

func sessionFlowExec() {
	for {
		AuthAccountInput = CreateCharacterInput.Auth

		session, ok := UsedSessionFlow.StartSession(AuthAccountInput)
		if ok && session.Success {
			Session = session
			SpawnCharacterInput.CharacterID = int(Character.ID)
			SpawnCharacterInput.CheckSessionInput = inputs.NewCheckSessionInput()
			SpawnCharacterInput.CheckSessionInput.Username = Account.Username
			SpawnCharacterInput.CheckSessionInput.Token = Session.Token

			break
		}
	}

	for {
		checkSessionOutput, ok := UsedSessionFlow.SpawnCharacter(SpawnCharacterInput)
		if ok && checkSessionOutput.Success {
			break
		}
	}
}

// BotFlow loop
func BotFlow() {
	createAccountFlowExec()
	createCharacterFlowExec()
	sessionFlowExec()
}
