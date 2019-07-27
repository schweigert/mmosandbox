package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DbSuite struct {
	suite.Suite
}

func (suite *DbSuite) TestConnect() {
	suite.NotPanics(func() {
		conn := Connect()
		defer conn.Close()
	})
}

func (suite *DbSuite) TestMigrate() {
	conn := Connect()
	defer conn.Close()

	Clear()

	suite.NotPanics(Migrate)

	suite.True(conn.HasTable("accounts"))
	suite.True(conn.HasTable("characters"))
}

func TestDbSuite(t *testing.T) {
	suite.Run(t, new(DbSuite))
}
