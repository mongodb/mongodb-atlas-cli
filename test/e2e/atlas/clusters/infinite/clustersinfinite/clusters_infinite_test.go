// Copyright 2026 MongoDB Inc
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

package clustersinfinite

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

// Atlas Infinite (disaggregated storage) clusters are requested via databaseEdition=INFINITE.
// They are gated behind Atlas feature flags (e.g. DISAGGREGATED_STORAGE_ATLAS and the Atlas Infinite
// public-preview flag), so the project used by this test must have Atlas Infinite enabled.
func TestClustersInfinite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	req := require.New(t)

	cliPath, err := internal.AtlasCLIBin()
	req.NoError(err)

	clusterName := g.Memory("clusterName", internal.Must(internal.RandClusterName())).(string)
	g.GenerateProject("ClustersInfinite")

	// databaseEdition is exposed on the createGroupCluster 2024-10-23 schema.
	const apiVersion = "2024-10-23"

	g.Run("Create Infinite Cluster via api subcommand", func(_ *testing.T) {
		// On the 2024-10-23 schema there is no top-level providerSettings; the provider is set
		// per regionConfig via providerName. Disk-related fields (diskSizeGB, diskIOPS,
		// ebsVolumeType, ...) are managed by Atlas Infinite and must not be set.
		// autoScaling.compute.enabled is required, and Atlas Infinite replica sets reject
		// nodeCount=3 electable nodes, so a single electable node is used.
		payload := map[string]any{
			"name":            clusterName,
			"clusterType":     "REPLICASET",
			"databaseEdition": "INFINITE",
			"replicationSpecs": []any{
				map[string]any{
					"zoneName": "Zone 1",
					"regionConfigs": []any{
						map[string]any{
							"providerName": "AWS",
							"regionName":   "US_EAST_1",
							"priority":     7,
							"electableSpecs": map[string]any{
								"instanceSize": "M10",
								"nodeCount":    1,
							},
							"autoScaling": map[string]any{
								"compute": map[string]any{"enabled": false},
							},
						},
					},
				},
			},
			"terminationProtectionEnabled": false,
		}

		payloadBytes, err := json.Marshal(payload)
		req.NoError(err)

		cmd := exec.Command(cliPath,
			"api",
			"clusters",
			"createCluster",
			"--groupId", g.ProjectID,
			"--version", apiVersion,
			"--watch",
			"-P", internal.ProfileName())
		cmd.Env = os.Environ()
		cmd.Stdin = bytes.NewReader(payloadBytes)

		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})

	g.Run("Get Infinite Cluster via api subcommand", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			"api",
			"clusters",
			"getGroupCluster",
			"--groupId", g.ProjectID,
			"--clusterName", clusterName,
			"--version", apiVersion,
			"-P", internal.ProfileName())
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		// The create/update databaseEdition defaults to CORE on the doc; the resolved edition is
		// reported via effectiveDatabaseEdition.
		var cluster struct {
			Name                     string `json:"name"`
			EffectiveDatabaseEdition string `json:"effectiveDatabaseEdition"`
		}
		req.NoError(json.Unmarshal(resp, &cluster))
		req.Equal(clusterName, cluster.Name)
		req.Equal("INFINITE", cluster.EffectiveDatabaseEdition)
	})

	g.Run("Delete Infinite Cluster via api subcommand", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			"api",
			"clusters",
			"deleteCluster",
			"--groupId", g.ProjectID,
			"--clusterName", clusterName,
			"--version", apiVersion,
			"-P", internal.ProfileName())
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})
}
