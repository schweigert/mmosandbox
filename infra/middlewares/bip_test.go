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
	suite.Equal("http.noname.root", bipStat("/"))
	suite.Equal("http.noname.auth", bipStat("/auth"))
	suite.Equal("http.noname.auth_google", bipStat("/auth/google"))
}

func TestBipSuite(t *testing.T) {
	suite.Run(t, new(BipSuite))
}
