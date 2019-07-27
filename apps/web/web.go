package main

import (
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/web"
)

func main() {
	dont.Panic(web.NewApp(config.Web().Addr()).Run())
}
