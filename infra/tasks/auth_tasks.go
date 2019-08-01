package tasks

import (
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
)

// AuthTasks call auth methods from domain over rpc
type AuthTasks struct{}

// NewAuthTasks constructor
func NewAuthTasks() *AuthTasks {
	return &AuthTasks{}
}

// AuthAccount and return the assigned account object
func (task *AuthTasks) AuthAccount(in inputs.AuthAccountInput, out *outputs.AuthAccountOutput) (err error) {
	return nil
}
