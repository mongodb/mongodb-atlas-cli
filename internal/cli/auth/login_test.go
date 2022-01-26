package auth

import (
	"testing"

	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/test"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		1,
		[]string{},
	)
}

func TestLoginBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LoginBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Page, flag.Limit},
	)
}
