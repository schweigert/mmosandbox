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

type CharacterRepositoryLoadCharacter struct {
	Character *entities.Character
}
