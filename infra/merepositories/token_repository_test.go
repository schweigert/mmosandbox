package merepositories

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/infra/cache"

	"github.com/stretchr/testify/suite"
)

type TokenRepositorySuite struct {
	suite.Suite
}

func (suite *TokenRepositorySuite) TestNewTokenRepository() {
	suite.NotPanics(func() {
		suite.NotNil(NewTokenRepository())
	})
}

func (suite *TokenRepositorySuite) TestUseTokenRepository() {
	suite.Nil(domain.TokenRepository)

	UseTokenRepository()
	suite.NotNil(domain.TokenRepository)
}

func (suite *TokenRepositorySuite) TestGenerateToken() {
	conn := cache.Connect()
	defer conn.Close()

	repository := NewTokenRepository()

	token := repository.GenerateToken("working suite")
	suite.NotEmpty(token)
	tokenCheck, err := conn.Get("session.working suite").Result()

	suite.NoError(err)
	suite.Equal(token, tokenCheck)
}

func TestTokenRepositorySuite(t *testing.T) {
	suite.Run(t, new(TokenRepositorySuite))
}
