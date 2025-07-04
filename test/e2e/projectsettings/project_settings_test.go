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
//go:build e2e || e2eSnap || (atlas && generic)

// TODO: fix the test and add snapshots

package projectsettings

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
	projectsEntity = "projects"
	settingsEntity = "settings"
)

func TestProjectSettings(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.GenerateProject("settings")

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			settingsEntity,
			"get",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
		var settings atlasv2.GroupSettings

		for range 10 { // try again for 10 seconds
			cmd := exec.Command(cliPath,
				projectsEntity,
				settingsEntity,
				"update",
				"--disableCollectDatabaseSpecificsStatistics",
				"--projectId",
				g.ProjectID,
				"-o=json")
			cmd.Env = os.Environ()
			resp, err := internal.RunAndGetStdOut(cmd)
			require.NoError(t, err, string(resp))
			require.NoError(t, json.Unmarshal(resp, &settings))

			if !settings.GetIsCollectDatabaseSpecificsStatisticsEnabled() {
				break
			}

			time.Sleep(time.Second)
		}

		a := assert.New(t)
		a.False(settings.GetIsCollectDatabaseSpecificsStatisticsEnabled())
		a.True(settings.GetIsSchemaAdvisorEnabled())
		a.True(settings.GetIsPerformanceAdvisorEnabled())
		a.True(settings.GetIsRealtimePerformancePanelEnabled())
		a.True(settings.GetIsDataExplorerEnabled())
	})
}
