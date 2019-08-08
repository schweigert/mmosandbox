package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/tasks"
	"github.com/schweigert/mmosandbox/lib/dont"
)

func configCache() {}

func configDb() {
	dbrepositories.UseCharacterRepository()
}

func configRPC() {
	dont.Panic(rpc.Register(tasks.NewSessionProxyTask()))
	dont.Panic(rpc.Register(tasks.NewGameTask()))
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
