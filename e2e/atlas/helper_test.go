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

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	atlasEntity    = "atlas"
	clustersEntity = "clusters"
	mongoCliPath   = "../../bin/mongocli"
)

var cliPath string

func init() {
	path, err := filepath.Abs(mongoCliPath)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(cliPath)
	if err != nil {
		panic(err)
	}

	cliPath = path
}

func getHostnameAndPort() (string, error) {
	cmd := exec.Command(cliPath,
		atlasEntity,
		"processes",
		"list")

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	var processes []*mongodbatlas.Process
	err = json.Unmarshal(resp, &processes)

	if err != nil {
		return "", err
	}

	if len(processes) == 0 {
		return "", fmt.Errorf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return processes[0].Hostname + ":" + strconv.Itoa(processes[0].Port), nil
}

// anyCluster returns true if there is at least a cluster is deployed, false otherwise
func anyCluster() bool {
	cmd := exec.Command(cliPath,
		atlasEntity,
		clustersEntity,
		"list")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return false
	}

	var clusters []mongodbatlas.Cluster
	err = json.Unmarshal(resp, &clusters)

	if err != nil {
		return false
	}

	return len(clusters) > 0
}

func deployCluster(clusterName string) error {
	cmd := exec.Command(cliPath,
		atlasEntity,
		clustersEntity,
		"create",
		clusterName,
		"--region=US_EAST_1",
		"--members=3",
		"--tier=M10",
		"--provider=AWS",
		"--mdbVersion=4.0",
		"--diskSizeGB=10")
	cmd.Env = os.Environ()
	err := cmd.Run()

	if err != nil {
		return err
	}

	cmd = exec.Command(cliPath,
		"atlas",
		clustersEntity,
		"watch",
		clusterName)
	cmd.Env = os.Environ()
	return cmd.Run()
}

func deleteCluster(clusterName string) error {
	cmd := exec.Command(cliPath, atlasEntity, "clusters", "delete", clusterName, "--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}

func getHostname() (string, error) {
	hostnamePort, err := getHostnameAndPort()
	if err != nil {
		return "", err
	}

	parts := strings.Split(hostnamePort, ":")
	return parts[0], nil
}
