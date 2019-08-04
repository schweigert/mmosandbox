package tasks

import (
	"fmt"

	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
)

// SessionTask call auth methods from domain over rpc
type SessionTask struct{}

// NewSessionTask constructor
func NewSessionTask() *SessionTask {
	return &SessionTask{}
}

// StartSession and return the assigned account object
func (task *SessionTask) StartSession(in inputs.AuthAccountInput, out *outputs.StartSessionOutput) (err error) {
	fmt.Println("SessionTask |> StartSession")

	domain.NewSessionRules().StartSession(&in, out)
	return
}
