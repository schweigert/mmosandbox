package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// TokenRepository interface
type TokenRepository struct {
	Client *rpc.Client
}

// NewTokenRepository constructor
func NewTokenRepository() repositories.TokenRepository {
	client, err := rpc.Dial("tcp", config.Service().Crud())
	dont.Panic(err)
	return &TokenRepository{Client: client}
}

// GenerateToken proxy method
func (repo *TokenRepository) GenerateToken(username string) string {
	out := &crudtaskio.TokenRepositoryGenerateTokenOutput{}

	err := repo.Client.Call("CrudTask.TokenRepositoryGenerateToken", username, out)
	dont.Panic(err)

	return out.Token
}

// CheckUsername proxy method
func (repo *TokenRepository) CheckUsername(username string, token string) bool {
	in := &crudtaskio.TokenRepositoryCheckUsernameInput{}
	out := &crudtaskio.TokenRepositoryCheckUsernameOutput{}

	in.Token = token
	in.Username = username

	err := repo.Client.Call("CrudTask.TokenRepositoryCheckUsername", in, out)
	dont.Panic(err)

	return out.Result
}
