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
// +build e2e atlas

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	atlasEntity    = "atlas"
	clustersEntity = "clusters"
	searchEntity   = "search"
	indexEntity    = "index"
	datalakeEntity = "datalake"
	alertsEntity   = "alerts"
	configEntity   = "settings"
	dbusersEntity  = "dbusers"
	certsEntity    = "certs"
)

func getHostnameAndPort() (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(cliPath,
		atlasEntity,
		"processes",
		"list",
		"-o=json")

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

func deployCluster() (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", fmt.Errorf("error creating cluster %w", err)
	}
	clusterName, err := RandClusterName()
	if err != nil {
		return "", err
	}
	create := exec.Command(cliPath,
		atlasEntity,
		clustersEntity,
		"create",
		clusterName,
		"--region=US_EAST_1",
		"--tier=M10",
		"--provider=AWS",
		"--diskSizeGB=10")
	create.Env = os.Environ()
	if err := create.Run(); err != nil {
		return "", fmt.Errorf("error creating cluster %w", err)
	}

	watch := exec.Command(cliPath,
		"atlas",
		clustersEntity,
		"watch",
		clusterName)
	watch.Env = os.Environ()
	if err := watch.Run(); err != nil {
		return "", fmt.Errorf("error watching cluster %w", err)
	}
	return clusterName, nil
}

func deleteCluster(clusterName string) error {
	cliPath, err := e2e.Bin()
	if err != nil {
		return err
	}
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

func RandClusterName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("e2e-cluster-%v", n), nil
}
