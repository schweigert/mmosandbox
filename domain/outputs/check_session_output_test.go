package outputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CheckSessionOutputSuite struct {
	suite.Suite
}

func (suite *CheckSessionOutputSuite) TestNewCheckSessionOutput() {
	output := NewCheckSessionOutput()
	suite.NotNil(output)
	suite.False(output.Success)
}

func TestCheckSessionOutputSuite(t *testing.T) {
	suite.Run(t, new(CheckSessionOutputSuite))
}
