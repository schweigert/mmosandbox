package client

import (
	"fmt"
	"log"
	"time"

	"github.com/logrusorgru/aurora"

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

	AuthAccountInput    inputs.AuthAccountInput
	SpawnCharacterInput inputs.SpawnCharacterInput
	UsedSessionFlow     SessionFlow
	Session             *outputs.StartSessionOutput

	UsedGameFlow       GameFlow
	MoveCharacterInput inputs.MoveCharacterInput
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

// GameFlow used play the game
type GameFlow interface {
	GameLoop() bool
	MoveCharacter(in inputs.MoveCharacterInput) (*outputs.CheckSessionOutput, bool)
	SendChat(in inputs.ChatInput) (*outputs.CheckSessionOutput, bool)
	ListenChat(in inputs.ChatInput) (*outputs.ListenMessagesOutput, bool)
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

func gameFlowUpdate() {
	for {
		in := *inputs.NewMoveCharacterInput()
		in.CharacterID = int(Character.ID)
		in.CheckSessionInput = inputs.NewCheckSessionInput()
		in.CheckSessionInput.Username = Account.Username
		in.CheckSessionInput.Token = Session.Token

		in.DeltaX = faker.RandomInt(-2, 2)
		in.DeltaY = faker.RandomInt(-2, 2)

		checkSessionOutput, ok := UsedGameFlow.MoveCharacter(in)

		if ok && checkSessionOutput.Success {
			break
		}
	}

	for {
		in := *inputs.NewChatInput()
		in.CharacterID = int(Character.ID)
		in.CheckSessionInput = inputs.NewCheckSessionInput()
		in.CheckSessionInput.Username = Account.Username
		in.CheckSessionInput.Token = Session.Token

		in.Body = faker.Hacker().SaySomethingSmart()

		checkSessionOutput, ok := UsedGameFlow.SendChat(in)
		if ok && checkSessionOutput.Success {
			break
		}
	}

	for {
		in := *inputs.NewChatInput()
		in.CharacterID = int(Character.ID)
		in.CheckSessionInput = inputs.NewCheckSessionInput()
		in.CheckSessionInput.Username = Account.Username
		in.CheckSessionInput.Token = Session.Token

		receiveMessagesOutput, ok := UsedGameFlow.ListenChat(in)
		if ok && receiveMessagesOutput.Success {
			for _, message := range receiveMessagesOutput.Messages {
				log.Printf("%s |> %s\n", aurora.Blue(message.CharacterName).Bold(), aurora.Yellow(message.Body))
			}

			break
		}
	}
}

func gameFlowExec() {
	gameFlowUpdate()
	for UsedGameFlow.GameLoop() {
		time.Sleep(time.Second)
		gameFlowUpdate()
	}
}

// BotFlow loop
func BotFlow() {
	createAccountFlowExec()
	createCharacterFlowExec()
	sessionFlowExec()
	gameFlowExec()
}
