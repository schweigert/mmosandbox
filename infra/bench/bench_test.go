package bench

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BenchSuite struct {
	suite.Suite
}

func (suite *BenchSuite) TestBip() {
	suite.NotPanics(func() {
		Bench("testing", func() {})
	})
}

func TestBenchSuite(t *testing.T) {
	suite.Run(t, new(BenchSuite))
}
