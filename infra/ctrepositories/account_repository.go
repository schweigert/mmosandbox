package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
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

// UsernameHasTaken proxy method
func (repo *AccountRepository) UsernameHasTaken(username string) bool {
	return true
}

// Create proxy method
func (repo *AccountRepository) Create(*entities.Account) bool {
	return true
}

// UsernameAndPasswordAreEqual proxy method
func (repo *AccountRepository) UsernameAndPasswordAreEqual(*entities.Account) bool {
	return true
}
