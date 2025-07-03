// Copyright 2025 MongoDB Inc
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

package setup

import (
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	defaultNodeCount = 3
	defaultPriority  = 7
)

func defaultDiskSizeGB(provider, tier string) float64 {
	return atlas.DefaultDiskSizeGB[strings.ToUpper(provider)][tier]
}

// newCluster creates a new cluster for the pinned API version.
func (opts *Opts) newCluster() *atlasClustersPinned.AdvancedClusterDescription {
	cluster := &atlasClustersPinned.AdvancedClusterDescription{
		GroupId:                      pointer.Get(opts.ConfigProjectID()),
		ClusterType:                  pointer.Get(replicaSet),
		Name:                         &opts.ClusterName,
		TerminationProtectionEnabled: &opts.EnableTerminationProtection,
		ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
			{
				NumShards: pointer.Get(1),
				ZoneName:  pointer.Get("Zone 1"),
				RegionConfigs: &[]atlasClustersPinned.CloudRegionConfig{
					opts.newAdvancedRegionConfig(),
				},
			},
		},
	}

	opts.setTags(cluster)

	if diskSizeGB := opts.getDiskSizeOverride(); diskSizeGB != nil {
		cluster.DiskSizeGB = diskSizeGB
	}
	if version := opts.getVersionOverride(); version != nil {
		cluster.MongoDBMajorVersion = version
	}

	return cluster
}

// setTags sets the tags for a cluster of the pinned API version.
func (opts *Opts) setTags(cluster *atlasClustersPinned.AdvancedClusterDescription) {
	if len(opts.Tag) > 0 {
		var tags []atlasClustersPinned.ResourceTag
		for k, v := range opts.Tag {
			if k != "" && v != "" {
				tags = append(tags, atlasClustersPinned.ResourceTag{Key: k, Value: v})
			}
		}
		cluster.Tags = &tags
	}
}

// newAdvancedRegionConfig creates a new advanced region config for the pinned API version.
func (opts *Opts) newAdvancedRegionConfig() atlasClustersPinned.CloudRegionConfig {
	providerName := opts.providerName()

	regionConfig := atlasClustersPinned.CloudRegionConfig{
		ProviderName: &providerName,
		Priority:     pointer.Get(defaultPriority),
		RegionName:   &opts.Region,
	}

	regionConfig.ElectableSpecs = &atlasClustersPinned.HardwareSpec{
		InstanceSize: &opts.Tier,
	}

	if providerName == tenant {
		regionConfig.BackingProviderName = &opts.Provider
	} else {
		regionConfig.ElectableSpecs.NodeCount = pointer.Get(defaultNodeCount)
	}

	return regionConfig
}

// getDiskSizeOverride returns the disk size override, defaults to nil if the provider is tenant or the disk size is 0.
func (opts *Opts) getDiskSizeOverride() *float64 {
	if opts.providerName() == tenant {
		return nil
	}

	diskSizeGB := defaultDiskSizeGB(opts.providerName(), opts.Tier)
	if diskSizeGB == 0 {
		return nil
	}

	return &diskSizeGB
}

// getVersionOverride returns the MongoDB major version override which is only applicable for non-tenant clusters.
// if the user has not provided a version override, the default MongoDB major version is returned.
func (opts *Opts) getVersionOverride() *string {
	if opts.providerName() == tenant {
		return nil
	}

	if opts.MDBVersion != "" {
		return &opts.MDBVersion
	}

	if mdbVersion, err := cli.DefaultMongoDBMajorVersion(); err == nil && mdbVersion != "" {
		return &mdbVersion
	}

	return nil
}

// newClusterLatest creates a new cluster for the latest API version.
func (opts *Opts) newClusterLatest() *atlasv2.ClusterDescription20240805 {
	cluster := &atlasv2.ClusterDescription20240805{
		GroupId:                      pointer.Get(opts.ConfigProjectID()),
		ClusterType:                  pointer.Get(replicaSet),
		Name:                         pointer.Get(opts.ClusterName),
		TerminationProtectionEnabled: &opts.EnableTerminationProtection,
		ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
			{
				ZoneName: pointer.Get("Zone 1"),
				RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
					opts.newAdvancedRegionConfigLatest(),
				},
			},
		},
	}

	opts.setTagsLatest(cluster)

	version := opts.getVersionOverride()
	if version != nil {
		cluster.MongoDBMajorVersion = version
	}

	return cluster
}

// newAdvancedRegionConfigLatest creates a new advanced region config for the latest API version.
func (opts *Opts) newAdvancedRegionConfigLatest() atlasv2.CloudRegionConfig20240805 {
	regionConfig := atlasv2.CloudRegionConfig20240805{
		ProviderName: pointer.Get(opts.providerName()),
		Priority:     pointer.Get(defaultPriority),
		RegionName:   pointer.Get(opts.Region),
		ElectableSpecs: &atlasv2.HardwareSpec20240805{
			InstanceSize: pointer.Get(opts.Tier),
		},
	}

	if opts.providerName() == tenant {
		regionConfig.BackingProviderName = &opts.Provider
	} else {
		regionConfig.ElectableSpecs.NodeCount = pointer.Get(defaultNodeCount)
	}

	// diskSize is now a field in the HardwareSpec20240805 struct
	diskSizeGB := opts.getDiskSizeOverride()
	if diskSizeGB != nil {
		regionConfig.ElectableSpecs.DiskSizeGB = diskSizeGB
	}

	return regionConfig
}

// setTagsLatest sets the tags for a cluster of the latest API version.
func (opts *Opts) setTagsLatest(cluster *atlasv2.ClusterDescription20240805) {
	if len(opts.Tag) > 0 {
		var tags []atlasv2.ResourceTag
		for k, v := range opts.Tag {
			if k != "" && v != "" {
				tags = append(tags, atlasv2.ResourceTag{Key: k, Value: v})
			}
		}
		cluster.Tags = &tags
	}
}
