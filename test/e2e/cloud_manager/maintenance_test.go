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

//go:build e2e || (generic && (cloudmanager || om60))

package cloud_manager_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestMaintenanceWindows(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)

	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	projectName := fmt.Sprintf("e2e-maintenance-proj-%v", n)
	projectID, err := e2e.CreateProject(projectName)
	require.NoError(t, err)
	t.Cleanup(func() {
		e2e.DeleteProjectWithRetry(t, projectID)
	})

	startDate := time.Now().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 0, 1).Format(time.RFC3339)
	var maintenanceWindowID string

	t.Run("create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			maintenanceEntity,
			"create",
			"--startDate",
			startDate,
			"--endDate",
			endDate,
			"--alertType",
			"REPLICA_SET",
			"-o",
			"json",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var maintenanceWindow opsmngr.MaintenanceWindow
		require.NoError(t, json.Unmarshal(resp, &maintenanceWindow))
		a.Equal(startDate, maintenanceWindow.StartDate)
		a.Equal(endDate, maintenanceWindow.EndDate)
		maintenanceWindowID = maintenanceWindow.ID
	})

	t.Run("describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			maintenanceEntity,
			"describe",
			maintenanceWindowID,
			"-o",
			"json",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var maintenanceWindow opsmngr.MaintenanceWindow
		require.NoError(t, json.Unmarshal(resp, &maintenanceWindow))
		assert.Equal(t, maintenanceWindowID, maintenanceWindow.ID)
	})

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			maintenanceEntity,
			"ls",
			"-o",
			"json",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var maintenanceWindows opsmngr.MaintenanceWindows
		require.NoError(t, json.Unmarshal(resp, &maintenanceWindows))
		a.NotEmpty(maintenanceWindows.Results)
	})

	t.Run("update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			maintenanceEntity,
			"update",
			maintenanceWindowID,
			"--alertType",
			"CLUSTER",
			"-o",
			"json",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var maintenanceWindow opsmngr.MaintenanceWindow
		require.NoError(t, json.Unmarshal(resp, &maintenanceWindow))
		a.Contains(maintenanceWindow.AlertTypeNames, "CLUSTER")
	})

	t.Run("delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			maintenanceEntity,
			"delete",
			maintenanceWindowID,
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)

		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Maintenance window '%s' deleted\n", maintenanceWindowID)
		a.Equal(expected, string(resp))
	})
}
