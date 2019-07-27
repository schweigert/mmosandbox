package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DbEnvsSuite struct {
	suite.Suite
}

func (suite *DbEnvsSuite) TestHost() {
	suite.NotNil(Db().Host())
}

func (suite *DbEnvsSuite) TestPort() {
	suite.NotNil(Db().Port())
}

func (suite *DbEnvsSuite) TestUser() {
	suite.NotNil(Db().User())
}

func (suite *DbEnvsSuite) TestName() {
	suite.NotNil(Db().Name())
}

func (suite *DbEnvsSuite) TestSSLMode() {
	suite.NotNil(Db().SSLMode())
}

func (suite *DbEnvsSuite) TestPassword() {
	suite.NotNil(Db().Password())
}

func (suite *DbEnvsSuite) TestAddr() {
	suite.NotNil(Db().Addr())
	suite.Equal("host=localhost port=5432 user=postgres dbname=test sslmode=disable password=postgres", Db().Addr())
}

func TestDbEnvsSuite(t *testing.T) {
	suite.Run(t, new(DbEnvsSuite))
}
