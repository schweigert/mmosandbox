package entities

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EntitiesSuite struct {
	suite.Suite
}

func (suite *EntitiesSuite) TestNewCharacter() {
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

func TestEntitiesSuite(t *testing.T) {
	suite.Run(t, new(EntitiesSuite))
}
