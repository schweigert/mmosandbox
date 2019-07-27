package env

import "os"

// Env with default value
func Env(key string, def string) string {
	value := os.Getenv(key)
	if isBlank(value) {
		return def
	}

	return value
}

func isBlank(value string) bool {
	return value == ""
}
