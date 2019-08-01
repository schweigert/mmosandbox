package dbrepositories

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/db"
)

// CharacterRepository implements domain repository
type CharacterRepository struct {
	conn *gorm.DB
}

// UseCharacterRepository in domain
func UseCharacterRepository() {
	domain.CharacterRepository = NewCharacterRepository()
}

// NewCharacterRepository constructor
func NewCharacterRepository() repositories.CharacterRepository {
	return &CharacterRepository{conn: db.Connect()}
}

// NameHasTaken over gorm
func (repo *CharacterRepository) NameHasTaken(username string) bool {
	model := &entities.Character{}
	repo.conn.Where("name = ?", username).First(model)

	return !repo.conn.NewRecord(model)
}

// CreateInAccount over gorm
func (repo *CharacterRepository) CreateInAccount(model *entities.Character, account *entities.Account) bool {
	repo.conn.Where("username = ?", account.Username).First(account)

	model.AccountID = account.ID
	repo.conn.Create(model)

	return !repo.conn.NewRecord(model)
}
