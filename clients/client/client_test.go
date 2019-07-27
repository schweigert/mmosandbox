package client

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
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

	Account = nil
	CreateAccountInput = nil

	suite.NotPanics(BotFlow)

	suite.NotNil(Account)
	suite.NotNil(CreateAccountInput)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
