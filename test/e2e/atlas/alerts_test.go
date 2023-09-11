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
//go:build e2e || (atlas && generic)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201006/admin"
)

const (
	closed = "CLOSED"
)

func TestAlerts(t *testing.T) {
	var alertID string

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// This test should be run before all other tests to grab an alert ID for all others tests
	t.Run("List with status CLOSED", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"list",
			"--status",
			closed,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alerts atlasv2.PaginatedAlert
			err := json.Unmarshal(resp, &alerts)
			a.NoError(err)
			a.NotEmpty(alerts.Results)
			alertID = *alerts.Results[0].Id
		}
	})

	t.Run("List with status OPEN", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"list",
			"--status",
			"OPEN",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})

	t.Run("List with no status", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"list",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"describe",
			alertID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alert atlasv2.AlertViewForNdsGroup
			err := json.Unmarshal(resp, &alert)
			a.NoError(err)
			a.Equal(alertID, *alert.Id)
			a.Equal(closed, *alert.Status)
		}
	})

	t.Run("Acknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"ack",
			alertID,
			"--until",
			time.Now().Format(time.RFC3339),
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alert atlasv2.AlertViewForNdsGroup
			err := json.Unmarshal(resp, &alert)
			a.NoError(err)
			a.Equal(alertID, *alert.Id)
		}
	})

	t.Run("Acknowledge Forever", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"ack",
			alertID,
			"--forever",
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alert atlasv2.AlertViewForNdsGroup
			err := json.Unmarshal(resp, &alert)
			a.NoError(err)
			a.Equal(alertID, *alert.Id)
		}
	})

	t.Run("UnAcknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			"unack",
			alertID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alert atlasv2.AlertViewForNdsGroup
			err := json.Unmarshal(resp, &alert)
			a.NoError(err)
			a.Equal(alertID, *alert.Id)
		}
	})
}
