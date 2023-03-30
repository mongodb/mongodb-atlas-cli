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

//go:build unit

package deployment

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/atlas/mongodbatlasv2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const resourceVersion = "x.y.z"

func TestBuildAtlasAdvancedDeployment(t *testing.T) {
	ctl := gomock.NewController(t)
	clusterStore := mocks.NewMockAtlasOperatorClusterStore(ctl)
	dictionary := resources.AtlasNameToKubernetesName()
	featureValidator := mocks.NewMockFeatureValidator(ctl)

	t.Run("Can import Advanced deployment", func(t *testing.T) {
		const projectName = "testProject-1"
		const clusterName = "testCluster-1"
		const targetNamespace = "test-namespace-1"
		const zoneName1 = "us-east-1"
		const zoneID1 = "TestReplicaID"
		const (
			firstLocation = "CA"
		)

		cluster := &mongodbatlas.AdvancedCluster{
			BackupEnabled: pointer.Get(true),
			BiConnector: &mongodbatlas.BiConnector{
				Enabled:        pointer.Get(true),
				ReadPreference: "TestRef",
			},
			ClusterType:              "REPLICASET",
			ConnectionStrings:        nil,
			DiskSizeGB:               pointer.Get[float64](20.4),
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
			Paused:              pointer.Get(false),
			PitEnabled:          pointer.Get(true),
			StateName:           "RUNNING",
			ReplicationSpecs: []*mongodbatlas.AdvancedReplicationSpec{
				{
					NumShards: 3,
					ID:        zoneID1,
					ZoneName:  zoneName1,
					RegionConfigs: []*mongodbatlas.AdvancedRegionConfig{
						{
							AnalyticsSpecs: &mongodbatlas.Specs{
								DiskIOPS:      pointer.Get[int64](10),
								EbsVolumeType: "TestEBSVolume",
								InstanceSize:  "M20",
								NodeCount:     pointer.Get(3),
							},
							ElectableSpecs: &mongodbatlas.Specs{
								DiskIOPS:      pointer.Get[int64](10),
								EbsVolumeType: "TestEBSVolume",
								InstanceSize:  "M20",
								NodeCount:     pointer.Get(3),
							},
							ReadOnlySpecs: &mongodbatlas.Specs{
								DiskIOPS:      pointer.Get[int64](10),
								EbsVolumeType: "TestEBSVolume",
								InstanceSize:  "M20",
								NodeCount:     pointer.Get(3),
							},
							AutoScaling: &mongodbatlas.AdvancedAutoScaling{
								DiskGB: &mongodbatlas.DiskGB{Enabled: pointer.Get(true)},
								Compute: &mongodbatlas.Compute{
									Enabled:          pointer.Get(true),
									ScaleDownEnabled: pointer.Get(true),
									MinInstanceSize:  "M20",
									MaxInstanceSize:  "M40",
								},
							},
							BackingProviderName: "AWS",
							Priority:            pointer.Get(1),
							ProviderName:        "AWS",
							RegionName:          "US_EAST_1",
						},
					},
				},
			},
			CreateDate:           "01-01-2022",
			RootCertType:         "TestRootCertType",
			VersionReleaseSystem: "TestReleaseSystem",
		}
		processArgs := &mongodbatlas.ProcessArgs{
			DefaultReadConcern:               "TestReadConcern",
			DefaultWriteConcern:              "TestWriteConcert",
			MinimumEnabledTLSProtocol:        "1.0",
			FailIndexKeyTooLong:              pointer.Get(true),
			JavascriptEnabled:                pointer.Get(true),
			NoTableScan:                      pointer.Get(true),
			OplogSizeMB:                      pointer.Get[int64](10),
			SampleSizeBIConnector:            pointer.Get[int64](10),
			SampleRefreshIntervalBIConnector: pointer.Get[int64](10),
			OplogMinRetentionHours:           pointer.Get[float64](10.1),
		}
		backupSchedule := &mongodbatlas.CloudProviderSnapshotBackupPolicy{
			ClusterID:             "testClusterID",
			ClusterName:           clusterName,
			ReferenceHourOfDay:    pointer.Get[int64](5),
			ReferenceMinuteOfHour: pointer.Get[int64](5),
			RestoreWindowDays:     pointer.Get[int64](5),
			UpdateSnapshots:       pointer.Get(true),
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
			AutoExportEnabled: pointer.Get(true),
			Export: &mongodbatlas.Export{
				ExportBucketID: "TestBucketID",
				FrequencyType:  "TestFreqType",
			},
			UseOrgAndGroupNamesInExportPrefix: pointer.Get(true),
			CopySettings: []mongodbatlas.CopySetting{
				{
					CloudProvider:     pointer.Get("AWS"),
					RegionName:        pointer.Get("US_EAST_1"),
					ReplicationSpecID: pointer.Get("123456"),
					ShouldCopyOplogs:  pointer.Get(false),
					Frequencies:       []string{"DAILY"},
				},
			},
		}
		globalCluster := &mongodbatlas.GlobalCluster{
			CustomZoneMapping: map[string]string{
				firstLocation: zoneID1,
			},
			ManagedNamespaces: []mongodbatlas.ManagedNamespace{
				{
					Db:                     "testDB",
					Collection:             "testCollection",
					CustomShardKey:         "testShardKey",
					IsCustomShardKeyHashed: pointer.Get(true),
					IsShardKeyUnique:       pointer.Get(true),
					NumInitialChunks:       4,
					PresplitHashedZones:    pointer.Get(true),
				},
			},
		}

		managedNamespace := globalCluster.ManagedNamespaces

		clusterStore.EXPECT().AtlasCluster(projectName, clusterName).Return(cluster, nil)
		clusterStore.EXPECT().AtlasClusterConfigurationOptions(projectName, clusterName).Return(processArgs, nil)
		clusterStore.EXPECT().GlobalCluster(projectName, clusterName).Return(globalCluster, nil)
		clusterStore.EXPECT().DescribeSchedule(projectName, clusterName).Return(backupSchedule, nil)

		expectCluster := &atlasV1.AtlasDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDeployment",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(clusterName),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
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
					Labels: map[string]string{
						features.ResourceVersion: resourceVersion,
					},
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
				Status: status.BackupPolicyStatus{},
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
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
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
				CopySettings: []atlasV1.CopySetting{
					{
						CloudProvider:     pointer.Get("AWS"),
						RegionName:        pointer.Get("US_EAST_1"),
						ReplicationSpecID: pointer.Get("123456"),
						ShouldCopyOplogs:  pointer.Get(false),
						Frequencies:       []string{"DAILY"},
					},
				},
			},
			Status: status.BackupScheduleStatus{},
		}

		expected := &AtlasDeploymentResult{
			Deployment:     expectCluster,
			BackupSchedule: expectSchedule,
			BackupPolicies: expectPolicies,
		}

		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasDeployment, featureProcessArgs).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasDeployment, featureBackupSchedule).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasDeployment, featureGlobalDeployments).Return(true)

		got, err := BuildAtlasAdvancedDeployment(clusterStore, featureValidator, projectName, projectName, clusterName, targetNamespace, dictionary, resourceVersion)
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

	ctl := gomock.NewController(t)
	clusterStore := mocks.NewMockAtlasOperatorClusterStore(ctl)
	dictionary := resources.AtlasNameToKubernetesName()

	featureValidator := mocks.NewMockFeatureValidator(ctl)

	t.Run("Can import Serverless deployment", func(t *testing.T) {
		speID := "TestPEId"
		speCloudProviderEndpointID := "TestCloudProviderID"
		speComment := "TestPEName"
		spePrivateEndpointIPAddress := ""

		spe := []mongodbatlasv2.ServerlessTenantEndpoint{
			mongodbatlasv2.ServerlessAzureTenantEndpointAsServerlessTenantEndpoint(
				&mongodbatlasv2.ServerlessAzureTenantEndpoint{
					Id:                       &speID,
					CloudProviderEndpointId:  &speCloudProviderEndpointID,
					Comment:                  &speComment,
					PrivateEndpointIpAddress: &spePrivateEndpointIPAddress,
				},
			),
		}

		cluster := &mongodbatlas.Cluster{
			AutoScaling: &mongodbatlas.AutoScaling{
				AutoIndexingEnabled: pointer.Get(true),
				Compute: &mongodbatlas.Compute{
					Enabled:          pointer.Get(true),
					ScaleDownEnabled: pointer.Get(true),
					MinInstanceSize:  "M20",
					MaxInstanceSize:  "M40",
				},
				DiskGBEnabled: pointer.Get(true),
			},
			BackupEnabled: nil,
			BiConnector: &mongodbatlas.BiConnector{
				Enabled:        pointer.Get(true),
				ReadPreference: "TestRef",
			},
			ClusterType:              "SERVERLESS",
			DiskSizeGB:               pointer.Get[float64](20),
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
					AutoIndexingEnabled: pointer.Get(true),
					Compute: &mongodbatlas.Compute{
						Enabled:          pointer.Get(true),
						ScaleDownEnabled: pointer.Get(true),
						MinInstanceSize:  "M20",
						MaxInstanceSize:  "M40",
					},
					DiskGBEnabled: pointer.Get(true),
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
		}

		clusterStore.EXPECT().ServerlessInstance(projectName, clusterName).Return(cluster, nil)
		clusterStore.EXPECT().ServerlessPrivateEndpoints(projectName, clusterName).Return(spe, nil)

		expected := &atlasV1.AtlasDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDeployment",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(cluster.Name),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
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
							Name:                     speComment,
							CloudProviderEndpointID:  speCloudProviderEndpointID,
							PrivateEndpointIPAddress: spePrivateEndpointIPAddress,
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

		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasDeployment, featureServerlessPrivateEndpoints).Return(true)

		got, err := BuildServerlessDeployments(clusterStore, featureValidator, projectName, projectName, clusterName, targetNamespace, dictionary, resourceVersion)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Serverless deployment mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}
