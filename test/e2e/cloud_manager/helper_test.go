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

//go:build e2e || cloudmanager || om60

package cloud_manager_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	alertsEntity      = "alerts"
	testedMDBVersion  = "5.0.0"
	testedMDBFCV      = "5.0"
	entity            = "cloud-manager"
	serversEntity     = "servers"
	projectsEntity    = "projects"
	orgsEntity        = "orgs"
	clustersEntity    = "clusters"
	maintenanceEntity = "maintenanceWindows"
	monitoringEntity  = "monitoring"
	processesEntity   = "processes"
	featurePolicies   = "featurePolicies"
	eventsEntity      = "events"
	agentsEntity      = "agents"
	backupEntity      = "backup"
	dbUsersEntity     = "dbusers"
	securityEntity    = "security"
)

var (
	ErrNoServers = errors.New("no server available")
	ErrNoHosts   = errors.New("no hosts available")
)

// automationServerHostname tries to list available server running the automation agent
// and returns the first available hostname for deployments.
func automationServerHostname(cliPath string) (string, error) {
	cmd := exec.Command(cliPath, entity, serversEntity, "list", "-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w\n %s", err, string(resp))
	}

	var servers *opsmngr.Agents
	if err := json.Unmarshal(resp, &servers); err != nil {
		return "", err
	}
	if servers.TotalCount == 0 {
		return "", ErrNoServers
	}
	sort.Sort(byLastConf(*servers))
	return servers.Results[0].Hostname, nil
}

type byLastConf opsmngr.Agents

func (s byLastConf) Len() int      { return len(s.Results) }
func (s byLastConf) Swap(i, j int) { s.Results[i], s.Results[j] = s.Results[j], s.Results[i] }
func (s byLastConf) Less(i, j int) bool {
	v1, _ := time.Parse(time.RFC3339, s.Results[i].LastConf)
	v2, _ := time.Parse(time.RFC3339, s.Results[j].LastConf)
	return v2.Before(v1)
}

func hostIDs(cliPath string) ([]string, error) {
	cmd := exec.Command(cliPath, entity, processesEntity, "list", "-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%w\n %s", err, string(resp))
	}

	var servers *opsmngr.Hosts
	if err := json.Unmarshal(resp, &servers); err != nil {
		return nil, err
	}
	if servers.TotalCount == 0 {
		return nil, ErrNoHosts
	}
	result := make([]string, len(servers.Results))
	for i, h := range servers.Results {
		result[i] = h.ID
	}
	return result, nil
}

func generateRSConfig(filename, hostname, clusterName, version, fcVersion string) error {
	configFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer configFile.Close()

	cluster := &convert.ClusterConfig{
		RSConfig: convert.RSConfig{
			FeatureCompatibilityVersion: fcVersion,
			Name:                        clusterName,
			Version:                     version,
			Processes: []*convert.ProcessConfig{
				{
					DBPath:   fmt.Sprintf("/data/%s/27000", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/27000/mongodb.log", clusterName),
					Port:     27000,
					Priority: pointer.Get[float64](1),
					Votes:    pointer.Get[float64](1),
					WiredTiger: &map[string]interface{}{
						"collectionConfig": map[string]interface{}{},
						"engineConfig": map[string]interface{}{
							"cacheSizeGB": 1,
						},
						"indexConfig": map[string]interface{}{},
					},
				},
				{
					DBPath:   fmt.Sprintf("/data/%s/27001", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/27001/mongodb.log", clusterName),
					Port:     27001,
					Priority: pointer.Get[float64](1),
					Votes:    pointer.Get[float64](1),
					WiredTiger: &map[string]interface{}{
						"collectionConfig": map[string]interface{}{},
						"engineConfig": map[string]interface{}{
							"cacheSizeGB": 1,
						},
						"indexConfig": map[string]interface{}{},
					},
				},
				{
					DBPath:   fmt.Sprintf("/data/%s/27002", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/27002/mongodb.log", clusterName),
					Port:     27002,
					Priority: pointer.Get[float64](1),
					Votes:    pointer.Get[float64](1),
					WiredTiger: &map[string]interface{}{
						"collectionConfig": map[string]interface{}{},
						"engineConfig": map[string]interface{}{
							"cacheSizeGB": 1,
						},
						"indexConfig": map[string]interface{}{},
					},
				},
			},
		},
	}

	jsonEncoder := json.NewEncoder(configFile)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder.Encode(cluster)
}

func generateRSConfigUpdate(filename string) error {
	defaultWriteConcernJ := true
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var cluster convert.ClusterConfig
	err = json.Unmarshal(jsonData, &cluster)
	if err != nil {
		return err
	}

	for i := 0; i < len(cluster.Processes); i++ {
		cluster.Processes[i].DefaultRWConcern = &convert.DefaultRWConcern{
			DefaultReadConcern: &convert.DefaultReadConcern{
				Level: "majority",
			},
			DefaultWriteConcern: &convert.DefaultWriteConcern{
				W:        1,
				J:        &defaultWriteConcernJ,
				Wtimeout: 0,
			},
		}
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()
	jsonEncoder := json.NewEncoder(out)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder.Encode(cluster)
}

func generateShardedConfig(filename, hostname, clusterName, version, fcVersion string) error {
	configFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer configFile.Close()

	cluster := &convert.ClusterConfig{
		RSConfig: convert.RSConfig{
			FeatureCompatibilityVersion: fcVersion,
			Name:                        clusterName,
			Version:                     version,
		},
		Config: &convert.RSConfig{
			Name: clusterName + "_configRS",
			Processes: []*convert.ProcessConfig{
				{
					DBPath:   fmt.Sprintf("/data/%s/29000", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/29000/mongodb.log", clusterName),
					Port:     29000,
					Priority: pointer.Get[float64](1),
					Votes:    pointer.Get[float64](1),
					WiredTiger: &map[string]interface{}{
						"collectionConfig": map[string]interface{}{},
						"engineConfig": map[string]interface{}{
							"cacheSizeGB": 1,
						},
						"indexConfig": map[string]interface{}{},
					},
				},
				{
					DBPath:   fmt.Sprintf("/data/%s/29001", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/29001/mongodb.log", clusterName),
					Port:     29001,
					Priority: pointer.Get[float64](1),
					Votes:    pointer.Get[float64](1),
					WiredTiger: &map[string]interface{}{
						"collectionConfig": map[string]interface{}{},
						"engineConfig": map[string]interface{}{
							"cacheSizeGB": 1,
						},
						"indexConfig": map[string]interface{}{},
					},
				},
				{
					DBPath:   fmt.Sprintf("/data/%s/29002", clusterName),
					Hostname: hostname,
					LogPath:  fmt.Sprintf("/data/%s/29002/mongodb.log", clusterName),
					Port:     29002,
					Priority: pointer.Get[float64](1),
					Votes:    pointer.Get[float64](1),
					WiredTiger: &map[string]interface{}{
						"collectionConfig": map[string]interface{}{},
						"engineConfig": map[string]interface{}{
							"cacheSizeGB": 1,
						},
						"indexConfig": map[string]interface{}{},
					},
				},
			},
		},
		Mongos: []*convert.ProcessConfig{
			{
				Hostname: hostname,
				LogPath:  fmt.Sprintf("/data/%s/30000/mongodb.log", clusterName),
				Port:     30000,
			},
		},
		Shards: []*convert.RSConfig{
			{
				Name: clusterName + "_myShard_0",
				Processes: []*convert.ProcessConfig{
					{
						DBPath:   fmt.Sprintf("/data/%s/27000", clusterName),
						Hostname: hostname,
						LogPath:  fmt.Sprintf("/data/%s/27000/mongodb.log", clusterName),
						Port:     27000,
						Priority: pointer.Get[float64](1),
						Votes:    pointer.Get[float64](1),
						WiredTiger: &map[string]interface{}{
							"collectionConfig": map[string]interface{}{},
							"engineConfig": map[string]interface{}{
								"cacheSizeGB": 0.5,
							},
							"indexConfig": map[string]interface{}{},
						},
					},
					{
						DBPath:   fmt.Sprintf("/data/%s/27001", clusterName),
						Hostname: hostname,
						LogPath:  fmt.Sprintf("/data/%s/27001/mongodb.log", clusterName),
						Port:     27001,
						Priority: pointer.Get[float64](1),
						Votes:    pointer.Get[float64](1),
						WiredTiger: &map[string]interface{}{
							"collectionConfig": map[string]interface{}{},
							"engineConfig": map[string]interface{}{
								"cacheSizeGB": 0.5,
							},
							"indexConfig": map[string]interface{}{},
						},
					},
					{
						DBPath:   fmt.Sprintf("/data/%s/27002", clusterName),
						Hostname: hostname,
						LogPath:  fmt.Sprintf("/data/%s/27002/mongodb.log", clusterName),
						Port:     27002,
						Priority: pointer.Get[float64](1),
						Votes:    pointer.Get[float64](1),
						WiredTiger: &map[string]interface{}{
							"collectionConfig": map[string]interface{}{},
							"engineConfig": map[string]interface{}{
								"cacheSizeGB": 0.5,
							},
							"indexConfig": map[string]interface{}{},
						},
					},
				},
			},
			{
				Name: clusterName + "_myShard_1",
				Processes: []*convert.ProcessConfig{
					{
						DBPath:   fmt.Sprintf("/data/%s/28000", clusterName),
						Hostname: hostname,
						LogPath:  fmt.Sprintf("/data/%s/28000/mongodb.log", clusterName),
						Port:     28000,
						Priority: pointer.Get[float64](1),
						Votes:    pointer.Get[float64](1),
						WiredTiger: &map[string]interface{}{
							"collectionConfig": map[string]interface{}{},
							"engineConfig": map[string]interface{}{
								"cacheSizeGB": 0.5,
							},
							"indexConfig": map[string]interface{}{},
						},
					},
					{
						DBPath:   fmt.Sprintf("/data/%s/28001", clusterName),
						Hostname: hostname,
						LogPath:  fmt.Sprintf("/data/%s/28001/mongodb.log", clusterName),
						Port:     28001,
						Priority: pointer.Get[float64](1),
						Votes:    pointer.Get[float64](1),
						WiredTiger: &map[string]interface{}{
							"collectionConfig": map[string]interface{}{},
							"engineConfig": map[string]interface{}{
								"cacheSizeGB": 0.5,
							},
							"indexConfig": map[string]interface{}{},
						},
					},
					{
						DBPath:   fmt.Sprintf("/data/%s/28002", clusterName),
						Hostname: hostname,
						LogPath:  fmt.Sprintf("/data/%s/28002/mongodb.log", clusterName),
						Port:     28002,
						Priority: pointer.Get[float64](1),
						Votes:    pointer.Get[float64](1),
						WiredTiger: &map[string]interface{}{
							"collectionConfig": map[string]interface{}{},
							"engineConfig": map[string]interface{}{
								"cacheSizeGB": 0.5,
							},
							"indexConfig": map[string]interface{}{},
						},
					},
				},
			},
		},
	}

	jsonEncoder := json.NewEncoder(configFile)
	jsonEncoder.SetIndent("", "  ")
	return jsonEncoder.Encode(cluster)
}

func watchAutomation(cliPath string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		cmd := exec.Command(cliPath,
			entity,
			"automation",
			"watch",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	}
}
