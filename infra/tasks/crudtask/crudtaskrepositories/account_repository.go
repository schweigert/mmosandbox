package crudtaskrepositories

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

// CharacterRepository interface
type CharacterRepository struct {
	Client *rpc.Client
}

// NewCharacterRepository constructor
func NewCharacterRepository() repositories.CharacterRepository {
	client, err := rpc.Dial("tcp", config.Service().Crud())
	dont.Panic(err)
	return &CharacterRepository{Client: client}
}

// NameHasTaken proxy method
func (repo *CharacterRepository) NameHasTaken(name string) bool {
	return true
}

// CreateInAccount proxy method
func (repo *CharacterRepository) CreateInAccount(*entities.Character, *entities.Account) bool {
	return true
}

// LoadCharacter proxy method
func (repo *CharacterRepository) LoadCharacter(id int) *entities.Character {
	return nil
}

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
func (repo *SessionRepository) StoreAccountToken(string, *entities.Account) {
	return
}

// FindAccountToken proxy method
func (repo *SessionRepository) FindAccountToken(*entities.Account) {
	return
}

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
