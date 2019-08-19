package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SGameSuite struct {
	suite.Suite
}

func (suite *SGameSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestSGameSuite(t *testing.T) {
	suite.Run(t, new(SGameSuite))
}
