package repositories

import "github.com/schweigert/mmosandbox/domain/entities"

// AccountRepository interface
type AccountRepository interface {
	UsernameHashTaken(username string) bool
	Create(*entities.Account) bool
}

// CharacterRepository interface
type CharacterRepository interface {
	NameHasTaken(name string) bool
	Create(*entities.Character) bool
}
