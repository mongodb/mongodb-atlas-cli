// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/templatewriter"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CmdValidator helps validate a cobra.Command, verifying the number of sub commands
// and the flags that are being defined for it.
func CmdValidator(t *testing.T, subject *cobra.Command, nSubCommands int, flags []string) {
	t.Helper()
	a := assert.New(t)
	if len(subject.Commands()) != nSubCommands {
		t.Errorf("Sub command count mismatch. Expected %d, got %d. Check the CmdValidator for your command.\n", nSubCommands, len(subject.Commands()))
	}
	if len(flags) == 0 {
		a.False(subject.HasAvailableFlags(), "expected command to not have flags but it does")
		return
	}
	a.True(subject.HasAvailableFlags(), "expected command to have flag but has none")
	for _, f := range flags {
		a.NotNilf(subject.Flags().Lookup(f), "command has no flag: %s", f)
	}
}

// VerifyOutputTemplate validates that the given template string is valid.
func VerifyOutputTemplate(t *testing.T, tmpl string, typeValue any) {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, templatewriter.Print(&buf, tmpl, typeValue))
}
