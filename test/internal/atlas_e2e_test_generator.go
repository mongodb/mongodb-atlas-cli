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

package internal

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312027/admin"
)

// AtlasE2ETestGenerator is about providing capabilities to provide projects and clusters for our e2e tests.
type AtlasE2ETestGenerator struct {
	ProjectID     string
	projectName   string
	ClusterName   string
	clusterRegion string
	Tier          string
	MDBVer        string
	enableBackup  bool
	firstProcess  *atlasv2.ApiHostViewAtlas
	t             *testing.T
	ClientID      string
	ClientSecret  string
}

// Log formats its arguments using default formatting, analogous to Println,
// and records the text in the error log. For tests, the text will be printed only if
// the test fails or the -test.v flag is set. For benchmarks, the text is always
// printed to avoid having performance depend on the value of the -test.v flag.
func (g *AtlasE2ETestGenerator) Log(args ...any) {
	g.t.Log(args...)
}

// Logf formats its arguments according to the format, analogous to Printf, and
// records the text in the error log. A final newline is added if not provided. For
// tests, the text will be printed only if the test fails or the -test.v flag is
// set. For benchmarks, the text is always printed to avoid having performance
// depend on the value of the -test.v flag.
func (g *AtlasE2ETestGenerator) Logf(format string, args ...any) {
	g.t.Logf(format, args...)
}

// newAtlasE2ETestGenerator creates a new instance of AtlasE2ETestGenerator struct.
func NewAtlasE2ETestGenerator(t *testing.T, opts ...func(g *AtlasE2ETestGenerator)) *AtlasE2ETestGenerator {
	t.Helper()

	g := &AtlasE2ETestGenerator{
		t: t,
	}

	for _, opt := range opts {
		opt(g)
	}
	return g
}

func WithBackup() func(g *AtlasE2ETestGenerator) {
	return func(g *AtlasE2ETestGenerator) {
		g.enableBackup = true
	}
}

func (g *AtlasE2ETestGenerator) Run(name string, f func(t *testing.T)) {
	g.t.Helper()

	g.t.Run(name, func(t *testing.T) {
		t.Helper()

		f(t)
	})
}

// GenerateProject generates a new project and also registers its deletion on test cleanup.
func (g *AtlasE2ETestGenerator) GenerateProject(prefix string) {
	g.t.Helper()

	if g.ProjectID != "" {
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

	g.ProjectID, err = createProject(g.projectName)
	if err != nil {
		g.t.Fatalf("unexpected error creating project: %v", err)
	}
	g.Logf("projectID=%s", g.ProjectID)
	g.Logf("projectName=%s", g.projectName)
	if g.ProjectID == "" {
		g.t.Fatal("projectID not created")
	}

	g.t.Cleanup(func() {
		g.Logf("Project cleanup %q", g.ProjectID)
		DeleteProjectWithRetry(g.t, g.ProjectID)
	})
}

func (g *AtlasE2ETestGenerator) generateClusterWithPrefixAndProvider(prefix, provider string) {
	g.t.Helper()

	if g.ProjectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	var err error
	if g.Tier == "" {
		g.Tier = E2eTier()
	}

	if g.MDBVer == "" {
		mdbVersion, e := MongoDBMajorVersion()
		require.NoError(g.t, e)

		g.MDBVer = mdbVersion
	}

	g.ClusterName = Must(RandClusterNameWithPrefix(prefix))

	// register the cleanup before deploying so a cluster that was created but failed to become ready is still deleted
	g.t.Cleanup(func() {
		err := DeleteClusterForProject(g.ProjectID, g.ClusterName)
		if err == nil {
			return
		}

		if strings.Contains(err.Error(), "CLUSTER_NOT_FOUND") || strings.Contains(err.Error(), "GROUP_NOT_FOUND") {
			return
		}

		g.t.Errorf("unexpected error deleting cluster: %v", err)
	})

	g.clusterRegion, err = deployClusterForProject(g.ProjectID, g.ClusterName, g.Tier, g.MDBVer, provider, g.enableBackup)
	if err != nil {
		g.Logf("projectID=%q, clusterName=%q", g.ProjectID, g.ClusterName)
		g.t.Fatalf("unexpected error deploying cluster: %v", err)
	}
}

func (g *AtlasE2ETestGenerator) generateClusterWithPrefix(prefix string) {
	g.generateClusterWithPrefixAndProvider(prefix, e2eClusterProvider)
}

// GenerateCluster generates a new cluster and also registers its deletion on test cleanup.
func (g *AtlasE2ETestGenerator) GenerateCluster() {
	g.generateClusterWithPrefix("cluster")
}

// GenerateAzureCluster generates a new Azure cluster and also registers its deletion on test cleanup.
func (g *AtlasE2ETestGenerator) GenerateAzureCluster() {
	g.generateClusterWithPrefixAndProvider("azureCluster", "AZURE")
}

// GenerateProjectAndCluster calls both generateProject and generateCluster.
func (g *AtlasE2ETestGenerator) GenerateProjectAndCluster(prefix string) {
	g.t.Helper()

	g.GenerateProject(prefix)
	g.generateClusterWithPrefix(prefix)
}

// GenerateOrgServiceAccount generates a new organization service account and also registers its deletion on test cleanup.
func (g *AtlasE2ETestGenerator) GenerateOrgServiceAccount(cliPath, name string) {
	g.t.Helper()

	if g.ClientID != "" && g.ClientSecret != "" {
		g.t.Fatal("unexpected error: service account was already generated")
	}

	var err error
	g.ClientID, g.ClientSecret, err = createOrgServiceAccount(cliPath, name)
	if err != nil {
		g.t.Fatalf("unexpected error creating org service account: %v", err)
	}

	g.Logf("clientID=%s", g.ClientID)
	if g.ClientID == "" {
		g.t.Fatal("clientID not created")
	}

	g.t.Cleanup(func() {
		g.Logf("Service Account cleanup %q", g.ClientID)
		deleteOrgServiceAccount(g.t, cliPath, g.ClientID)
	})
}

// NewAvailableRegion returns the first region for the provider/tier.
func (g *AtlasE2ETestGenerator) NewAvailableRegion(tier, provider string) (string, error) {
	g.t.Helper()

	if g.ProjectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	return NewAvailableRegion(g.ProjectID, tier, provider)
}

// GetHostnameAndPort returns hostname:port from the first process.
func (g *AtlasE2ETestGenerator) GetHostnameAndPort() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return *p.Id, nil
}

// GetHostname returns the hostname of first process.
func (g *AtlasE2ETestGenerator) GetHostname() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	return *p.Hostname, nil
}

// getFirstProcess returns the first process of the project.
func (g *AtlasE2ETestGenerator) getFirstProcess() (*atlasv2.ApiHostViewAtlas, error) {
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
func (g *AtlasE2ETestGenerator) getProcesses() ([]atlasv2.ApiHostViewAtlas, error) {
	g.t.Helper()

	if g.ProjectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	resp, err := g.RunCommand(
		processesEntity,
		"list",
		"--projectId",
		g.ProjectID,
		"-o=json",
		"-P",
		ProfileName(),
	)
	if err != nil {
		return nil, err
	}

	var processes *atlasv2.PaginatedHostViewAtlas
	if err := json.Unmarshal(resp, &processes); err != nil {
		g.t.Fatalf("unexpected error: project must be generated %s - %s", err, resp)
	}

	if len(processes.GetResults()) == 0 {
		g.t.Fatalf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	return processes.GetResults(), nil
}

// RunCommand runs a command on atlascli.
func (g *AtlasE2ETestGenerator) RunCommand(args ...string) ([]byte, error) {
	g.t.Helper()

	cliPath, err := AtlasCLIBin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath, args...)

	cmd.Env = os.Environ()
	return RunAndGetStdOut(cmd)
}
