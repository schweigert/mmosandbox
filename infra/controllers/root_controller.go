package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/infra/routes"
)

// RootController handle /
type RootController struct{}

// NewRootController constructor
func NewRootController() *RootController {
	return &RootController{}
}

// Routes for web lib
func (controller *RootController) Routes(engine *gin.Engine) {
	engine.GET(routes.Root, controller.Index)
}

// Index GET /
func (controller *RootController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
