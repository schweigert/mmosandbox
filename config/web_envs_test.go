package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebEnvsSuite struct {
	suite.Suite
}

func (suite *WebEnvsSuite) TestInterface() {
	suite.NotEmpty(Web().Interface())
}

func (suite *WebEnvsSuite) TestPort() {
	suite.NotEmpty(Web().Port())
}

func (suite *WebEnvsSuite) TestAddr() {
	suite.NotEmpty(Web().Addr())
}

func TestWebEnvsSuite(t *testing.T) {
	suite.Run(t, new(WebEnvsSuite))
}
