package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/infra/ctrepositories"
	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/schweigert/mmosandbox/infra/tasks/sessiontask"
	"github.com/schweigert/mmosandbox/lib/dont"
)

func configCache() {
	ctrepositories.UseTokenRepository()
}

func configDb() {
	ctrepositories.UseAccountRepository()
	ctrepositories.UseCharacterRepository()

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
