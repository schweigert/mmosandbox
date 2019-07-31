package dbrepositories

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/db"
)

// AccountRepository implements domain repository
type AccountRepository struct {
	conn *gorm.DB
}

// UseAccountRepository in domain
func UseAccountRepository() {
	domain.AccountRepository = NewAccountRepository()
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

// UsernameAndPasswordAreEqual over gorm
func (repo *AccountRepository) UsernameAndPasswordAreEqual(model *entities.Account) bool {
	query := &entities.Account{Username: model.Username, SecurePassword: model.SecurePassword}

	repo.conn.Where("username = ? and secure_password = ?", query.Username, query.SecurePassword).Take(query)

	return !repo.conn.NewRecord(query)
}
