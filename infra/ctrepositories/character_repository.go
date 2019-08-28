package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/lib/dont"
)

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
