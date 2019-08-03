package dont

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorSuite struct {
	suite.Suite
}

func (suite *ErrorSuite) TestPanic() {
	panics := errors.New("PANIC")

	suite.Panics(func() { Panic(panics) })
	suite.NotPanics(func() { Panic(nil) })
}

func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}
