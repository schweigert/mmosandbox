package presenters

import "strings"

// WorldNamePresenter replace spaces and store as lower case
func WorldNamePresenter(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	return strings.ToLower(name)
}
