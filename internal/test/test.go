package test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// CmdValidator helps validate a cobra.Command, verifying the number of sub commands
// and the flags that are being defined for it
func CmdValidator(t *testing.T, subject *cobra.Command, nSubCommands int, flags []string) {
	t.Helper()
	a := assert.New(t)
	a.Len(subject.Commands(), nSubCommands)
	if len(flags) == 0 {
		a.False(subject.HasAvailableFlags())
		return
	}
	a.True(subject.HasAvailableFlags())
	for _, f := range flags {
		a.NotNil(subject.Flags().Lookup(f))
	}
}
