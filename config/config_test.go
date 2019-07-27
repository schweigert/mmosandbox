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

func (suite *ConfigSuite) TestDb() {
	suite.NotNil(Db())
}

func (suite *ConfigSuite) TestClient() {
	suite.NotNil(Client())
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}
