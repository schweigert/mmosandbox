package crudtask

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CrudTaskSuite struct {
	suite.Suite
}

func (suite *CrudTaskSuite) TestNewCrudTask() {
	suite.NotNil(NewCrudTask())
}

func TestCrudTaskSuite(t *testing.T) {
	suite.Run(t, new(CrudTaskSuite))
}
