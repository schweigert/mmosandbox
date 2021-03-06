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
	CreateInAccount(*entities.Character, *entities.Account) bool
	LoadCharacter(id int) *entities.Character
}

// SessionRepository interface
type SessionRepository interface {
	StoreAccountToken(string, *entities.Account)
	FindAccountToken(*entities.Account)
}

// TokenRepository interface
type TokenRepository interface {
	GenerateToken(username string) string
	CheckUsername(username string, token string) bool
}
