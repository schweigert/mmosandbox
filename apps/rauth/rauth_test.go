package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RAuthSuite struct {
	suite.Suite
}

func (suite *RAuthSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestRAuthSuite(t *testing.T) {
	suite.Run(t, new(RAuthSuite))
}
