package dbweb

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/domain/repositories"
	"github.com/schweigert/mmosandbox/infra/dbrepositories"
	"github.com/schweigert/mmosandbox/infra/routes"
)

// AccountRepositoryHandler struct
type AccountRepositoryHandler struct {
	repo repositories.AccountRepository
}

// NewAccountRepositoryHandler constructor
func NewAccountRepositoryHandler() *AccountRepositoryHandler {
	return &AccountRepositoryHandler{
		repo: dbrepositories.NewAccountRepository(),
	}
}

// Routes for web lib
func (controller *AccountRepositoryHandler) Routes(engine *gin.Engine) {
	engine.POST(routes.AccountRepositoryUsernameHasTaken, controller.UsernameHasTaken)
}

// UsernameHasTaken handler
func (controller *AccountRepositoryHandler) UsernameHasTaken(c *gin.Context) {
	value := ""
	if c.BindJSON(&value) == nil {
		c.JSON(http.StatusOK, controller.repo.UsernameHasTaken(value))
	}
}
