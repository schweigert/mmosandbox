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

// Auth $SERVICE_AUTH localhost:3101
func (service *ServiceEnvs) Auth() string {
	return env.Env("SERVICE_AUTH", "localhost:3101")
}

// Game $SERVICE_GAME localhost:3101
func (service *ServiceEnvs) Game() string {
	return env.Env("SERVICE_GAME", "localhost:3100")
}
