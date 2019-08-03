package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MeEnvsSuite struct {
	suite.Suite
}

func (suite *MeEnvsSuite) TestHost() {
	suite.NotEmpty(Me().Host())
}

func (suite *MeEnvsSuite) TestPort() {
	suite.NotEmpty(Me().Port())
}

func (suite *MeEnvsSuite) TestPassword() {
	suite.NotEmpty(Me().Password())
}

func (suite *MeEnvsSuite) TestAddr() {
	suite.NotEmpty(Me().Addr())
}

func TestMeEnvsSuite(t *testing.T) {
	suite.Run(t, new(MeEnvsSuite))
}
