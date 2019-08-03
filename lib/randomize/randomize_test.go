package randomize

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RandomizeSuite struct {
	suite.Suite
}

func (suite *RandomizeSuite) TestNewRandomize() {
	suite.Equal(256, len(RandStringRunes(256)))
}

func TestRandomizeSuite(t *testing.T) {
	suite.Run(t, new(RandomizeSuite))
}
