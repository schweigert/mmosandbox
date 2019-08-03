package config

import (
	"fmt"

	"github.com/schweigert/mmosandbox/lib/env"
)

// MeEnvs configurator
type MeEnvs struct{}

// Host $ME_HOST 0.0.0.0
func (me *MeEnvs) Host() string {
	return env.Env("ME_HOST", "localhost")
}

// Port $ME_PORT 6379
func (me *MeEnvs) Port() string {
	return env.Env("ME_PORT", "6379")
}

// Password $ME_PASSWORD nopasswd
func (me *MeEnvs) Password() string {
	return env.Env("ME_PASSWORD", "nopasswd")
}

// Addr $ME_INTERFACE:$ME_PORT
func (me *MeEnvs) Addr() string {
	return fmt.Sprintf("%s:%s", me.Host(), me.Port())
}
