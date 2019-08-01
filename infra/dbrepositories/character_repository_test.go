package dbrepositories

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"

	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/stretchr/testify/suite"
)

type CharacterRepositorySuite struct {
	suite.Suite
}

func (suite *CharacterRepositorySuite) TestUseCharacterRepository() {
	suite.Nil(domain.CharacterRepository)
	UseCharacterRepository()
	suite.NotNil(domain.CharacterRepository)
}

func (suite *CharacterRepositorySuite) TestNewCharacterRepository() {
	db.Clear()

	suite.NotPanics(func() {
		suite.NotNil(NewCharacterRepository())
	})
}

func (suite *CharacterRepositorySuite) TestCreateInAccount() {
	db.Clear()

	account := entities.NewAccount()
	account.Email = "testing@kk.onion"
	account.Username = "testing"
	account.SecurePassword = "secure hash"

	suite.True(NewAccountRepository().Create(account))

	character := entities.NewCharacter()
	character.Name = "testing now! ¬¬"

	suite.True(NewCharacterRepository().CreateInAccount(character, account))

	suite.Equal(character.AccountID, account.ID)
	suite.NotEqual(0, account.ID)
}

func (suite *CharacterRepositorySuite) TestNameHasTaken() {
	db.Clear()

	account := entities.NewAccount()
	account.Email = "testing@kk.onion"
	account.Username = "testing"
	account.SecurePassword = "secure hash"

	suite.True(NewAccountRepository().Create(account))
	suite.False(NewCharacterRepository().NameHasTaken("testing now! ¬¬"))

	character := entities.NewCharacter()
	character.Name = "testing now! ¬¬"

	suite.True(NewCharacterRepository().CreateInAccount(character, account))

	suite.True(NewCharacterRepository().NameHasTaken("testing now! ¬¬"))
}

func TestCharacterRepositorySuite(t *testing.T) {
	suite.Run(t, new(CharacterRepositorySuite))
}
