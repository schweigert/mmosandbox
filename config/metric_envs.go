package config

import (
	"github.com/schweigert/mmosandbox/lib/env"
)

// MetricEnvs configurator
type MetricEnvs struct{}

// Host $METRIC_HOST localhost
func (db *MetricEnvs) Host() string {
	return env.Env("METRIC_HOST", "localhost")
}

// Port $METRIC_PORT 5432
func (db *MetricEnvs) Port() string {
	return env.Env("METRIC_PORT", "2003")
}
