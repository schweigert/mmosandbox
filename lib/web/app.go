package web

import "github.com/gin-gonic/gin"

// App struct
type App struct {
	Server *gin.Engine
	Addr   string
}

// NewApp constructor
func NewApp(addr string) *App {
	return &App{Server: gin.Default(), Addr: addr}
}

// Setup many controllers
func (app *App) Setup(controllers ...Controller) {
	for _, controller := range controllers {
		controller.Routes(app.Server)
	}
}

// Run this web app
func (app *App) Run() error {
	return app.Server.Run(app.Addr)
}
