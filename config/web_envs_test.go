package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebEnvsConfig struct {
	suite.Suite
}

func (suite *WebEnvsConfig) TestInterface() {
	suite.NotEmpty(Web().Interface())
}

func (suite *WebEnvsConfig) TestPort() {
	suite.NotEmpty(Web().Port())
}

func (suite *WebEnvsConfig) TestAddr() {
	suite.NotEmpty(Web().Addr())
}

func TestWebEnvsConfig(t *testing.T) {
	suite.Run(t, new(WebEnvsConfig))
}
