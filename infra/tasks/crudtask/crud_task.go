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
func (task *CrudTask) AccountRepositoryUsernameHasTaken(username string, out *crudtaskio.AccountRepositoryUsernameHasTaken) error {
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
