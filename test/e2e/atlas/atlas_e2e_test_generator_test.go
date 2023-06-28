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
	"strings"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

const (
	maxRetryAttempts = 5
)

// atlasE2ETestGenerator is about providing capabilities to provide projects and clusters for our e2e tests.
type atlasE2ETestGenerator struct {
	projectID      string
	projectName    string
	clusterName    string
	clusterRegion  string
	serverlessName string
	teamName       string
	teamID         string
	teamUser       string
	dbUser         string
	tier           string
	enableBackup   bool
	firstProcess   *atlasv2.ApiHostViewAtlas
	t              *testing.T
}

// Log formats its arguments using default formatting, analogous to Println,
// and records the text in the error log. For tests, the text will be printed only if
// the test fails or the -test.v flag is set. For benchmarks, the text is always
// printed to avoid having performance depend on the value of the -test.v flag.
func (g *atlasE2ETestGenerator) Log(args ...any) {
	g.t.Log(args...)
}

// Logf formats its arguments according to the format, analogous to Printf, and
// records the text in the error log. A final newline is added if not provided. For
// tests, the text will be printed only if the test fails or the -test.v flag is
// set. For benchmarks, the text is always printed to avoid having performance
// depend on the value of the -test.v flag.
func (g *atlasE2ETestGenerator) Logf(format string, args ...any) {
	g.t.Logf(format, args...)
}

// newAtlasE2ETestGenerator creates a new instance of atlasE2ETestGenerator struct.
func newAtlasE2ETestGenerator(t *testing.T) *atlasE2ETestGenerator {
	t.Helper()
	return &atlasE2ETestGenerator{t: t}
}

func newAtlasE2ETestGeneratorWithBackup(t *testing.T) *atlasE2ETestGenerator {
	t.Helper()
	return &atlasE2ETestGenerator{t: t, enableBackup: true}
}

func (g *atlasE2ETestGenerator) generateTeam(prefix string) {
	g.t.Helper()

	if g.teamID != "" {
		g.t.Fatal("unexpected error: team was already generated")
	}

	var err error
	if prefix == "" {
		g.teamName, err = RandTeamName()
	} else {
		g.teamName, err = RandTeamNameWithPrefix(prefix)
	}
	if err != nil {
		g.t.Fatalf("unexpected error: %v", err)
	}

	g.teamUser, err = getFirstOrgUser()
	if err != nil {
		g.t.Fatalf("unexpected error retrieving org user: %v", err)
	}
	g.teamID, err = createTeam(g.teamName, g.teamUser)
	if err != nil {
		g.t.Fatalf("unexpected error creating team: %v", err)
	}
	g.Logf("teamID=%s", g.teamID)
	g.Logf("teamName=%s", g.teamName)
	if g.teamID == "" {
		g.t.Fatal("teamID not created")
	}
	g.t.Cleanup(func() {
		deleteTeamWithRetry(g.t, g.teamID)
	})
}

// generateProject generates a new project and also registers its deletion on test cleanup.
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
		g.t.Fatalf("unexpected error creating project: %v", err)
	}
	g.Logf("projectID=%s", g.projectID)
	g.Logf("projectName=%s", g.projectName)
	if g.projectID == "" {
		g.t.Fatal("projectID not created")
	}

	g.t.Cleanup(func() {
		deleteProjectWithRetry(g.t, g.projectID)
	})
}

func (g *atlasE2ETestGenerator) generateEmptyProject(prefix string) {
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

	g.projectID, err = createProjectWithoutAlertSettings(g.projectName)
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

func (g *atlasE2ETestGenerator) generateDBUser(prefix string) {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	if g.dbUser != "" {
		g.t.Fatal("unexpected error: DBUser was already generated")
	}

	var err error
	if prefix == "" {
		g.dbUser, err = RandTeamName()
	} else {
		g.dbUser, err = RandTeamNameWithPrefix(prefix)
	}
	if err != nil {
		g.t.Fatalf("unexpected error: %v", err)
	}

	err = createDBUserWithCert(g.projectID, g.dbUser)
	if err != nil {
		g.dbUser = ""
		g.t.Fatalf("unexpected error: %v", err)
	}
	g.t.Logf("dbUser=%s", g.dbUser)
}

func deleteTeamWithRetry(t *testing.T, teamID string) {
	t.Helper()
	deleted := false
	backoff := 1
	for attempts := 1; attempts <= maxRetryAttempts; attempts++ {
		e := deleteTeam(teamID)
		if e == nil {
			t.Logf("team %q successfully deleted", teamID)
			deleted = true
			break
		}
		t.Logf("%d/%d attempts - trying again in %d seconds: unexpected error while deleting the team %q: %v", attempts, maxRetryAttempts, backoff, teamID, e)
		time.Sleep(time.Duration(backoff) * time.Second)
		backoff *= 2
	}

	if !deleted {
		t.Errorf("we could not delete the team %q", teamID)
	}
}

func deleteProjectWithRetry(t *testing.T, projectID string) {
	t.Helper()
	deleted := false
	backoff := 1
	for attempts := 1; attempts <= maxRetryAttempts; attempts++ {
		e := deleteProject(projectID)
		if e == nil {
			t.Logf("project %q successfully deleted", projectID)
			deleted = true
			break
		}
		t.Logf("%d/%d attempts - trying again in %d seconds: unexpected error while deleting the project %q: %v", attempts, maxRetryAttempts, backoff, projectID, e)
		if strings.Contains(e.Error(), "CANNOT_CLOSE_GROUP_ACTIVE_ATLAS_CLUSTERS") {
			cliPath, err := e2e.AtlasCLIBin()
			if err != nil {
				t.Errorf("%s: invalid bin", err)
			}
			deleteClustersForProject(t, cliPath, projectID)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
		backoff *= 2
	}
	if !deleted {
		t.Errorf("we could not delete the project %q", projectID)
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
		g.t.Errorf("unexpected error deploying serverless instance: %v", err)
	}
	g.t.Logf("serverlessName=%s", g.serverlessName)

	g.t.Cleanup(func() {
		if e := deleteServerlessInstanceForProject(g.projectID, g.serverlessName); e != nil {
			g.t.Errorf("unexpected error deleting serverless instance: %v", e)
		}
	})
}

// generateCluster generates a new cluster and also registers its deletion on test cleanup.
func (g *atlasE2ETestGenerator) generateCluster() {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	var err error
	if g.tier == "" {
		g.tier = e2eTier()
	}

	g.clusterName, g.clusterRegion, err = deployClusterForProject(g.projectID, g.tier, g.enableBackup)
	if err != nil {
		g.Logf("projectID=%q, clusterName=%q", g.projectID, g.clusterName)
		g.t.Errorf("unexpected error deploying cluster: %v", err)
	}
	g.t.Logf("clusterName=%s", g.clusterName)

	g.t.Cleanup(func() {
		g.Logf("Cluster cleanup %q\n", g.projectID)
		if e := deleteClusterForProject(g.projectID, g.clusterName); e != nil {
			g.t.Errorf("unexpected error deleting cluster: %v", e)
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
	return *p.Id, nil
}

// getHostname returns the hostname of first process.
func (g *atlasE2ETestGenerator) getHostname() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	return *p.Hostname, nil
}

// getFirstProcess returns the first process of the project.
func (g *atlasE2ETestGenerator) getFirstProcess() (*atlasv2.ApiHostViewAtlas, error) {
	g.t.Helper()

	if g.firstProcess != nil {
		return g.firstProcess, nil
	}

	processes, err := g.getProcesses()
	if err != nil {
		return nil, err
	}
	g.firstProcess = &processes[0]

	return g.firstProcess, nil
}

// getHostname returns the list of processes.
func (g *atlasE2ETestGenerator) getProcesses() ([]atlasv2.ApiHostViewAtlas, error) {
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

	var processes *atlasv2.PaginatedHostViewAtlas
	if err := json.Unmarshal(resp, &processes); err != nil {
		g.t.Fatalf("unexpected error: project must be generated %s - %s", err, resp)
	}

	if len(processes.Results) == 0 {
		g.t.Fatalf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	return processes.Results, nil
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
