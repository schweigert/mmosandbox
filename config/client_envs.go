package config

import (
	"fmt"

	"github.com/schweigert/mmosandbox/lib/env"
)

// ClientEnvs configurator
type ClientEnvs struct{}

// WebHost $CLIENT_WHOST localhost
func (client *ClientEnvs) WebHost() string {
	return env.Env("CLIENT_WHOST", "localhost")
}

// WebProtocol $CLIENT_WPROTOCOL http
func (client *ClientEnvs) WebProtocol() string {
	return env.Env("CLIENT_WPROTOCOL", "http")
}

// WebPort $CLIENT_WPORT 3000
func (client *ClientEnvs) WebPort() string {
	return env.Env("CLIENT_WPORT", "3000")
}

// Addr host=%s port=%s user=%s clientname=%s sslmode=%s password=%s
func (client *ClientEnvs) Addr() string {
	return fmt.Sprintf(
		"%s://%s:%s",
		client.WebProtocol(),
		client.WebHost(),
		client.WebPort(),
	)
}
