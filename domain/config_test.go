package domain

import "github.com/schweigert/mmosandbox/domain/entities"

// AccountRepository interface mock
type AccountRepositoryMock struct {
	UsernameHasTakenResult bool
	CreateResult           bool
}

func (mock *AccountRepositoryMock) UsernameHasTaken(string) bool {
	return mock.UsernameHasTakenResult
}

func (mock *AccountRepositoryMock) Create(*entities.Account) bool {
	return mock.CreateResult
}
