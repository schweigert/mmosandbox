package security

import (
	"crypto/sha256"
	"fmt"
)

// Sum data
func Sum(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}
