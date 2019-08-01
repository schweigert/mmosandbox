package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RPCEnvsSuite struct {
	suite.Suite
}

func (suite *RPCEnvsSuite) TestInterface() {
	suite.NotEmpty(RPC().Interface())
}

func (suite *RPCEnvsSuite) TestPort() {
	suite.NotEmpty(RPC().Port())
}

func (suite *RPCEnvsSuite) TestAddr() {
	suite.NotEmpty(RPC().Addr())
}

func TestRPCEnvsSuite(t *testing.T) {
	suite.Run(t, new(RPCEnvsSuite))
}
