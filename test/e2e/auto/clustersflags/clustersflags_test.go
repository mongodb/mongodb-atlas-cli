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

// Code generated using `make gen-arazzo-tests`. DO NOT EDIT.

package clustersflags

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal/apitest"
	"github.com/stretchr/testify/require"
)

func TestClustersFlags(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	orgId := apitest.OrgID(t)

	rand := apitest.RandomSuffix(t)

	service := apitest.Service()

	clusterName := "cluster-" + rand

	projectName := "clustersFlags-" + rand

	regionUsageRestrictions := ""

	if service == "cloudgov" {

		regionUsageRestrictions = "GOV_REGIONS_ONLY"

	}

	tier := ""

	if service == "cloud" {

		tier = "M10"

	}

	if service == "cloudgov" {

		tier = "M20"

	}

	var (
		createCluster_clusterName                               string
		createProject_projectId                                 string
		discoverMongoDBMajorVersion_mongoDBMajorVersion         string
		discoverMongoDBMajorVersion_previousMongoDBMajorVersion string
		discoverRegion_region                                   string
	)

	// Create the project that owns the test cluster.

	t.Run("createProject", func(t *testing.T) {

		createProjectBody := map[string]any{"name": projectName, "orgId": orgId}
		if regionUsageRestrictions != "" {
			createProjectBody["regionUsageRestrictions"] = regionUsageRestrictions
		}

		createProjectOut := apitest.Run(t, apitest.Marshal(t, createProjectBody), []string{"projects", "createGroup", "--version", "2023-01-01"})
		createProjectResp := apitest.Unmarshal(t, createProjectOut)

		createProject_projectId = apitest.JSONString(createProjectResp, "id")

		require.NotNil(t, apitest.JSONValue(createProjectResp, "id"))
	})

	apitest.CleanupWithRetry(t, []string{"projects", "deleteGroup", "--groupId", createProject_projectId})

	// Find an available region for the requested provider and tier.

	t.Run("discoverRegion", func(t *testing.T) {

		discoverRegionOut := apitest.Run(t, "", []string{"clusters", "listClusterProviderRegions", "--groupId", createProject_projectId, "--providers", "AWS", "--tier", tier})
		discoverRegionResp := apitest.Unmarshal(t, discoverRegionOut)

		discoverRegion_region = apitest.JSONString(discoverRegionResp, "results/0/instanceSizes/0/availableRegions/0/name")

		require.NotNil(t, apitest.JSONValue(discoverRegionResp, "results/0/instanceSizes/0/availableRegions/0/name"))
	})

	// Discover the default MongoDB major version for the provider and tier.
	// Replaces the private defaultMongoDBMajorVersion endpoint used by the
	// Go test with the public project versions endpoint.

	t.Run("discoverMongoDBMajorVersion", func(t *testing.T) {

		discoverMongoDBMajorVersionOut := apitest.Run(t, "", []string{"projects", "getMongoDbVersions", "--groupId", createProject_projectId, "--cloudProvider", "AWS", "--instanceSize", tier, "--defaultStatus", "DEFAULT"})
		discoverMongoDBMajorVersionResp := apitest.Unmarshal(t, discoverMongoDBMajorVersionOut)

		discoverMongoDBMajorVersion_mongoDBMajorVersion = apitest.JSONString(discoverMongoDBMajorVersionResp, "results/0/version")

		previousMongoDBMajorVersion, err := apitest.PreviousMajorVersion(discoverMongoDBMajorVersion_mongoDBMajorVersion)

		require.NoError(t, err)

		discoverMongoDBMajorVersion_previousMongoDBMajorVersion = previousMongoDBMajorVersion

		require.NotNil(t, apitest.JSONValue(discoverMongoDBMajorVersionResp, "results/0/version"))
	})

	// Create a replica set with termination protection enabled.

	t.Run("createCluster", func(t *testing.T) {

		createClusterBody := map[string]any{"clusterType": "REPLICASET", "diskSizeGB": 30, "mongoDBMajorVersion": discoverMongoDBMajorVersion_previousMongoDBMajorVersion, "name": clusterName, "replicationSpecs": []any{map[string]any{"numShards": 1, "regionConfigs": []any{map[string]any{"electableSpecs": map[string]any{"instanceSize": tier, "nodeCount": 3}, "priority": 7, "providerName": "AWS", "readOnlySpecs": map[string]any{"instanceSize": tier, "nodeCount": 0}, "regionName": discoverRegion_region}}, "zoneName": "Zone 1"}}, "tags": []any{map[string]any{"key": "env", "value": "test"}}, "terminationProtectionEnabled": true}

		createClusterOut := apitest.Run(t, apitest.Marshal(t, createClusterBody), []string{"clusters", "createCluster", "--groupId", createProject_projectId, "--version", "2023-02-01", "--watch", "--watchTimeout", "1800"})
		createClusterResp := apitest.Unmarshal(t, createClusterOut)

		createCluster_clusterName = apitest.JSONString(createClusterResp, "name")

		require.Equal(t, clusterName, apitest.JSONString(createClusterResp, "name"))

		require.Equal(t, discoverMongoDBMajorVersion_previousMongoDBMajorVersion, apitest.JSONString(createClusterResp, "mongoDBMajorVersion"))

		require.Equal(t, "30", apitest.JSONString(createClusterResp, "diskSizeGB"))

		require.Equal(t, "true", apitest.JSONString(createClusterResp, "terminationProtectionEnabled"))
	})

	apitest.Cleanup(t, []string{"clusters", "deleteCluster", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName, "--watch", "--watchTimeout", "1200"})

	// Load the sample dataset into the cluster.

	t.Run("loadSampleData", func(t *testing.T) {

		loadSampleDataOut := apitest.Run(t, "", []string{"clusters", "requestSampleDatasetLoad", "--groupId", createProject_projectId, "--name", createCluster_clusterName})
		loadSampleDataResp := apitest.Unmarshal(t, loadSampleDataOut)

		require.Equal(t, clusterName, apitest.JSONString(loadSampleDataResp, "clusterName"))
	})

	// List the project clusters.

	t.Run("listClusters", func(t *testing.T) {

		listClustersOut := apitest.Run(t, "", []string{"clusters", "listClusters", "--groupId", createProject_projectId, "--includeCount", "--itemsPerPage", "100", "--pageNum", "1"})
		listClustersResp := apitest.Unmarshal(t, listClustersOut)

		require.NotNil(t, apitest.JSONValue(listClustersResp, "results/0/id"))
	})

	// Describe the cluster.

	t.Run("describeCluster", func(t *testing.T) {

		describeClusterOut := apitest.Run(t, "", []string{"clusters", "getCluster", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName})
		describeClusterResp := apitest.Unmarshal(t, describeClusterOut)

		require.Equal(t, clusterName, apitest.JSONString(describeClusterResp, "name"))
	})

	// Read the standard and SRV connection strings from the cluster description.

	t.Run("describeConnectionString", func(t *testing.T) {

		describeConnectionStringOut := apitest.Run(t, "", []string{"clusters", "getCluster", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName})
		describeConnectionStringResp := apitest.Unmarshal(t, describeConnectionStringOut)

		require.NotNil(t, apitest.JSONValue(describeConnectionStringResp, "connectionStrings/standard"))

		require.NotNil(t, apitest.JSONValue(describeConnectionStringResp, "connectionStrings/standardSrv"))
	})

	// Set the cluster default write concern.

	t.Run("updateAdvancedSettings", func(t *testing.T) {

		updateAdvancedSettingsBody := map[string]any{"defaultWriteConcern": "majority"}

		updateAdvancedSettingsOut := apitest.Run(t, apitest.Marshal(t, updateAdvancedSettingsBody), []string{"clusters", "updateProcessArgs", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName, "--version", "2023-01-01"})
		updateAdvancedSettingsResp := apitest.Unmarshal(t, updateAdvancedSettingsOut)

		require.Equal(t, "majority", apitest.JSONString(updateAdvancedSettingsResp, "defaultWriteConcern"))
	})

	// Verify the advanced configuration values.

	t.Run("describeAdvancedSettings", func(t *testing.T) {

		describeAdvancedSettingsOut := apitest.Run(t, "", []string{"clusters", "getProcessArgs", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName})
		describeAdvancedSettingsResp := apitest.Unmarshal(t, describeAdvancedSettingsOut)

		require.Equal(t, "majority", apitest.JSONString(describeAdvancedSettingsResp, "defaultWriteConcern"))

		require.NotNil(t, apitest.JSONValue(describeAdvancedSettingsResp, "minimumEnabledTlsProtocol"))
	})

	// Create a rolling index on the cluster.

	t.Run("createRollingIndex", func(t *testing.T) {

		createRollingIndexBody := map[string]any{"collection": "tes", "db": "tes", "keys": []any{map[string]any{"name": "1"}}, "options": map[string]any{"sparse": false}}

		apitest.Run(t, apitest.Marshal(t, createRollingIndexBody), []string{"rollingIndex", "createRollingIndex", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName, "--version", "2023-01-01"})
	})

	// Verify that deletion fails while termination protection is enabled.
	// The API returns 400 for this request.

	t.Run("failDeleteWithTerminationProtection", func(t *testing.T) {

		apitest.RunExpectError(t, "", []string{"clusters", "deleteCluster", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName})
	})

	// Update disk size, MongoDB version and disable termination protection.
	// The CLI reads the cluster and PATCHes the merged body; this step sends
	// the equivalent partial PATCH.

	t.Run("updateCluster", func(t *testing.T) {

		updateClusterBody := map[string]any{"diskSizeGB": 40, "mongoDBMajorVersion": discoverMongoDBMajorVersion_mongoDBMajorVersion, "terminationProtectionEnabled": false}

		updateClusterOut := apitest.Run(t, apitest.Marshal(t, updateClusterBody), []string{"clusters", "updateCluster", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName, "--version", "2023-02-01"})
		updateClusterResp := apitest.Unmarshal(t, updateClusterOut)

		require.Equal(t, "40", apitest.JSONString(updateClusterResp, "diskSizeGB"))

		require.Equal(t, discoverMongoDBMajorVersion_mongoDBMajorVersion, apitest.JSONString(updateClusterResp, "mongoDBMajorVersion"))

		require.Equal(t, "false", apitest.JSONString(updateClusterResp, "terminationProtectionEnabled"))
	})

	// Delete the cluster.

	t.Run("deleteCluster", func(t *testing.T) {

		apitest.Run(t, "", []string{"clusters", "deleteCluster", "--groupId", createProject_projectId, "--clusterName", createCluster_clusterName, "--watch", "--watchTimeout", "1200"})
	})

	// Delete the test project.

	t.Run("deleteProject", func(t *testing.T) {

		apitest.Run(t, "", []string{"projects", "deleteGroup", "--groupId", createProject_projectId})
	})
}
