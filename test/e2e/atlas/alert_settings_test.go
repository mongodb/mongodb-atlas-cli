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
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var alert admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &alert))
		a := assert.New(t)
		a.Equal(eventTypeName, alert.GetEventTypeName())
		a.NotEmpty(alert.GetNotifications())
		a.Equal(delayMin, alert.GetNotifications()[0].GetDelayMin())
		a.Equal(group, alert.GetNotifications()[0].GetTypeName())
		a.Equal(intervalMin, alert.GetNotifications()[0].GetIntervalMin())
		a.False(alert.GetNotifications()[0].GetSmsEnabled())
		alertID = alert.GetId()
	})
	if alertID == "" {
		assert.FailNow(t, "Failed to create alert setting")
	}

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"get",
			alertID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var config admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &config))
		assert.Equal(t, alertID, config.GetId())
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var config admin.PaginatedAlertConfig
		require.NoError(t, json.Unmarshal(resp, &config))
		assert.NotEmpty(t, config.Results)
	})

	t.Run("List Compact", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var config []admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &config))
		assert.NotEmpty(t, config)
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
		resp, err := e2e.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var alert admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.False(alert.GetEnabled())
		a.NotEmpty(alert.GetNotifications())
		a.True(alert.GetNotifications()[0].GetSmsEnabled())
		a.True(alert.GetNotifications()[0].GetEmailEnabled())
	})

	t.Run("Update Setting using file input", func(t *testing.T) {
		n, err := e2e.RandInt(1000)
		require.NoError(t, err)
		fileName := fmt.Sprintf("%d_alerts.json", n.Int64())
		fileContent := fmt.Sprintf(`{
			"eventTypeName": %q,
			"id": "%s",
			"enabled": false,
			"notifications": [
			  {
				"typeName": "%s",
				"intervalMin": %d,
				"delayMin": %d,
				"emailEnabled": true,
				"smsEnabled": true
			  }
			]
		}`, eventTypeName, alertID,
			group, intervalMin, delayMin)

		require.NoError(t, os.WriteFile(fileName, []byte(fileContent), 0600))

		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"update",
			alertID,
			"--file", fileName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var alert admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.False(alert.GetEnabled())
		a.NotEmpty(alert.GetNotifications())
		a.True(alert.GetNotifications()[0].GetSmsEnabled())
		a.True(alert.GetNotifications()[0].GetEmailEnabled())
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, alertsEntity, configEntity, "delete", alertID, "--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

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
