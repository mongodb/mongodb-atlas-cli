// Copyright 2026 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pluginfirstclass

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type telemetryEvent struct {
	Properties map[string]any `json:"properties"`
}

// TestPluginLocalTelemetry verifies that `atlas local list --output json` reports
// the full subcommand path and flags, both when the plugin is installed on the
// fly (first-class plugin command) and on subsequent runs (installed plugin
// command). Flag parsing is disabled in both cases, so the command and flags
// must be recovered from the forwarded arguments.
func TestPluginLocalTelemetry(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	_ = internal.TempConfigFolder(t)
	t.Setenv("DO_NOT_TRACK", "")
	t.Setenv("MONGODB_ATLAS_TELEMETRY_ENABLED", "true")

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	removeFirstClassPlugin(t, "atlas-local-plugin", cliPath)

	t.Run("install path", func(t *testing.T) {
		runLocalList(t, cliPath)
		assertLocalListTelemetry(t)
	})

	t.Run("installed path", func(t *testing.T) {
		runLocalList(t, cliPath)
		assertLocalListTelemetry(t)
	})
}

func runLocalList(t *testing.T, cliPath string) {
	t.Helper()

	// Clear any cached event so the assertions only see this run.
	_ = os.Remove(telemetryCachePath(t))

	cmd := exec.Command(cliPath, "local", "list", "--output", "json")
	// Report the cloud service but point the API at an unreachable address so the
	// telemetry event is cached locally instead of being sent to the backend.
	cmd.Env = append(os.Environ(),
		"MONGODB_ATLAS_SERVICE=cloud",
		"MONGODB_ATLAS_OPS_MANAGER_URL=http://127.0.0.1:1",
	)
	resp, _ := cmd.CombinedOutput()
	t.Logf("atlas local list output: %s", resp)
}

func assertLocalListTelemetry(t *testing.T) {
	t.Helper()

	events := cachedTelemetryEvents(t)

	var event *telemetryEvent
	for i := range events {
		if events[i].Properties["command"] == "atlas-local-list" {
			event = &events[i]
			break
		}
	}

	require.NotNil(t, event, "expected a telemetry event reporting the atlas-local-list command")
	assert.Contains(t, event.Properties["flags"], "output")
}

func telemetryCachePath(t *testing.T) string {
	t.Helper()

	cacheDir, err := os.UserCacheDir()
	require.NoError(t, err)

	return filepath.Join(cacheDir, "atlascli", "telemetry")
}

func cachedTelemetryEvents(t *testing.T) []telemetryEvent {
	t.Helper()

	data, err := os.ReadFile(telemetryCachePath(t))
	require.NoError(t, err, "telemetry cache file should exist after a failed send")

	var events []telemetryEvent
	decoder := json.NewDecoder(bytes.NewReader(data))
	for decoder.More() {
		var event telemetryEvent
		require.NoError(t, decoder.Decode(&event))
		events = append(events, event)
	}

	return events
}
