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
// +build e2e atlas,generic

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestMaintenanceWindows(t *testing.T) {
	n, err := e2e.RandInt(255)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projectName := fmt.Sprintf("e2e-maintenance-proj-%v", n)
	projectID, err := createProject(projectName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			maintenanceEntity,
			"update",
			"--dayOfWeek",
			"1",
			"--hourOfDay",
			"1",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			expected := "Maintenance window updated.\n"
			a.Equal(expected, string(resp))
		}
	})

	t.Run("describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			maintenanceEntity,
			"describe",
			"-o",
			"json",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var maintenanceWindow mongodbatlas.MaintenanceWindow
		if err := json.Unmarshal(resp, &maintenanceWindow); a.NoError(err) {
			a.Equal(1, maintenanceWindow.DayOfWeek)
			a.Equal(1, *maintenanceWindow.HourOfDay)
		}
	})

	t.Run("clear", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			maintenanceEntity,
			"clear",
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			expected := "Maintenance window removed.\n"
			a.Equal(expected, string(resp))
		}
	})
}
