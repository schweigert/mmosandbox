package dbrepositories

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/db"
)

// AccountRepository implements domain repository
type AccountRepository struct {
	conn *gorm.DB
}

// NewAccountRepository constructor
func NewAccountRepository() repositories.AccountRepository {
	return &AccountRepository{conn: db.Connect()}
}

// UsernameHasTaken over gorm
func (repo *AccountRepository) UsernameHasTaken(username string) bool {
	model := &entities.Account{}
	repo.conn.Where("username = ?", username).First(model)

	return !repo.conn.NewRecord(model)
}

// Create over gorm
func (repo *AccountRepository) Create(model *entities.Account) bool {
	repo.conn.Create(model)
	return !repo.conn.NewRecord(model)
}
