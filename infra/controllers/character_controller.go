package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/domain"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/infra/routes"
)

// CharacterController handle /character
type CharacterController struct{}

// NewCharacterController constructor
func NewCharacterController() *CharacterController {
	return &CharacterController{}
}

// Routes for web lib
func (controller *CharacterController) Routes(engine *gin.Engine) {
	engine.PUT(routes.Character, controller.Create)
}

// Create PUT /character/
func (controller *CharacterController) Create(c *gin.Context) {
	input := inputs.NewCreateCharacterInput()
	if c.BindJSON(input) == nil {
		out := domain.NewAccountRules().CreateCharacter(input)
		c.JSON(out.HTTPCode(), out)
	} else {
		c.JSON(http.StatusInternalServerError, input)
	}
}
