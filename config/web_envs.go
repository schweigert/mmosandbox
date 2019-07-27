package config

import (
	"fmt"

	"github.com/schweigert/mmosandbox/lib/env"
)

// WebEnvs configurator
type WebEnvs struct{}

// Interface $WEB_INTERFACE 0.0.0.0
func (web *WebEnvs) Interface() string {
	return env.Env("WEB_INTERFACE", "0.0.0.0")
}

// Port $WEB_PORT 3000
func (web *WebEnvs) Port() string {
	return env.Env("WEB_PORT", "3000")
}

// Addr $WEB_INTERFACE:$WEB_PORT
func (web *WebEnvs) Addr() string {
	return fmt.Sprintf("%s:%s", web.Interface(), web.Port())
}
