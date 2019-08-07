package tasks

import "github.com/schweigert/mmosandbox/domain/inputs"

// GameTask handle game requests
type GameTask struct{}

// NewGameTask constructor
func NewGameTask() *GameTask {
	return &GameTask{}
}

// SpawnCharacter task
func (task *GameTask) SpawnCharacter(in inputs.SpawnCharacterInput, ok bool) error {
	return nil
}
