package tasks

import (
	"os"
	"testing"

	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/merepositories"
	"github.com/stretchr/testify/suite"
)

type SessionProxyTaskSuite struct {
	suite.Suite
}

func (suite *SessionProxyTaskSuite) TestNewSessionProxyTask() {
	SessionProxyTask := NewSessionProxyTask()
	suite.NotNil(SessionProxyTask)
}

func (suite *SessionProxyTaskSuite) TestStartSession() {
	os.Setenv("DB_NAME", "development")
	defer os.Setenv("DB_NAME", "test")
	db.Clear()

	dbrepositories.UseAccountRepository()
	merepositories.UseTokenRepository()

	account := entities.NewAccount()
	account.Username = "testing new sessions"
	account.SecurePassword = "testing new sessions"

	suite.True(domain.AccountRepository.Create(account))

	in := *inputs.NewAuthAccountInput()
	in.Username = account.Username
	in.SecurePassword = account.SecurePassword

	task := NewSessionProxyTask()
	out := outputs.NewStartSessionOutput()

	err := task.StartSession(in, out)

	suite.NoError(err)
	suite.True(out.Success)
	suite.NotEmpty(out.Token)
}

func TestSessionProxyTaskSuite(t *testing.T) {
	suite.Run(t, new(SessionProxyTaskSuite))
}
