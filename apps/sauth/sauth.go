package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/merepositories"
	"github.com/schweigert/mmosandbox/infra/tasks/sessiontask"
	"github.com/schweigert/mmosandbox/lib/dont"
)

func configCache() {
	merepositories.UseTokenRepository()
}

func configDb() {
	dbrepositories.UseAccountRepository()
	dbrepositories.UseCharacterRepository()

	db.Migrate()
}

func configRPC() {
	dont.Panic(rpc.Register(sessiontask.New()))
}

func startRPC() {
	listener, err := net.Listen("tcp", config.RPC().Addr())
	dont.Panic(err)

	configRPC()

	fmt.Println("RPC Starting at", config.RPC().Addr(), "...")
	defer fmt.Println("RPC finished")
	rpc.Accept(listener)
}

func main() {
	configCache()
	configDb()
	startRPC()
}
