package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SWebSuite struct {
	suite.Suite
}

func (suite *SWebSuite) TestMain() {
	os.Setenv("WEB_INTERFACE", "localhost")
	os.Setenv("WEB_PORT", "-1")

	suite.Panics(main)
}

func TestSWebSuite(t *testing.T) {
	suite.Run(t, new(SWebSuite))
}
