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

type StartSessionFlowMock struct {
	StartSessionBoolResult bool
}

func (mock *StartSessionFlowMock) StartSession(in *inputs.AuthAccountInput) (*outputs.StartSessionOutput, bool) {
	return &outputs.StartSessionOutput{Success: true, Token: "blefe"}, mock.StartSessionBoolResult
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

	UsedStartSessionFlow = &StartSessionFlowMock{
		StartSessionBoolResult: true,
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
