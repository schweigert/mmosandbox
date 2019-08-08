package domain

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/entities"

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

func TestWorldRulesSuite(t *testing.T) {
	suite.Run(t, new(WorldRulesSuite))
}
