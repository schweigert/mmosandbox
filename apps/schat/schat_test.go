package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SChatSuite struct {
	suite.Suite
}

func (suite *SChatSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestSChatSuite(t *testing.T) {
	suite.Run(t, new(SChatSuite))
}
