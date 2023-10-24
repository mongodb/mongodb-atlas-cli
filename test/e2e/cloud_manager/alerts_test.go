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

//go:build e2e || (cloudmanager && generic)

package cloud_manager_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	closed = "CLOSED"
)

func TestAlerts(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var alertID string
	a := assert.New(t)

	t.Run("List with status CLOSED", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"list",
			"--status",
			closed,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err)
		require.NoError(t, err, string(resp))
		var alerts opsmngr.AlertsResponse
		require.NoError(t, json.Unmarshal(resp, &alerts))
		a.NotEmpty(alerts.Results)
		alertID = alerts.Results[0].ID
	})

	t.Run("List with status OPEN", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"list",
			"--status",
			"OPEN",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"describe",
			alertID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var alert opsmngr.Alert
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.Equal(alertID, alert.ID)
		a.Equal(closed, alert.Status)
	})

	t.Run("List with no status", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"list",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var alerts opsmngr.AlertsResponse
		require.NoError(t, json.Unmarshal(resp, &alerts), string(resp))
	})

	t.Run("List with status CLOSED", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"list",
			"--status",
			closed,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var alerts opsmngr.AlertsResponse
		require.NoError(t, json.Unmarshal(resp, &alerts), string(resp))
	})

	t.Run("Acknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"ack",
			alertID,
			"--until",
			time.Now().Format(time.RFC3339),
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var alert opsmngr.Alert
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.Equal(alertID, alert.ID)
	})

	t.Run("Acknowledge Forever", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"ack",
			alertID,
			"--forever",
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var alert opsmngr.Alert
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.Equal(alertID, alert.ID)
	})

	t.Run("UnAcknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"unack",
			alertID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var alert opsmngr.Alert
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.Equal(alertID, alert.ID)
	})
}
