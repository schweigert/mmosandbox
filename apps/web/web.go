package main

import (
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/infra/controllers"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/web"
)

var app *web.App
var usedControllers = []web.Controller{
	controllers.NewRootController(),
}

func main() {
	app = web.NewApp(config.Web().Addr())
	app.Setup(usedControllers...)
	dont.Panic(app.Run())
}
