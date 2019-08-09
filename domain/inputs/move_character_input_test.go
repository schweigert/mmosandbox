package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MoveCharacterInputSuite struct {
	suite.Suite
}

func (suite *MoveCharacterInputSuite) TestNewMoveCharacterInput() {
	input := NewMoveCharacterInput()
	suite.NotNil(input)
	suite.Zero(input.CharacterID)
	suite.Nil(input.CheckSessionInput)
	suite.Zero(input.DeltaX)
	suite.Zero(input.DeltaY)
}

func TestMoveCharacterInputSuite(t *testing.T) {
	suite.Run(t, new(MoveCharacterInputSuite))
}
