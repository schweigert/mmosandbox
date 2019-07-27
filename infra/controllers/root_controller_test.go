package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmosandbox/infra/routes"

	"github.com/stretchr/testify/suite"
)

type RootControllerSuite struct {
	suite.Suite
}

func (suite *RootControllerSuite) TestNewRootController() {
	suite.NotNil(NewRootController())
}

func (suite *RootControllerSuite) TestRoutes() {
	engine := gin.Default()
	NewRootController().Routes(engine)

	suite.Equal(1, len(engine.Routes()))
}

func (suite *RootControllerSuite) TestIndex() {
	req, err := http.NewRequest(http.MethodGet, routes.Root, strings.NewReader(""))

	suite.NoError(err)

	rec := httptest.NewRecorder()
	engine := gin.Default()
	NewRootController().Routes(engine)
	engine.ServeHTTP(rec, req)

	suite.Equal(http.StatusOK, rec.Code)
}

func TestRootControllerSuite(t *testing.T) {
	suite.Run(t, new(RootControllerSuite))
}
