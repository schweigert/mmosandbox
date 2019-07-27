package env

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvSuite struct {
	suite.Suite
}

func (suite *EnvSuite) TestEnvFunc() {
	suite.Equal("testing", Env("NOT_GOPATH", "testing"))
	suite.NotEqual("testing", Env("GOPATH", "testing"))
}

func TestEnvSuite(t *testing.T) {
	suite.Run(t, new(EnvSuite))
}
