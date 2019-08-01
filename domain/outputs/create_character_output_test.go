package outputs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateCharacterOutputSuite struct {
	suite.Suite
}

func (suite *CreateCharacterOutputSuite) TestNewCreateCharacterOutput() {
	out := NewCreateCharacterOutput()
	suite.NotNil(out)
	suite.Nil(out.Account)
	suite.Nil(out.Character)
	suite.False(out.Success)
}

func (suite *CreateCharacterOutputSuite) TestHTTPCode() {
	out := NewCreateCharacterOutput()

	suite.False(out.Auth)
	suite.Equal(http.StatusBadRequest, out.HTTPCode())

	out.Auth = true
	suite.False(out.Success)
	suite.Equal(http.StatusBadRequest, out.HTTPCode())

	out.Success = true
	suite.Equal(http.StatusCreated, out.HTTPCode())
}

func TestCreateCharacterOutputSuite(t *testing.T) {
	suite.Run(t, new(CreateCharacterOutputSuite))
}
