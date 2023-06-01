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
	"strconv"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/admin"
)

func TestAlertConfig(t *testing.T) {
	var alertID string

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"create",
			"--event",
			eventTypeName,
			"--enabled=true",
			"--notificationType",
			group,
			"--notificationIntervalMin",
			strconv.Itoa(intervalMin),
			"--notificationDelayMin",
			strconv.Itoa(delayMin),
			"--notificationSmsEnabled=false",
			"--notificationEmailEnabled=true",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alert admin.AlertConfigViewForNdsGroup
			if err := json.Unmarshal(resp, &alert); a.NoError(err) {
				a.Equal(eventTypeName, alert.GetEventTypeName())
				a.NotEmpty(alert.Notifications)
				a.Equal(delayMin, alert.Notifications[0].GroupNotification.GetDelayMin())
				a.Equal(group, alert.Notifications[0].GroupNotification.GetTypeName())
				a.Equal(intervalMin, alert.Notifications[0].GroupNotification.GetIntervalMin())
				a.False(alert.Notifications[0].GroupNotification.GetSmsEnabled())
				alertID = alert.GetId()
			}
		}
	})
	if alertID == "" {
		assert.FailNow(t, "Failed to create alert setting")
	}

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"update",
			alertID,
			"--event",
			eventTypeName,
			"--notificationType",
			group,
			"--notificationIntervalMin",
			strconv.Itoa(intervalMin),
			"--notificationDelayMin",
			strconv.Itoa(delayMin),
			"--notificationSmsEnabled=true",
			"--notificationEmailEnabled=true",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var alert admin.AlertConfigViewForNdsGroup
			if err := json.Unmarshal(resp, &alert); a.NoError(err) {
				a.False(alert.GetEnabled())
				a.NotEmpty(alert.Notifications)
				a.True(alert.Notifications[0].GroupNotification.GetSmsEnabled())
				a.True(alert.Notifications[0].GroupNotification.GetEmailEnabled())
			}
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, alertsEntity, configEntity, "delete", alertID, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("List Matcher Fields", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"fields",
			"type",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err)

		var fields []string
		if err := json.Unmarshal(resp, &fields); err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := []string{
			"TYPE_NAME",
			"HOSTNAME",
			"PORT",
			"HOSTNAME_AND_PORT",
			"REPLICA_SET_NAME",
			"SHARD_NAME",
			"CLUSTER_NAME",
			"APPLICATION_ID",
		}
		assert.ElementsMatch(t, fields, expected)
	})
}
