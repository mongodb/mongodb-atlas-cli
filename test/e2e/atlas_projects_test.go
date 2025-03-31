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

//go:build e2e || e2eSnap || (iam && atlas)

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestAtlasProjects(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n := g.memoryRand("rand", 1000)
	projectName := fmt.Sprintf("e2e-proj-%d", n)

	var projectID string
	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		// This depends on a ORG_ID ENV
		cmd := exec.Command(cliPath,
			projectsEntity,
			"create",
			projectName,
			"--tag", "env=e2e",
			"--tag", "prod=false",
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var project admin.Group
		require.NoError(t, json.Unmarshal(resp, &project))
		if project.GetName() != projectName {
			t.Errorf("got=%#v\nwant=%#v\n", project.Name, projectName)
		}
		projectID = project.GetId()
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			"ls",
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			"describe",
			projectID,
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	g.Run("Tags", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			"describe",
			projectID,
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var project admin.Group
		if err := json.Unmarshal(resp, &project); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		a := assert.New(t)
		expectedTags := []admin.ResourceTag{{Key: "env", Value: "e2e"}, {Key: "prod", Value: "false"}}
		gotTags := project.GetTags()
		slices.SortFunc(gotTags, func(a, b admin.ResourceTag) int {
			if a.Key < b.Key {
				return -1
			}

			if a.Key > b.Key {
				return 1
			}

			return 0
		})
		a.ElementsMatch(expectedTags, gotTags)
	})

	updatedProjectName := projectName + "-updated"
	updateTests := []struct {
		name                string
		filename            string
		SetProjectName      *string
		expectedProjectName string
		expectedProjectTags []admin.ResourceTag
	}{
		{
			"setNameAndTags",
			"data/update_project_name_and_tags.json",
			&updatedProjectName,
			updatedProjectName,
			[]admin.ResourceTag{{Key: "app", Value: "cli"}, {Key: "env", Value: "e2e"}},
		},
		{
			"resetTags",
			"data/update_project_reset_tags.json",
			nil,
			updatedProjectName,
			[]admin.ResourceTag{},
		},
	}
	for _, tt := range updateTests {
		g.Run("Update_"+tt.name, func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
			filename := fmt.Sprintf("update_project_%s.json", tt.name)
			testTmpl, err := os.ReadFile(tt.filename)
			tpl := template.Must(template.New("").Parse(string(testTmpl)))

			require.NoError(t, err)
			file, err := os.Create(filename)
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, os.Remove(filename))
			})

			require.NoError(t, tpl.Execute(file, tt))

			cmd := exec.Command(cliPath,
				projectsEntity,
				"update",
				projectID,
				"--file",
				filename,
				"-o=json")
			cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
			resp, err := RunAndGetStdOut(cmd)
			require.NoError(t, err, string(resp))

			cmd = exec.Command(cliPath,
				projectsEntity,
				"describe",
				projectID,
				"-o=json")
			cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
			resp, err = RunAndGetStdOut(cmd)
			require.NoError(t, err, string(resp))
			var project admin.Group
			if err := json.Unmarshal(resp, &project); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			a := assert.New(t)
			a.Equal(tt.expectedProjectName, project.Name)
			gotTags := project.GetTags()
			slices.SortFunc(gotTags, func(a, b admin.ResourceTag) int {
				if a.Key < b.Key {
					return -1
				}

				if a.Key > b.Key {
					return 1
				}

				return 0
			})
			a.ElementsMatch(tt.expectedProjectTags, gotTags)
		})
	}

	g.Run("Users", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			usersEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			"delete",
			projectID,
			"--force")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Project '%s' deleted\n", projectID)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})
}
