package client

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
}

type CreateAccountFlowMock struct {
	CreateAccountOperationMock bool
}

func (mock *CreateAccountFlowMock) CreateAccountOperation(*inputs.CreateAccountInput) bool {
	return mock.CreateAccountOperationMock
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
		CreateAccountOperationMock: true,
	}

	suite.NotPanics(BotFlow)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
