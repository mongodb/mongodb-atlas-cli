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
//go:build e2e || (atlas && processes)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

func TestProcesses(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProjectAndCluster("processes")

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	var processes *atlasv2.PaginatedHostViewAtlas

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			processesEntity,
			"list",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if err := json.Unmarshal(resp, &processes); assert.NoError(t, err) {
			require.NotEmpty(t, processes.Results)
		}
	})

	t.Run("list compact", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			processesEntity,
			"list",
			"-c",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var hostViewsCompact []atlasv2.ApiHostViewAtlas

		if err := json.Unmarshal(resp, &hostViewsCompact); assert.NoError(t, err) {
			require.NotEmpty(t, hostViewsCompact)
		}
	})

	t.Run("describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			processesEntity,
			"describe",
			processes.Results[0].GetId(),
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var process *atlasv2.ApiHostViewAtlas
		if err := json.Unmarshal(resp, &process); assert.NoError(t, err) {
			assert.Equal(t, *process.Id, *processes.Results[0].Id)
		}
	})
}
