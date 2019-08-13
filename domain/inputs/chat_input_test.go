package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChatInputSuite struct {
	suite.Suite
}

func (suite *ChatInputSuite) TestNewChatInput() {
	input := NewChatInput()
	suite.NotNil(input)
	suite.Zero(input.CharacterID)
	suite.Nil(input.CheckSessionInput)
	suite.Empty(input.Body)
}

func TestChatInputSuite(t *testing.T) {
	suite.Run(t, new(ChatInputSuite))
}
