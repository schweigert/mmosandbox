package web

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
}

func (suite *AppSuite) Routes(engine *gin.Engine) {
	suite.NotNil(engine)
}

func (suite *AppSuite) TestNewApp() {
	app := NewApp("localhost:3000")

	suite.NotEmpty(app)
	suite.NotNil(app.Server)
	suite.Equal("localhost:3000", app.Addr)
}

func (suite *AppSuite) TestSetup() {
	app := NewApp("localhost:3000")

	suite.NotPanics(func() {
		app.Setup(suite, suite, suite)
	})
}

func (suite *AppSuite) TestRun() {
	app := NewApp("localhost:-1")
	suite.Error(app.Run())
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}
