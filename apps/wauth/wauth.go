package main

import (
	"net"
	"net/rpc"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/infra/tasks"
	"github.com/schweigert/mmosandbox/lib/dont"
)

func main() {
	listener, err := net.Listen("tcp", config.RPC().Addr())
	dont.Panic(err)

	dont.Panic(rpc.Register(tasks.NewSessionTask()))

	rpc.Accept(listener)
}
