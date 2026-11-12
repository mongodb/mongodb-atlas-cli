// Copyright 2026 MongoDB Inc
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

package plugin

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestTelemetryForForwardedArgs(t *testing.T) {
	rootCmd := &cobra.Command{Use: "atlas"}
	localCmd := &cobra.Command{Use: "local"}
	rootCmd.AddCommand(localCmd)

	tests := []struct {
		name            string
		args            []string
		expectedCommand string
		expectedFlags   []string
	}{
		{
			name:            "subcommand and value flag",
			args:            []string{"list", "--output", "json"},
			expectedCommand: "atlas-local-list",
			expectedFlags:   []string{"output"},
		},
		{
			name:            "subcommand only",
			args:            []string{"list"},
			expectedCommand: "atlas-local-list",
			expectedFlags:   nil,
		},
		{
			name:            "no args",
			args:            nil,
			expectedCommand: "atlas-local",
			expectedFlags:   nil,
		},
		{
			name:            "nested subcommands with equals flag and boolean flag",
			args:            []string{"deployments", "list", "--output=json", "--force"},
			expectedCommand: "atlas-local-deployments-list",
			expectedFlags:   []string{"output", "force"},
		},
		{
			// Subcommands are assumed to precede flags, so a subcommand after a
			// flag is not recovered.
			name:            "subcommand after flag",
			args:            []string{"--output", "json", "list"},
			expectedCommand: "atlas-local",
			expectedFlags:   []string{"output"},
		},
		{
			name:            "short flag",
			args:            []string{"list", "-o", "json"},
			expectedCommand: "atlas-local-list",
			expectedFlags:   []string{"o"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := telemetry.Event{Properties: map[string]any{}}
			TelemetryForForwardedArgs(localCmd, tt.args)(event)

			assert.Equal(t, tt.expectedCommand, event.Properties["command"])
			if tt.expectedFlags == nil {
				assert.NotContains(t, event.Properties, "flags")
			} else {
				assert.Equal(t, tt.expectedFlags, event.Properties["flags"])
			}
		})
	}
}
