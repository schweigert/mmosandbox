package ctrepositories

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
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
	out := &crudtaskio.CharacterRepositoryNameHasTakenOutput{}
	err := repo.Client.Call("CrudTask.CharacterRepositoryNameHasTaken", name, out)

	dont.Panic(err)

	return out.Result
}

// CreateInAccount proxy method
func (repo *CharacterRepository) CreateInAccount(character *entities.Character, account *entities.Account) bool {
	out := &crudtaskio.CharacterRepositoryCreateInAccountOutput{}
	in := crudtaskio.CharacterRepositoryCreateInAccountInput{}

	in.Account = account
	in.Character = character

	err := repo.Client.Call("CrudTask.CharacterRepositoryCreateInAccount", in, out)

	dont.Panic(err)

	return out.Result
}

// LoadCharacter proxy method
func (repo *CharacterRepository) LoadCharacter(id int) *entities.Character {
	out := &crudtaskio.CharacterRepositoryLoadCharacterOutput{}
	err := repo.Client.Call("CrudTask.CharacterRepositoryLoadCharacter", id, out)

	dont.Panic(err)

	return out.Character
}
