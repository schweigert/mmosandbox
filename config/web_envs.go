package config

import "github.com/schweigert/mmosandbox/lib/env"

// WebEnvs configurator
type WebEnvs struct{}

// Interface $WEB_INTERFACE 0.0.0.0
func (web *WebEnvs) Interface() string {
	return env.Env("WEB_INTERFACE", "0.0.0.0")
}
