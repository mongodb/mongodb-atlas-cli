// Copyright 2021 MongoDB Inc
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
//go:build e2e || atlas

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	maxRetryAttempts   = 10
	sleepTimeInSeconds = 30
)

// atlasE2ETestGenerator is about providing capabilities to provide projects and clusters for our e2e tests.
type atlasE2ETestGenerator struct {
	projectID      string
	projectName    string
	clusterName    string
	serverlessName string
	firstProcess   *mongodbatlas.Process
	t              *testing.T
}

// newAtlasE2ETestGenerator creates a new instance of atlasE2ETestGenerator struct.
func newAtlasE2ETestGenerator(t *testing.T) *atlasE2ETestGenerator {
	t.Helper()

	return &atlasE2ETestGenerator{t: t}
}

// generateProject generates a new project and also registers it's deletion on test cleanup.
func (g *atlasE2ETestGenerator) generateProject(prefix string) {
	g.t.Helper()

	if g.projectID != "" {
		g.t.Fatal("unexpected error: project was already generated")
	}

	var err error
	if prefix == "" {
		g.projectName, err = RandProjectName()
	} else {
		g.projectName, err = RandProjectNameWithPrefix(prefix)
	}
	if err != nil {
		g.t.Fatalf("unexpected error: %v", err)
	}

	g.projectID, err = createProject(g.projectName)
	if err != nil {
		g.t.Fatalf("unexpected error: %v", err)
	}
	g.t.Logf("projectID=%s", g.projectID)
	g.t.Logf("projectName=%s", g.projectName)
	if g.projectID == "" {
		g.t.Fatal("projectID not created")
	}

	g.t.Cleanup(func() {
		deleteProjectWithRetry(g.t, g.projectID)
	})
}

func deleteProjectWithRetry(t *testing.T, projectID string) {
	t.Helper()
	deleted := false
	for attempts := 1; attempts <= maxRetryAttempts; attempts++ {
		if e := deleteProject(projectID); e != nil {
			t.Logf("%d/%d attempts - trying again in %d seconds: unexpected error while deleting the project '%s': %v", attempts, maxRetryAttempts, sleepTimeInSeconds, projectID, e)
			time.Sleep(sleepTimeInSeconds * time.Second)
		} else {
			t.Logf("project '%s' successfully deleted", projectID)
			deleted = true
			break
		}
	}

	if !deleted {
		t.Errorf("we could not delete the project '%s'", projectID)
	}
}

func (g *atlasE2ETestGenerator) generateServerlessCluster() {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	var err error
	g.serverlessName, err = deployServerlessInstanceForProject(g.projectID)
	if err != nil {
		g.t.Errorf("unexpected error: %v", err)
	}
	g.t.Logf("serverlessName=%s", g.serverlessName)

	g.t.Cleanup(func() {
		if e := deleteServerlessInstanceForProject(g.projectID, g.serverlessName); e != nil {
			g.t.Errorf("unexpected error: %v", e)
		}
	})
}

// generateCluster generates a new cluster and also registers it's deletion on test cleanup.
func (g *atlasE2ETestGenerator) generateCluster() {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	var err error
	g.clusterName, err = deployClusterForProject(g.projectID)
	if err != nil {
		g.t.Errorf("unexpected error: %v", err)
	}
	g.t.Logf("clusterName=%s", g.clusterName)

	g.t.Cleanup(func() {
		if e := deleteClusterForProject(g.projectID, g.clusterName); e != nil {
			g.t.Errorf("unexpected error: %v", e)
		}
	})
}

// generateProjectAndCluster calls both generateProject and generateCluster.
func (g *atlasE2ETestGenerator) generateProjectAndCluster(prefix string) {
	g.t.Helper()

	g.generateProject(prefix)
	g.generateCluster()
}

// newAvailableRegion returns the first region for the provider/tier.
func (g *atlasE2ETestGenerator) newAvailableRegion(tier, provider string) (string, error) {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	return newAvailableRegion(g.projectID, tier, provider)
}

// getHostnameAndPort returns hostname:port from the first process.
func (g *atlasE2ETestGenerator) getHostnameAndPort() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return p.ID, nil
}

// getHostname returns the hostname of first process.
func (g *atlasE2ETestGenerator) getHostname() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	return p.Hostname, nil
}

// getFirstProcess returns the first process of the project.
func (g *atlasE2ETestGenerator) getFirstProcess() (*mongodbatlas.Process, error) {
	g.t.Helper()

	if g.firstProcess != nil {
		return g.firstProcess, nil
	}

	processes, err := g.getProcesses()
	if err != nil {
		return nil, err
	}
	g.firstProcess = processes[0]

	return g.firstProcess, nil
}

// getHostname returns the list of processes.
func (g *atlasE2ETestGenerator) getProcesses() ([]*mongodbatlas.Process, error) {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	resp, err := g.runCommand(
		processesEntity,
		"list",
		"--projectId",
		g.projectID,
		"-o=json",
	)
	if err != nil {
		return nil, err
	}

	var processes []*mongodbatlas.Process
	if err := json.Unmarshal(resp, &processes); err != nil {
		g.t.Fatalf("unexpected error: project must be generated %s - %s", err, resp)
	}

	if len(processes) == 0 {
		g.t.Fatalf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	return processes, nil
}

// runCommand runs a command on mongocli tool.
func (g *atlasE2ETestGenerator) runCommand(args ...string) ([]byte, error) {
	g.t.Helper()

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath, args...)

	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}
