package client

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
}

type CreateAccountFlowMock struct {
	CreateAccountOperationBoolMock bool
}

func (mock *CreateAccountFlowMock) CreateAccountOperation(input *inputs.CreateAccountInput) (*entities.Account, bool) {
	return input.BuildAccount(), mock.CreateAccountOperationBoolMock
}

type CreateCharacterFlowMock struct {
	CreateCharacterOperationBoolMock bool
}

func (mock *CreateCharacterFlowMock) CreateCharacterOperation(input *inputs.CreateCharacterInput) (*entities.Character, bool) {
	return input.BuildCharacter(), mock.CreateCharacterOperationBoolMock
}

type SessionFlowMock struct {
	StartSessionBoolResult bool
	SpawnCharacterResult   bool
}

func (mock *SessionFlowMock) StartSession(in inputs.AuthAccountInput) (*outputs.StartSessionOutput, bool) {
	return &outputs.StartSessionOutput{Success: true, Token: "blefe"}, mock.StartSessionBoolResult
}

func (mock *SessionFlowMock) SpawnCharacter(in inputs.SpawnCharacterInput) (*outputs.CheckSessionOutput, bool) {
	return &outputs.CheckSessionOutput{Success: true}, mock.SpawnCharacterResult
}

type GameFlowMock struct {
	GameLoopResult      bool
	MoveCharacterResult bool
}

func (mock *GameFlowMock) GameLoop() bool {
	return mock.GameLoopResult
}

func (mock *GameFlowMock) MoveCharacter(in inputs.MoveCharacterInput) (*outputs.CheckSessionOutput, bool) {
	return &outputs.CheckSessionOutput{Success: true}, mock.MoveCharacterResult
}

func (suite *ClientSuite) TestNewCreateAccountInput() {
	FakeCreateAccountInput()

	suite.NotNil(CreateAccountInput)
	suite.NotEmpty(CreateAccountInput.Username)
	suite.NotEmpty(CreateAccountInput.Password)
	suite.NotEmpty(CreateAccountInput.Email)
}

func (suite *ClientSuite) TestBotFlow() {
	UsedCreateAccountFlow = &CreateAccountFlowMock{
		CreateAccountOperationBoolMock: true,
	}

	UsedCreateCharacterFlow = &CreateCharacterFlowMock{
		CreateCharacterOperationBoolMock: true,
	}

	UsedSessionFlow = &SessionFlowMock{
		StartSessionBoolResult: true,
		SpawnCharacterResult:   true,
	}

	UsedGameFlow = &GameFlowMock{
		GameLoopResult:      false,
		MoveCharacterResult: true,
	}

	Account = nil
	CreateAccountInput = nil

	suite.NotPanics(BotFlow)

	suite.NotNil(Account)
	suite.NotNil(CreateAccountInput)

	suite.NotNil(Character)
	suite.NotNil(CreateCharacterInput)

	suite.NotNil(Session)
	suite.True(Session.Success)
	suite.NotNil(Session.Token)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
