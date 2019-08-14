package entities

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MessageSuite struct {
	suite.Suite
}

func (suite *MessageSuite) TestNewMessageFromCharacter() {
	character := NewCharacter()
	character.ID = 25
	character.MapYPosition = 25
	character.MapXPosition = 33
	character.Name = "name of this character"

	message := NewMessageFromCharacter(character, "I send this message!")

	suite.Equal(message.CharacterID, character.ID)
	suite.Equal(message.CharacterName, character.Name)
	suite.Equal(message.MapXPosition, character.MapXPosition)
	suite.Equal(message.MapYPosition, character.MapYPosition)
	suite.Equal(message.Body, "I send this message!")
}

func TestMessageSuite(t *testing.T) {
	suite.Run(t, new(MessageSuite))
}
