package entities

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CharacterSuite struct {
	suite.Suite
}

func (suite *CharacterSuite) TestNewCharacter() {
	character := NewCharacter()
	suite.Equal(uint(1), character.Level)
	suite.Equal(uint(0), character.Exp)
	suite.Equal(uint(5), character.StrengthStatus)
	suite.Equal(uint(5), character.IntelligenceStatus)
	suite.Equal(uint(5), character.AgilityStatus)
	suite.Equal(uint(5), character.VitalityStatus)
	suite.Equal(uint(0), character.MapIndex)
	suite.Equal(0, character.MapXPosition)
	suite.Equal(0, character.MapYPosition)
}

func (suite *CharacterSuite) TestDistanceTo() {
	character := NewCharacter()

	suite.Equal(0, character.MapXPosition)
	suite.Equal(0, character.MapYPosition)

	suite.Equal(float64(0), character.DistanceTo(0, 0))
	suite.Equal(float64(1), character.DistanceTo(1, 0))
	suite.Equal(float64(1), character.DistanceTo(-1, 0))
	suite.Equal(float64(1), character.DistanceTo(0, 1))
	suite.Equal(float64(1), character.DistanceTo(0, -1))

	suite.Equal(float64(1.4142135623730951), character.DistanceTo(1, 1))
}

func TestCharacterSuite(t *testing.T) {
	suite.Run(t, new(CharacterSuite))
}
