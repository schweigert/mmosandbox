package tasks

import (
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/lib/bench"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// GameTask handle game requests
type GameTask struct {
	SessionConn *rpc.Client
}

// NewGameTask constructor
func NewGameTask() *GameTask {
	conn, err := rpc.Dial("tcp", config.Service().Auth())
	dont.Panic(err)

	return &GameTask{SessionConn: conn}
}

// SpawnCharacter task
func (task *GameTask) SpawnCharacter(in inputs.SpawnCharacterInput, ok bool) error {
	return bench.Bench("spawn_character", func() (err error) {
		checkOut := outputs.NewCheckSessionOutput()
		err = task.SessionConn.Call("SessionTask.StartSession", in.CheckSessionInput, checkOut)

		if err == nil && checkOut.Success {
			// Spawn it!
		}

		return
	})
}
