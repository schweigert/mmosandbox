package tasks

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SessionTaskSuite struct {
	suite.Suite
}

func (suite *SessionTaskSuite) TestNewSessionTask() {
	SessionTask := NewSessionTask()
	suite.NotNil(SessionTask)
}

func TestSessionTaskSuite(t *testing.T) {
	suite.Run(t, new(SessionTaskSuite))
}
