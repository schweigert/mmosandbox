package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/infra/routes"
)

// AccountController handle /
type AccountController struct{}

// NewAccountController constructor
func NewAccountController() *AccountController {
	return &AccountController{}
}

// Routes for web lib
func (controller *AccountController) Routes(engine *gin.Engine) {
	engine.PUT(routes.Account, controller.Create)
}

// Create PUT /account/
func (controller *AccountController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
