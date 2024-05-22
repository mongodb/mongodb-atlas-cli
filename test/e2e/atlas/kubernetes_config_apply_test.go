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

//go:build e2e || (atlas && cluster && kubernetes && apply)

package atlas_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestKubernetesConfigApply(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("should fail to apply resources when namespace do not exist", func(t *testing.T) {
		operator := setupCluster(t, "k8s-config-apply-wrong-ns", defaultOperatorNamespace)
		err = operator.installOperator(defaultOperatorNamespace, features.LatestOperatorMajorVersion)
		require.NoError(t, err)

		g := newAtlasE2ETestGenerator(t)
		g.generateProject("k8sConfigApplyWrongNs")

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--targetNamespace",
			"a-wrong-namespace",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.Error(t, err, string(resp))
		assert.Equal(t, "Error: namespaces \"a-wrong-namespace\" not found\n", string(resp))
	})

	t.Run("should fail to apply resources when unable to autodetect parameters", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)

		setupCluster(t, "k8s-config-apply-no-auto-detect", defaultOperatorNamespace)

		g.generateProject("k8sConfigApplyNoAutoDetect")

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--targetNamespace", "e2e-autodetect-parameters",
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.Error(t, err, string(resp))
		assert.Equal(t, "Error: unable to auto detect params: couldn't find an operator installed in any accessible namespace\n", string(resp))
	})

	t.Run("should fail to apply resources when unable to autodetect operator version", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)

		operator := setupCluster(t, "k8s-config-apply-fail-version", defaultOperatorNamespace)
		err = operator.installOperator(defaultOperatorNamespace, features.LatestOperatorMajorVersion)
		require.NoError(t, err)

		operator.emulateCertifiedOperator()
		g.t.Cleanup(func() {
			operator.restoreOperatorImage()
		})

		g.generateProject("k8sConfigApplyFailVersion")

		e2eNamespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "e2e-autodetect-operator-version",
			},
		}
		require.NoError(t, operator.createK8sObject(e2eNamespace))
		g.t.Cleanup(func() {
			require.NoError(t, operator.deleteK8sObject(e2eNamespace))
		})

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--targetNamespace", defaultOperatorNamespace,
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.Error(t, err, string(resp))
		assert.Equal(t, "Error: unable to auto detect operator version. you should explicitly set operator version if you are running an openshift certified installation\n", string(resp))
	})

	t.Run("export and apply atlas resource to kubernetes cluster", func(t *testing.T) {
		operator := setupCluster(t, "k8s-config-apply", defaultOperatorNamespace)
		err = operator.installOperator(defaultOperatorNamespace, features.LatestOperatorMajorVersion)
		require.NoError(t, err)

		// we don't want the operator to do reconcile and avoid conflict with cli actions
		operator.stopOperator()
		t.Cleanup(func() {
			operator.startOperator()
		})

		g := setupAtlasResources(t)

		e2eNamespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "e2e-export-atlas-resource",
			},
		}
		require.NoError(t, operator.createK8sObject(e2eNamespace))
		g.t.Cleanup(func() {
			require.NoError(t, operator.deleteK8sObject(e2eNamespace))
		})

		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"apply",
			"--targetNamespace", "e2e-export-atlas-resource",
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		t.Log(string(resp))
		g.t.Cleanup(func() {
			operator.cleanUpResources()
		})

		akoProject := akov2.AtlasProject{}
		require.NoError(
			t,
			operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(g.projectName), Namespace: e2eNamespace.Name},
				&akoProject,
				true,
			),
		)
		assert.NotEmpty(t, akoProject.Spec.AlertConfigurations)
		akoProject.Spec.AlertConfigurations = nil
		assert.Equal(t, referenceExportedProject(g.projectName, g.teamName, &akoProject).Spec, akoProject.Spec)

		// Assert Database User
		akoDBUser := akov2.AtlasDatabaseUser{}
		require.NoError(
			t,
			operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s", g.projectName, g.dbUser)), Namespace: e2eNamespace.Name},
				&akoDBUser,
				true,
			),
		)
		assert.Equal(t, referenceExportedDBUser(g.projectName, g.dbUser, e2eNamespace.Name).Spec, akoDBUser.Spec)

		// Assert Team
		akoTeam := akov2.AtlasTeam{}
		require.NoError(
			t,
			operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-team-%s", g.projectName, g.teamName)), Namespace: e2eNamespace.Name},
				&akoTeam,
				true,
			),
		)
		assert.Equal(t, referenceExportedTeam(g.teamName, g.teamUser).Spec, akoTeam.Spec)

		// Assert Backup Policy
		akoBkpPolicy := akov2.AtlasBackupPolicy{}
		require.NoError(
			t,
			operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s-backuppolicy", g.projectName, g.clusterName)), Namespace: e2eNamespace.Name},
				&akoBkpPolicy,
				true,
			),
		)
		assert.Equal(t, referenceExportedBackupPolicy().Spec, akoBkpPolicy.Spec)

		// Assert Backup Schedule
		akoBkpSchedule := akov2.AtlasBackupSchedule{}
		require.NoError(
			t,
			operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s-backupschedule", g.projectName, g.clusterName)), Namespace: e2eNamespace.Name},
				&akoBkpSchedule,
				true,
			),
		)
		assert.Equal(
			t,
			referenceExportedBackupSchedule(g.projectName, g.clusterName, e2eNamespace.Name, akoBkpSchedule.Spec.ReferenceHourOfDay, akoBkpSchedule.Spec.ReferenceMinuteOfHour).Spec,
			akoBkpSchedule.Spec,
		)

		// Assert Deployment
		akoDeployment := akov2.AtlasDeployment{}
		require.NoError(
			t,
			operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s", g.projectName, g.clusterName)), Namespace: e2eNamespace.Name},
				&akoDeployment,
				true,
			),
		)
		assert.Equal(t, referenceExportedDeployment(g.projectName, g.clusterName, e2eNamespace.Name).Spec, akoDeployment.Spec)
	})
}

func setupAtlasResources(t *testing.T) *atlasE2ETestGenerator {
	t.Helper()

	g := newAtlasE2ETestGeneratorWithBackup(t)
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

const credSuffix = "-credentials"

func referenceExportedProject(projectName, teamName string, expectedProject *akov2.AtlasProject) *akov2.AtlasProject {
	return &akov2.AtlasProject{
		Spec: akov2.AtlasProjectSpec{
			Name: projectName,
			ConnectionSecret: &akov2common.ResourceRefNamespaced{
				Name: prepareK8sName(projectName + credSuffix),
			},
			WithDefaultAlertsSettings: true,
			Settings: &akov2.ProjectSettings{
				IsCollectDatabaseSpecificsStatisticsEnabled: pointer.Get(true),
				IsDataExplorerEnabled:                       pointer.Get(true),
				IsPerformanceAdvisorEnabled:                 pointer.Get(true),
				IsRealtimePerformancePanelEnabled:           pointer.Get(true),
				IsSchemaAdvisorEnabled:                      pointer.Get(true),
			},
			EncryptionAtRest: &akov2.EncryptionAtRest{
				AwsKms: akov2.AwsKms{
					Enabled: pointer.Get(false),
					Valid:   pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      expectedProject.Spec.EncryptionAtRest.AwsKms.SecretRef.Name,
						Namespace: expectedProject.Spec.EncryptionAtRest.AwsKms.SecretRef.Namespace,
					},
				},
				AzureKeyVault: akov2.AzureKeyVault{
					Enabled: pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      expectedProject.Spec.EncryptionAtRest.AzureKeyVault.SecretRef.Name,
						Namespace: expectedProject.Spec.EncryptionAtRest.AzureKeyVault.SecretRef.Namespace,
					},
				},
				GoogleCloudKms: akov2.GoogleCloudKms{
					Enabled: pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      expectedProject.Spec.EncryptionAtRest.GoogleCloudKms.SecretRef.Name,
						Namespace: expectedProject.Spec.EncryptionAtRest.GoogleCloudKms.SecretRef.Namespace,
					},
				},
			},
			Auditing: &akov2.Auditing{
				AuditAuthorizationSuccess: false,
				Enabled:                   false,
			},
			Teams: []akov2.Team{
				{
					TeamRef: akov2common.ResourceRefNamespaced{
						Namespace: expectedProject.Namespace,
						Name:      prepareK8sName(fmt.Sprintf("%s-team-%s", projectName, teamName)),
					},
					Roles: []akov2.TeamRole{
						"GROUP_OWNER",
					},
				},
			},
			RegionUsageRestrictions: "NONE",
		},
	}
}

func referenceExportedDBUser(projectName, dbUser, namespace string) *akov2.AtlasDatabaseUser {
	return &akov2.AtlasDatabaseUser{
		Spec: akov2.AtlasDatabaseUserSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      prepareK8sName(projectName),
				Namespace: namespace,
			},
			Roles: []akov2.RoleSpec{
				{
					RoleName:     "readAnyDatabase",
					DatabaseName: "admin",
				},
			},
			Username:     dbUser,
			OIDCAuthType: "NONE",
			AWSIAMType:   "NONE",
			X509Type:     "MANAGED",
			DatabaseName: "$external",
		},
	}
}

func referenceExportedTeam(teamName, username string) *akov2.AtlasTeam {
	return &akov2.AtlasTeam{
		Spec: akov2.TeamSpec{
			Name: teamName,
			Usernames: []akov2.TeamUser{
				akov2.TeamUser(username),
			},
		},
	}
}

func referenceExportedBackupSchedule(projectName, clusterName, namespace string, refHour, refMin int64) *akov2.AtlasBackupSchedule {
	return &akov2.AtlasBackupSchedule{
		Spec: akov2.AtlasBackupScheduleSpec{
			PolicyRef: akov2common.ResourceRefNamespaced{
				Name:      prepareK8sName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName)),
				Namespace: namespace,
			},
			AutoExportEnabled:     false,
			ReferenceHourOfDay:    refHour,
			ReferenceMinuteOfHour: refMin,
			RestoreWindowDays:     7,
		},
	}
}

func referenceExportedBackupPolicy() *akov2.AtlasBackupPolicy {
	return &akov2.AtlasBackupPolicy{
		Spec: akov2.AtlasBackupPolicySpec{
			Items: []akov2.AtlasBackupPolicyItem{
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
				{
					FrequencyType:     "yearly",
					FrequencyInterval: 12,
					RetentionUnit:     "years",
					RetentionValue:    1,
				},
			},
		},
	}
}

func referenceExportedDeployment(projectName, clusterName, namespace string) *akov2.AtlasDeployment {
	return &akov2.AtlasDeployment{
		Spec: akov2.AtlasDeploymentSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      prepareK8sName(projectName),
				Namespace: namespace,
			},
			BackupScheduleRef: akov2common.ResourceRefNamespaced{
				Name:      prepareK8sName(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName)),
				Namespace: namespace,
			},
			DeploymentSpec: &akov2.AdvancedDeploymentSpec{
				Name:          clusterName,
				BackupEnabled: pointer.Get(true),
				BiConnector: &akov2.BiConnectorSpec{
					Enabled:        pointer.Get(false),
					ReadPreference: "secondary",
				},
				ClusterType:              "REPLICASET",
				EncryptionAtRestProvider: "NONE",
				Paused:                   pointer.Get(false),
				PitEnabled:               pointer.Get(true),
				ReplicationSpecs: []*akov2.AdvancedReplicationSpec{
					{
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*akov2.AdvancedRegionConfig{
							{
								AnalyticsSpecs: &akov2.Specs{
									DiskIOPS:      pointer.Get[int64](3000),
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     pointer.Get(0),
								},
								ElectableSpecs: &akov2.Specs{
									DiskIOPS:      pointer.Get[int64](3000),
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     pointer.Get(3),
								},
								ReadOnlySpecs: &akov2.Specs{
									DiskIOPS:      pointer.Get[int64](3000),
									EbsVolumeType: "STANDARD",
									InstanceSize:  "M10",
									NodeCount:     pointer.Get(0),
								},
								AutoScaling: &akov2.AdvancedAutoScalingSpec{
									DiskGB: &akov2.DiskGB{
										Enabled: pointer.Get(false),
									},
									Compute: &akov2.ComputeSpec{
										Enabled:          pointer.Get(false),
										ScaleDownEnabled: pointer.Get(false),
									},
								},
								Priority:     pointer.Get(7),
								ProviderName: "AWS",
								RegionName:   "US_EAST_1",
							},
						},
					},
				},
				RootCertType:         "ISRGROOTX1",
				VersionReleaseSystem: "LTS",
			},
			ProcessArgs: &akov2.ProcessArgs{
				MinimumEnabledTLSProtocol: "TLS1_2",
				JavascriptEnabled:         pointer.Get(true),
				NoTableScan:               pointer.Get(false),
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
	require.NoError(t, err, string(resp))
}

func prepareK8sName(pattern string) string {
	return resources.NormalizeAtlasName(pattern, resources.AtlasNameToKubernetesName())
}
