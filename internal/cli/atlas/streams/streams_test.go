package streams

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/test"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		1,
		[]string{},
	)
}
