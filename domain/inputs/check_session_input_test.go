package inputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CheckSessionInputSuite struct {
	suite.Suite
}

func (suite *CheckSessionInputSuite) TestNewCheckSessionInput() {
	input := NewCheckSessionInput()
	suite.NotNil(input)
	suite.Empty(input.Token)
	suite.Empty(input.Username)
}

func TestCheckSessionInputSuite(t *testing.T) {
	suite.Run(t, new(CheckSessionInputSuite))
}
