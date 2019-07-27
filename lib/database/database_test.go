package database

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
}

func (suite *DatabaseSuite) TestConnect() {
	_, err := Connect("")
	suite.Error(err)
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}
