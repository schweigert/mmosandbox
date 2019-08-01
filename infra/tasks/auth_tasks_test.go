package tasks

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthTasksSuite struct {
	suite.Suite
}

func (suite *AuthTasksSuite) TestNewAuthTasks() {
	authTasks := NewAuthTasks()
	suite.NotNil(authTasks)
}

func TestAuthTasksSuite(t *testing.T) {
	suite.Run(t, new(AuthTasksSuite))
}
