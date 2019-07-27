package config

import (
	"fmt"

	"github.com/schweigert/mmosandbox/lib/env"
)

// DbEnvs configurator
type DbEnvs struct{}

// Host $DB_HOST 0.0.0.0
func (db *DbEnvs) Host() string {
	return env.Env("DB_HOST", "localhost")
}

// Port $DB_PORT 5432
func (db *DbEnvs) Port() string {
	return env.Env("DB_PORT", "5432")
}

// User $DB_USER postgres
func (db *DbEnvs) User() string {
	return env.Env("DB_USER", "postgres")
}

// Name $DB_NAME test
func (db *DbEnvs) Name() string {
	return env.Env("DB_NAME", "test")
}

// SSLMode $DB_SSLMODE disable
func (db *DbEnvs) SSLMode() string {
	return env.Env("DB_SSLMODE", "disable")
}

// Password $DB_PASSWORD postgres
func (db *DbEnvs) Password() string {
	return env.Env("DB_PASSWORD", "postgres")
}

// Addr host=%s port=%s user=%s dbname=%s sslmode=%s password=%s
func (db *DbEnvs) Addr() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		db.Host(),
		db.Port(),
		db.User(),
		db.Name(),
		db.SSLMode(),
		db.Password(),
	)
}
