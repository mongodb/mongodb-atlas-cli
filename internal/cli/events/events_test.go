package events

import (
	"testing"

	"github.com/mongodb/mongocli/internal/test"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		3,
		[]string{},
	)
}
