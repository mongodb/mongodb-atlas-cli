// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAdvancedRegionConfigs_AreEqual(t *testing.T) {
	testCases := []struct {
		name     string
		provider string
		region   string
		tier     string
	}{
		{
			name:     "AWS M10 cluster",
			provider: "AWS",
			region:   "us-west-2",
			tier:     "M10",
		},
		{
			name:     "GCP M20 cluster",
			provider: "GCP",
			region:   "us-central1",
			tier:     "M20",
		},
		{
			name:     "Azure M30 cluster",
			provider: "AZURE",
			region:   "eastus",
			tier:     "M30",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opts := &Opts{
				Provider: tc.provider,
				Region:   tc.region,
				Tier:     tc.tier,
			}

			pinnedConfig := opts.newAdvancedRegionConfig()
			latestConfig := opts.newAdvancedRegionConfigLatest()

			// Verify both methods produce equivalent basic configurations
			assert.Equal(t, *pinnedConfig.ProviderName, *latestConfig.ProviderName)
			assert.Equal(t, *pinnedConfig.RegionName, *latestConfig.RegionName)
			assert.Equal(t, *pinnedConfig.Priority, *latestConfig.Priority)
			assert.Equal(t, *pinnedConfig.ElectableSpecs.InstanceSize, *latestConfig.ElectableSpecs.InstanceSize)
		})
	}
}

func TestNewAdvancedRegionConfig_TenantVsNonTenant(t *testing.T) {
	testCases := []struct {
		name                string
		tier                string
		expectTenant        bool
		shouldHaveNodeCount bool
	}{
		{
			name:                "M2 should be tenant",
			tier:                "M2",
			expectTenant:        true,
			shouldHaveNodeCount: false,
		},
		{
			name:                "M5 should be tenant",
			tier:                "M5",
			expectTenant:        true,
			shouldHaveNodeCount: false,
		},
		{
			name:                "M10 should be non-tenant",
			tier:                "M10",
			expectTenant:        false,
			shouldHaveNodeCount: true,
		},
		{
			name:                "M20 should be non-tenant",
			tier:                "M20",
			expectTenant:        false,
			shouldHaveNodeCount: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opts := &Opts{
				Provider: "AWS",
				Region:   "us-east-1",
				Tier:     tc.tier,
			}

			pinnedConfig := opts.newAdvancedRegionConfig()
			latestConfig := opts.newAdvancedRegionConfigLatest()

			// Both methods should handle tenant logic identically
			if tc.expectTenant {
				// Tenant clusters use TENANT provider and have backing provider
				assert.Equal(t, "TENANT", *pinnedConfig.ProviderName)
				assert.Equal(t, "TENANT", *latestConfig.ProviderName)
				assert.Equal(t, "AWS", *pinnedConfig.BackingProviderName)
				assert.Equal(t, "AWS", *latestConfig.BackingProviderName)
			} else {
				// Non-tenant clusters use actual provider, no backing provider
				assert.Equal(t, "AWS", *pinnedConfig.ProviderName)
				assert.Equal(t, "AWS", *latestConfig.ProviderName)
				assert.Nil(t, pinnedConfig.BackingProviderName)
				assert.Nil(t, latestConfig.BackingProviderName)
			}

			// Node count behavior should match
			if tc.shouldHaveNodeCount {
				assert.NotNil(t, pinnedConfig.ElectableSpecs.NodeCount)
				assert.NotNil(t, latestConfig.ElectableSpecs.NodeCount)
				assert.Equal(t, defaultNodeCount, *pinnedConfig.ElectableSpecs.NodeCount)
				assert.Equal(t, defaultNodeCount, *latestConfig.ElectableSpecs.NodeCount)
			} else {
				assert.Nil(t, pinnedConfig.ElectableSpecs.NodeCount)
				assert.Nil(t, latestConfig.ElectableSpecs.NodeCount)
			}
		})
	}
}

func TestNewAdvancedRegionConfig_DefaultValues(t *testing.T) {
	opts := &Opts{
		Provider: "AWS",
		Region:   "us-east-1",
		Tier:     "M10",
	}

	pinnedConfig := opts.newAdvancedRegionConfig()
	latestConfig := opts.newAdvancedRegionConfigLatest()

	// Both should set default priority
	assert.Equal(t, defaultPriority, *pinnedConfig.Priority)
	assert.Equal(t, defaultPriority, *latestConfig.Priority)

	// Both should set the tier as instance size
	assert.Equal(t, "M10", *pinnedConfig.ElectableSpecs.InstanceSize)
	assert.Equal(t, "M10", *latestConfig.ElectableSpecs.InstanceSize)
}

func TestNewClusters_AreEqual(t *testing.T) {
	testCases := []struct {
		name string
		opts *Opts
	}{
		{
			name: "Basic non-tenant cluster",
			opts: &Opts{
				ClusterName:                 "test-cluster",
				Provider:                    "AWS",
				Region:                      "us-east-1",
				Tier:                        "M10",
				EnableTerminationProtection: true,
			},
		},
		{
			name: "Tenant cluster",
			opts: &Opts{
				ClusterName:                 "tenant-cluster",
				Provider:                    "GCP",
				Region:                      "us-central1",
				Tier:                        "M2",
				EnableTerminationProtection: false,
			},
		},
		{
			name: "Cluster with version and tags",
			opts: &Opts{
				ClusterName:                 "tagged-cluster",
				Provider:                    "AZURE",
				Region:                      "eastus",
				Tier:                        "M20",
				EnableTerminationProtection: true,
				MDBVersion:                  "7.0",
				Tag: map[string]string{
					"env":  "test",
					"team": "platform",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pinnedCluster := tc.opts.newCluster()
			latestCluster := tc.opts.newClusterLatest()

			// Basic cluster properties should match
			assert.Equal(t, *pinnedCluster.Name, *latestCluster.Name)
			assert.Equal(t, *pinnedCluster.GroupId, *latestCluster.GroupId)
			assert.Equal(t, *pinnedCluster.ClusterType, *latestCluster.ClusterType)
			assert.Equal(t, *pinnedCluster.TerminationProtectionEnabled, *latestCluster.TerminationProtectionEnabled)

			// Both should have one replication spec
			require.NotNil(t, pinnedCluster.ReplicationSpecs)
			require.NotNil(t, latestCluster.ReplicationSpecs)
			require.Len(t, *pinnedCluster.ReplicationSpecs, 1)
			require.Len(t, *latestCluster.ReplicationSpecs, 1)

			pinnedSpec := (*pinnedCluster.ReplicationSpecs)[0]
			latestSpec := (*latestCluster.ReplicationSpecs)[0]

			// Replication specs should have equivalent configuration
			require.NotNil(t, pinnedSpec.RegionConfigs)
			require.NotNil(t, latestSpec.RegionConfigs)
			require.Len(t, *pinnedSpec.RegionConfigs, 1)
			require.Len(t, *latestSpec.RegionConfigs, 1)

			pinnedRegion := (*pinnedSpec.RegionConfigs)[0]
			latestRegion := (*latestSpec.RegionConfigs)[0]

			// Region configurations should be equivalent
			assert.Equal(t, *pinnedRegion.ProviderName, *latestRegion.ProviderName)
			assert.Equal(t, *pinnedRegion.RegionName, *latestRegion.RegionName)
			assert.Equal(t, *pinnedRegion.Priority, *latestRegion.Priority)

			// ElectableSpecs should be equivalent
			require.NotNil(t, pinnedRegion.ElectableSpecs)
			require.NotNil(t, latestRegion.ElectableSpecs)
			assert.Equal(t, *pinnedRegion.ElectableSpecs.InstanceSize, *latestRegion.ElectableSpecs.InstanceSize)

			// Node count and backing provider logic should match
			if pinnedRegion.BackingProviderName != nil {
				require.NotNil(t, latestRegion.BackingProviderName)
				assert.Equal(t, *pinnedRegion.BackingProviderName, *latestRegion.BackingProviderName)
				assert.Nil(t, pinnedRegion.ElectableSpecs.NodeCount)
				assert.Nil(t, latestRegion.ElectableSpecs.NodeCount)
			} else {
				assert.Nil(t, latestRegion.BackingProviderName)
				require.NotNil(t, pinnedRegion.ElectableSpecs.NodeCount)
				require.NotNil(t, latestRegion.ElectableSpecs.NodeCount)
				assert.Equal(t, *pinnedRegion.ElectableSpecs.NodeCount, *latestRegion.ElectableSpecs.NodeCount)
			}

			// Disk size handling - different locations but same logical value
			if tc.opts.providerName() != "TENANT" {
				// Non-tenant clusters should have disk size
				expectedDiskSize := defaultDiskSizeGB(tc.opts.providerName(), tc.opts.Tier)

				// Pinned version stores at cluster level
				require.NotNil(t, pinnedCluster.DiskSizeGB)
				assert.InEpsilon(t, expectedDiskSize, *pinnedCluster.DiskSizeGB, 0.000001)

				// Latest version stores in ElectableSpecs
				require.NotNil(t, latestRegion.ElectableSpecs.DiskSizeGB)
				assert.InEpsilon(t, expectedDiskSize, *latestRegion.ElectableSpecs.DiskSizeGB, 0.000001)
			}

			// MongoDB version handling
			if tc.opts.MDBVersion != "" {
				require.NotNil(t, pinnedCluster.MongoDBMajorVersion)
				require.NotNil(t, latestCluster.MongoDBMajorVersion)
				assert.Equal(t, *pinnedCluster.MongoDBMajorVersion, *latestCluster.MongoDBMajorVersion)
				assert.Equal(t, tc.opts.MDBVersion, *pinnedCluster.MongoDBMajorVersion)
			}

			// Tags handling
			if len(tc.opts.Tag) > 0 {
				require.NotNil(t, pinnedCluster.Tags)
				require.NotNil(t, latestCluster.Tags)
				assert.Len(t, *pinnedCluster.Tags, len(tc.opts.Tag))
				assert.Len(t, *latestCluster.Tags, len(tc.opts.Tag))

				// Convert tags to maps for easier comparison
				pinnedTagsMap := make(map[string]string)
				for _, tag := range *pinnedCluster.Tags {
					pinnedTagsMap[tag.Key] = tag.Value
				}

				latestTagsMap := make(map[string]string)
				for _, tag := range *latestCluster.Tags {
					latestTagsMap[tag.Key] = tag.Value
				}

				assert.Equal(t, pinnedTagsMap, latestTagsMap)
				assert.Equal(t, tc.opts.Tag, pinnedTagsMap)
			}
		})
	}
}
