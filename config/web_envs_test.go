package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebEnvsConfig struct {
	suite.Suite
}

func (suite *WebEnvsConfig) TestInterface() {
	suite.NotNil(Web().Interface())
}

func (suite *WebEnvsConfig) TestPort() {
	suite.NotNil(Web().Port())
}

func (suite *WebEnvsConfig) TestAddr() {
	suite.NotNil(Web().Addr())
}

func TestWebEnvsConfig(t *testing.T) {
	suite.Run(t, new(WebEnvsConfig))
}
