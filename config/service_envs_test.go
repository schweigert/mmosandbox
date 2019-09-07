package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceEnvsSuite struct {
	suite.Suite
}

func (suite *ServiceEnvsSuite) TestName() {
	suite.NotEmpty(Service().Name())
}

func (suite *ServiceEnvsSuite) TestAuth() {
	suite.NotEmpty(Service().Auth())
}

func (suite *ServiceEnvsSuite) TestGame() {
	suite.NotEmpty(Service().Game())
}

func (suite *ServiceEnvsSuite) TestChat() {
	suite.NotEmpty(Service().Chat())
}

func TestServiceEnvsSuite(t *testing.T) {
	suite.Run(t, new(ServiceEnvsSuite))
}
