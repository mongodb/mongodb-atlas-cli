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
// +build e2e cloudmanager

package cloud_manager_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/mongodb/mongocli/internal/convert"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	entity        = "cloud-manager"
	serversEntity = "servers"
)

// automationServerHostname tries to list available server running the automation agent
// and returns the first available hostname for deployments
func automationServerHostname(cliPath string) (string, error) {
	cmd := exec.Command(cliPath, entity, serversEntity, "list")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	var servers *opsmngr.Agents
	if err := json.Unmarshal(resp, &servers); err != nil {
		return "", err
	}
	if servers.TotalCount == 0 {
		return "", errors.New("no server available")
	}
	return servers.Results[0].Hostname, nil
}

func generateConfig(filename, hostname, clusterName, version, fcVersion string) error {
	feedFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer feedFile.Close()

	var one float64 = 1
	downloadArchive := &convert.ClusterConfig{
		RSConfig: convert.RSConfig{
			FCVersion: fcVersion,
			Name:      clusterName,
			Version:   version,
			ProcessConfigs: []*convert.ProcessConfig{
				{
					DBPath:   fmt.Sprintf("/data/%s/27000", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/27000/mongodb.log", clusterName),
					Port:     27000,
					Priority: &one,
					Votes:    &one,
				},
				{
					DBPath:   fmt.Sprintf("/data/%s/27001", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/27001/mongodb.log", clusterName),
					Port:     27001,
					Priority: &one,
					Votes:    &one,
				},
				{
					DBPath:   fmt.Sprintf("/data/%s/27002", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/27002/mongodb.log", clusterName),
					Port:     27002,
					Priority: &one,
					Votes:    &one,
				},
			},
		},
	}

	jsonEncoder := json.NewEncoder(feedFile)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder.Encode(downloadArchive)
}
