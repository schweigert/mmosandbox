package outputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthAccountOutputSuite struct {
	suite.Suite
}

func (suite *AuthAccountOutputSuite) TestNewAuthAccountOutput() {
	output := NewAuthAccountOutput()
	suite.NotNil(output)
	suite.Empty(output.JWT)
	suite.False(output.Success)
}

func TestAuthAccountOutputSuite(t *testing.T) {
	suite.Run(t, new(AuthAccountOutputSuite))
}
