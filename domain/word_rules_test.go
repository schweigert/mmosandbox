package domain

import (
	"testing"

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

func TestWorldRulesSuite(t *testing.T) {
	suite.Run(t, new(WorldRulesSuite))
}
