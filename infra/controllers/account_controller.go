package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/infra/routes"
)

// AccountController handle /account
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
	input := inputs.NewCreateAccountInput()
	if c.BindJSON(input) == nil {
		out := domain.NewSessionRules().CreateAccount(input)
		c.JSON(out.HTTPCode(), out)
	} else {
		c.JSON(http.StatusInternalServerError, input)
	}
}
