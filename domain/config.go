package domain

import (
	"github.com/schweigert/mmosandbox/domain/repositories"
)

// Repositories
var (
	TokenRepository     repositories.TokenRepository
	SessionRepository   repositories.SessionRepository
	AccountRepository   repositories.AccountRepository
	CharacterRepository repositories.CharacterRepository
)
