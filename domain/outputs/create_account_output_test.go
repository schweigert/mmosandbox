package outputs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateAccountOutputSuite struct {
	suite.Suite
}

func (suite *CreateAccountOutputSuite) TestNewCreateAccountOutput() {
	output := NewCreateAccountOutput()
	suite.NotNil(output)
	suite.Nil(output.Account)
	suite.False(output.Success)
}

func (suite *CreateAccountOutputSuite) TestValid() {
	output := NewCreateAccountOutput()
	output.Success = true

	suite.Equal(http.StatusCreated, output.HTTPCode())

	output.Success = false
	suite.Equal(http.StatusBadRequest, output.HTTPCode())
}

func TestCreateAccountOutputSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountOutputSuite))
}
