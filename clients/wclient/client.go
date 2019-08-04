package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"time"

	"github.com/imroc/req"
	"github.com/schweigert/mmosandbox/clients/client"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/infra/routes"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// CreateAccountFlow implementation for client.CreateAccountFlow
type CreateAccountFlow struct {
}

// NewCreateAccountFlow constructor
func NewCreateAccountFlow() client.CreateAccountFlow {
	return &CreateAccountFlow{}
}

// CreateAccountOperation for willson flow
func (flow *CreateAccountFlow) CreateAccountOperation(in *inputs.CreateAccountInput) (*entities.Account, bool) {
	rec := req.New()
	rec.SetTimeout(5 * time.Second)

	resp, err := rec.Put(routes.URL(config.Client().Addr(), routes.Account), req.BodyJSON(in))
	dont.Panic(err)

	out := outputs.NewCreateAccountOutput()

	if resp.Response().StatusCode != http.StatusCreated {
		return nil, false
	}

	err = resp.ToJSON(out)
	dont.Panic(err)

	fmt.Println("Account", out.Account.ID, "|>", out.Account)

	return out.Account, out.Success
}

// CreateCharacterFlow implementation for client.CreateCharacterFlow
type CreateCharacterFlow struct {
}

// NewCreateCharacterFlow constructor
func NewCreateCharacterFlow() client.CreateCharacterFlow {
	return &CreateCharacterFlow{}
}

// CreateCharacterOperation for willson flow
func (flow *CreateCharacterFlow) CreateCharacterOperation(in *inputs.CreateCharacterInput) (*entities.Character, bool) {
	rec := req.New()
	rec.SetTimeout(5 * time.Second)

	resp, err := rec.Put(routes.URL(config.Client().Addr(), routes.Character), req.BodyJSON(in))
	dont.Panic(err)

	out := outputs.NewCreateCharacterOutput()

	if resp.Response().StatusCode != http.StatusCreated {
		return nil, false
	}

	err = resp.ToJSON(out)
	dont.Panic(err)

	fmt.Println("Character", out.Character.ID, "|>", out.Character)

	return out.Character, out.Success
}

// StartSessionFlow used in client
type StartSessionFlow struct {
	Conn *rpc.Client
}

// NewStartSessionFlow constructor
func NewStartSessionFlow() client.StartSessionFlow {
	conn, err := rpc.Dial("tcp", config.Service().Game())
	dont.Panic(err)

	return &StartSessionFlow{Conn: conn}
}

// StartSession for willson flow
func (flow *StartSessionFlow) StartSession(in inputs.AuthAccountInput) (*outputs.StartSessionOutput, bool) {
	out := outputs.NewStartSessionOutput()

	err := flow.Conn.Call("SessionProxyTask.StartSession", in, out)
	return out, err == nil
}

func main() {
	client.UsedCreateAccountFlow = NewCreateAccountFlow()
	client.UsedCreateCharacterFlow = NewCreateCharacterFlow()
	client.UsedStartSessionFlow = NewStartSessionFlow()

	for {
		client.BotFlow()
		time.Sleep(time.Minute)
	}
}
