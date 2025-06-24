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

//go:build unit

package clusters

import (
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

func TestRemoveReadOnlyAttributes(t *testing.T) {
	var (
		id        = "Test"
		testVar   = "test"
		specID    = "22"
		shards    = 2
		zone      = "1"
		timeStamp = time.Now()
	)
	tests := []struct {
		name string
		args *atlasClustersPinned.AdvancedClusterDescription
		want *atlasClustersPinned.AdvancedClusterDescription
	}{
		{
			name: "One AdvancedReplicationSpec",
			args: &atlasClustersPinned.AdvancedClusterDescription{
				Id:             &id,
				MongoDBVersion: &testVar,
				StateName:      &testVar,
				ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
				CreateDate: &timeStamp,
			},
			want: &atlasClustersPinned.AdvancedClusterDescription{
				ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
		},
		{
			name: "More AdvancedReplicationSpecs",
			args: &atlasClustersPinned.AdvancedClusterDescription{
				Id:             &id,
				MongoDBVersion: &testVar,
				StateName:      &testVar,
				ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
				CreateDate: &timeStamp,
			},
			want: &atlasClustersPinned.AdvancedClusterDescription{
				ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
		},
		{
			name: "Nothing to remove",
			args: &atlasClustersPinned.AdvancedClusterDescription{
				ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
			want: &atlasClustersPinned.AdvancedClusterDescription{
				ReplicationSpecs: &[]atlasClustersPinned.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		name := tt.name
		arg := tt.args
		want := tt.want
		t.Run(name, func(t *testing.T) {
			removeReadOnlyAttributes(arg)
			if diff := deep.Equal(arg, want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestRemoveReadOnlyAttributesLatest(t *testing.T) {
	var (
		id           = "Test"
		testVar      = "test"
		specID       = "22"
		diskSizeGB   = 30.0
		priority     = 7
		providerName = "AWS"
		regionName   = "US_EAST_1"
		timeStamp    = time.Now()
	)
	tests := []struct {
		name string
		args *atlasv2.ClusterDescription20240805
		want *atlasv2.ClusterDescription20240805
	}{
		{
			name: "One ReplicationSpec",
			args: &atlasv2.ClusterDescription20240805{
				Id:             &id,
				MongoDBVersion: &testVar,
				StateName:      &testVar,
				ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
					{
						Id: &specID,
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
								Priority:     &priority,
								ProviderName: &providerName,
								RegionName:   &regionName,
							},
						},
					},
				},
				CreateDate: &timeStamp,
			},
			want: &atlasv2.ClusterDescription20240805{
				ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
					{
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
								Priority:     &priority,
								ProviderName: &providerName,
								RegionName:   &regionName,
							},
						},
					},
				},
			},
		},
		{
			name: "More ReplicationSpecs",
			args: &atlasv2.ClusterDescription20240805{
				Id:             &id,
				MongoDBVersion: &testVar,
				StateName:      &testVar,
				ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
					{
						Id: &specID,
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
								Priority:     &priority,
								ProviderName: &providerName,
								RegionName:   &regionName,
							},
						},
					},
					{
						Id: &specID,
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
							},
						},
					},
				},
				CreateDate: &timeStamp,
			},
			want: &atlasv2.ClusterDescription20240805{
				ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
					{
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
								Priority:     &priority,
								ProviderName: &providerName,
								RegionName:   &regionName,
							},
						},
					},
					{
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Tenant cluster",
			args: &atlasv2.ClusterDescription20240805{
				Id: &id,
				ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
					{
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: &diskSizeGB,
								},
								Priority:     &priority,
								ProviderName: pointer.Get(tenant),
								RegionName:   &regionName,
							},
						},
					},
				},
			},
			want: &atlasv2.ClusterDescription20240805{
				ReplicationSpecs: &[]atlasv2.ReplicationSpec20240805{
					{
						RegionConfigs: &[]atlasv2.CloudRegionConfig20240805{
							{
								ElectableSpecs: &atlasv2.HardwareSpec20240805{
									DiskSizeGB: nil,
								},
								Priority:     &priority,
								ProviderName: pointer.Get(tenant),
								RegionName:   &regionName,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		name := tt.name
		arg := tt.args
		want := tt.want
		t.Run(name, func(t *testing.T) {
			removeReadOnlyAttributesLatest(arg)
			if diff := deep.Equal(arg, want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
