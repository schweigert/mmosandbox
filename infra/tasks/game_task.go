package tasks

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/domain"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/lib/bench"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// GameTask handle game requests
type GameTask struct {
	SessionConn *rpc.Client
	WorldRules  *domain.WorldRules
}

// NewGameTask constructor
func NewGameTask() *GameTask {
	conn, err := rpc.Dial("tcp", config.Service().Auth())
	dont.Panic(err)

	return &GameTask{
		SessionConn: conn,
		WorldRules:  domain.NewWorldRules(),
	}
}

// SpawnCharacter task
func (task *GameTask) SpawnCharacter(in inputs.SpawnCharacterInput, out *outputs.CheckSessionOutput) error {
	return bench.Bench("spawn_character", func() (err error) {
		checkOut := outputs.NewCheckSessionOutput()
		err = task.SessionConn.Call("SessionTask.CheckSession", in.CheckSessionInput, checkOut)

		out.Success = checkOut.Success

		if err == nil && out.Success {
			task.WorldRules.SpawnCharacter(in.CharacterID)
		}

		return err
	})
}

// MoveCharacter task
func (task *GameTask) MoveCharacter(in inputs.MoveCharacterInput, out *outputs.CheckSessionOutput) error {
	return bench.Bench("move_character", func() (err error) {
		checkOut := outputs.NewCheckSessionOutput()
		err = task.SessionConn.Call("SessionTask.CheckSession", in.CheckSessionInput, checkOut)

		out.Success = checkOut.Success

		if err == nil && out.Success {
			task.WorldRules.MoveCharacter(&in)
		}

		return err
	})
}
