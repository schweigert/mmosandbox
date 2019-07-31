package metrics

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MetronomeSuite struct {
	suite.Suite
}

func (suite *MetronomeSuite) TestConnect() {
	metronome, err := Connect("localhost", "2003")
	suite.NoError(err)
	suite.NotNil(metronome)

	metronome, err = Connect("localhost", "-1")
	suite.Error(err)
	suite.NotNil(metronome)

	metronome, err = Connect("localhost", "kk")
	suite.Error(err)
	suite.Nil(metronome)
}

func (suite *MetronomeSuite) TestBip() {
	metronome, err := Connect("localhost", "2003")
	suite.NoError(err)
	suite.NotNil(metronome)

	metronome.Bip("testing now", "-1")
}

func TestMetronomeSuite(t *testing.T) {
	suite.Run(t, new(MetronomeSuite))
}
