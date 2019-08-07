package tasks

import (
	"testing"

	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"

	"github.com/schweigert/mmosandbox/domain/entities"

	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/infra/db"

	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/merepositories"
	"github.com/stretchr/testify/suite"
)

type SessionTaskSuite struct {
	suite.Suite
}

func (suite *SessionTaskSuite) TestNewSessionTask() {
	SessionTask := NewSessionTask()
	suite.NotNil(SessionTask)
}

func (suite *SessionTaskSuite) TestStartSession() {
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

	task := NewSessionTask()
	out := outputs.NewStartSessionOutput()

	err := task.StartSession(in, out)
	suite.NoError(err)

	suite.True(out.Success)
	suite.NotEmpty(out.Token)
}

func (suite *SessionTaskSuite) TestCheckSession() {
	db.Clear()

	dbrepositories.UseAccountRepository()
	merepositories.UseTokenRepository()

	account := entities.NewAccount()
	account.Username = "secured username"
	account.SecurePassword = "super secret password"

	suite.True(domain.AccountRepository.Create(account))

	in := *inputs.NewAuthAccountInput()
	in.Username = account.Username
	in.SecurePassword = account.SecurePassword

	task := NewSessionTask()
	out := outputs.NewStartSessionOutput()

	err := task.StartSession(in, out)
	suite.NoError(err)

	suite.True(out.Success)
	suite.NotEmpty(out.Token)

	checkIn := *inputs.NewCheckSessionInput()
	checkIn.Username = in.Username
	checkIn.Token = out.Token

	checkOut := outputs.NewCheckSessionOutput()
	suite.NoError(task.CheckSession(checkIn, checkOut))

	suite.True(checkOut.Success)
}

func TestSessionTaskSuite(t *testing.T) {
	suite.Run(t, new(SessionTaskSuite))
}
