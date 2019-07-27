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

func TestWebEnvsConfig(t *testing.T) {
	suite.Run(t, new(WebEnvsConfig))
}
