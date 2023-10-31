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

//go:build e2e || (atlas && cluster && kubernetes)

package atlas_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	akov1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/util/toptr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestKubernetesConfigApply(t *testing.T) {
	a := assert.New(t)
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	t.Log(cliPath)
	req.NoError(err)

	t.Run("should failed to apply resources when namespace doesn't exist", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)
		g.generateProject("k8sConfigApplyWrongNs")

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--targetNamespace", "a-wrong-namespace",
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err, string(resp))
		a.Equal("Error: namespaces \"a-wrong-namespace\" not found\n", string(resp))
	})

	t.Run("should failed to apply resources when unable to autodetect parameters", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)
		g.generateProject("k8sConfigApplyNoAutoDetect")

		operator, err := newOperatorHelper(t)
		req.NoError(err)
		operator.deleteOperator()
		g.t.Cleanup(func() {
			operator.restoreOperator()
		})

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err, string(resp))
		a.Equal("Error: unable to auto detect params: couldn't find an operator installed in any accessible namespace\n", string(resp))
	})

	t.Run("should failed to apply resources when unable to autodetect operator version", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)
		g.generateProject("k8sConfigApplyFailVersion")

		operator, err := newOperatorHelper(t)
		req.NoError(err)
		operator.emulateCertifiedOperator()
		g.t.Cleanup(func() {
			operator.restoreOperatorImage()
		})

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err, string(resp))
		a.Equal("Error: unable to auto detect operator version. you should explicitly set operator version if you are running an openshift certified installation\n", string(resp))
	})

	t.Run("export and apply atlas resource to kubernetes cluster", func(t *testing.T) {
		g := setupAtlasResources(t)

		operator, err := newOperatorHelper(t)
		req.NoError(err)
		// we don't want the operator to do reconcile and avoid conflict with cli actions
		operator.stopOperator()
		g.t.Cleanup(func() {
			operator.startOperator()
		})

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		t.Log(string(resp))
		g.t.Cleanup(func() {
			operator.cleanUpResources()
		})

		namespace := "mongodb-atlas-system"
		akoProject := akov1.AtlasProject{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(g.projectName), Namespace: namespace},
			&akoProject,
			true,
		)
		req.NoError(err)
		a.NotEmpty(akoProject.Spec.AlertConfigurations)
		akoProject.Spec.AlertConfigurations = nil
		a.Equal(referenceExportedProject(g.projectName, g.teamName).Spec, akoProject.Spec)

		// Assert Database User
		akoDBUser := akov1.AtlasDatabaseUser{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s", g.projectName, g.dbUser)), Namespace: namespace},
			&akoDBUser,
			true,
		)
		req.NoError(err)
		a.Equal(referenceExportedDBUser(g.projectName, g.dbUser).Spec, akoDBUser.Spec)

		// Assert Team
		akoTeam := akov1.AtlasTeam{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-team-%s", g.projectName, g.teamName)), Namespace: namespace},
			&akoTeam,
			true,
		)
		req.NoError(err)
		a.Equal(referenceExportedTeam(g.teamName, g.teamUser).Spec, akoTeam.Spec)

		// Assert Backup Policy
		akoBkpPolicy := akov1.AtlasBackupPolicy{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s-backuppolicy", g.projectName, g.clusterName)), Namespace: namespace},
			&akoBkpPolicy,
			true,
		)
		req.NoError(err)
		a.Equal(referenceExportedBackupPolicy().Spec, akoBkpPolicy.Spec)

		// Assert Backup Schedule
		akoBkpSchedule := akov1.AtlasBackupSchedule{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s-backupschedule", g.projectName, g.clusterName)), Namespace: namespace},
			&akoBkpSchedule,
			true,
		)
		req.NoError(err)
		a.Equal(
			referenceExportedBackupSchedule(g.projectName, g.clusterName, akoBkpSchedule.Spec.ReferenceHourOfDay, akoBkpSchedule.Spec.ReferenceMinuteOfHour).Spec,
			akoBkpSchedule.Spec,
		)

		// Assert Deployment
		akoDeployment := akov1.AtlasDeployment{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s", g.projectName, g.clusterName)), Namespace: namespace},
			&akoDeployment,
			true,
		)
		req.NoError(err)
		a.Equal(referenceExportedDeployment(g.projectName, g.clusterName).Spec, akoDeployment.Spec)
	})
}

func setupAtlasResources(t *testing.T) *atlasE2ETestGenerator {
	t.Helper()

	g := newAtlasE2ETestGenerator(t)
	g.enableBackup = true
	g.generateProject("k8sConfigApply")
	g.generateCluster()
	g.generateTeam("k8sConfigApply")
	g.generateDBUser("k8sConfigApply")

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	cmd := exec.Command(cliPath,
		projectsEntity,
		teamsEntity,
		"add",
		g.teamID,
		"--role",
		"GROUP_OWNER",
		"--projectId",
		g.projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	require.NoError(t, err, string(resp))
	g.t.Cleanup(func() {
		deleteTeamFromProject(g.t, cliPath, g.projectID, g.teamID)
	})

	return g
}

func referenceExportedProject(projectName, teamName string) *akov1.AtlasProject {
	return &akov1.AtlasProject{
		Spec: akov1.AtlasProjectSpec{
			Name: projectName,
			ConnectionSecret: &common.ResourceRefNamespaced{
				Name: prepareK8sName(fmt.Sprintf("%s-credentials", projectName)),
			},
			WithDefaultAlertsSettings: true,
			Settings: &akov1.ProjectSettings{
				IsCollectDatabaseSpecificsStatisticsEnabled: toptr.MakePtr(true),
				IsDataExplorerEnabled:                       toptr.MakePtr(true),
				IsPerformanceAdvisorEnabled:                 toptr.MakePtr(true),
				IsRealtimePerformancePanelEnabled:           toptr.MakePtr(true),
				IsSchemaAdvisorEnabled:                      toptr.MakePtr(true),
			},
			EncryptionAtRest: &akov1.EncryptionAtRest{
				AwsKms: akov1.AwsKms{
					Enabled: toptr.MakePtr(false),
					Valid:   toptr.MakePtr(false),
				},
				AzureKeyVault: akov1.AzureKeyVault{
					Enabled: toptr.MakePtr(false),
				},
				GoogleCloudKms: akov1.GoogleCloudKms{
					Enabled: toptr.MakePtr(false),
				},
			},
			Auditing: &akov1.Auditing{
				AuditAuthorizationSuccess: false,
				Enabled:                   false,
			},
			Teams: []akov1.Team{
				{
					TeamRef: common.ResourceRefNamespaced{
						Namespace: "mongodb-atlas-system",
						Name:      prepareK8sName(fmt.Sprintf("%s-team-%s", projectName, teamName)),
					},
					Roles: []akov1.TeamRole{
						"GROUP_OWNER",
					},
				},
			},
			RegionUsageRestrictions: "NONE",
		},
	}
}

func referenceExportedDBUser(projectName, dbUser string) *akov1.AtlasDatabaseUser {
	return &akov1.AtlasDatabaseUser{
		Spec: akov1.AtlasDatabaseUserSpec{
			Project: common.ResourceRefNamespaced{
				Name:      prepareK8sName(projectName),
				Namespace: "mongodb-atlas-system",
			},
			Roles: []akov1.RoleSpec{
				{
					RoleName:     "readAnyDatabase",
					DatabaseName: "admin",
				},
			},
			Username:     dbUser,
			X509Type:     "MANAGED",
			DatabaseName: "$external",
		},
	}
}

func referenceExportedTeam(teamName, username string) *akov1.AtlasTeam {
	return &akov1.AtlasTeam{
		Spec: akov1.TeamSpec{
			Name: teamName,
			Usernames: []akov1.TeamUser{
				akov1.TeamUser(username),
			},
		},
	}
}

func referenceExportedBackupSchedule(projectName, clusterName string, refHour, refMin int64) *akov1.AtlasBackupSchedule {
	return &akov1.AtlasBackupSchedule{
		Spec: akov1.AtlasBackupScheduleSpec{
			PolicyRef: common.ResourceRefNamespaced{
				Name:      prepareK8sName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName)),
				Namespace: "mongodb-atlas-system",
			},
			AutoExportEnabled:     false,
			ReferenceHourOfDay:    refHour,
			ReferenceMinuteOfHour: refMin,
			RestoreWindowDays:     7,
		},
	}
}

func referenceExportedBackupPolicy() *akov1.AtlasBackupPolicy {
	return &akov1.AtlasBackupPolicy{
		Spec: akov1.AtlasBackupPolicySpec{
			Items: []akov1.AtlasBackupPolicyItem{
				{
					FrequencyType:     "hourly",
					FrequencyInterval: 6,
					RetentionUnit:     "days",
					RetentionValue:    7,
				},
				{
					FrequencyType:     "daily",
					FrequencyInterval: 1,
					RetentionUnit:     "days",
					RetentionValue:    7,
				},
				{
					FrequencyType:     "weekly",
					FrequencyInterval: 6,
					RetentionUnit:     "weeks",
					RetentionValue:    4,
				},
				{
					FrequencyType:     "monthly",
					FrequencyInterval: 40,
					RetentionUnit:     "months",
					RetentionValue:    12,
				},
			},
		},
	}
}

func referenceExportedDeployment(projectName, clusterName string) *akov1.AtlasDeployment {
	return &akov1.AtlasDeployment{
		Spec: akov1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      prepareK8sName(projectName),
				Namespace: "mongodb-atlas-system",
			},
			BackupScheduleRef: common.ResourceRefNamespaced{
				Name:      prepareK8sName(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName)),
				Namespace: "mongodb-atlas-system",
			},
			AdvancedDeploymentSpec: &akov1.AdvancedDeploymentSpec{
				Name:          clusterName,
				BackupEnabled: toptr.MakePtr(true),
				BiConnector: &akov1.BiConnectorSpec{
					Enabled:        toptr.MakePtr(false),
					ReadPreference: "secondary",
				},
				ClusterType:              "REPLICASET",
				EncryptionAtRestProvider: "NONE",
				Labels: []common.LabelSpec{
					{
						Key:   "Infrastructure Tool",
						Value: "Atlas CLI",
					},
				},
				Paused:     toptr.MakePtr(false),
				PitEnabled: toptr.MakePtr(true),
				ReplicationSpecs: []*akov1.AdvancedReplicationSpec{
					{
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*akov1.AdvancedRegionConfig{
							{
								AnalyticsSpecs: &akov1.Specs{
									DiskIOPS:      toptr.MakePtr(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     toptr.MakePtr(0),
								},
								ElectableSpecs: &akov1.Specs{
									DiskIOPS:      toptr.MakePtr(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     toptr.MakePtr(3),
								},
								ReadOnlySpecs: &akov1.Specs{
									DiskIOPS:      toptr.MakePtr(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     toptr.MakePtr(0),
								},
								AutoScaling: &akov1.AdvancedAutoScalingSpec{
									DiskGB: &akov1.DiskGB{
										Enabled: toptr.MakePtr(false),
									},
									Compute: &akov1.ComputeSpec{
										Enabled:          toptr.MakePtr(false),
										ScaleDownEnabled: toptr.MakePtr(false),
									},
								},
								Priority:     toptr.MakePtr(7),
								ProviderName: "AWS",
								RegionName:   "US_EAST_1",
							},
						},
					},
				},
				RootCertType:         "ISRGROOTX1",
				VersionReleaseSystem: "LTS",
			},
			ProcessArgs: &akov1.ProcessArgs{
				MinimumEnabledTLSProtocol: "TLS1_2",
				JavascriptEnabled:         toptr.MakePtr(true),
				NoTableScan:               toptr.MakePtr(false),
			},
		},
	}
}

func deleteTeamFromProject(t *testing.T, cliPath, projectID, teamID string) {
	t.Helper()

	cmd := exec.Command(cliPath,
		projectsEntity,
		teamsEntity,
		"delete",
		teamID,
		"--projectId",
		projectID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("%s (%v)", string(resp), err)
	}
}

func prepareK8sName(pattern string) string {
	return resources.NormalizeAtlasName(pattern, resources.AtlasNameToKubernetesName())
}
