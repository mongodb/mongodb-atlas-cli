package e2e

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	iamEntity      = "iam"
	projectsEntity = "projects"
)

func CreateProject(projectName string) (string, error) {
	cliPath, err := Bin()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		projectsEntity,
		"create",
		projectName,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	var project mongodbatlas.Project
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", err
	}

	return project.ID, nil
}

func deleteProject(projectID string) error {
	cliPath, err := Bin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
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
		if e := deleteProject(projectID); e != nil {
			t.Logf("%d/%d attempts - trying again in %d seconds: unexpected error while deleting the project %q: %v", attempts, maxRetryAttempts, sleepTimeInSeconds, projectID, e)
			time.Sleep(sleepTimeInSeconds * time.Second)
		} else {
			t.Logf("project %q successfully deleted", projectID)
			deleted = true
			break
		}
	}

	if !deleted {
		t.Errorf("we could not delete the project %q", projectID)
	}
}
