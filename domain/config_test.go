package domain

import (
	"github.com/schweigert/mmosandbox/domain/entities"
)

// AccountRepository interface mock
type AccountRepositoryMock struct {
	UsernameHasTakenResult            bool
	UsernameAndPasswordAreEqualResult bool
	CreateResult                      bool
}

func (mock *AccountRepositoryMock) UsernameHasTaken(string) bool {
	return mock.UsernameHasTakenResult
}

func (mock *AccountRepositoryMock) UsernameAndPasswordAreEqual(*entities.Account) bool {
	return mock.UsernameAndPasswordAreEqualResult
}

func (mock *AccountRepositoryMock) Create(*entities.Account) bool {
	return mock.CreateResult
}

type CharacterRepositoryMock struct {
	CreateInAccountResult bool
	NameHasTakenResult    bool
	LoadCharacterResult   *entities.Character
}

func (mock *CharacterRepositoryMock) NameHasTaken(string) bool {
	return mock.NameHasTakenResult
}

func (mock *CharacterRepositoryMock) CreateInAccount(*entities.Character, *entities.Account) bool {
	return mock.CreateInAccountResult
}

func (mock *CharacterRepositoryMock) LoadCharacter(id int) *entities.Character {
	return mock.LoadCharacterResult
}

type TokenRepositoryMock struct {
	GenerateTokenResult string
	CheckUsernameResult bool
}

func (mock *TokenRepositoryMock) GenerateToken(username string) string {
	return mock.GenerateTokenResult
}

func (mock *TokenRepositoryMock) CheckUsername(username, token string) bool {
	return mock.CheckUsernameResult
}
