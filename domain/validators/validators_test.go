package validators

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidatorsSuite struct {
	suite.Suite
}

func (suite *ValidatorsSuite) TestEmail() {
	suite.True(Email("marlon.henry@magrathealabs.com"))
	suite.False(Email("testing@testing@"))
}

func (suite *ValidatorsSuite) TestStringSize() {
	suite.True(StringSize("abcdefghi", 2, 20))
	suite.False(StringSize("abcdefghi", 2, 5))
	suite.False(StringSize("a", 2, 5))
}

func TestValidatorsSuite(t *testing.T) {
	suite.Run(t, new(ValidatorsSuite))
}
