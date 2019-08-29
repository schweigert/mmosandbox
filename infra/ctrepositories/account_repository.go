package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// AccountRepository interface
type AccountRepository struct {
	Client *rpc.Client
}

// NewAccountRepository constructor
func NewAccountRepository() repositories.AccountRepository {
	client, err := rpc.Dial("tcp", config.Service().Crud())
	dont.Panic(err)
	return &AccountRepository{Client: client}
}

// UseAccountRepository in domain
func UseAccountRepository() {
	domain.AccountRepository = NewAccountRepository()
}

// UsernameHasTaken proxy method
func (repo *AccountRepository) UsernameHasTaken(username string) bool {
	out := &crudtaskio.AccountRepositoryUsernameHasTakenOutput{}
	err := repo.Client.Call("CrudTask.AccountRepositoryUsernameHasTaken", username, out)

	dont.Panic(err)

	return out.Result
}

// Create proxy method
func (repo *AccountRepository) Create(account *entities.Account) bool {
	out := &crudtaskio.CrudCreateAccountOutput{}
	err := repo.Client.Call("CrudTask.AccountRepositoryCreate", account, out)

	dont.Panic(err)

	return out.Result
}

// UsernameAndPasswordAreEqual proxy method
func (repo *AccountRepository) UsernameAndPasswordAreEqual(account *entities.Account) bool {
	out := &crudtaskio.UsernameAndPasswordAreEqualOutput{}
	err := repo.Client.Call("CrudTask.AccountRepositoryUsernameAndPasswordAreEqual", account, out)

	dont.Panic(err)

	return out.Result
}
