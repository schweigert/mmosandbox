package security

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SecuritySuite struct {
	suite.Suite
}

func (suite *SecuritySuite) TestSum() {
	suite.Equal(
		"40bbd2844b115b29953c28e36e74dd78a693383ff021bea1b96270678cc25e8e",
		Sum("working hard"),
	)
}

func TestSecuritySuite(t *testing.T) {
	suite.Run(t, new(SecuritySuite))
}
