package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WGameSuite struct {
	suite.Suite
}

func (suite *WGameSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestWGameSuite(t *testing.T) {
	suite.Run(t, new(WGameSuite))
}
