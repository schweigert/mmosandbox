package crudtask

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/merepositories"
	"github.com/schweigert/mmosandbox/infra/tasks/crudtask/crudtaskio"
	"github.com/stretchr/testify/suite"
)

type CrudTaskSuite struct {
	suite.Suite
}

func (suite *CrudTaskSuite) SetupTest() {
	merepositories.UseTokenRepository()
	dbrepositories.UseAccountRepository()
	dbrepositories.UseCharacterRepository()
}

func (suite *CrudTaskSuite) TestNew() {
	suite.NotNil(New())
}

func (suite *CrudTaskSuite) TestAccountRepositoryUsernameHasTaken() {
	defer db.Clear()

	out := crudtaskio.AccountRepositoryUsernameHasTakenOutput{Result: true}
	err := New().AccountRepositoryUsernameHasTaken("testing", &out)
	suite.NoError(err)
	suite.False(out.Result)
}

func (suite *CrudTaskSuite) TestAccountRepositoryCreate() {
	defer db.Clear()

	out := crudtaskio.CrudCreateAccountOutput{Result: false}
	account := entities.Account{Username: "testing"}
	err := New().AccountRepositoryCreate(account, &out)
	suite.NoError(err)
	suite.True(out.Result)

	outTaken := crudtaskio.AccountRepositoryUsernameHasTakenOutput{Result: false}
	err = New().AccountRepositoryUsernameHasTaken("testing", &outTaken)
	suite.NoError(err)
	suite.True(outTaken.Result)
}

func TestCrudTaskSuite(t *testing.T) {
	suite.Run(t, new(CrudTaskSuite))
}
