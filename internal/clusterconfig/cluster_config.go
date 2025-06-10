// Copyright 2020 MongoDB Inc
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
package clusterconfig

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312003/admin"
)

const (
	ClusterWideScalingFlag       = "clusterWideScaling"
	IndependentShardScalingShard = "independentShardScaling"
	atlasM0                      = "M0"
	atlasM2                      = "M2"
	atlasM5                      = "M5"
)

// SetTags sets the tags for a cluster of the pinned API version.
func SetTags(cluster *atlasClustersPinned.AdvancedClusterDescription, providedTags map[string]string) {
	if len(providedTags) < 0 {
		return
	}

	// Cover the case where the cluster already has tags
	var tags []atlasClustersPinned.ResourceTag
	if cluster.Tags != nil {
		tags = *cluster.Tags
	}

	for k, v := range providedTags {
		if k != "" && v != "" {
			tags = append(tags, atlasClustersPinned.ResourceTag{Key: k, Value: v})
		}
	}
	cluster.Tags = &tags
}

// SetTagsLatest sets the tags for a cluster of the latest API version.
func SetTagsLatest(cluster *atlasv2.ClusterDescription20240805, providedTags map[string]string) {
	if len(providedTags) < 0 {
		return
	}

	// Cover the case where the cluster already has tags
	var tags []atlasv2.ResourceTag
	if cluster.Tags != nil {
		tags = *cluster.Tags
	}

	for k, v := range providedTags {
		if k != "" && v != "" {
			tags = append(tags, atlasv2.ResourceTag{Key: k, Value: v})
		}
	}
	cluster.Tags = &tags
}

const (
	priority = 7
	tenant   = "TENANT"
)

func NewAdvancedRegionConfig(providerName, region, tier, provider string, members int) atlasClustersPinned.CloudRegionConfig {
	regionConfig := atlasClustersPinned.CloudRegionConfig{
		Priority:     pointer.Get(priority),
		RegionName:   &region,
		ProviderName: &providerName,
		ElectableSpecs: &atlasClustersPinned.HardwareSpec{
			InstanceSize: &tier,
		},
	}

	if providerName == tenant {
		regionConfig.BackingProviderName = &provider
	} else {
		regionConfig.ElectableSpecs.NodeCount = &members
	}

	return regionConfig
}
