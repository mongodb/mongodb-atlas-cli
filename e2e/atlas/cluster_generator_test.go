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

func (pg *clusterGenerator) generateProject(prefix string) {
	pg.t.Helper()

	if pg.projectID != "" {
		pg.t.Errorf("generateProject() may be called only once per test")
	}

	var err error
	if prefix == "" {
		pg.projectName, err = RandProjectName()
	} else {
		pg.projectName, err = RandProjectNameWithPrefix(prefix)
	}
	if err != nil {
		pg.t.Errorf("unexpected error: %v", err)
	}

	pg.projectID, err = createProject(pg.projectName)
	if err != nil {
		pg.t.Errorf("unexpected error: %v", err)
	}
	pg.t.Cleanup(func() {
		if e := deleteProject(pg.projectID); e != nil {
			pg.t.Errorf("unexpected error: %v", e)
		}
	})
}

func (pg *clusterGenerator) generateCluster() {
	pg.t.Helper()

	if pg.projectID == "" {
		pg.t.Errorf("unexpected error: generateProject() must be called before generateCluster()")
	}

	var err error
	pg.clusterName, err = deployClusterForProject(pg.projectID)
	if err != nil {
		pg.t.Errorf("unexpected error: %v", err)
	}

	pg.t.Cleanup(func() {
		if e := deleteClusterForProject(pg.projectID, pg.clusterName); e != nil {
			pg.t.Errorf("unexpected error: %v", e)
		}
	})
}

func (pg *clusterGenerator) generateProjectAndCluster(prefix string) {
	pg.t.Helper()

	pg.generateProject(prefix)
	pg.generateCluster()
}

func (pg *clusterGenerator) getHostname() (string, error) {
	pg.t.Helper()

	processes, err := pg.getProcesses()
	if err != nil {
		return "", err
	}

	return processes[0].Hostname, nil
}

func (pg *clusterGenerator) getProcesses() ([]*mongodbatlas.Process, error) {
	pg.t.Helper()

	resp, err := pg.runCommand(atlasEntity,
		processesEntity,
		"list",
		"--projectId",
		pg.projectID,
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

func (pg *clusterGenerator) runCommand(args ...string) ([]byte, error) {
	pg.t.Helper()

	cliPath, err := e2e.Bin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath, args...)

	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}
