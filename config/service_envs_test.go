package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceEnvsSuite struct {
	suite.Suite
}

func (suite *ServiceEnvsSuite) TestService() {
	suite.NotEmpty(Service().Name())
}

func TestServiceEnvsSuite(t *testing.T) {
	suite.Run(t, new(ServiceEnvsSuite))
}
