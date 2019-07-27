package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientEnvsSuite struct {
	suite.Suite
}

func (suite *ClientEnvsSuite) TestWebHost() {
	suite.NotEmpty(Client().WebHost())
}

func (suite *ClientEnvsSuite) TestPort() {
	suite.NotEmpty(Client().WebPort())
}

func (suite *ClientEnvsSuite) TestWebProtocol() {
	suite.NotEmpty(Client().WebProtocol())
}

func (suite *ClientEnvsSuite) TestAddr() {
	suite.NotEmpty(Client().Addr())
	suite.Equal("http://localhost:3000", Client().Addr())
}

func TestClientEnvsSuite(t *testing.T) {
	suite.Run(t, new(ClientEnvsSuite))
}
