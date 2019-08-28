package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/repositories"
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
	return ""
}

// CheckUsername proxy method
func (repo *TokenRepository) CheckUsername(username string, token string) bool {
	return true
}
