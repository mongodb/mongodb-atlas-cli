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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2provider "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/provider"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	featureProcessArgs                = "processArgs"
	featureBackupSchedule             = "backupRef"
	featureServerlessPrivateEndpoints = "serverlessSpec.privateEndpoints"
	featureGlobalDeployments          = "deploymentSpec.customZoneMapping"
	DeletingState                     = "DELETING"
	DeletedState                      = "DELETED"
)

type AtlasDeploymentResult struct {
	Deployment     *akov2.AtlasDeployment
	BackupSchedule *akov2.AtlasBackupSchedule
	BackupPolicies []*akov2.AtlasBackupPolicy
}

func BuildAtlasAdvancedDeployment(deploymentStore store.OperatorClusterStore, validator features.FeatureValidator, projectID, projectName, clusterID, targetNamespace string, dictionary map[string]string, version string) (*AtlasDeploymentResult, error) {
	deployment, err := deploymentStore.AtlasCluster(projectID, clusterID)
	if err != nil {
		return nil, err
	}

	if !isAdvancedDeploymentExportable(deployment) {
		return nil, nil
	}

	var advancedSpec *akov2.AdvancedDeploymentSpec

	convertBiConnector := func(biConnector *atlasv2.BiConnector) *akov2.BiConnectorSpec {
		if biConnector == nil {
			return nil
		}
		return &akov2.BiConnectorSpec{
			Enabled:        biConnector.Enabled,
			ReadPreference: biConnector.GetReadPreference(),
		}
	}

	convertLabels := func(labels []atlasv2.ComponentLabel) []akov2common.LabelSpec {
		result := make([]akov2common.LabelSpec, 0, len(labels))

		for _, atlasLabel := range labels {
			result = append(result, akov2common.LabelSpec{
				Key:   atlasLabel.GetKey(),
				Value: atlasLabel.GetValue(),
			})
		}
		return result
	}

	convertTags := func(tags []atlasv2.ResourceTag) []*akov2.TagSpec {
		result := make([]*akov2.TagSpec, 0, len(tags))

		for _, atlasTag := range tags {
			result = append(result, &akov2.TagSpec{
				Key:   atlasTag.GetKey(),
				Value: atlasTag.GetValue(),
			})
		}
		return result
	}

	replicationSpec := buildReplicationSpec(deployment.GetReplicationSpecs())

	// TODO: DiskSizeGB field skipped on purpose. See https://jira.mongodb.org/browse/CLOUDP-146469
	advancedSpec = &akov2.AdvancedDeploymentSpec{
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
		Tags:                     convertTags(deployment.GetTags()),
	}

	atlasDeployment := &akov2.AtlasDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, clusterID), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: akov2.AtlasDeploymentSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: targetNamespace,
			},
			DeploymentSpec: advancedSpec,
			ServerlessSpec: nil,
			ProcessArgs:    nil,
		},
		Status: akov2status.AtlasDeploymentStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
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
		var backupScheduleRef akov2common.ResourceRefNamespaced
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
		customZoneMapping, managedNamespaces, err := buildGlobalDeployment(deployment.GetReplicationSpecs(), deploymentStore, projectID, clusterID)
		if err != nil {
			return nil, err
		}
		advancedSpec.CustomZoneMapping = customZoneMapping
		advancedSpec.ManagedNamespaces = managedNamespaces
	}

	if hasTenantRegionConfig(atlasDeployment) {
		atlasDeployment.Spec.DeploymentSpec.BiConnector = nil
		atlasDeployment.Spec.DeploymentSpec.EncryptionAtRestProvider = ""
		atlasDeployment.Spec.DeploymentSpec.DiskSizeGB = nil
		atlasDeployment.Spec.DeploymentSpec.MongoDBMajorVersion = ""
		atlasDeployment.Spec.DeploymentSpec.PitEnabled = nil
		atlasDeployment.Spec.DeploymentSpec.BackupEnabled = nil
	}

	return deploymentResult, nil
}

func hasTenantRegionConfig(out *akov2.AtlasDeployment) bool {
	if out.Spec.DeploymentSpec == nil {
		return false
	}
	for _, spec := range out.Spec.DeploymentSpec.ReplicationSpecs {
		if spec == nil {
			continue
		}
		for _, c := range spec.RegionConfigs {
			if c != nil && c.ProviderName == "TENANT" {
				return true
			}
		}
	}
	return false
}

func buildGlobalDeployment(atlasRepSpec []atlasv2.ReplicationSpec, globalDeploymentProvider store.GlobalClusterDescriber, projectID, clusterID string) ([]akov2.CustomZoneMapping, []akov2.ManagedNamespace, error) {
	globalCluster, err := globalDeploymentProvider.GlobalCluster(projectID, clusterID)
	if err != nil {
		return nil, nil, err
	}
	var customZoneMapping []akov2.CustomZoneMapping
	if globalCluster.CustomZoneMapping != nil {
		// create map ID -> Name for zones
		zoneMap := make(map[string]string, len(atlasRepSpec))
		for _, rc := range atlasRepSpec {
			zoneMap[rc.GetId()] = rc.GetZoneName()
		}

		customZoneMapping = make([]akov2.CustomZoneMapping, 0, len(globalCluster.GetCustomZoneMapping()))
		for location, zoneID := range globalCluster.GetCustomZoneMapping() {
			customZoneMapping = append(customZoneMapping, akov2.CustomZoneMapping{
				Zone:     zoneMap[zoneID],
				Location: location,
			})
		}
	}

	if globalCluster.ManagedNamespaces == nil {
		return customZoneMapping, nil, nil
	}

	managedNamespace := make([]akov2.ManagedNamespace, len(globalCluster.GetManagedNamespaces()))
	for i, ns := range globalCluster.GetManagedNamespaces() {
		managedNamespace[i] = akov2.ManagedNamespace{
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
func buildProcessArgs(configOptsProvider store.AtlasClusterConfigurationOptionsDescriber, projectID, clusterName string) (*akov2.ProcessArgs, error) {
	pArgs, err := configOptsProvider.AtlasClusterConfigurationOptions(projectID, clusterName)
	if err != nil {
		return nil, err
	}

	// TODO: OplogMinRetentionHours is not exported due to a bug https://jira.mongodb.org/browse/CLOUDP-146481
	return &akov2.ProcessArgs{
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

func isServerlessExportable(deployment *atlasv2.ServerlessInstanceDescription) bool {
	stateName := deployment.GetStateName()
	if stateName == DeletingState || stateName == DeletedState {
		return false
	}
	return true
}

func buildBackups(backupsProvider store.ScheduleDescriber, projectName, projectID, clusterName, targetNamespace, version string, dictionary map[string]string) (*akov2.AtlasBackupSchedule, []*akov2.AtlasBackupPolicy) {
	bs, err := backupsProvider.DescribeSchedule(projectID, clusterName)
	if err != nil {
		return nil, nil
	}

	// Although we have a for loop here, there should be only one policy per schedule. See Atlas API implementation
	policies := make([]*akov2.AtlasBackupPolicy, 0, len(bs.GetPolicies()))
	for _, p := range bs.GetPolicies() {
		items := make([]akov2.AtlasBackupPolicyItem, 0, len(p.GetPolicyItems()))
		for _, pItem := range p.GetPolicyItems() {
			items = append(items, akov2.AtlasBackupPolicyItem{
				FrequencyType:     pItem.FrequencyType,
				FrequencyInterval: pItem.FrequencyInterval,
				RetentionUnit:     pItem.RetentionUnit,
				RetentionValue:    pItem.RetentionValue,
			})
		}
		policies = append(policies, &akov2.AtlasBackupPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasBackupPolicy",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: version,
				},
			},
			Spec: akov2.AtlasBackupPolicySpec{
				Items: items,
			},
			Status: akov2status.BackupPolicyStatus{},
		})
	}

	var export *akov2.AtlasBackupExportSpec
	if bs.Export != nil {
		export = &akov2.AtlasBackupExportSpec{
			ExportBucketID: bs.Export.GetExportBucketId(),
			FrequencyType:  bs.Export.GetFrequencyType(),
		}
	}

	schedule := &akov2.AtlasBackupSchedule{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasBackupSchedule",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: akov2.AtlasBackupScheduleSpec{
			AutoExportEnabled: bs.GetAutoExportEnabled(),
			Export:            export,
			PolicyRef: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(policies[0].Name, dictionary),
				Namespace: policies[0].Namespace,
			},
			ReferenceHourOfDay:                int64(bs.GetReferenceHourOfDay()),
			ReferenceMinuteOfHour:             int64(bs.GetReferenceMinuteOfHour()),
			RestoreWindowDays:                 int64(bs.GetRestoreWindowDays()),
			UpdateSnapshots:                   bs.GetUpdateSnapshots(),
			UseOrgAndGroupNamesInExportPrefix: bs.GetUseOrgAndGroupNamesInExportPrefix(),
		},
		Status: akov2status.BackupScheduleStatus{},
	}

	if len(bs.GetCopySettings()) > 0 {
		copySettings := make([]akov2.CopySetting, 0, len(bs.GetCopySettings()))

		for _, copySetting := range bs.GetCopySettings() {
			copySettings = append(
				copySettings,
				akov2.CopySetting{
					CloudProvider:    copySetting.CloudProvider,
					RegionName:       copySetting.RegionName,
					ShouldCopyOplogs: copySetting.ShouldCopyOplogs,
					Frequencies:      copySetting.GetFrequencies(),
				},
			)
		}

		schedule.Spec.CopySettings = copySettings
	}

	return schedule, policies
}

func buildReplicationSpec(atlasRepSpec []atlasv2.ReplicationSpec) []*akov2.AdvancedReplicationSpec {
	result := make([]*akov2.AdvancedReplicationSpec, 0, len(atlasRepSpec))
	for _, rs := range atlasRepSpec {
		replicationSpec := &akov2.AdvancedReplicationSpec{
			NumShards:     rs.GetNumShards(),
			ZoneName:      rs.GetZoneName(),
			RegionConfigs: nil,
		}

		if rs.RegionConfigs == nil {
			result = append(result, replicationSpec)
			continue
		}

		replicationSpec.RegionConfigs = make([]*akov2.AdvancedRegionConfig, 0, len(replicationSpec.RegionConfigs))
		for _, rc := range rs.GetRegionConfigs() {
			var analyticsSpecs *akov2.Specs
			if rc.AnalyticsSpecs != nil {
				analyticsSpecs = &akov2.Specs{
					DiskIOPS:      convert.IntToInt64(rc.AnalyticsSpecs.DiskIOPS),
					EbsVolumeType: rc.AnalyticsSpecs.GetEbsVolumeType(),
					InstanceSize:  rc.AnalyticsSpecs.GetInstanceSize(),
					NodeCount:     rc.AnalyticsSpecs.NodeCount,
				}
			}
			var electableSpecs *akov2.Specs
			if rc.ElectableSpecs != nil {
				electableSpecs = &akov2.Specs{
					DiskIOPS:      convert.IntToInt64(rc.ElectableSpecs.DiskIOPS),
					EbsVolumeType: rc.ElectableSpecs.GetEbsVolumeType(),
					InstanceSize:  rc.ElectableSpecs.GetInstanceSize(),
					NodeCount:     rc.ElectableSpecs.NodeCount,
				}
			}

			var readOnlySpecs *akov2.Specs
			if rc.ReadOnlySpecs != nil {
				readOnlySpecs = &akov2.Specs{
					DiskIOPS:      convert.IntToInt64(rc.ReadOnlySpecs.DiskIOPS),
					EbsVolumeType: rc.ReadOnlySpecs.GetEbsVolumeType(),
					InstanceSize:  rc.ReadOnlySpecs.GetInstanceSize(),
					NodeCount:     rc.ReadOnlySpecs.NodeCount,
				}
			}

			var autoscalingSpec *akov2.AdvancedAutoScalingSpec
			if rc.AutoScaling != nil {
				var compute *akov2.ComputeSpec
				if rc.AutoScaling.Compute != nil {
					compute = &akov2.ComputeSpec{
						Enabled:          rc.AutoScaling.Compute.Enabled,
						ScaleDownEnabled: rc.AutoScaling.Compute.ScaleDownEnabled,
						MinInstanceSize:  rc.AutoScaling.Compute.GetMinInstanceSize(),
						MaxInstanceSize:  rc.AutoScaling.Compute.GetMaxInstanceSize(),
					}
				}

				var diskGB *akov2.DiskGB
				if rc.AutoScaling.DiskGB != nil {
					diskGB = &akov2.DiskGB{Enabled: rc.AutoScaling.DiskGB.Enabled}
				}
				autoscalingSpec = &akov2.AdvancedAutoScalingSpec{
					DiskGB:  diskGB,
					Compute: compute,
				}
			}
			replicationSpec.RegionConfigs = append(replicationSpec.RegionConfigs, &akov2.AdvancedRegionConfig{
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

func BuildServerlessDeployments(deploymentStore store.OperatorClusterStore, validator features.FeatureValidator, projectID, projectName, clusterID, targetNamespace string, dictionary map[string]string, version string) (*akov2.AtlasDeployment, error) {
	deployment, err := deploymentStore.GetServerlessInstance(projectID, clusterID)
	if err != nil {
		return nil, err
	}

	if !isServerlessExportable(deployment) {
		return nil, nil
	}

	providerSettings := &akov2.ServerlessProviderSettingsSpec{
		BackingProviderName: deployment.ProviderSettings.BackingProviderName,
		ProviderName:        akov2provider.ProviderName(deployment.ProviderSettings.GetProviderName()),
		RegionName:          deployment.ProviderSettings.RegionName,
	}

	serverlessSpec := &akov2.ServerlessSpec{
		Name:             deployment.GetName(),
		ProviderSettings: providerSettings,
	}

	atlasName := fmt.Sprintf("%s-%s", projectName, deployment.GetName())
	atlasDeployment := &akov2.AtlasDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(atlasName, dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: akov2.AtlasDeploymentSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: targetNamespace,
			},
			BackupScheduleRef: akov2common.ResourceRefNamespaced{},
			ServerlessSpec:    serverlessSpec,
			ProcessArgs:       nil,
		},
		Status: akov2status.AtlasDeploymentStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}

	if validator.FeatureExist(features.ResourceAtlasDeployment, featureServerlessPrivateEndpoints) {
		privateEndpoints, err := buildServerlessPrivateEndpoints(deploymentStore, projectID, deployment.GetName())
		if err != nil {
			return nil, err
		}
		atlasDeployment.Spec.ServerlessSpec.PrivateEndpoints = privateEndpoints
	}

	return atlasDeployment, nil
}

func buildServerlessPrivateEndpoints(deploymentStore store.ServerlessPrivateEndpointsLister, projectID, clusterName string) ([]akov2.ServerlessPrivateEndpoint, error) {
	endpoints, err := deploymentStore.ServerlessPrivateEndpoints(projectID, clusterName)
	if err != nil {
		return nil, err
	}

	result := make([]akov2.ServerlessPrivateEndpoint, 0, len(endpoints))

	for i := range endpoints {
		endpoint := endpoints[i]

		switch endpoint.GetProviderName() {
		case "AWS":
			result = append(result, akov2.ServerlessPrivateEndpoint{
				Name:                     endpoint.GetComment(),
				CloudProviderEndpointID:  endpoint.GetCloudProviderEndpointId(),
				PrivateEndpointIPAddress: "",
			})
		case "AZURE":
			result = append(result, akov2.ServerlessPrivateEndpoint{
				Name:                     endpoint.GetComment(),
				CloudProviderEndpointID:  endpoint.GetCloudProviderEndpointId(),
				PrivateEndpointIPAddress: endpoint.GetPrivateEndpointIpAddress(),
			})
		}
	}
	return result, nil
}
