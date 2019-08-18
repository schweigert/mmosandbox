package middlewares

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BipSuite struct {
	suite.Suite
}

func (suite *BipSuite) TestBip() {
	middleware := Bip()
	suite.NotNil(middleware)
}

func (suite *BipSuite) TestPrivateBipStat() {
	suite.Equal("routes.root", bipStat("/"))
	suite.Equal("routes.auth", bipStat("/auth"))
	suite.Equal("routes.auth_google", bipStat("/auth/google"))
}

func TestBipSuite(t *testing.T) {
	suite.Run(t, new(BipSuite))
}
