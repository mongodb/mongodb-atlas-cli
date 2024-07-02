// Copyright 2023 MongoDB Inc
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

package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	projectsEntity = "projects"
)

func CreateProject(projectName string) (string, error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(cliPath,
		projectsEntity,
		"create",
		projectName,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	var project atlasv2.Group
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("%w: %s", err, resp)
	}

	return project.GetId(), nil
}

func deleteProject(projectID string) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		projectsEntity,
		"delete",
		projectID,
		"--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}

const (
	maxRetryAttempts   = 10
	sleepTimeInSeconds = 30
)

func DeleteProjectWithRetry(t *testing.T, projectID string) {
	t.Helper()
	deleted := false
	for attempts := 1; attempts <= maxRetryAttempts; attempts++ {
		e := deleteProject(projectID)
		if e == nil {
			t.Logf("project %q successfully deleted", projectID)
			deleted = true
			break
		}
		t.Logf("%d/%d attempts - trying again in %d seconds: unexpected error while deleting the project %q: %v", attempts, maxRetryAttempts, sleepTimeInSeconds, projectID, e)
		time.Sleep(sleepTimeInSeconds * time.Second)
	}

	if !deleted {
		t.Errorf("we could not delete the project %q", projectID)
	}
}

func RunAndGetStdOut(cmd *exec.Cmd) ([]byte, error) {
	cmd.Stderr = os.Stderr
	var b bytes.Buffer
	cmd.Stdout = &b

	err := cmd.Run()
	resp := b.Bytes()

	if err != nil {
		return nil, fmt.Errorf("%s (%w)", string(resp), err)
	}

	return resp, nil
}
