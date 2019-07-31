package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MetricEnvsSuite struct {
	suite.Suite
}

func (suite *MetricEnvsSuite) TestHost() {
	suite.NotEmpty(Metric().Host())
}

func (suite *MetricEnvsSuite) TestPort() {
	suite.NotEmpty(Metric().Port())
}

func TestMetricEnvsSuite(t *testing.T) {
	suite.Run(t, new(MetricEnvsSuite))
}
