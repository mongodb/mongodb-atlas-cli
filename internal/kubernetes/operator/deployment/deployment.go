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

package deployment

import (
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/pointers"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MaxItems = 500
)

type AtlasDeploymentResult struct {
	Deployment     *atlasV1.AtlasDeployment
	BackupSchedule *atlasV1.AtlasBackupSchedule
	BackupPolicies []*atlasV1.AtlasBackupPolicy
}

func BuildAtlasAdvancedDeployment(deploymentStore store.AtlasOperatorClusterStore, projectID, projectName, clusterID, targetNamespace string) (*AtlasDeploymentResult, error) {
	deployment, err := deploymentStore.AtlasCluster(projectID, clusterID)
	if err != nil {
		return nil, err
	}

	var advancedSpec *atlasV1.AdvancedDeploymentSpec

	convertBiConnector := func(biConnector *mongodbatlas.BiConnector) *atlasV1.BiConnectorSpec {
		if biConnector == nil {
			return nil
		}
		return &atlasV1.BiConnectorSpec{
			Enabled:        biConnector.Enabled,
			ReadPreference: biConnector.ReadPreference,
		}
	}

	convertLabels := func(labels []mongodbatlas.Label) []common.LabelSpec {
		result := make([]common.LabelSpec, 0, len(labels))

		for _, atlasLabel := range labels {
			result = append(result, common.LabelSpec{
				Key:   atlasLabel.Key,
				Value: atlasLabel.Value,
			})
		}
		return result
	}

	replicationSpec := buildReplicationSpec(deployment.ReplicationSpecs)

	processArgs, err := buildProcessArgs(deploymentStore, projectID, clusterID)
	if err != nil {
		return nil, err
	}

	customZoneMapping, ManagedNamespaces, err := buildGlobalDeployment(deployment.ReplicationSpecs, deploymentStore, projectID, clusterID)

	// TODO: DiskSizeGB field skipped on purpose. See https://jira.mongodb.org/browse/CLOUDP-146469
	advancedSpec = &atlasV1.AdvancedDeploymentSpec{
		BackupEnabled:            deployment.BackupEnabled,
		BiConnector:              convertBiConnector(deployment.BiConnector),
		ClusterType:              deployment.ClusterType,
		EncryptionAtRestProvider: deployment.EncryptionAtRestProvider,
		Labels:                   convertLabels(deployment.Labels),
		Name:                     deployment.Name,
		Paused:                   deployment.Paused,
		PitEnabled:               deployment.PitEnabled,
		ReplicationSpecs:         replicationSpec,
		RootCertType:             deployment.RootCertType,
		VersionReleaseSystem:     deployment.VersionReleaseSystem,
		ManagedNamespaces:        ManagedNamespaces,
		CustomZoneMapping:        customZoneMapping,
	}

	var backupScheduleRef common.ResourceRefNamespaced
	backupSchedule, backupPolicies := buildBackups(deploymentStore, projectID, clusterID, targetNamespace)
	if backupSchedule != nil {
		backupScheduleRef.Name = backupSchedule.Name
		backupScheduleRef.Namespace = backupSchedule.Namespace
	}

	atlasDeployment := &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      strings.ToLower(clusterID),
			Namespace: targetNamespace,
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      strings.ToLower(projectName),
				Namespace: targetNamespace,
			},
			DeploymentSpec:         nil,
			AdvancedDeploymentSpec: advancedSpec,
			ServerlessSpec:         nil,
			ProcessArgs:            processArgs,
			BackupScheduleRef:      backupScheduleRef,
		},
		Status: status.AtlasDeploymentStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}

	result := &AtlasDeploymentResult{
		Deployment:     atlasDeployment,
		BackupSchedule: backupSchedule,
		BackupPolicies: backupPolicies,
	}

	return result, nil
}

func buildGlobalDeployment(atlasRepSpec []*mongodbatlas.AdvancedReplicationSpec, globalDeploymentProvider store.GlobalDeploymentDescriber, projectID, clusterID string) ([]atlasV1.CustomZoneMapping, []atlasV1.ManagedNamespace, error) {
	globalCluster, err := globalDeploymentProvider.GlobalDeployment(projectID, clusterID)
	if err != nil {
		return nil, nil, err
	}
	var customZoneMapping []atlasV1.CustomZoneMapping
	if globalCluster.CustomZoneMapping != nil {
		// create map ID -> Name for zones
		zoneMap := make(map[string]string, len(atlasRepSpec))
		for _, rc := range atlasRepSpec {
			zoneMap[rc.ID] = rc.ZoneName
		}

		customZoneMapping = make([]atlasV1.CustomZoneMapping, 0, len(globalCluster.CustomZoneMapping))
		for location, zoneID := range globalCluster.CustomZoneMapping {
			customZoneMapping = append(customZoneMapping, atlasV1.CustomZoneMapping{
				Zone:     zoneMap[zoneID],
				Location: location,
			})
		}
	}

	var managedNamespace []atlasV1.ManagedNamespace
	if globalCluster.ManagedNamespaces != nil {

		// TODO: wait until https://github.com/mongodb/go-client-mongodb-atlas/pull/337 is merged
		for _, ns := range globalCluster.ManagedNamespaces {
			managedNamespace = append(managedNamespace, atlasV1.ManagedNamespace{
				Db:             ns.Db,
				Collection:     ns.Collection,
				CustomShardKey: ns.CustomShardKey,
				// NumInitialChunks:       ns.NumInitialChunks,
				// PresplitHashedZones:    ns.PreSplitHashedZones,
				IsCustomShardKeyHashed: ns.IsCustomShardKeyHashed,
				IsShardKeyUnique:       ns.IsShardKeyUnique,
			})
		}
	}

	return customZoneMapping, managedNamespace, nil

}
func buildProcessArgs(configOptsProvider store.AtlasClusterConfigurationOptionsDescriber, projectID, clusterName string) (*atlasV1.ProcessArgs, error) {
	pArgs, err := configOptsProvider.AtlasClusterConfigurationOptions(projectID, clusterName)
	if err != nil {
		return nil, err
	}

	// TODO: OplogMinRetentionHours is not exported due to a bug https://jira.mongodb.org/browse/CLOUDP-146481
	return &atlasV1.ProcessArgs{
		DefaultReadConcern:               pArgs.DefaultReadConcern,
		DefaultWriteConcern:              pArgs.DefaultWriteConcern,
		MinimumEnabledTLSProtocol:        pArgs.MinimumEnabledTLSProtocol,
		FailIndexKeyTooLong:              pArgs.FailIndexKeyTooLong,
		JavascriptEnabled:                pArgs.JavascriptEnabled,
		NoTableScan:                      pArgs.NoTableScan,
		OplogSizeMB:                      pArgs.OplogSizeMB,
		SampleSizeBIConnector:            pArgs.SampleSizeBIConnector,
		SampleRefreshIntervalBIConnector: pArgs.SampleRefreshIntervalBIConnector,
	}, nil
}

func buildBackups(backupsProvider store.ScheduleDescriber, projectID, clusterName, targetNamespace string) (*atlasV1.AtlasBackupSchedule, []*atlasV1.AtlasBackupPolicy) {
	bs, err := backupsProvider.DescribeSchedule(projectID, clusterName)
	if err != nil {
		return nil, nil
	}

	// Although we have a for loop here, there should be only one policy per schedule. See Atlas API implementation
	policies := make([]*atlasV1.AtlasBackupPolicy, 0, len(bs.Policies))
	for _, p := range bs.Policies {
		items := make([]atlasV1.AtlasBackupPolicyItem, 0, len(p.PolicyItems))
		for _, pItem := range p.PolicyItems {
			items = append(items, atlasV1.AtlasBackupPolicyItem{
				FrequencyType:     pItem.FrequencyType,
				FrequencyInterval: pItem.FrequencyInterval,
				RetentionUnit:     pItem.RetentionUnit,
				RetentionValue:    pItem.RetentionValue,
			})
		}
		policies = append(policies, &atlasV1.AtlasBackupPolicy{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasBackupPolicy",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-backuppolicy", clusterName)),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasBackupPolicySpec{
				Items: items,
			},
			Status: atlasV1.AtlasBackupPolicyStatus{},
		})
	}

	export := &atlasV1.AtlasBackupExportSpec{}
	if bs.Export != nil {
		export.ExportBucketID = bs.Export.ExportBucketID
		export.FrequencyType = bs.Export.FrequencyType
	}

	schedule := &atlasV1.AtlasBackupSchedule{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasBackupSchedule",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      strings.ToLower(fmt.Sprintf("%s-backupschedule", clusterName)),
			Namespace: targetNamespace,
		},
		Spec: atlasV1.AtlasBackupScheduleSpec{
			AutoExportEnabled: pointers.PtrValOrDefault(bs.AutoExportEnabled, false),
			Export:            export,
			PolicyRef: common.ResourceRefNamespaced{
				Name:      policies[0].Name,
				Namespace: policies[0].Namespace,
			},
			ReferenceHourOfDay:                pointers.PtrValOrDefault(bs.ReferenceHourOfDay, 0),
			ReferenceMinuteOfHour:             pointers.PtrValOrDefault(bs.ReferenceMinuteOfHour, 0),
			RestoreWindowDays:                 pointers.PtrValOrDefault(bs.RestoreWindowDays, 0),
			UpdateSnapshots:                   pointers.PtrValOrDefault(bs.UpdateSnapshots, false),
			UseOrgAndGroupNamesInExportPrefix: pointers.PtrValOrDefault(bs.UseOrgAndGroupNamesInExportPrefix, false),
		},
		Status: atlasV1.AtlasBackupScheduleStatus{},
	}
	return schedule, policies
}

func buildReplicationSpec(atlasRepSpec []*mongodbatlas.AdvancedReplicationSpec) []*atlasV1.AdvancedReplicationSpec {
	result := make([]*atlasV1.AdvancedReplicationSpec, 0, len(atlasRepSpec))
	for _, rs := range atlasRepSpec {
		if rs == nil {
			continue
		}

		replicationSpec := &atlasV1.AdvancedReplicationSpec{
			NumShards:     rs.NumShards,
			ZoneName:      rs.ZoneName,
			RegionConfigs: nil,
		}

		if rs.RegionConfigs == nil {
			result = append(result, replicationSpec)
			continue
		}

		replicationSpec.RegionConfigs = make([]*atlasV1.AdvancedRegionConfig, 0, len(replicationSpec.RegionConfigs))
		for _, rc := range rs.RegionConfigs {
			if rc == nil {
				continue
			}

			var analyticsSpecs *atlasV1.Specs
			if rc.AnalyticsSpecs != nil {
				analyticsSpecs = &atlasV1.Specs{
					DiskIOPS:      rc.AnalyticsSpecs.DiskIOPS,
					EbsVolumeType: rc.AnalyticsSpecs.EbsVolumeType,
					InstanceSize:  rc.AnalyticsSpecs.InstanceSize,
					NodeCount:     rc.AnalyticsSpecs.NodeCount,
				}
			}
			var electableSpecs *atlasV1.Specs
			if rc.ElectableSpecs != nil {
				electableSpecs = &atlasV1.Specs{
					DiskIOPS:      rc.ElectableSpecs.DiskIOPS,
					EbsVolumeType: rc.ElectableSpecs.EbsVolumeType,
					InstanceSize:  rc.ElectableSpecs.InstanceSize,
					NodeCount:     rc.ElectableSpecs.NodeCount,
				}
			}

			var readOnlySpecs *atlasV1.Specs
			if rc.ReadOnlySpecs != nil {
				readOnlySpecs = &atlasV1.Specs{
					DiskIOPS:      rc.ReadOnlySpecs.DiskIOPS,
					EbsVolumeType: rc.ReadOnlySpecs.EbsVolumeType,
					InstanceSize:  rc.ReadOnlySpecs.InstanceSize,
					NodeCount:     rc.ReadOnlySpecs.NodeCount,
				}
			}

			var autoscalingSpec *atlasV1.AdvancedAutoScalingSpec
			if rc.AutoScaling != nil {
				var compute *atlasV1.ComputeSpec
				if rc.AutoScaling.Compute != nil {
					compute = &atlasV1.ComputeSpec{
						Enabled:          rc.AutoScaling.Compute.Enabled,
						ScaleDownEnabled: rc.AutoScaling.Compute.ScaleDownEnabled,
						MinInstanceSize:  rc.AutoScaling.Compute.MinInstanceSize,
						MaxInstanceSize:  rc.AutoScaling.Compute.MaxInstanceSize,
					}
				}

				var diskGB *atlasV1.DiskGB
				if rc.AutoScaling.DiskGB != nil {
					diskGB = &atlasV1.DiskGB{Enabled: rc.AutoScaling.DiskGB.Enabled}
				}
				autoscalingSpec = &atlasV1.AdvancedAutoScalingSpec{
					DiskGB:  diskGB,
					Compute: compute,
				}
			}

			replicationSpec.RegionConfigs = append(replicationSpec.RegionConfigs, &atlasV1.AdvancedRegionConfig{
				AnalyticsSpecs:      analyticsSpecs,
				ElectableSpecs:      electableSpecs,
				ReadOnlySpecs:       readOnlySpecs,
				AutoScaling:         autoscalingSpec,
				BackingProviderName: rc.BackingProviderName,
				Priority:            rc.Priority,
				ProviderName:        rc.ProviderName,
				RegionName:          rc.RegionName,
			})
		}
		result = append(result, replicationSpec)
	}
	return result
}

func BuildServerlessDeployments(deploymentStore store.AtlasOperatorClusterStore, projectID, projectName, clusterID, targetNamespace string) (*atlasV1.AtlasDeployment, error) {
	deployment, err := deploymentStore.ServerlessInstance(projectID, clusterID)
	if err != nil {
		return nil, err
	}

	privateEndpoints, err := buildServerlessPrivateEndpoints(deploymentStore, projectID, deployment.Name)
	if err != nil {
		return nil, err
	}

	var providerSettings *atlasV1.ProviderSettingsSpec

	if deployment.ProviderSettings != nil {
		var autoscalingSpec *atlasV1.AutoScalingSpec

		if deployment.AutoScaling != nil {
			var computeSpec *atlasV1.ComputeSpec

			if deployment.AutoScaling.Compute != nil {
				computeSpec = &atlasV1.ComputeSpec{
					Enabled:          deployment.AutoScaling.Compute.Enabled,
					ScaleDownEnabled: deployment.AutoScaling.Compute.ScaleDownEnabled,
					MinInstanceSize:  deployment.AutoScaling.Compute.MinInstanceSize,
					MaxInstanceSize:  deployment.AutoScaling.Compute.MaxInstanceSize,
				}
			}
			autoscalingSpec = &atlasV1.AutoScalingSpec{
				AutoIndexingEnabled: deployment.AutoScaling.AutoIndexingEnabled,
				DiskGBEnabled:       deployment.AutoScaling.DiskGBEnabled,
				Compute:             computeSpec,
			}
		}

		providerSettings = &atlasV1.ProviderSettingsSpec{
			BackingProviderName: deployment.ProviderSettings.BackingProviderName,
			DiskIOPS:            deployment.ProviderSettings.DiskIOPS,
			DiskTypeName:        deployment.ProviderSettings.DiskTypeName,
			EncryptEBSVolume:    deployment.ProviderSettings.EncryptEBSVolume,
			InstanceSizeName:    deployment.ProviderSettings.InstanceSizeName,
			ProviderName:        provider.ProviderName(deployment.ProviderSettings.ProviderName),
			RegionName:          deployment.ProviderSettings.RegionName,
			VolumeType:          deployment.ProviderSettings.VolumeType,
			AutoScaling:         autoscalingSpec,
		}
	}

	serverlessSpec := &atlasV1.ServerlessSpec{
		Name:             deployment.Name,
		ProviderSettings: providerSettings,
		PrivateEndpoints: privateEndpoints,
	}

	atlasDeployment := &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      strings.ToLower(deployment.Name),
			Namespace: targetNamespace,
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      strings.ToLower(projectName),
				Namespace: targetNamespace,
			},
			BackupScheduleRef: common.ResourceRefNamespaced{},
			ServerlessSpec:    serverlessSpec,
			ProcessArgs:       nil,
		},
		Status: status.AtlasDeploymentStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}

	return atlasDeployment, nil
}

func buildServerlessPrivateEndpoints(deploymentStore store.ServerlessPrivateEndpointsLister, projectID, clusterName string) ([]atlasV1.ServerlessPrivateEndpoint, error) {
	endpoints, err := deploymentStore.ServerlessPrivateEndpoints(projectID, clusterName, &mongodbatlas.ListOptions{ItemsPerPage: MaxItems})
	if err != nil {
		return nil, err
	}

	result := make([]atlasV1.ServerlessPrivateEndpoint, 0, len(endpoints))

	for i := range endpoints {
		result = append(result, atlasV1.ServerlessPrivateEndpoint{
			Name:                     endpoints[i].Comment,
			CloudProviderEndpointID:  endpoints[i].CloudProviderEndpointID,
			PrivateEndpointIPAddress: endpoints[i].PrivateEndpointIPAddress,
		})
	}
	return result, nil
}
