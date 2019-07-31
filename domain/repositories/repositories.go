package repositories

import "github.com/schweigert/mmosandbox/domain/entities"

// AccountRepository interface
type AccountRepository interface {
	UsernameHasTaken(username string) bool
	Create(*entities.Account) bool
	UsernameAndPasswordAreEqual(*entities.Account) bool
}

// CharacterRepository interface
type CharacterRepository interface {
	NameHasTaken(name string) bool
	Create(*entities.Character) bool
}
