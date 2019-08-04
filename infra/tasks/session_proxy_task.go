package tasks

import (
	"fmt"
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// SessionProxyTask call auth methods from domain over rpc
type SessionProxyTask struct {
	SessionConn *rpc.Client
}

// NewSessionProxyTask constructor
func NewSessionProxyTask() *SessionProxyTask {
	conn, err := rpc.Dial("tcp", config.Service().Auth())
	dont.Panic(err)

	return &SessionProxyTask{SessionConn: conn}
}

// StartSession and return the assigned account object
func (task *SessionProxyTask) StartSession(in inputs.AuthAccountInput, out *outputs.StartSessionOutput) (err error) {
	fmt.Println("SessionProxyTask |> StartSession")

	err = task.SessionConn.Call("SessionTask.StartSession", in, out)
	return
}
