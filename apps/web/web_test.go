package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebSuite struct {
	suite.Suite
}

func (suite *WebSuite) TestMain() {
	os.Setenv("WEB_INTERFACE", "localhost")
	os.Setenv("WEB_PORT", "-1")

	suite.Panics(main)
}

func TestWebSuite(t *testing.T) {
	suite.Run(t, new(WebSuite))
}
