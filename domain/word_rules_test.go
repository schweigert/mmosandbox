package domain

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/outputs"

	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/stretchr/testify/suite"
)

type WorldRulesSuite struct {
	suite.Suite
}

func (suite *WorldRulesSuite) TestNewWorldRulesSuite() {
	rules := NewWorldRules()
	suite.NotNil(rules)
	suite.Zero(len(rules.Characters))
	suite.NotEmpty(rules.WorldName)
	suite.NotContains(rules.WorldName, " ")
}

func (suite *WorldRulesSuite) TestSpawnCharacter() {
	character := entities.NewCharacter()

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: character,
	}

	rules := NewWorldRules()

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(character.ID))
	})

	suite.NotZero(len(rules.Characters))
	suite.Equal(character, rules.Characters[0])
}

func (suite *WorldRulesSuite) TestMoveCharacter() {
	character := entities.NewCharacter()

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: character,
	}

	rules := NewWorldRules()

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(character.ID))
	})

	suite.NotZero(len(rules.Characters))
	suite.Equal(character, rules.Characters[0])

	in := inputs.NewMoveCharacterInput()
	in.CharacterID = int(character.ID)
	in.DeltaX = -1
	in.DeltaY = -1

	suite.Equal(0, character.MapXPosition)
	suite.Equal(0, character.MapYPosition)
	err := rules.MoveCharacter(in)
	suite.NoError(err)

	suite.Equal(-1, character.MapXPosition)
	suite.Equal(-1, character.MapYPosition)

	in.DeltaY = 0

	err = rules.MoveCharacter(in)
	suite.NoError(err)

	suite.Equal(-2, character.MapXPosition)
	suite.Equal(-1, character.MapYPosition)

	rules = NewWorldRules()
	err = rules.MoveCharacter(in)
	suite.Error(err)
}

func (suite *WorldRulesSuite) TestPrivateFindCharactersInFieldOfVision() {
	characterOne := entities.NewCharacter()
	characterOne.ID = 1
	characterTwo := entities.NewCharacter()
	characterTwo.ID = 2

	rules := NewWorldRules()

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: characterOne,
	}

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(characterOne.ID))
	})

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: characterTwo,
	}

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(characterTwo.ID))
	})

	characterOne.MapXPosition = 0
	characterOne.MapYPosition = 0

	characterTwo.MapXPosition = 25
	characterTwo.MapYPosition = 0

	queriedChars := rules.findCharactersInFieldOfVision(characterTwo)
	suite.Equal(1, len(queriedChars))
	suite.Equal(characterOne.ID, queriedChars[0].ID)

	characterOne.MapXPosition = 0
	characterOne.MapYPosition = -25

	queriedChars = rules.findCharactersInFieldOfVision(characterTwo)
	suite.Zero(len(queriedChars))
}

func (suite *WorldRulesSuite) TestCharacterSpoke() {
	characterOne := entities.NewCharacter()
	characterOne.ID = 1
	characterTwo := entities.NewCharacter()
	characterTwo.ID = 2

	rules := NewWorldRules()

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: characterOne,
	}

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(characterOne.ID))
	})

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: characterTwo,
	}

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(characterTwo.ID))
	})

	characterOne.MapXPosition = 0
	characterOne.MapYPosition = 0

	characterTwo.MapXPosition = 10
	characterTwo.MapYPosition = 10

	in := inputs.NewChatInput()
	in.CharacterID = 1
	in.Body = "Hello 2!"

	err := rules.CharacterSpoke(in)
	suite.NoError(err)

	messages := characterTwo.MessageBox.Read()

	suite.Equal(1, len(messages))
	suite.Equal(uint(1), messages[0].CharacterID)
	suite.Equal("Hello 2!", messages[0].Body)

	messages = characterOne.MessageBox.Read()

	suite.Zero(len(messages))
}

func (suite *WorldRulesSuite) TestCharacterListen() {
	characterOne := entities.NewCharacter()
	characterOne.ID = 1

	rules := NewWorldRules()

	CharacterRepository = &CharacterRepositoryMock{
		LoadCharacterResult: characterOne,
	}

	message := entities.NewMessage()
	message.Body = "Testing"

	suite.NotPanics(func() {
		rules.SpawnCharacter(int(characterOne.ID))
	})

	characterOne.MessageBox.Append(message)

	in := inputs.NewChatInput()
	in.CharacterID = int(characterOne.ID)

	out := outputs.NewListenMessagesOutput()

	err := rules.CharacterListen(in, out)
	suite.NoError(err)

	suite.Equal(1, len(out.Messages))
	suite.Equal(message.Body, out.Messages[0].Body)
}

func TestWorldRulesSuite(t *testing.T) {
	suite.Run(t, new(WorldRulesSuite))
}
