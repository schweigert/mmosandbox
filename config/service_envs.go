package config

import (
	"github.com/schweigert/mmosandbox/lib/env"
)

// ServiceEnvs configurator
type ServiceEnvs struct{}

// Name $SERVICE_NAME NONAME
func (service *ServiceEnvs) Name() string {
	return env.Env("SERVICE_NAME", "NONAME")
}
