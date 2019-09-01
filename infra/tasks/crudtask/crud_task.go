package crudtask

import (
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
	"github.com/schweigert/mmosandbox/lib/bench"
)

// CrudTask struct
type CrudTask struct {
	AccountRepository   repositories.AccountRepository
	CharacterRepository repositories.CharacterRepository
	SessionRepository   repositories.SessionRepository
	TokenRepository     repositories.TokenRepository
}

// New constructor
func New() *CrudTask {
	return &CrudTask{
		AccountRepository:   domain.AccountRepository,
		CharacterRepository: domain.CharacterRepository,
		SessionRepository:   domain.SessionRepository,
		TokenRepository:     domain.TokenRepository,
	}
}

// AccountRepositoryUsernameHasTaken method
func (task *CrudTask) AccountRepositoryUsernameHasTaken(username string, out *crudtaskio.AccountRepositoryUsernameHasTakenOutput) error {
	return bench.Bench("account_repository_username_has_taken", func() error {
		out.Result = task.AccountRepository.UsernameHasTaken(username)
		return nil
	})
}

// AccountRepositoryCreate method
func (task *CrudTask) AccountRepositoryCreate(account entities.Account, out *crudtaskio.CrudCreateAccountOutput) error {
	return bench.Bench("account_repository_username_has_taken", func() error {
		out.Result = task.AccountRepository.Create(&account)
		return nil
	})
}

// AccountRepositoryUsernameAndPasswordAreEqual method
func (task *CrudTask) AccountRepositoryUsernameAndPasswordAreEqual(account entities.Account, out *crudtaskio.UsernameAndPasswordAreEqualOutput) error {
	return bench.Bench("account_repository_username_and_password_are_equal", func() error {
		out.Result = task.AccountRepository.UsernameAndPasswordAreEqual(&account)
		return nil
	})
}

// CharacterRepositoryNameHasTaken method
func (task *CrudTask) CharacterRepositoryNameHasTaken(name string, out *crudtaskio.CharacterRepositoryNameHasTakenOutput) error {
	return bench.Bench("account_repository_name_has_taken", func() error {
		out.Result = task.CharacterRepository.NameHasTaken(name)
		return nil
	})
}

// CharacterRepositoryCreateInAccount method
func (task *CrudTask) CharacterRepositoryCreateInAccount(in crudtaskio.CharacterRepositoryCreateInAccountInput, out *crudtaskio.CharacterRepositoryCreateInAccountOutput) error {
	return bench.Bench("account_repository_create_in_account", func() error {
		out.Result = task.CharacterRepository.CreateInAccount(in.Character, in.Account)
		return nil
	})
}

// CharacterRepositoryLoadCharacter method
func (task *CrudTask) CharacterRepositoryLoadCharacter(id int, out *crudtaskio.CharacterRepositoryLoadCharacterOutput) error {
	return bench.Bench("account_repository_load_character", func() error {
		out.Character = task.CharacterRepository.LoadCharacter(id)
		return nil
	})
}

// SessionRepositoryStoreAccountToken method
func (task *CrudTask) SessionRepositoryStoreAccountToken(in crudtaskio.SessionRepositoryStoreAccountTokenInput, ok *bool) error {
	return bench.Bench("account_repository_store_acount_token", func() error {
		task.SessionRepository.StoreAccountToken(in.Token, in.Account)
		return nil
	})
}

// SessionRepositoryFindAccountToken method
func (task *CrudTask) SessionRepositoryFindAccountToken(in crudtaskio.SessionRepositoryFindAccountTokenInput, out *crudtaskio.SessionRepositoryFindAccountTokenInput) error {
	return bench.Bench("account_repository_find_account_token", func() error {
		task.SessionRepository.FindAccountToken(in.Account)
		out.Account = in.Account
		return nil
	})
}

// TokenRepositoryGenerateToken method
func (task *CrudTask) TokenRepositoryGenerateToken(username string, out *crudtaskio.TokenRepositoryGenerateTokenOutput) error {
	return bench.Bench("account_repository_generate_token", func() error {
		out.Token = task.TokenRepository.GenerateToken(username)
		return nil
	})
}

// TokenRepositoryCheckUsername method
func (task *CrudTask) TokenRepositoryCheckUsername(in crudtaskio.TokenRepositoryCheckUsernameInput, out *crudtaskio.TokenRepositoryCheckUsernameOutput) error {
	return bench.Bench("account_repository_check_username", func() error {
		out.Result = task.TokenRepository.CheckUsername(in.Username, in.Token)
		return nil
	})
}
