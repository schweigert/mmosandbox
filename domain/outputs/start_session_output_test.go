package outputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StartSessionOutputSuite struct {
	suite.Suite
}

func (suite *StartSessionOutputSuite) TestNewStartSessionOutput() {
	output := NewStartSessionOutput()
	suite.NotNil(output)
	suite.Empty(output.Token)
	suite.False(output.Success)
}

func TestStartSessionOutputSuite(t *testing.T) {
	suite.Run(t, new(StartSessionOutputSuite))
}
