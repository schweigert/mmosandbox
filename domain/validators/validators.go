package validators

import "github.com/badoux/checkmail"

// StringSize validator
func StringSize(str string, min uint, max uint) bool {
	size := uint(len(str))
	return size >= min && size <= max
}

// Email validator
func Email(email string) bool {
	return checkmail.ValidateFormat(email) == nil
}
