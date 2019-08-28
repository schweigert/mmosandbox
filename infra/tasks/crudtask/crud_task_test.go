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

func (suite *CrudTaskSuite) TestNewCrudTask() {
	suite.NotNil(NewCrudTask())
}

func (suite *CrudTaskSuite) TestAccountRepositoryUsernameHasTaken() {
	defer db.Clear()

	out := crudtaskio.AccountRepositoryUsernameHasTakenOutput{Result: true}
	NewCrudTask().AccountRepositoryUsernameHasTaken("testing", &out)
	suite.False(out.Result)
}

func (suite *CrudTaskSuite) TestAccountRepositoryCreate() {
	defer db.Clear()

	out := crudtaskio.CrudCreateAccountOutput{Result: false}
	account := entities.Account{Username: "testing"}
	NewCrudTask().AccountRepositoryCreate(account, &out)
	suite.True(out.Result)

	outTaken := crudtaskio.AccountRepositoryUsernameHasTakenOutput{Result: false}
	NewCrudTask().AccountRepositoryUsernameHasTaken("testing", &outTaken)
	suite.True(outTaken.Result)
}

func TestCrudTaskSuite(t *testing.T) {
	suite.Run(t, new(CrudTaskSuite))
}
