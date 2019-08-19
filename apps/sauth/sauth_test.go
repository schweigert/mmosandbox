package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SAuthSuite struct {
	suite.Suite
}

func (suite *SAuthSuite) TestMain() {
	os.Setenv("RPC_INTERFACE", "localhost")
	os.Setenv("RPC_PORT", "-1")

	suite.Panics(main)
}

func TestSAuthSuite(t *testing.T) {
	suite.Run(t, new(SAuthSuite))
}
