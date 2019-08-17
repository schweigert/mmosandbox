package outputs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ListenMessagesOutputSuite struct {
	suite.Suite
}

func (suite *ListenMessagesOutputSuite) TestNewListenMessagesOutput() {
	out := NewListenMessagesOutput()
	suite.NotNil(out)
	suite.Empty(out.Messages)
	suite.False(out.Success)
}

func TestListenMessagesOutputSuite(t *testing.T) {
	suite.Run(t, new(ListenMessagesOutputSuite))
}
