package merepositories

import (
	"fmt"
	"time"

	"github.com/dmgk/faker"

	"github.com/go-redis/redis"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/cache"
	"github.com/schweigert/mmosandbox/lib/dont"
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
	repository.set(fmt.Sprintf("session.%s", username), token)

	return token
}

func (repository *TokenRepository) randomToken() string {
	return faker.RandomString(256)
}

func (repository *TokenRepository) set(key string, value string) string {
	_, err := repository.Conn.Set(key, value, time.Hour*128).Result()
	dont.Panic(err)

	return value
}