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

//go:build e2e || e2eSnap || (atlas && generic)

package alertsettings

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312004/admin"
)

const (
	alertsEntity   = "alerts"
	settingsEntity = "settings"

	// AlertConfig constants.
	group         = "GROUP"
	eventTypeName = "NO_PRIMARY"
	intervalMin   = 5
	delayMin      = 0
)

func TestAlertConfig(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	var alertID string

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			settingsEntity,
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
		resp, err := internal.RunAndGetStdOut(cmd)
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

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			settingsEntity,
			"get",
			alertID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var config admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &config))
		assert.Equal(t, alertID, config.GetId())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			settingsEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var config admin.PaginatedAlertConfig
		require.NoError(t, json.Unmarshal(resp, &config))
		assert.NotEmpty(t, config.Results)
	})

	g.Run("List Compact", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			settingsEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var config []admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &config))
		assert.NotEmpty(t, config)
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			settingsEntity,
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
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var alert admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.False(alert.GetEnabled())
		a.NotEmpty(alert.GetNotifications())
		a.True(alert.GetNotifications()[0].GetSmsEnabled())
		a.True(alert.GetNotifications()[0].GetEmailEnabled())
	})

	g.Run("Update Setting using file input", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.MemoryRand("rand", 1000)
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
			settingsEntity,
			"update",
			alertID,
			"--file", fileName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var alert admin.GroupAlertsConfig
		require.NoError(t, json.Unmarshal(resp, &alert))
		a.False(alert.GetEnabled())
		a.NotEmpty(alert.GetNotifications())
		a.True(alert.GetNotifications()[0].GetSmsEnabled())
		a.True(alert.GetNotifications()[0].GetEmailEnabled())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath, alertsEntity, settingsEntity, "delete", alertID, "--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("List Matcher Fields", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			alertsEntity,
			settingsEntity,
			"fields",
			"type",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
