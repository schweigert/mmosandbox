package main

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/imroc/req"
	"github.com/schweigert/mmosandbox/clients/client"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/infra/routes"
	"github.com/schweigert/mmosandbox/lib/bench"
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
	out := outputs.NewCreateAccountOutput()

	_ = bench.Bench("create_account_operation", func() error {
		rec := req.New()

		resp, err := rec.Put(routes.URL(config.Client().Addr(), routes.Account), req.BodyJSON(in))
		dont.Panic(err)

		if resp.Response().StatusCode != http.StatusCreated {
			out.Account = nil
			out.Success = false
			return nil
		}

		err = resp.ToJSON(out)
		dont.Panic(err)

		fmt.Println("Account", out.Account.ID, "|>", out.Account)
		return nil
	})
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
	out := outputs.NewCreateCharacterOutput()

	_ = bench.Bench("create_character_operation", func() error {
		rec := req.New()

		resp, err := rec.Put(routes.URL(config.Client().Addr(), routes.Character), req.BodyJSON(in))
		dont.Panic(err)

		if resp.Response().StatusCode != http.StatusCreated {
			out.Character = nil
			out.Success = false
			return nil
		}

		err = resp.ToJSON(out)
		dont.Panic(err)

		fmt.Println("Character", out.Character.ID, "|>", out.Character)
		return nil
	})

	return out.Character, out.Success
}

// SessionFlow used in client
type SessionFlow struct {
	Conn *rpc.Client
}

// NewSessionFlow constructor
func NewSessionFlow() client.SessionFlow {
	conn, err := rpc.Dial("tcp", config.Service().Game())
	dont.Panic(err)

	return &SessionFlow{Conn: conn}
}

// StartSession for willson flow
func (flow *SessionFlow) StartSession(in inputs.AuthAccountInput) (*outputs.StartSessionOutput, bool) {
	out := outputs.NewStartSessionOutput()

	err := bench.Bench("start_session", func() error {
		return flow.Conn.Call("SessionProxyTask.StartSession", in, out)
	})

	return out, err == nil
}

// SpawnCharacter in the world
func (flow *SessionFlow) SpawnCharacter(in inputs.SpawnCharacterInput) (*outputs.CheckSessionOutput, bool) {
	out := outputs.NewCheckSessionOutput()

	err := bench.Bench("spawn_character", func() error {
		return flow.Conn.Call("GameTask.SpawnCharacter", in, out)
	})

	return out, err == nil
}

// GameFlow used in client
type GameFlow struct {
	Conn *rpc.Client
}

// NewGameFlow constructor
func NewGameFlow() client.GameFlow {
	conn, err := rpc.Dial("tcp", config.Service().Game())
	dont.Panic(err)

	return &GameFlow{Conn: conn}
}

// MoveCharacter in game flow
func (flow *GameFlow) MoveCharacter(in inputs.MoveCharacterInput) (*outputs.CheckSessionOutput, bool) {
	out := outputs.NewCheckSessionOutput()

	err := bench.Bench("move_character", func() error {
		return flow.Conn.Call("GameTask.MoveCharacter", in, out)
	})

	return out, err == nil
}

// SendChat in game flow
func (flow *GameFlow) SendChat(in inputs.ChatInput) (*outputs.CheckSessionOutput, bool) {
	out := outputs.NewCheckSessionOutput()

	err := bench.Bench("send_chat", func() error {
		return flow.Conn.Call("GameTask.SendChat", in, out)
	})

	return out, err == nil
}

// ListenChat in game flow
func (flow *GameFlow) ListenChat(in inputs.ChatInput) (*outputs.ListenMessagesOutput, bool) {
	out := outputs.NewListenMessagesOutput()

	err := bench.Bench("listen_chat", func() error {
		return flow.Conn.Call("GameTask.ListenChat", in, out)
	})

	return out, err == nil
}

// GameLoop !
func (flow *GameFlow) GameLoop() bool {
	return true
}

func main() {
	client.UsedCreateAccountFlow = NewCreateAccountFlow()
	client.UsedCreateCharacterFlow = NewCreateCharacterFlow()
	client.UsedSessionFlow = NewSessionFlow()
	client.UsedGameFlow = NewGameFlow()

	// for {
	client.BotFlow()
	// }
}
