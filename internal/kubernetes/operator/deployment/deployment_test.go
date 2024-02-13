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
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	atlasmocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2provider "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/provider"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115006/admin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const resourceVersion = "x.y.z"

func TestBuildAtlasAdvancedDeployment(t *testing.T) {
	ctl := gomock.NewController(t)
	clusterStore := atlasmocks.NewMockOperatorClusterStore(ctl)
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

		cluster := &atlasv2.AdvancedClusterDescription{
			BackupEnabled: pointer.Get(true),
			BiConnector: &atlasv2.BiConnector{
				Enabled:        pointer.Get(true),
				ReadPreference: pointer.Get("TestRef"),
			},
			ClusterType:              pointer.Get("REPLICASET"),
			ConnectionStrings:        nil,
			DiskSizeGB:               pointer.Get[float64](20.4),
			EncryptionAtRestProvider: pointer.Get("TestProvider"),
			GroupId:                  pointer.Get("TestGroupID"),
			Id:                       pointer.Get("TestID"),
			Labels: &[]atlasv2.ComponentLabel{
				{
					Key:   pointer.Get("TestKey"),
					Value: pointer.Get("TestValue"),
				},
			},
			Tags: &[]atlasv2.ResourceTag{
				{
					Key:   pointer.Get("TestTagKey"),
					Value: pointer.Get("TestTagValue"),
				},
			},
			MongoDBMajorVersion: pointer.Get("5.0"),
			MongoDBVersion:      pointer.Get("5.0"),
			Name:                pointer.Get(clusterName),
			Paused:              pointer.Get(false),
			PitEnabled:          pointer.Get(true),
			StateName:           pointer.Get("RUNNING"),
			ReplicationSpecs: &[]atlasv2.ReplicationSpec{
				{
					NumShards: pointer.Get(3),
					Id:        pointer.Get(zoneID1),
					ZoneName:  pointer.Get(zoneName1),
					RegionConfigs: &[]atlasv2.CloudRegionConfig{
						{
							AnalyticsSpecs: &atlasv2.DedicatedHardwareSpec{
								DiskIOPS:      pointer.Get(10),
								EbsVolumeType: pointer.Get("TestEBSVolume"),
								InstanceSize:  pointer.Get("M20"),
								NodeCount:     pointer.Get(3),
							},
							ElectableSpecs: &atlasv2.HardwareSpec{
								DiskIOPS:      pointer.Get(10),
								EbsVolumeType: pointer.Get("TestEBSVolume"),
								InstanceSize:  pointer.Get("M20"),
								NodeCount:     pointer.Get(3),
							},
							ReadOnlySpecs: &atlasv2.DedicatedHardwareSpec{
								DiskIOPS:      pointer.Get(10),
								EbsVolumeType: pointer.Get("TestEBSVolume"),
								InstanceSize:  pointer.Get("M20"),
								NodeCount:     pointer.Get(3),
							},
							AutoScaling: &atlasv2.AdvancedAutoScalingSettings{
								DiskGB: &atlasv2.DiskGBAutoScaling{Enabled: pointer.Get(true)},
								Compute: &atlasv2.AdvancedComputeAutoScaling{
									Enabled:          pointer.Get(true),
									ScaleDownEnabled: pointer.Get(true),
									MinInstanceSize:  pointer.Get("M20"),
									MaxInstanceSize:  pointer.Get("M40"),
								},
							},
							Priority:     pointer.Get(1),
							ProviderName: pointer.Get("AWS"),
							RegionName:   pointer.Get("US_EAST_1"),
						},
					},
				},
			},
			CreateDate:           &time.Time{},
			RootCertType:         pointer.Get("TestRootCertType"),
			VersionReleaseSystem: pointer.Get("TestReleaseSystem"),
		}
		processArgs := &atlasv2.ClusterDescriptionProcessArgs{
			DefaultReadConcern:               pointer.Get("TestReadConcern"),
			DefaultWriteConcern:              pointer.Get("TestWriteConcert"),
			MinimumEnabledTlsProtocol:        pointer.Get("1.0"),
			FailIndexKeyTooLong:              pointer.Get(true),
			JavascriptEnabled:                pointer.Get(true),
			NoTableScan:                      pointer.Get(true),
			SampleSizeBIConnector:            pointer.Get[int](10),
			SampleRefreshIntervalBIConnector: pointer.Get[int](10),
		}
		processArgs.OplogSizeMB = pointer.Get(10)
		processArgs.OplogMinRetentionHours = pointer.Get(float64(10.1))
		backupSchedule := &atlasv2.DiskBackupSnapshotSchedule{
			ClusterId:             pointer.Get("testClusterID"),
			ClusterName:           pointer.Get(clusterName),
			ReferenceHourOfDay:    pointer.Get[int](5),
			ReferenceMinuteOfHour: pointer.Get[int](5),
			RestoreWindowDays:     pointer.Get[int](5),
			UpdateSnapshots:       pointer.Get(true),
			NextSnapshot:          pointer.Get(time.Now()),
			Policies: &[]atlasv2.AdvancedDiskBackupSnapshotSchedulePolicy{
				{
					Id: pointer.Get("1"),
					PolicyItems: &[]atlasv2.DiskBackupApiPolicyItem{
						{
							Id:                pointer.Get("1"),
							FrequencyInterval: 10,
							FrequencyType:     "DAYS",
							RetentionUnit:     "WEEKS",
							RetentionValue:    1,
						},
					},
				},
			},
			AutoExportEnabled: pointer.Get(true),
			Export: &atlasv2.AutoExportPolicy{
				ExportBucketId: pointer.Get("TestBucketID"),
				FrequencyType:  pointer.Get("TestFreqType"),
			},
			UseOrgAndGroupNamesInExportPrefix: pointer.Get(true),
			CopySettings: &[]atlasv2.DiskBackupCopySetting{
				{
					CloudProvider:     pointer.Get("AWS"),
					RegionName:        pointer.Get("US_EAST_1"),
					ReplicationSpecId: pointer.Get("123456"),
					ShouldCopyOplogs:  pointer.Get(false),
					Frequencies:       &[]string{"DAILY"},
				},
			},
		}
		globalCluster := &atlasv2.GeoSharding{
			CustomZoneMapping: &map[string]string{
				firstLocation: zoneID1,
			},
			ManagedNamespaces: &[]atlasv2.ManagedNamespaces{
				{
					Db:                     "testDB",
					Collection:             "testCollection",
					CustomShardKey:         "testShardKey",
					IsCustomShardKeyHashed: pointer.Get(true),
					IsShardKeyUnique:       pointer.Get(true),
					NumInitialChunks:       pointer.Get(int64(4)),
					PresplitHashedZones:    pointer.Get(true),
				},
			},
		}

		managedNamespace := globalCluster.GetManagedNamespaces()

		clusterStore.EXPECT().AtlasCluster(projectName, clusterName).Return(cluster, nil)
		clusterStore.EXPECT().AtlasClusterConfigurationOptions(projectName, clusterName).Return(processArgs, nil)
		clusterStore.EXPECT().GlobalCluster(projectName, clusterName).Return(globalCluster, nil)
		clusterStore.EXPECT().DescribeSchedule(projectName, clusterName).Return(backupSchedule, nil)

		expectCluster := &akov2.AtlasDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDeployment",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s", projectName, clusterName)),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: akov2.AtlasDeploymentSpec{
				Project: akov2common.ResourceRefNamespaced{
					Name:      strings.ToLower(projectName),
					Namespace: targetNamespace,
				},
				DeploymentSpec: &akov2.AdvancedDeploymentSpec{
					BackupEnabled: cluster.BackupEnabled,
					CustomZoneMapping: []akov2.CustomZoneMapping{
						{
							Location: firstLocation,
							Zone:     *cluster.GetReplicationSpecs()[0].ZoneName,
						},
					},
					ManagedNamespaces: []akov2.ManagedNamespace{
						{
							Db:                     managedNamespace[0].Db,
							Collection:             managedNamespace[0].Collection,
							CustomShardKey:         managedNamespace[0].CustomShardKey,
							IsCustomShardKeyHashed: managedNamespace[0].IsCustomShardKeyHashed,
							IsShardKeyUnique:       managedNamespace[0].IsShardKeyUnique,
							NumInitialChunks:       int(managedNamespace[0].GetNumInitialChunks()),
							PresplitHashedZones:    managedNamespace[0].PresplitHashedZones,
						},
					},
					BiConnector: &akov2.BiConnectorSpec{
						Enabled:        cluster.BiConnector.Enabled,
						ReadPreference: *cluster.BiConnector.ReadPreference,
					},
					ClusterType:              *cluster.ClusterType,
					EncryptionAtRestProvider: *cluster.EncryptionAtRestProvider,
					Labels: []akov2common.LabelSpec{
						{
							Key:   *cluster.GetLabels()[0].Key,
							Value: *cluster.GetLabels()[0].Value,
						},
					},
					Tags: []*akov2.TagSpec{
						{
							Key:   *cluster.GetTags()[0].Key,
							Value: *cluster.GetTags()[0].Value,
						},
					},
					Name:       clusterName,
					Paused:     cluster.Paused,
					PitEnabled: cluster.PitEnabled,
					ReplicationSpecs: []*akov2.AdvancedReplicationSpec{
						{
							NumShards: *cluster.GetReplicationSpecs()[0].NumShards,
							ZoneName:  *cluster.GetReplicationSpecs()[0].ZoneName,
							RegionConfigs: []*akov2.AdvancedRegionConfig{
								{
									AnalyticsSpecs: &akov2.Specs{
										DiskIOPS:      pointer.Get(int64(*cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AnalyticsSpecs.DiskIOPS)),
										EbsVolumeType: *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AnalyticsSpecs.EbsVolumeType,
										InstanceSize:  *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AnalyticsSpecs.InstanceSize,
										NodeCount:     cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AnalyticsSpecs.NodeCount,
									},
									ElectableSpecs: &akov2.Specs{
										DiskIOPS:      pointer.Get(int64(*cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.DiskIOPS)),
										EbsVolumeType: *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.EbsVolumeType,
										InstanceSize:  *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.InstanceSize,
										NodeCount:     cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.NodeCount,
									},
									ReadOnlySpecs: &akov2.Specs{
										DiskIOPS:      pointer.Get(int64(*cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ReadOnlySpecs.DiskIOPS)),
										EbsVolumeType: *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ReadOnlySpecs.EbsVolumeType,
										InstanceSize:  *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ReadOnlySpecs.InstanceSize,
										NodeCount:     cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ReadOnlySpecs.NodeCount,
									},
									AutoScaling: &akov2.AdvancedAutoScalingSpec{
										DiskGB: &akov2.DiskGB{Enabled: cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AutoScaling.DiskGB.Enabled},
										Compute: &akov2.ComputeSpec{
											Enabled:          cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AutoScaling.Compute.Enabled,
											ScaleDownEnabled: cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AutoScaling.Compute.ScaleDownEnabled,
											MinInstanceSize:  *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AutoScaling.Compute.MinInstanceSize,
											MaxInstanceSize:  *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].AutoScaling.Compute.MaxInstanceSize,
										},
									},
									Priority:     cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].Priority,
									ProviderName: *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ProviderName,
									RegionName:   *cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].RegionName,
								},
							},
						},
					},
					RootCertType:         *cluster.RootCertType,
					VersionReleaseSystem: *cluster.VersionReleaseSystem,
				},
				BackupScheduleRef: akov2common.ResourceRefNamespaced{
					Name:      strings.ToLower(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName)),
					Namespace: targetNamespace,
				},
				ServerlessSpec: nil,
				ProcessArgs: &akov2.ProcessArgs{
					DefaultReadConcern:               processArgs.GetDefaultReadConcern(),
					DefaultWriteConcern:              processArgs.GetDefaultWriteConcern(),
					MinimumEnabledTLSProtocol:        processArgs.GetMinimumEnabledTlsProtocol(),
					FailIndexKeyTooLong:              processArgs.FailIndexKeyTooLong,
					JavascriptEnabled:                processArgs.JavascriptEnabled,
					NoTableScan:                      processArgs.NoTableScan,
					OplogSizeMB:                      pointer.Get(int64(processArgs.GetOplogSizeMB())),
					SampleSizeBIConnector:            pointer.Get(int64(processArgs.GetSampleSizeBIConnector())),
					SampleRefreshIntervalBIConnector: pointer.Get(int64(processArgs.GetSampleRefreshIntervalBIConnector())),
				},
			},
			Status: akov2status.AtlasDeploymentStatus{
				Common: akov2status.Common{
					Conditions: []akov2status.Condition{},
				},
			},
		}

		expectPolicies := []*akov2.AtlasBackupPolicy{
			{
				TypeMeta: v1.TypeMeta{
					Kind:       "AtlasBackupPolicy",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      strings.ToLower(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName)),
					Namespace: targetNamespace,
					Labels: map[string]string{
						features.ResourceVersion: resourceVersion,
					},
				},
				Spec: akov2.AtlasBackupPolicySpec{
					Items: []akov2.AtlasBackupPolicyItem{
						{
							FrequencyType:     backupSchedule.GetPolicies()[0].GetPolicyItems()[0].GetFrequencyType(),
							FrequencyInterval: backupSchedule.GetPolicies()[0].GetPolicyItems()[0].GetFrequencyInterval(),
							RetentionUnit:     backupSchedule.GetPolicies()[0].GetPolicyItems()[0].GetRetentionUnit(),
							RetentionValue:    backupSchedule.GetPolicies()[0].GetPolicyItems()[0].GetRetentionValue(),
						},
					},
				},
				Status: akov2status.BackupPolicyStatus{},
			},
		}

		expectSchedule := &akov2.AtlasBackupSchedule{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasBackupSchedule",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName)),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: akov2.AtlasBackupScheduleSpec{
				AutoExportEnabled: *backupSchedule.AutoExportEnabled,
				Export: &akov2.AtlasBackupExportSpec{
					ExportBucketID: backupSchedule.Export.GetExportBucketId(),
					FrequencyType:  backupSchedule.Export.GetFrequencyType(),
				},
				PolicyRef: akov2common.ResourceRefNamespaced{
					Name:      strings.ToLower(expectPolicies[0].Name),
					Namespace: expectPolicies[0].Namespace,
				},
				ReferenceHourOfDay:                int64(backupSchedule.GetReferenceHourOfDay()),
				ReferenceMinuteOfHour:             int64(backupSchedule.GetReferenceMinuteOfHour()),
				RestoreWindowDays:                 int64(backupSchedule.GetRestoreWindowDays()),
				UpdateSnapshots:                   backupSchedule.GetUpdateSnapshots(),
				UseOrgAndGroupNamesInExportPrefix: backupSchedule.GetUseOrgAndGroupNamesInExportPrefix(),
				CopySettings: []akov2.CopySetting{
					{
						CloudProvider:    pointer.Get("AWS"),
						RegionName:       pointer.Get("US_EAST_1"),
						ShouldCopyOplogs: pointer.Get(false),
						Frequencies:      []string{"DAILY"},
					},
				},
			},
			Status: akov2status.BackupScheduleStatus{},
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
	clusterStore := atlasmocks.NewMockOperatorClusterStore(ctl)
	dictionary := resources.AtlasNameToKubernetesName()

	featureValidator := mocks.NewMockFeatureValidator(ctl)

	t.Run("Can import Serverless deployment", func(t *testing.T) {
		speID := "TestPEId"
		speCloudProviderEndpointID := "TestCloudProviderID"
		speComment := "TestPEName"
		spePrivateEndpointIPAddress := ""

		spe := []atlasv2.ServerlessTenantEndpoint{
			{
				Id:                       &speID,
				CloudProviderEndpointId:  &speCloudProviderEndpointID,
				Comment:                  &speComment,
				PrivateEndpointIpAddress: &spePrivateEndpointIPAddress,
				ProviderName:             pointer.Get("AZURE"),
			},
		}

		cluster := &atlasv2.ServerlessInstanceDescription{
			Id:             pointer.Get("TestClusterID"),
			GroupId:        pointer.Get("TestGroupID"),
			MongoDBVersion: pointer.Get("5.0"),
			Name:           pointer.Get(clusterName),
			CreateDate:     pointer.Get(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)),
			ProviderSettings: atlasv2.ServerlessProviderSettings{
				BackingProviderName: "AWS",
				ProviderName:        pointer.Get("AWS"),
				RegionName:          "US_EAST_1",
			},
			StateName:               pointer.Get(""),
			ServerlessBackupOptions: nil,
			ConnectionStrings:       nil,
			Links:                   nil,
		}

		clusterStore.EXPECT().GetServerlessInstance(projectName, clusterName).Return(cluster, nil)
		clusterStore.EXPECT().ServerlessPrivateEndpoints(projectName, clusterName).Return(spe, nil)

		expected := &akov2.AtlasDeployment{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDeployment",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s", projectName, clusterName)),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: akov2.AtlasDeploymentSpec{
				Project: akov2common.ResourceRefNamespaced{
					Name:      strings.ToLower(projectName),
					Namespace: targetNamespace,
				},
				BackupScheduleRef: akov2common.ResourceRefNamespaced{},
				ServerlessSpec: &akov2.ServerlessSpec{
					Name: atlas.StringOrEmpty(cluster.Name),
					ProviderSettings: &akov2.ProviderSettingsSpec{
						BackingProviderName: cluster.ProviderSettings.BackingProviderName,
						ProviderName:        akov2provider.ProviderName(atlas.StringOrEmpty(cluster.ProviderSettings.ProviderName)),
						RegionName:          cluster.ProviderSettings.RegionName,
					},
					PrivateEndpoints: []akov2.ServerlessPrivateEndpoint{
						{
							Name:                     speComment,
							CloudProviderEndpointID:  speCloudProviderEndpointID,
							PrivateEndpointIPAddress: spePrivateEndpointIPAddress,
						},
					},
				},
				ProcessArgs: nil,
			},
			Status: akov2status.AtlasDeploymentStatus{
				Common: akov2status.Common{
					Conditions: []akov2status.Condition{},
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
