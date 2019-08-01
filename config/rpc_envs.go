package config

import (
	"fmt"

	"github.com/schweigert/mmosandbox/lib/env"
)

// RPCEnvs configurator
type RPCEnvs struct{}

// Interface $RPC_INTERFACE 0.0.0.0
func (web *RPCEnvs) Interface() string {
	return env.Env("RPC_INTERFACE", "0.0.0.0")
}

// Port $RPC_PORT 3000
func (web *RPCEnvs) Port() string {
	return env.Env("RPC_PORT", "3000")
}

// Addr $RPC_INTERFACE:$RPC_PORT
func (web *RPCEnvs) Addr() string {
	return fmt.Sprintf("%s:%s", web.Interface(), web.Port())
}
