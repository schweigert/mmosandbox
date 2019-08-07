package tasks

import (
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/lib/bench"
)

// SessionTask call auth methods from domain over rpc
type SessionTask struct{}

// NewSessionTask constructor
func NewSessionTask() *SessionTask {
	return &SessionTask{}
}

// StartSession and return the assigned account object
func (task *SessionTask) StartSession(in inputs.AuthAccountInput, out *outputs.StartSessionOutput) (err error) {
	return bench.Bench("start_session", func() error {
		domain.NewSessionRules().StartSession(&in, out)
		return nil
	})
}

// CheckSession in session database
func (task *SessionTask) CheckSession(in inputs.CheckSessionInput, out *outputs.CheckSessionOutput) (err error) {
	return bench.Bench("check_session", func() error {
		out.Success = domain.NewSessionRules().CheckSession(&in)
		return nil
	})
}
