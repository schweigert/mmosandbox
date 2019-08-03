package cache

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CacheSuite struct {
	suite.Suite
}

func (suite *CacheSuite) TestConnect() {
	suite.NotPanics(func() {
		conn := Connect()
		defer conn.Close()
	})
}

func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}
