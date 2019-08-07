package merepositories

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/cache"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/randomize"
)

// TokenRepository struct
type TokenRepository struct {
	Conn *redis.Client
}

// NewTokenRepository constructor
func NewTokenRepository() repositories.TokenRepository {
	return &TokenRepository{Conn: cache.Connect()}
}

// UseTokenRepository in domain
func UseTokenRepository() {
	domain.TokenRepository = NewTokenRepository()
}

// GenerateToken based in username from account
func (repository *TokenRepository) GenerateToken(username string) string {
	token := repository.randomToken()
	repository.set(repository.keyPattern(username), token)

	return token
}

// CheckUsername in redis
func (repository *TokenRepository) CheckUsername(username string, token string) bool {
	correctToken, err := repository.Conn.Get(repository.keyPattern(username)).Result()
	dont.Panic(err)

	return correctToken == token
}

func (repository *TokenRepository) keyPattern(username string) string {
	return fmt.Sprintf("session.%s", username)
}

func (repository *TokenRepository) randomToken() string {
	return randomize.RandStringRunes(256)
}

func (repository *TokenRepository) set(key string, value string) string {
	_, err := repository.Conn.Set(key, value, time.Hour*128).Result()
	dont.Panic(err)

	return value
}
