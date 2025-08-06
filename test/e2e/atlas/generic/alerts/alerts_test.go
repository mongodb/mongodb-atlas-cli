// Copyright 2020 MongoDB Inc
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

package alerts

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	alertsEntity = "alerts"
)

const (
	closed = "CLOSED"
)

func TestAlerts(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)
	var alertID string
	// This test should be run before all other tests to grab an alert ID for all others tests
	g.Run("List with status CLOSED", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			"list",
			"--status",
			closed,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var alerts atlasv2.PaginatedAlert
		require.NoError(t, json.Unmarshal(resp, &alerts))
		if len(alerts.GetResults()) != 0 {
			alertID = alerts.GetResults()[0].GetId()
		}
	})

	g.Run("List with status OPEN", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			"list",
			"--status",
			"OPEN",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("List with no status", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			"list",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		if alertID == "" {
			t.Skip("no alert found")
		}
		cmd := exec.Command(cliPath,
			alertsEntity,
			"describe",
			alertID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var alert atlasv2.AlertViewForNdsGroup
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.Equal(alertID, alert.GetId())
		a.Equal(closed, alert.GetStatus())
	})

	g.Run("Acknowledge", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		if alertID == "" {
			t.Skip("no alert found")
		}
		cmd := exec.Command(cliPath,
			alertsEntity,
			"ack",
			alertID,
			"--until",
			time.Now().Format(time.RFC3339),
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var alert atlasv2.AlertViewForNdsGroup
		require.NoError(t, json.Unmarshal(resp, &alert))
		assert.Equal(t, alertID, alert.GetId())
	})

	g.Run("Acknowledge Forever", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		if alertID == "" {
			t.Skip("no alert found")
		}
		cmd := exec.Command(cliPath,
			alertsEntity,
			"ack",
			alertID,
			"--forever",
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var alert atlasv2.AlertViewForNdsGroup
		require.NoError(t, json.Unmarshal(resp, &alert))
		assert.Equal(t, alertID, alert.GetId())
	})

	g.Run("UnAcknowledge", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		if alertID == "" {
			t.Skip("no alert found")
		}
		cmd := exec.Command(cliPath,
			alertsEntity,
			"unack",
			alertID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var alert atlasv2.AlertViewForNdsGroup
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.Equal(alertID, alert.GetId())
	})
}
