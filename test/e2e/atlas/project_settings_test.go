// Copyright 2022 MongoDB Inc
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
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestProjectSettings(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	n, err := e2e.RandInt(1000)
	require.NoError(t, err)

	projectName := fmt.Sprintf("e2e-proj-%v", n)
	projectID, err := createProject(projectName)
	require.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			settingsEntity,
			"get",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var settings map[string]any
		if err := json.Unmarshal(resp, &settings); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		keys := [5]string{"isCollectDatabaseSpecificsStatisticsEnabled", "isDataExplorerEnabled", "isPerformanceAdvisorEnabled", "isRealtimePerformancePanelEnabled", "isSchemaAdvisorEnabled"}
		for _, k := range keys {
			if _, ok := settings[k]; !ok {
				t.Errorf("expected %v, to have key %s\n", settings, k)
			}
		}
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			settingsEntity,
			"update",
			"--disableCollectDatabaseSpecificsStatistics",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var settings atlasv2.GroupSettings
		require.NoError(t, json.Unmarshal(resp, &settings))
		a := assert.New(t)
		a.False(settings.GetIsCollectDatabaseSpecificsStatisticsEnabled())
		a.True(settings.GetIsSchemaAdvisorEnabled())
		a.True(settings.GetIsPerformanceAdvisorEnabled())
		a.True(settings.GetIsRealtimePerformancePanelEnabled())
		a.True(settings.GetIsDataExplorerEnabled())
	})
}
