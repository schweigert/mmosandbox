package routes

import "fmt"

// URL with base
func URL(base string, route string) string {
	return fmt.Sprintf("%s%s", base, route)
}
