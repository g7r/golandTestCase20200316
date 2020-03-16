package golandTestCase20200316

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type S struct {
	suite.Suite
}

func TestS(t *testing.T) {
	suite.Run(t, &S{})
}

func (s *S) TestFoo() {}
func (s *S) TestBar() {}