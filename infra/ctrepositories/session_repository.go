package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/ulule/deepcopier"
)

// SessionRepository interface
type SessionRepository struct {
	Client *rpc.Client
}

// NewSessionRepository constructor
func NewSessionRepository() repositories.SessionRepository {
	client, err := rpc.Dial("tcp", config.Service().Crud())
	dont.Panic(err)
	return &SessionRepository{Client: client}
}

// StoreAccountToken proxy method
func (repo *SessionRepository) StoreAccountToken(token string, account *entities.Account) {
	in := crudtaskio.SessionRepositoryStoreAccountTokenInput{}
	out := true

	in.Account = account
	in.Token = token

	err := repo.Client.Call("CrudTask.SessionRepositoryStoreAccountToken", in, &out)
	dont.Panic(err)
}

// FindAccountToken proxy method
func (repo *SessionRepository) FindAccountToken(account *entities.Account) {
	in := crudtaskio.SessionRepositoryFindAccountTokenInput{}
	out := &crudtaskio.SessionRepositoryFindAccountTokenInput{}

	err := repo.Client.Call("CrudTask.SessionRepositoryFindAccountToken", in, out)
	dont.Panic(err)

	err = deepcopier.Copy(out.Account).To(account)
	dont.Panic(err)
}
