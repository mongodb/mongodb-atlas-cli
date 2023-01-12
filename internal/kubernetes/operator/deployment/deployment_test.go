// Copyright 2022 MongoDB Inc
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

///go:build unit

package deployment

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/pointers"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockAtlasOperatorClusterStore struct {
	projectToAdvancedClusters             map[string]map[string]*mongodbatlas.AdvancedCluster
	projectToClusterToProcessArgs         map[string]map[string]*mongodbatlas.ProcessArgs
	projectToClusterToSchedule            map[string]map[string]*mongodbatlas.CloudProviderSnapshotBackupPolicy
	projectToServerlessClusters           map[string]map[string]*mongodbatlas.Cluster
	projectIDToServerlessPrivateEndpoints map[string]map[string][]mongodbatlas.ServerlessPrivateEndpointConnection
	projectIDToGlobalCluster              map[string]map[string]*mongodbatlas.GlobalCluster
}

func (m *MockAtlasOperatorClusterStore) GlobalCluster(projectID string, instanceName string) (*mongodbatlas.GlobalCluster, error) {
	return m.projectIDToGlobalCluster[projectID][instanceName], nil
}

func (m *MockAtlasOperatorClusterStore) ServerlessInstances(projectID string, _ *mongodbatlas.ListOptions) (*mongodbatlas.ClustersResponse, error) {
	clusters := make([]*mongodbatlas.Cluster, 0, len(m.projectToServerlessClusters[projectID]))

	for clusterName := range m.projectToServerlessClusters[projectID] {
		clusters = append(clusters, m.projectToServerlessClusters[projectID][clusterName])
	}
	return &mongodbatlas.ClustersResponse{
		Links:      nil,
		Results:    clusters,
		TotalCount: 0,
	}, nil
}

func (m *MockAtlasOperatorClusterStore) ProjectClusters(projectID string, _ *mongodbatlas.ListOptions) (interface{}, error) {
	clusterNames := make([]string, 0, len(m.projectToAdvancedClusters[projectID])+len(m.projectToServerlessClusters[projectID]))
	for k := range m.projectToAdvancedClusters[projectID] {
		clusterNames = append(clusterNames, m.projectToAdvancedClusters[projectID][k].Name)
	}
	for k := range m.projectToServerlessClusters[projectID] {
		clusterNames = append(clusterNames, m.projectToServerlessClusters[projectID][k].Name)
	}

	clusters := make([]*mongodbatlas.AdvancedCluster, 0, len(clusterNames))
	for i := range clusterNames {
		clusters = append(clusters, &mongodbatlas.AdvancedCluster{
			Name:    clusterNames[i],
			GroupID: projectID,
		})
	}
	return &mongodbatlas.AdvancedClustersResponse{
		Links:      nil,
		Results:    clusters,
		TotalCount: 0,
	}, nil
}

func (m *MockAtlasOperatorClusterStore) ServerlessPrivateEndpoints(projectID, instanceName string, _ *mongodbatlas.ListOptions) ([]mongodbatlas.ServerlessPrivateEndpointConnection, error) {
	return m.projectIDToServerlessPrivateEndpoints[projectID][instanceName], nil
}

func (m *MockAtlasOperatorClusterStore) AtlasCluster(projectName, clusterName string) (*mongodbatlas.AdvancedCluster, error) {
	return m.projectToAdvancedClusters[projectName][clusterName], nil
}

func (m *MockAtlasOperatorClusterStore) AtlasClusterConfigurationOptions(projectName, clusterName string) (*mongodbatlas.ProcessArgs, error) {
	return m.projectToClusterToProcessArgs[projectName][clusterName], nil
}

func (m *MockAtlasOperatorClusterStore) DescribeSchedule(projectName, clusterName string) (*mongodbatlas.CloudProviderSnapshotBackupPolicy, error) {
	return m.projectToClusterToSchedule[projectName][clusterName], nil
}

func (m *MockAtlasOperatorClusterStore) ServerlessInstance(projectName, clusterName string) (*mongodbatlas.Cluster, error) {
	return m.projectToServerlessClusters[projectName][clusterName], nil
}

func TestBuildAtlasAdvancedDeployment(t *testing.T) {
	t.Run("Can import Advanced deployment", func(t *testing.T) {
		const projectName = "testProject-1"
		const clusterName = "testCluster-1"
		const targetNamespace = "test-namespace-1"
		const zoneName1 = "us-east-1"
		const zoneID1 = "TestReplicaID"
		const (
			firstLocation  = "CA"
			secondLocation = "US"
		)

		clusterStore := &MockAtlasOperatorClusterStore{
			projectToAdvancedClusters: map[string]map[string]*mongodbatlas.AdvancedCluster{
				projectName: {
					clusterName: &mongodbatlas.AdvancedCluster{
						BackupEnabled: pointers.MakePtr(true),
						BiConnector: &mongodbatlas.BiConnector{
							Enabled:        pointers.MakePtr(true),
							ReadPreference: "TestRef",
						},
						ClusterType:              "REPLICASET",
						ConnectionStrings:        nil,
						DiskSizeGB:               pointers.MakePtr[float64](20.4),
						EncryptionAtRestProvider: "TestProvider",
						GroupID:                  "TestGroupID",
						ID:                       "TestID",
						Labels: []mongodbatlas.Label{
							{
								Key:   "TestKey",
								Value: "TestValue",
							},
						},
						MongoDBMajorVersion: "5.0",
						MongoDBVersion:      "5.0",
						Name:                clusterName,
						Paused:              pointers.MakePtr(false),
						PitEnabled:          pointers.MakePtr(true),
						StateName:           "RUNNING",
						ReplicationSpecs: []*mongodbatlas.AdvancedReplicationSpec{
							{
								NumShards: 3,
								ID:        zoneID1,
								ZoneName:  zoneName1,
								RegionConfigs: []*mongodbatlas.AdvancedRegionConfig{
									{
										AnalyticsSpecs: &mongodbatlas.Specs{
											DiskIOPS:      pointers.MakePtr[int64](10),
											EbsVolumeType: "TestEBSVolume",
											InstanceSize:  "M20",
											NodeCount:     pointers.MakePtr(3),
										},
										ElectableSpecs: &mongodbatlas.Specs{
											DiskIOPS:      pointers.MakePtr[int64](10),
											EbsVolumeType: "TestEBSVolume",
											InstanceSize:  "M20",
											NodeCount:     pointers.MakePtr(3),
										},
										ReadOnlySpecs: &mongodbatlas.Specs{
											DiskIOPS:      pointers.MakePtr[int64](10),
											EbsVolumeType: "TestEBSVolume",
											InstanceSize:  "M20",
											NodeCount:     pointers.MakePtr(3),
										},
										AutoScaling: &mongodbatlas.AdvancedAutoScaling{
											DiskGB: &mongodbatlas.DiskGB{Enabled: pointers.MakePtr(true)},
											Compute: &mongodbatlas.Compute{
												Enabled:          pointers.MakePtr(true),
												ScaleDownEnabled: pointers.MakePtr(true),
												MinInstanceSize:  "M20",
												MaxInstanceSize:  "M40",
											},
										},
										BackingProviderName: "AWS",
										Priority:            pointers.MakePtr(1),
										ProviderName:        "AWS",
										RegionName:          "US_EAST_1",
									},
								},
							},
						},
						CreateDate:           "01-01-2022",
						RootCertType:         "TestRootCertType",
						VersionReleaseSystem: "TestReleaseSystem",
					},
				},
			},
			projectToClusterToProcessArgs: map[string]map[string]*mongodbatlas.ProcessArgs{
				projectName: {
					clusterName: &mongodbatlas.ProcessArgs{
						DefaultReadConcern:               "TestReadConcern",
						DefaultWriteConcern:              "TestWriteConcert",
						MinimumEnabledTLSProtocol:        "1.0",
						FailIndexKeyTooLong:              pointers.MakePtr(true),
						JavascriptEnabled:                pointers.MakePtr(true),
						NoTableScan:                      pointers.MakePtr(true),
						OplogSizeMB:                      pointers.MakePtr[int64](10),
						SampleSizeBIConnector:            pointers.MakePtr[int64](10),
						SampleRefreshIntervalBIConnector: pointers.MakePtr[int64](10),
						OplogMinRetentionHours:           pointers.MakePtr[float64](10.1),
					},
				},
			},
			projectToClusterToSchedule: map[string]map[string]*mongodbatlas.CloudProviderSnapshotBackupPolicy{
				projectName: {
					clusterName: &mongodbatlas.CloudProviderSnapshotBackupPolicy{
						ClusterID:             "testClusterID",
						ClusterName:           clusterName,
						ReferenceHourOfDay:    pointers.MakePtr[int64](5),
						ReferenceMinuteOfHour: pointers.MakePtr[int64](5),
						RestoreWindowDays:     pointers.MakePtr[int64](5),
						UpdateSnapshots:       pointers.MakePtr(true),
						NextSnapshot:          "",
						Policies: []mongodbatlas.Policy{
							{
								ID: "1",
								PolicyItems: []mongodbatlas.PolicyItem{
									{
										ID:                "1",
										FrequencyInterval: 10,
										FrequencyType:     "DAYS",
										RetentionUnit:     "WEEKS",
										RetentionValue:    1,
									},
								},
							},
						},
						AutoExportEnabled: pointers.MakePtr(true),
						Export: &mongodbatlas.Export{
							ExportBucketID: "TestBucketID",
							FrequencyType:  "TestFreqType",
						},
						UseOrgAndGroupNamesInExportPrefix: pointers.MakePtr(true),
					},
				},
			},
			projectToServerlessClusters: nil,
			projectIDToGlobalCluster: map[string]map[string]*mongodbatlas.GlobalCluster{
				projectName: {
					clusterName: &mongodbatlas.GlobalCluster{
						CustomZoneMapping: map[string]string{
							secondLocation: zoneID1,
							firstLocation:  zoneID1,
						},
						ManagedNamespaces: []mongodbatlas.ManagedNamespace{
							{
								Db:                     "testDB",
								Collection:             "testCollection",
								CustomShardKey:         "testShardKey",
								IsCustomShardKeyHashed: pointers.MakePtr(true),
								IsShardKeyUnique:       pointers.MakePtr(true),
								NumInitialChunks:       4,
								PresplitHashedZones:    pointers.MakePtr(true),
							},
						},
					},
				},
			},
		}

		cluster := clusterStore.projectToAdvancedClusters[projectName][clusterName]
		processArgs := clusterStore.projectToClusterToProcessArgs[projectName][clusterName]
		backupSchedule := clusterStore.projectToClusterToSchedule[projectName][clusterName]
		managedNamespace := clusterStore.projectIDToGlobalCluster[projectName][clusterName].ManagedNamespaces

		expectCluster := &atlasV1.AtlasDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDeployment",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(clusterName),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasDeploymentSpec{
				Project: common.ResourceRefNamespaced{
					Name:      strings.ToLower(projectName),
					Namespace: targetNamespace,
				},
				DeploymentSpec: nil,
				AdvancedDeploymentSpec: &atlasV1.AdvancedDeploymentSpec{
					BackupEnabled: cluster.BackupEnabled,
					CustomZoneMapping: []atlasV1.CustomZoneMapping{
						{
							Location: secondLocation,
							Zone:     cluster.ReplicationSpecs[0].ZoneName,
						},
						{
							Location: firstLocation,
							Zone:     cluster.ReplicationSpecs[0].ZoneName,
						},
					},
					ManagedNamespaces: []atlasV1.ManagedNamespace{
						{
							Db:                     managedNamespace[0].Db,
							Collection:             managedNamespace[0].Collection,
							CustomShardKey:         managedNamespace[0].CustomShardKey,
							IsCustomShardKeyHashed: managedNamespace[0].IsCustomShardKeyHashed,
							IsShardKeyUnique:       managedNamespace[0].IsShardKeyUnique,
							NumInitialChunks:       managedNamespace[0].NumInitialChunks,
							PresplitHashedZones:    managedNamespace[0].PresplitHashedZones,
						},
					},
					BiConnector: &atlasV1.BiConnectorSpec{
						Enabled:        cluster.BiConnector.Enabled,
						ReadPreference: cluster.BiConnector.ReadPreference,
					},
					ClusterType:              cluster.ClusterType,
					EncryptionAtRestProvider: cluster.EncryptionAtRestProvider,
					Labels: []common.LabelSpec{
						{
							Key:   cluster.Labels[0].Key,
							Value: cluster.Labels[0].Value,
						},
					},
					Name:       clusterName,
					Paused:     cluster.Paused,
					PitEnabled: cluster.PitEnabled,
					ReplicationSpecs: []*atlasV1.AdvancedReplicationSpec{
						{
							NumShards: cluster.ReplicationSpecs[0].NumShards,
							ZoneName:  cluster.ReplicationSpecs[0].ZoneName,
							RegionConfigs: []*atlasV1.AdvancedRegionConfig{
								{
									AnalyticsSpecs: &atlasV1.Specs{
										DiskIOPS:      cluster.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs.DiskIOPS,
										EbsVolumeType: cluster.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs.EbsVolumeType,
										InstanceSize:  cluster.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs.InstanceSize,
										NodeCount:     cluster.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs.NodeCount,
									},
									ElectableSpecs: &atlasV1.Specs{
										DiskIOPS:      cluster.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs.DiskIOPS,
										EbsVolumeType: cluster.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs.EbsVolumeType,
										InstanceSize:  cluster.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs.InstanceSize,
										NodeCount:     cluster.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs.NodeCount,
									},
									ReadOnlySpecs: &atlasV1.Specs{
										DiskIOPS:      cluster.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs.DiskIOPS,
										EbsVolumeType: cluster.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs.EbsVolumeType,
										InstanceSize:  cluster.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs.InstanceSize,
										NodeCount:     cluster.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs.NodeCount,
									},
									AutoScaling: &atlasV1.AdvancedAutoScalingSpec{
										DiskGB: &atlasV1.DiskGB{Enabled: cluster.ReplicationSpecs[0].RegionConfigs[0].AutoScaling.DiskGB.Enabled},
										Compute: &atlasV1.ComputeSpec{
											Enabled:          cluster.ReplicationSpecs[0].RegionConfigs[0].AutoScaling.Compute.Enabled,
											ScaleDownEnabled: cluster.ReplicationSpecs[0].RegionConfigs[0].AutoScaling.Compute.ScaleDownEnabled,
											MinInstanceSize:  cluster.ReplicationSpecs[0].RegionConfigs[0].AutoScaling.Compute.MinInstanceSize,
											MaxInstanceSize:  cluster.ReplicationSpecs[0].RegionConfigs[0].AutoScaling.Compute.MaxInstanceSize,
										},
									},
									BackingProviderName: cluster.ReplicationSpecs[0].RegionConfigs[0].BackingProviderName,
									Priority:            cluster.ReplicationSpecs[0].RegionConfigs[0].Priority,
									ProviderName:        cluster.ReplicationSpecs[0].RegionConfigs[0].ProviderName,
									RegionName:          cluster.ReplicationSpecs[0].RegionConfigs[0].RegionName,
								},
							},
						},
					},
					RootCertType:         cluster.RootCertType,
					VersionReleaseSystem: cluster.VersionReleaseSystem,
				},
				BackupScheduleRef: common.ResourceRefNamespaced{
					Name:      strings.ToLower(fmt.Sprintf("%s-backupschedule", clusterName)),
					Namespace: targetNamespace,
				},
				ServerlessSpec: nil,
				ProcessArgs: &atlasV1.ProcessArgs{
					DefaultReadConcern:               processArgs.DefaultReadConcern,
					DefaultWriteConcern:              processArgs.DefaultWriteConcern,
					MinimumEnabledTLSProtocol:        processArgs.MinimumEnabledTLSProtocol,
					FailIndexKeyTooLong:              processArgs.FailIndexKeyTooLong,
					JavascriptEnabled:                processArgs.JavascriptEnabled,
					NoTableScan:                      processArgs.NoTableScan,
					OplogSizeMB:                      processArgs.OplogSizeMB,
					SampleSizeBIConnector:            processArgs.SampleSizeBIConnector,
					SampleRefreshIntervalBIConnector: processArgs.SampleRefreshIntervalBIConnector,
				},
			},
			Status: status.AtlasDeploymentStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		expectPolicies := []*atlasV1.AtlasBackupPolicy{
			{
				TypeMeta: v1.TypeMeta{
					Kind:       "AtlasBackupPolicy",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      strings.ToLower(fmt.Sprintf("%s-backuppolicy", clusterName)),
					Namespace: targetNamespace,
				},
				Spec: atlasV1.AtlasBackupPolicySpec{
					Items: []atlasV1.AtlasBackupPolicyItem{
						{
							FrequencyType:     backupSchedule.Policies[0].PolicyItems[0].FrequencyType,
							FrequencyInterval: backupSchedule.Policies[0].PolicyItems[0].FrequencyInterval,
							RetentionUnit:     backupSchedule.Policies[0].PolicyItems[0].RetentionUnit,
							RetentionValue:    backupSchedule.Policies[0].PolicyItems[0].RetentionValue,
						},
					},
				},
				Status: atlasV1.AtlasBackupPolicyStatus{},
			},
		}

		expectSchedule := &atlasV1.AtlasBackupSchedule{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasBackupSchedule",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-backupschedule", clusterName)),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasBackupScheduleSpec{
				AutoExportEnabled: *backupSchedule.AutoExportEnabled,
				Export: &atlasV1.AtlasBackupExportSpec{
					ExportBucketID: backupSchedule.Export.ExportBucketID,
					FrequencyType:  backupSchedule.Export.FrequencyType,
				},
				PolicyRef: common.ResourceRefNamespaced{
					Name:      strings.ToLower(expectPolicies[0].Name),
					Namespace: expectPolicies[0].Namespace,
				},
				ReferenceHourOfDay:                *backupSchedule.ReferenceHourOfDay,
				ReferenceMinuteOfHour:             *backupSchedule.ReferenceMinuteOfHour,
				RestoreWindowDays:                 *backupSchedule.RestoreWindowDays,
				UpdateSnapshots:                   *backupSchedule.UpdateSnapshots,
				UseOrgAndGroupNamesInExportPrefix: *backupSchedule.UseOrgAndGroupNamesInExportPrefix,
			},
			Status: atlasV1.AtlasBackupScheduleStatus{},
		}

		expected := &AtlasDeploymentResult{
			Deployment:     expectCluster,
			BackupSchedule: expectSchedule,
			BackupPolicies: expectPolicies,
		}

		got, err := BuildAtlasAdvancedDeployment(clusterStore, projectName, projectName, clusterName, targetNamespace)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !reflect.DeepEqual(expected, got) {
			expJs, _ := json.MarshalIndent(expected, "", " ")
			gotJs, _ := json.MarshalIndent(got, "", " ")
			fmt.Printf("E:%s\r\n; G:%s\r\n", expJs, gotJs)
			t.Fatalf("Advanced deployment mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func TestBuildServerlessDeployments(t *testing.T) {
	const projectName = "testProject-2"
	const clusterName = "testCluster-2"
	const targetNamespace = "test-namespace-2"
	t.Run("Can import Serverless deployment", func(t *testing.T) {
		clusterStore := &MockAtlasOperatorClusterStore{
			projectToAdvancedClusters:     nil,
			projectToClusterToProcessArgs: nil,
			projectToClusterToSchedule:    nil,
			projectIDToServerlessPrivateEndpoints: map[string]map[string][]mongodbatlas.ServerlessPrivateEndpointConnection{
				projectName: {
					clusterName: {
						{
							ID:                           "TestPEId",
							CloudProviderEndpointID:      "TestCloudProviderID",
							Comment:                      "TestPEName",
							EndpointServiceName:          "",
							ErrorMessage:                 "",
							Status:                       "",
							ProviderName:                 "",
							PrivateEndpointIPAddress:     "",
							PrivateLinkServiceResourceID: "",
						},
					},
				},
			},
			projectToServerlessClusters: map[string]map[string]*mongodbatlas.Cluster{
				projectName: {
					clusterName: &mongodbatlas.Cluster{
						AutoScaling: &mongodbatlas.AutoScaling{
							AutoIndexingEnabled: pointers.MakePtr(true),
							Compute: &mongodbatlas.Compute{
								Enabled:          pointers.MakePtr(true),
								ScaleDownEnabled: pointers.MakePtr(true),
								MinInstanceSize:  "M20",
								MaxInstanceSize:  "M40",
							},
							DiskGBEnabled: pointers.MakePtr(true),
						},
						BackupEnabled: nil,
						BiConnector: &mongodbatlas.BiConnector{
							Enabled:        pointers.MakePtr(true),
							ReadPreference: "TestRef",
						},
						ClusterType:              "SERVERLESS",
						DiskSizeGB:               pointers.MakePtr[float64](20),
						EncryptionAtRestProvider: "TestProvider",
						Labels:                   nil,
						ID:                       "TestClusterID",
						GroupID:                  "TestGroupID",
						MongoDBVersion:           "5.0",
						MongoDBMajorVersion:      "5.0",
						MongoURI:                 "",
						MongoURIUpdated:          "",
						MongoURIWithOptions:      "",
						Name:                     clusterName,
						CreateDate:               "01-01-2021",
						NumShards:                nil,
						Paused:                   nil,
						PitEnabled:               nil,
						ProviderBackupEnabled:    nil,
						ProviderSettings: &mongodbatlas.ProviderSettings{
							BackingProviderName: "AWS",
							DiskIOPS:            nil,
							DiskTypeName:        "TestDiskName",
							EncryptEBSVolume:    nil,
							InstanceSizeName:    "M20",
							ProviderName:        "AWS",
							RegionName:          "US_EAST_1",
							VolumeType:          "",
							AutoScaling: &mongodbatlas.AutoScaling{
								AutoIndexingEnabled: pointers.MakePtr(true),
								Compute: &mongodbatlas.Compute{
									Enabled:          pointers.MakePtr(true),
									ScaleDownEnabled: pointers.MakePtr(true),
									MinInstanceSize:  "M20",
									MaxInstanceSize:  "M40",
								},
								DiskGBEnabled: pointers.MakePtr(true),
							},
						},
						ReplicationFactor:       nil,
						ReplicationSpec:         nil,
						ReplicationSpecs:        nil,
						SrvAddress:              "",
						StateName:               "",
						ServerlessBackupOptions: nil,
						ConnectionStrings:       nil,
						Links:                   nil,
						VersionReleaseSystem:    "",
					},
				},
			},
		}

		cluster := clusterStore.projectToServerlessClusters[projectName][clusterName]

		expected := &atlasV1.AtlasDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDeployment",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(cluster.Name),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasDeploymentSpec{
				Project: common.ResourceRefNamespaced{
					Name:      strings.ToLower(projectName),
					Namespace: targetNamespace,
				},
				BackupScheduleRef: common.ResourceRefNamespaced{},
				ServerlessSpec: &atlasV1.ServerlessSpec{
					Name: cluster.Name,
					ProviderSettings: &atlasV1.ProviderSettingsSpec{
						BackingProviderName: cluster.ProviderSettings.BackingProviderName,
						DiskIOPS:            cluster.ProviderSettings.DiskIOPS,
						DiskTypeName:        cluster.ProviderSettings.DiskTypeName,
						EncryptEBSVolume:    cluster.ProviderSettings.EncryptEBSVolume,
						InstanceSizeName:    cluster.ProviderSettings.InstanceSizeName,
						ProviderName:        provider.ProviderName(cluster.ProviderSettings.ProviderName),
						RegionName:          cluster.ProviderSettings.RegionName,
						VolumeType:          cluster.ProviderSettings.VolumeType,
						AutoScaling: &atlasV1.AutoScalingSpec{
							AutoIndexingEnabled: cluster.ProviderSettings.AutoScaling.AutoIndexingEnabled,
							DiskGBEnabled:       cluster.ProviderSettings.AutoScaling.DiskGBEnabled,
							Compute: &atlasV1.ComputeSpec{
								Enabled:          cluster.ProviderSettings.AutoScaling.Compute.Enabled,
								ScaleDownEnabled: cluster.ProviderSettings.AutoScaling.Compute.ScaleDownEnabled,
								MinInstanceSize:  cluster.ProviderSettings.AutoScaling.Compute.MinInstanceSize,
								MaxInstanceSize:  cluster.ProviderSettings.AutoScaling.Compute.MaxInstanceSize,
							},
						},
					},
					PrivateEndpoints: []atlasV1.ServerlessPrivateEndpoint{
						{
							Name:                     clusterStore.projectIDToServerlessPrivateEndpoints[projectName][clusterName][0].Comment,
							CloudProviderEndpointID:  clusterStore.projectIDToServerlessPrivateEndpoints[projectName][clusterName][0].CloudProviderEndpointID,
							PrivateEndpointIPAddress: clusterStore.projectIDToServerlessPrivateEndpoints[projectName][clusterName][0].PrivateEndpointIPAddress,
						},
					},
				},
				ProcessArgs: nil,
			},
			Status: status.AtlasDeploymentStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		got, err := BuildServerlessDeployments(clusterStore, projectName, projectName, clusterName, targetNamespace)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Serverless deployment mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}
