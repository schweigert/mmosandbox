package crudtask

import (
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
)

// CrudTask struct
type CrudTask struct {
	AccountRepository   repositories.AccountRepository
	CharacterRepository repositories.CharacterRepository
	SessionRepository   repositories.SessionRepository
	TokenRepository     repositories.TokenRepository
}

// NewCrudTask constructor
func NewCrudTask() *CrudTask {
	return &CrudTask{
		AccountRepository:   domain.AccountRepository,
		CharacterRepository: domain.CharacterRepository,
		SessionRepository:   domain.SessionRepository,
		TokenRepository:     domain.TokenRepository,
	}
}

// AccountRepositoryUsernameHasTaken method
func (task *CrudTask) AccountRepositoryUsernameHasTaken(username string, out *crudtaskio.AccountRepositoryUsernameHasTakenOutput) error {
	out.Result = task.AccountRepository.UsernameHasTaken(username)
	return nil
}

// AccountRepositoryCreate method
func (task *CrudTask) AccountRepositoryCreate(account entities.Account, out *crudtaskio.CrudCreateAccountOutput) error {
	out.Result = task.AccountRepository.Create(&account)
	return nil
}

// AccountRepositoryUsernameAndPasswordAreEqual method
func (task *CrudTask) AccountRepositoryUsernameAndPasswordAreEqual(account entities.Account, out *crudtaskio.UsernameAndPasswordAreEqualOutput) error {
	out.Result = task.AccountRepository.UsernameAndPasswordAreEqual(&account)
	return nil
}

// CharacterRepositoryNameHasTaken method
func (task *CrudTask) CharacterRepositoryNameHasTaken(name string, out *crudtaskio.CharacterRepositoryNameHasTakenOutput) error {
	out.Result = task.CharacterRepository.NameHasTaken(name)
	return nil
}

// CharacterRepositoryCreateInAccount method
func (task *CrudTask) CharacterRepositoryCreateInAccount(in crudtaskio.CharacterRepositoryCreateInAccountInput, out *crudtaskio.CharacterRepositoryCreateInAccountOutput) error {
	out.Result = task.CharacterRepository.CreateInAccount(in.Character, in.Account)
	return nil
}

// CharacterRepositoryLoadCharacter method
func (task *CrudTask) CharacterRepositoryLoadCharacter(id int, out *crudtaskio.CharacterRepositoryLoadCharacterOutput) error {
	out.Character = task.CharacterRepository.LoadCharacter(id)
	return nil
}

// FindAccountToken(*entities.Account)

// SessionRepositoryStoreAccountToken method
func (task *CrudTask) SessionRepositoryStoreAccountToken(in crudtaskio.SessionRepositoryStoreAccountTokenInput, ok *bool) error {
	task.SessionRepository.StoreAccountToken(in.Token, in.Account)
	return nil
}

// SessionRepositoryFindAccountToken method
func (task *CrudTask) SessionRepositoryFindAccountToken(in crudtaskio.SessionRepositoryFindAccountTokenInput, out *crudtaskio.SessionRepositoryFindAccountTokenInput) error {
	task.SessionRepository.FindAccountToken(in.Account)
	out.Account = in.Account
	return nil
}
