package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WAuthSuite struct {
	suite.Suite
}

func (suite *WAuthSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestWAuthSuite(t *testing.T) {
	suite.Run(t, new(WAuthSuite))
}
