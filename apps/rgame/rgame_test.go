package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RGameSuite struct {
	suite.Suite
}

func (suite *RGameSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestRGameSuite(t *testing.T) {
	suite.Run(t, new(RGameSuite))
}
