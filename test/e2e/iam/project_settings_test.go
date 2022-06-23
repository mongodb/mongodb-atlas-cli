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
//go:build e2e || iam

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/require"
	exec "golang.org/x/sys/execabs"
)

func TestProjectSettings(t *testing.T) {
	cliPath, err := e2e.Bin()
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
			iamEntity,
			projectsEntity,
			settingsEntity,
			"get",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var settings map[string]interface{}
		if err := json.Unmarshal(resp, &settings); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := settings["isCollectDatabaseSpecificsStatisticsEnabled"]; !ok {
			t.Errorf("expected %v, to have key %s\n", settings, "isCollectDatabaseSpecificsStatisticsEnabled")
		}
		if _, ok := settings["isDataExplorerEnabled"]; !ok {
			t.Errorf("expected %v, to have key %s\n", settings, "isDataExplorerEnabled")
		}
		if _, ok := settings["isPerformanceAdvisorEnabled"]; !ok {
			t.Errorf("expected %v, to have key %s\n", settings, "isPerformanceAdvisorEnabled")
		}
		if _, ok := settings["isRealtimePerformancePanelEnabled"]; !ok {
			t.Errorf("expected %v, to have key %s\n", settings, "isRealtimePerformancePanelEnabled")
		}
		if _, ok := settings["isSchemaAdvisorEnabled"]; !ok {
			t.Errorf("expected %v, to have key %s\n", settings, "isSchemaAdvisorEnabled")
		}
	})
}
