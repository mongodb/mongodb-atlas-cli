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
// +build e2e

package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestAtlasIAMProjects(t *testing.T) {
	cliPath, err := filepath.Abs("../bin/mcli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	iamEntity := "iam"
	projectEntity := "projects"
	u, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	projectName := "e2e-test-" + u.String()

	var projectID string
	t.Run("Create", func(t *testing.T) {
		// This depends on a ORG_ID ENV
		cmd := exec.Command(cliPath, iamEntity, projectEntity, "create", projectName)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		project := mongodbatlas.Project{}
		err = json.Unmarshal(resp, &project)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if project.Name != projectName {
			t.Errorf("got=%#v\nwant=%#v\n", project.Name, projectName)
		}
		projectID = project.ID
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, iamEntity, projectEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, iamEntity, projectEntity, "delete", projectID, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if string(resp) != fmt.Sprintf("Project '%s' deleted\n", projectID) {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), fmt.Sprintf("Project '%s' deleted\n", projectID))
		}
	})
}
