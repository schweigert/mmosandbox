package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SpawnCharacterInputSuite struct {
	suite.Suite
}

func (suite *SpawnCharacterInputSuite) TestNewSpawnCharacterInput() {
	input := NewSpawnCharacterInput()
	suite.NotNil(input)
	suite.Zero(input.CharacterID)
	suite.Nil(input.CheckSessionInput)
}

func TestSpawnCharacterInputSuite(t *testing.T) {
	suite.Run(t, new(SpawnCharacterInputSuite))
}
