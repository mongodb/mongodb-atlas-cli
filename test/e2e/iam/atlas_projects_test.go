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

//go:build e2e || (iam && atlas)

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestAtlasProjects(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	projectName := fmt.Sprintf("e2e-proj-%v", n)

	var projectID string
	t.Run("Create", func(t *testing.T) {
		// This depends on a ORG_ID ENV
		cmd := exec.Command(cliPath,
			projectsEntity,
			"create",
			projectName,
			"--tag", "env=e2e",
			"--tag", "prod=false",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		require.NoError(t, err, string(resp))

		var project admin.Group
		require.NoError(t, json.Unmarshal(resp, &project))
		if project.GetName() != projectName {
			t.Errorf("got=%#v\nwant=%#v\n", project.Name, projectName)
		}
		projectID = project.GetId()
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		require.NoError(t, err, string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			"describe",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		require.NoError(t, err, string(resp))
	})

	t.Run("Tags", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			"describe",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var project admin.Group
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

	t.Run("Users", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			usersEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			"delete",
			projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Project '%s' deleted\n", projectID)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})
}
