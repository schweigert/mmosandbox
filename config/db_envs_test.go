package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DbEnvsSuite struct {
	suite.Suite
}

func (suite *DbEnvsSuite) TestHost() {
	suite.NotEmpty(Db().Host())
}

func (suite *DbEnvsSuite) TestPort() {
	suite.NotEmpty(Db().Port())
}

func (suite *DbEnvsSuite) TestUser() {
	suite.NotEmpty(Db().User())
}

func (suite *DbEnvsSuite) TestName() {
	suite.NotEmpty(Db().Name())
}

func (suite *DbEnvsSuite) TestSSLMode() {
	suite.NotEmpty(Db().SSLMode())
}

func (suite *DbEnvsSuite) TestPassword() {
	suite.NotEmpty(Db().Password())
}

func (suite *DbEnvsSuite) TestAddr() {
	suite.NotEmpty(Db().Addr())
	suite.Equal("host=localhost port=5432 user=postgres dbname=test sslmode=disable password=postgres", Db().Addr())
}

func TestDbEnvsSuite(t *testing.T) {
	suite.Run(t, new(DbEnvsSuite))
}
