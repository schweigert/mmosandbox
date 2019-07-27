package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func (suite *ConfigSuite) TestWeb() {
	suite.NotNil(Web())
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
