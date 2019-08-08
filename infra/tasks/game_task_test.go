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

type GameTaskTest struct {
	suite.Suite
}

func (suite *GameTaskTest) TestNewGameTask() {
	task := NewGameTask()
	suite.NotNil(task)
}

func (suite *GameTaskTest) TestStartSession() {
	os.Setenv("DB_NAME", "development")
	defer os.Setenv("DB_NAME", "test")
	db.Clear()

	dbrepositories.UseAccountRepository()
	dbrepositories.UseCharacterRepository()
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

	character := entities.NewCharacter()
	character.Name = "testingName"

	suite.True(domain.CharacterRepository.CreateInAccount(character, account))

	gameTask := NewGameTask()

	spawnCharacterInput := *inputs.NewSpawnCharacterInput()
	spawnCharacterInput.CheckSessionInput = inputs.NewCheckSessionInput()
	spawnCharacterInput.CheckSessionInput.Token = out.Token
	spawnCharacterInput.CheckSessionInput.Username = account.Username

	spawnCharacterInput.CharacterID = int(character.ID)

	spawnCharacterOutput := outputs.NewCheckSessionOutput()
	suite.NotPanics(func() {
		gameTask.SpawnCharacter(spawnCharacterInput, spawnCharacterOutput)
	})

	suite.True(spawnCharacterOutput.Success)

	suite.NotZero(len(gameTask.WorldRules.Characters))
	suite.Equal("testingName", gameTask.WorldRules.Characters[0].Name)
}

func TestGameTaskTest(t *testing.T) {
	suite.Run(t, new(GameTaskTest))
}
