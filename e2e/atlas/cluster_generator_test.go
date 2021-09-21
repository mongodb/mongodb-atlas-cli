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
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

type clusterGenerator struct {
	projectID   string
	projectName string
	clusterName string
	t           *testing.T
}

func newClusterGenerator(t *testing.T) *clusterGenerator {
	t.Helper()

	return &clusterGenerator{t: t}
}

func (g *clusterGenerator) generateProject(prefix string) {
	g.t.Helper()

	if g.projectID != "" {
		g.t.Errorf("generateProject() may be called only once per test")
	}

	var err error
	if prefix == "" {
		g.projectName, err = RandProjectName()
	} else {
		g.projectName, err = RandProjectNameWithPrefix(prefix)
	}
	if err != nil {
		g.t.Errorf("unexpected error: %v", err)
	}

	g.projectID, err = createProject(g.projectName)
	if err != nil {
		g.t.Errorf("unexpected error: %v", err)
	}
	g.t.Logf("projectID=%s", g.projectID)

	g.t.Cleanup(func() {
		if e := deleteProject(g.projectID); e != nil {
			g.t.Errorf("unexpected error: %v", e)
		}
	})
}

func (g *clusterGenerator) generateCluster() {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Errorf("unexpected error: generateProject() must be called before generateCluster()")
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

func (g *clusterGenerator) generateProjectAndCluster(prefix string) {
	g.t.Helper()

	g.generateProject(prefix)
	g.generateCluster()
}

func (g *clusterGenerator) newAvailableRegion(tier, provider string) (string, error) {
	g.t.Helper()

	return newAvailableRegion(g.projectID, tier, provider)
}

func (g *clusterGenerator) getHostnameAndPort() (string, error) {
	g.t.Helper()

	processes, err := g.getProcesses()
	if err != nil {
		return "", err
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return processes[0].Hostname + ":" + strconv.Itoa(processes[0].Port), nil
}

func (g *clusterGenerator) getHostname() (string, error) {
	g.t.Helper()

	processes, err := g.getProcesses()
	if err != nil {
		return "", err
	}

	return processes[0].Hostname, nil
}

func (g *clusterGenerator) getProcesses() ([]*mongodbatlas.Process, error) {
	g.t.Helper()

	resp, err := g.runCommand(atlasEntity,
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
	err = json.Unmarshal(resp, &processes)

	if err != nil {
		return nil, err
	}

	if len(processes) == 0 {
		return nil, fmt.Errorf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	return processes, nil
}

func (g *clusterGenerator) runCommand(args ...string) ([]byte, error) {
	g.t.Helper()

	cliPath, err := e2e.Bin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath, args...)

	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}
