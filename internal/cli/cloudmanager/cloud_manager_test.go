package cloudmanager

import (
	"testing"

	"github.com/mongodb/mongocli/internal/cli"
)

func TestBuilder(t *testing.T) {
	cli.CmdValidator(
		t,
		Builder(),
		16,
		[]string{},
	)
}
