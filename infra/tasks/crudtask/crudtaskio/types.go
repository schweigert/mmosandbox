package crudtaskio

import "github.com/schweigert/mmosandbox/domain/entities"

// CrudCreateAccountOutput struct
type CrudCreateAccountOutput struct {
	Result bool
}

// UsernameAndPasswordAreEqualOutput struct
type UsernameAndPasswordAreEqualOutput struct {
	Result bool
}

// AccountRepositoryUsernameHasTakenOutput struct
type AccountRepositoryUsernameHasTakenOutput struct {
	Result bool
}

// CharacterRepositoryNameHasTakenOutput struct
type CharacterRepositoryNameHasTakenOutput struct {
	Result bool
}

// CharacterRepositoryCreateInAccountInput struct
type CharacterRepositoryCreateInAccountInput struct {
	Character *entities.Character
	Account   *entities.Account
}

// CharacterRepositoryCreateInAccountOutput struct
type CharacterRepositoryCreateInAccountOutput struct {
	Result bool
}

// CharacterRepositoryLoadCharacterOutput struct
type CharacterRepositoryLoadCharacterOutput struct {
	Character *entities.Character
}

// SessionRepositoryStoreAccountTokenInput struct
type SessionRepositoryStoreAccountTokenInput struct {
	Token   string
	Account *entities.Account
}

// SessionRepositoryFindAccountTokenInput struct
type SessionRepositoryFindAccountTokenInput struct {
	Account *entities.Account
}

// SessionRepositoryFindAccountTokenOutput struct
type SessionRepositoryFindAccountTokenOutput SessionRepositoryFindAccountTokenInput

// TokenRepositoryGenerateTokenOutput struct
type TokenRepositoryGenerateTokenOutput struct {
	Token string
}

// TokenRepositoryCheckUsernameInput struct
type TokenRepositoryCheckUsernameInput struct {
	Username string
	Token    string
}

// TokenRepositoryCheckUsernameOutput struct
type TokenRepositoryCheckUsernameOutput struct {
	Result bool
}
