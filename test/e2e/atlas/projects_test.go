// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestProjects(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	projectName, err := RandProjectName()
	require.NoError(t, err)

	projectID, err := createProject(projectName)
	require.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Validate tags", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			"describe",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var project atlasv2.Group
		if err := json.Unmarshal(resp, &project); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		require.Len(t, project.GetTags(), 2)

		expectedTags := map[string]string{"env": "e2e", "prod": "false"}
		for _, tag := range project.GetTags() {
			expectedValue, ok := expectedTags[tag.GetKey()]
			if !ok {
				t.Errorf("unexpected tag key %s in tags: %v, expected tags: %v\n", tag.GetKey(), project.Tags, expectedTags)
			}

			require.Equal(t, expectedValue, tag.GetValue())
		}
	})
}
