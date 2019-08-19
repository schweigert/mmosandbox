package main

import (
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/infra/controllers"
	"github.com/schweigert/mmosandbox/infra/db"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/middlewares"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/web"
)

func configDb() {
	dbrepositories.UseAccountRepository()
	dbrepositories.UseCharacterRepository()

	db.Migrate()
}

func startWeb() {
	controllers := []web.Controller{
		controllers.NewRootController(),
		controllers.NewAccountController(),
		controllers.NewCharacterController(),
	}

	app := web.NewApp(config.Web().Addr())
	app.Server.Use(middlewares.Bip())

	app.Setup(controllers...)

	dont.Panic(app.Run())
}

func main() {
	configDb()
	startWeb()
}
