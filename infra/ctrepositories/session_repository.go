package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/lib/dont"
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
func (repo *SessionRepository) StoreAccountToken(string, *entities.Account) {}

// FindAccountToken proxy method
func (repo *SessionRepository) FindAccountToken(*entities.Account) {}
