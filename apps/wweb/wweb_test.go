package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WWebSuite struct {
	suite.Suite
}

func (suite *WWebSuite) TestMain() {
	os.Setenv("WEB_INTERFACE", "localhost")
	os.Setenv("WEB_PORT", "-1")

	suite.Panics(main)
}

func TestWWebSuite(t *testing.T) {
	suite.Run(t, new(WWebSuite))
}
