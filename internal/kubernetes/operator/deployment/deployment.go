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

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201007/admin"
	"go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MaxItems                          = 500
	featureProcessArgs                = "processArgs"
	featureBackupSchedule             = "backupRef"
	featureServerlessPrivateEndpoints = "serverlessSpec.privateEndpoints"
	featureGlobalDeployments          = "advancedDeploymentSpec.customZoneMapping"
	DeletingState                     = "DELETING"
	DeletedState                      = "DELETED"
)

type AtlasDeploymentResult struct {
	Deployment     *atlasV1.AtlasDeployment
	BackupSchedule *atlasV1.AtlasBackupSchedule
	BackupPolicies []*atlasV1.AtlasBackupPolicy
}

func BuildAtlasAdvancedDeployment(deploymentStore store.AtlasOperatorClusterStore, validator features.FeatureValidator, projectID, projectName, clusterID, targetNamespace string, dictionary map[string]string, version string) (*AtlasDeploymentResult, error) {
	deployment, err := deploymentStore.AtlasCluster(projectID, clusterID)
	if err != nil {
		return nil, err
	}

	if !isAdvancedDeploymentExportable(deployment) {
		return nil, nil
	}

	var advancedSpec *atlasV1.AdvancedDeploymentSpec

	convertBiConnector := func(biConnector *atlasv2.BiConnector) *atlasV1.BiConnectorSpec {
		if biConnector == nil {
			return nil
		}
		return &atlasV1.BiConnectorSpec{
			Enabled:        biConnector.Enabled,
			ReadPreference: biConnector.GetReadPreference(),
		}
	}

	convertLabels := func(labels []atlasv2.ComponentLabel) []common.LabelSpec {
		result := make([]common.LabelSpec, 0, len(labels))

		for _, atlasLabel := range labels {
			result = append(result, common.LabelSpec{
				Key:   atlasLabel.GetKey(),
				Value: atlasLabel.GetValue(),
			})
		}
		return result
	}

	replicationSpec := buildReplicationSpec(deployment.GetReplicationSpecs())

	// TODO: DiskSizeGB field skipped on purpose. See https://jira.mongodb.org/browse/CLOUDP-146469
	advancedSpec = &atlasV1.AdvancedDeploymentSpec{
		BackupEnabled:            deployment.BackupEnabled,
		BiConnector:              convertBiConnector(deployment.BiConnector),
		ClusterType:              deployment.GetClusterType(),
		EncryptionAtRestProvider: deployment.GetEncryptionAtRestProvider(),
		Labels:                   convertLabels(deployment.GetLabels()),
		Name:                     deployment.GetName(),
		Paused:                   deployment.Paused,
		PitEnabled:               deployment.PitEnabled,
		ReplicationSpecs:         replicationSpec,
		RootCertType:             deployment.GetRootCertType(),
		VersionReleaseSystem:     deployment.GetVersionReleaseSystem(),
	}

	atlasDeployment := &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, clusterID), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: targetNamespace,
			},
			DeploymentSpec:         nil,
			AdvancedDeploymentSpec: advancedSpec,
			ServerlessSpec:         nil,
			ProcessArgs:            nil,
		},
		Status: status.AtlasDeploymentStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}

	deploymentResult := &AtlasDeploymentResult{
		Deployment:     atlasDeployment,
		BackupSchedule: nil,
		BackupPolicies: nil,
	}

	if validator.FeatureExist(features.ResourceAtlasDeployment, featureProcessArgs) {
		processArgs, err := buildProcessArgs(deploymentStore, projectID, clusterID)
		if err != nil {
			return nil, err
		}
		atlasDeployment.Spec.ProcessArgs = processArgs
	}

	if validator.FeatureExist(features.ResourceAtlasDeployment, featureBackupSchedule) {
		var backupScheduleRef common.ResourceRefNamespaced
		backupSchedule, backupPolicies := buildBackups(deploymentStore, projectName, projectID, clusterID, targetNamespace, version, dictionary)
		if backupSchedule != nil {
			backupScheduleRef.Name = backupSchedule.Name
			backupScheduleRef.Namespace = backupSchedule.Namespace
		}
		deploymentResult.BackupSchedule = backupSchedule
		deploymentResult.BackupPolicies = backupPolicies
		atlasDeployment.Spec.BackupScheduleRef = backupScheduleRef
	}

	if validator.FeatureExist(features.ResourceAtlasDeployment, featureGlobalDeployments) {
		customZoneMapping, managedNamespaces, err := buildGlobalDeployment(deployment.ReplicationSpecs, deploymentStore, projectID, clusterID)
		if err != nil {
			return nil, err
		}
		advancedSpec.CustomZoneMapping = customZoneMapping
		advancedSpec.ManagedNamespaces = managedNamespaces
	}

	return deploymentResult, nil
}

func buildGlobalDeployment(atlasRepSpec []atlasv2.ReplicationSpec, globalDeploymentProvider store.GlobalClusterDescriber, projectID, clusterID string) ([]atlasV1.CustomZoneMapping, []atlasV1.ManagedNamespace, error) {
	globalCluster, err := globalDeploymentProvider.GlobalCluster(projectID, clusterID)
	if err != nil {
		return nil, nil, err
	}
	var customZoneMapping []atlasV1.CustomZoneMapping
	if globalCluster.CustomZoneMapping != nil {
		// create map ID -> Name for zones
		zoneMap := make(map[string]string, len(atlasRepSpec))
		for _, rc := range atlasRepSpec {
			zoneMap[rc.GetId()] = rc.GetZoneName()
		}

		customZoneMapping = make([]atlasV1.CustomZoneMapping, 0, len(globalCluster.GetCustomZoneMapping()))
		for location, zoneID := range globalCluster.GetCustomZoneMapping() {
			customZoneMapping = append(customZoneMapping, atlasV1.CustomZoneMapping{
				Zone:     zoneMap[zoneID],
				Location: location,
			})
		}
	}

	if globalCluster.ManagedNamespaces == nil {
		return customZoneMapping, nil, nil
	}

	managedNamespace := make([]atlasV1.ManagedNamespace, len(globalCluster.ManagedNamespaces))
	for i, ns := range globalCluster.ManagedNamespaces {
		managedNamespace[i] = atlasV1.ManagedNamespace{
			Db:                     ns.Db,
			Collection:             ns.Collection,
			CustomShardKey:         ns.CustomShardKey,
			NumInitialChunks:       int(ns.GetNumInitialChunks()),
			PresplitHashedZones:    ns.PresplitHashedZones,
			IsCustomShardKeyHashed: ns.IsCustomShardKeyHashed,
			IsShardKeyUnique:       ns.IsShardKeyUnique,
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
		DefaultReadConcern:               pArgs.GetDefaultReadConcern(),
		DefaultWriteConcern:              pArgs.GetDefaultWriteConcern(),
		MinimumEnabledTLSProtocol:        pArgs.GetMinimumEnabledTlsProtocol(),
		FailIndexKeyTooLong:              pArgs.FailIndexKeyTooLong,
		JavascriptEnabled:                pArgs.JavascriptEnabled,
		NoTableScan:                      pArgs.NoTableScan,
		OplogSizeMB:                      pointer.GetNonZeroValue(int64(pArgs.GetOplogSizeMB())),
		SampleSizeBIConnector:            pointer.GetNonZeroValue(int64(pArgs.GetSampleSizeBIConnector())),
		SampleRefreshIntervalBIConnector: pointer.GetNonZeroValue(int64(pArgs.GetSampleRefreshIntervalBIConnector())),
	}, nil
}

func isAdvancedDeploymentExportable(deployments *atlasv2.AdvancedClusterDescription) bool {
	if deployments.GetStateName() == DeletingState || deployments.GetStateName() == DeletedState {
		return false
	}
	return true
}

func isServerlessExportable(deployment *mongodbatlas.Cluster) bool {
	if deployment.StateName == DeletingState || deployment.StateName == DeletedState {
		return false
	}
	return true
}

func buildBackups(backupsProvider store.ScheduleDescriber, projectName, projectID, clusterName, targetNamespace, version string, dictionary map[string]string) (*atlasV1.AtlasBackupSchedule, []*atlasV1.AtlasBackupPolicy) {
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
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: version,
				},
			},
			Spec: atlasV1.AtlasBackupPolicySpec{
				Items: items,
			},
			Status: status.BackupPolicyStatus{},
		})
	}

	var export *atlasV1.AtlasBackupExportSpec
	if bs.Export != nil {
		export = &atlasV1.AtlasBackupExportSpec{
			ExportBucketID: bs.Export.GetExportBucketId(),
			FrequencyType:  bs.Export.GetFrequencyType(),
		}
	}

	schedule := &atlasV1.AtlasBackupSchedule{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasBackupSchedule",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: atlasV1.AtlasBackupScheduleSpec{
			AutoExportEnabled: pointer.GetOrDefault(bs.AutoExportEnabled, false),
			Export:            export,
			PolicyRef: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(policies[0].Name, dictionary),
				Namespace: policies[0].Namespace,
			},
			ReferenceHourOfDay:                int64(pointer.GetOrDefault(bs.ReferenceHourOfDay, 0)),
			ReferenceMinuteOfHour:             int64(pointer.GetOrDefault(bs.ReferenceMinuteOfHour, 0)),
			RestoreWindowDays:                 int64(pointer.GetOrDefault(bs.RestoreWindowDays, 0)),
			UpdateSnapshots:                   pointer.GetOrDefault(bs.UpdateSnapshots, false),
			UseOrgAndGroupNamesInExportPrefix: pointer.GetOrDefault(bs.UseOrgAndGroupNamesInExportPrefix, false),
		},
		Status: status.BackupScheduleStatus{},
	}

	if len(bs.CopySettings) > 0 {
		copySettings := make([]atlasV1.CopySetting, 0, len(bs.CopySettings))

		for _, copySetting := range bs.CopySettings {
			copySettings = append(
				copySettings,
				atlasV1.CopySetting{
					CloudProvider:     copySetting.CloudProvider,
					RegionName:        copySetting.RegionName,
					ReplicationSpecID: copySetting.ReplicationSpecId,
					ShouldCopyOplogs:  copySetting.ShouldCopyOplogs,
					Frequencies:       copySetting.Frequencies,
				},
			)
		}

		schedule.Spec.CopySettings = copySettings
	}

	return schedule, policies
}

func buildReplicationSpec(atlasRepSpec []atlasv2.ReplicationSpec) []*atlasV1.AdvancedReplicationSpec {
	result := make([]*atlasV1.AdvancedReplicationSpec, 0, len(atlasRepSpec))
	for _, rs := range atlasRepSpec {
		replicationSpec := &atlasV1.AdvancedReplicationSpec{
			NumShards:     rs.GetNumShards(),
			ZoneName:      rs.GetZoneName(),
			RegionConfigs: nil,
		}

		if rs.RegionConfigs == nil {
			result = append(result, replicationSpec)
			continue
		}

		replicationSpec.RegionConfigs = make([]*atlasV1.AdvancedRegionConfig, 0, len(replicationSpec.RegionConfigs))
		for _, rc := range rs.RegionConfigs {
			var analyticsSpecs *atlasV1.Specs
			if rc.AnalyticsSpecs != nil {
				analyticsSpecs = &atlasV1.Specs{
					DiskIOPS:      pointer.Get(int64(rc.AnalyticsSpecs.GetDiskIOPS())),
					EbsVolumeType: rc.AnalyticsSpecs.GetEbsVolumeType(),
					InstanceSize:  rc.AnalyticsSpecs.GetInstanceSize(),
					NodeCount:     rc.AnalyticsSpecs.NodeCount,
				}
			}
			var electableSpecs *atlasV1.Specs
			if rc.ElectableSpecs != nil {
				electableSpecs = &atlasV1.Specs{
					DiskIOPS:      pointer.Get(int64(rc.ElectableSpecs.GetDiskIOPS())),
					EbsVolumeType: rc.ElectableSpecs.GetEbsVolumeType(),
					InstanceSize:  rc.ElectableSpecs.GetInstanceSize(),
					NodeCount:     rc.ElectableSpecs.NodeCount,
				}
			}

			var readOnlySpecs *atlasV1.Specs
			if rc.ReadOnlySpecs != nil {
				readOnlySpecs = &atlasV1.Specs{
					DiskIOPS:      pointer.Get(int64(rc.ReadOnlySpecs.GetDiskIOPS())),
					EbsVolumeType: rc.ReadOnlySpecs.GetEbsVolumeType(),
					InstanceSize:  rc.ReadOnlySpecs.GetInstanceSize(),
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
						MinInstanceSize:  rc.AutoScaling.Compute.GetMinInstanceSize(),
						MaxInstanceSize:  rc.AutoScaling.Compute.GetMaxInstanceSize(),
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
				BackingProviderName: rc.GetBackingProviderName(),
				Priority:            rc.Priority,
				ProviderName:        rc.GetProviderName(),
				RegionName:          rc.GetRegionName(),
			})
		}
		result = append(result, replicationSpec)
	}
	return result
}

func BuildServerlessDeployments(deploymentStore store.AtlasOperatorClusterStore, validator features.FeatureValidator, projectID, projectName, clusterID, targetNamespace string, dictionary map[string]string, version string) (*atlasV1.AtlasDeployment, error) {
	deployment, err := deploymentStore.ServerlessInstance(projectID, clusterID)
	if err != nil {
		return nil, err
	}

	if !isServerlessExportable(deployment) {
		return nil, nil
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
	}

	atlasDeployment := &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, deployment.Name), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
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

	if validator.FeatureExist(features.ResourceAtlasDeployment, featureServerlessPrivateEndpoints) {
		privateEndpoints, err := buildServerlessPrivateEndpoints(deploymentStore, projectID, deployment.Name)
		if err != nil {
			return nil, err
		}
		atlasDeployment.Spec.ServerlessSpec.PrivateEndpoints = privateEndpoints
	}

	return atlasDeployment, nil
}

func buildServerlessPrivateEndpoints(deploymentStore store.ServerlessPrivateEndpointsLister, projectID, clusterName string) ([]atlasV1.ServerlessPrivateEndpoint, error) {
	endpoints, err := deploymentStore.ServerlessPrivateEndpoints(projectID, clusterName)
	if err != nil {
		return nil, err
	}

	result := make([]atlasV1.ServerlessPrivateEndpoint, 0, len(endpoints))

	for i := range endpoints {
		endpoint := endpoints[i]

		switch endpoint.GetProviderName() {
		case "AWS":
			result = append(result, atlasV1.ServerlessPrivateEndpoint{
				Name:                     endpoint.GetComment(),
				CloudProviderEndpointID:  endpoint.GetCloudProviderEndpointId(),
				PrivateEndpointIPAddress: "",
			})
		case "AZURE":
			result = append(result, atlasV1.ServerlessPrivateEndpoint{
				Name:                     endpoint.GetComment(),
				CloudProviderEndpointID:  endpoint.GetCloudProviderEndpointId(),
				PrivateEndpointIPAddress: endpoint.GetPrivateEndpointIpAddress(),
			})
		}
	}
	return result, nil
}
