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
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2project "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/project"
	akov2provider "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/provider"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115008/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

const targetNamespace = "importer-namespace"

var expectedLabels = map[string]string{
	features.ResourceVersion: features.LatestOperatorMajorVersion,
}

func getK8SEntities(data []byte) ([]runtime.Object, error) {
	b := bufio.NewReader(bytes.NewReader(data))
	r := yaml.NewYAMLReader(b)

	var result []runtime.Object

	for {
		doc, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		d := scheme.Codecs.UniversalDeserializer()
		obj, _, err := d.Decode(doc, nil, nil)
		if err != nil {
			// if document is not a K8S object, skip it
			continue
		}
		if obj != nil {
			result = append(result, obj)
		}
	}
	return result, nil
}

type KubernetesConfigGenerateProjectSuite struct {
	t               *testing.T
	assertions      *assert.Assertions
	generator       *atlasE2ETestGenerator
	expectedProject *akov2.AtlasProject
	cliPath         string
}

func InitialSetupWithTeam(t *testing.T) KubernetesConfigGenerateProjectSuite {
	t.Helper()
	s := KubernetesConfigGenerateProjectSuite{
		t: t,
	}
	s.generator = newAtlasE2ETestGenerator(t)
	s.generator.generateTeam("Kubernetes")
	s.generator.generateEmptyProject(fmt.Sprintf("Kubernetes-%s", s.generator.projectName))
	s.expectedProject = referenceProject(s.generator.projectName, targetNamespace, expectedLabels)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	s.cliPath = cliPath

	s.assertions = assert.New(t)

	// always register atlas entities
	require.NoError(t, akov2.AddToScheme(scheme.Scheme))
	return s
}

func InitialSetup(t *testing.T) KubernetesConfigGenerateProjectSuite {
	t.Helper()
	s := KubernetesConfigGenerateProjectSuite{
		t: t,
	}
	s.generator = newAtlasE2ETestGenerator(t)
	s.generator.generateEmptyProject(fmt.Sprintf("Kubernetes-%s", s.generator.projectName))
	s.expectedProject = referenceProject(s.generator.projectName, targetNamespace, expectedLabels)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	s.cliPath = cliPath

	s.assertions = assert.New(t)

	// always register atlas entities
	require.NoError(t, akov2.AddToScheme(scheme.Scheme))
	return s
}

func TestEmptyProject(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	t.Run("Generate valid resources of ONE project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			found := false
			var secret *corev1.Secret
			var ok bool
			for i := range objects {
				secret, ok = objects[i].(*corev1.Secret)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, targetNamespace, secret.Namespace)
		})
	})
}

func TestProjectWithNonDefaultSettings(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions
	expectedProject.Spec.Settings.IsCollectDatabaseSpecificsStatisticsEnabled = pointer.Get(false)

	t.Run("Change project settings and generate", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			settingsEntity,
			"update",
			"--disableCollectDatabaseSpecificsStatistics",
			"-o=json",
			"--projectId",
			generator.projectID)
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err)
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithNonDefaultAlertConf(t *testing.T) {
	dictionary := resources.AtlasNameToKubernetesName()
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	newAlertConfig := akov2.AlertConfiguration{
		Threshold:       &akov2.Threshold{},
		MetricThreshold: &akov2.MetricThreshold{},
		EventTypeName:   eventTypeName,
		Enabled:         true,
		Notifications: []akov2.Notification{
			{
				TypeName:     group,
				IntervalMin:  intervalMin,
				DelayMin:     pointer.Get(delayMin),
				SMSEnabled:   pointer.Get(false),
				EmailEnabled: pointer.Get(true),
				APITokenRef: akov2common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-api-token-0", expectedProject.Name), dictionary),
					Namespace: targetNamespace,
				},
				DatadogAPIKeyRef: akov2common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-datadog-api-key-0", expectedProject.Name), dictionary),
					Namespace: targetNamespace,
				},
				OpsGenieAPIKeyRef: akov2common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-ops-genie-api-key-0", expectedProject.Name), dictionary),
					Namespace: targetNamespace,
				},
				ServiceKeyRef: akov2common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-service-key-0", expectedProject.Name), dictionary),
					Namespace: targetNamespace,
				},
				VictorOpsSecretRef: akov2common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-victor-ops-credentials-0", expectedProject.Name), dictionary),
					Namespace: targetNamespace,
				},
			},
		},
	}
	expectedProject.Spec.AlertConfigurations = []akov2.AlertConfiguration{
		newAlertConfig,
	}

	t.Run("Change project alert config and generate", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			alertsEntity,
			configEntity,
			"create",
			"--projectId",
			generator.projectID,
			"--event",
			newAlertConfig.EventTypeName,
			fmt.Sprintf("--enabled=%s", strconv.FormatBool(newAlertConfig.Enabled)),
			"--notificationType",
			newAlertConfig.Notifications[0].TypeName,
			"--notificationIntervalMin",
			strconv.Itoa(newAlertConfig.Notifications[0].IntervalMin),
			"--notificationDelayMin",
			strconv.Itoa(*newAlertConfig.Notifications[0].DelayMin),
			fmt.Sprintf("--notificationSmsEnabled=%s", strconv.FormatBool(*newAlertConfig.Notifications[0].SMSEnabled)),
			fmt.Sprintf("--notificationEmailEnabled=%s", strconv.FormatBool(*newAlertConfig.Notifications[0].EmailEnabled)),
			"-o=json")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err)
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithAccessList(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	entry := "192.168.0.10"
	newIPAccess := akov2project.IPAccessList{
		IPAddress: entry,
		Comment:   "test",
	}
	expectedProject.Spec.ProjectIPAccessList = []akov2project.IPAccessList{
		newIPAccess,
	}

	t.Run("Add access list to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			accessListEntity,
			"create",
			newIPAccess.IPAddress,
			fmt.Sprintf("--comment=%s", newIPAccess.Comment),
			"--projectId",
			generator.projectID,
			"--type",
			"ipAddress",
			"-o=json")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithAccessRole(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	newIPAccess := akov2.CloudProviderAccessRole{
		ProviderName: string(akov2provider.ProviderAWS),
	}
	expectedProject.Spec.CloudProviderAccessRoles = []akov2.CloudProviderAccessRole{
		newIPAccess,
	}

	t.Run("Add access role to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			cloudProvidersEntity,
			accessRolesEntity,
			awsEntity,
			"create",
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err)
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithCustomRole(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	newCustomRole := akov2.CustomRole{
		Name: "test-role",
		Actions: []akov2.Action{
			{
				Name: "FIND",
				Resources: []akov2.Resource{
					{
						Database:   pointer.Get("test-db	"),
						Collection: pointer.Get(""),
						Cluster:    pointer.Get(false),
					},
				},
			},
		},
	}
	expectedProject.Spec.CustomRoles = []akov2.CustomRole{
		newCustomRole,
	}

	t.Run("Add custom role to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"create",
			newCustomRole.Name,
			"--privilege",
			fmt.Sprintf("%s@%s", newCustomRole.Actions[0].Name, *newCustomRole.Actions[0].Resources[0].Database),
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithIntegration(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	datadogKey := "00000000000000000000000000000012"
	newIntegration := akov2project.Integration{
		Type:   datadogEntity,
		Region: "US", // it's a default value
		APIKeyRef: akov2common.ResourceRefNamespaced{
			Namespace: targetNamespace,
			Name:      fmt.Sprintf("%s-integration-%s", generator.projectID, strings.ToLower(datadogEntity)),
		},
	}
	expectedProject.Spec.Integrations = []akov2project.Integration{
		newIntegration,
	}

	t.Run("Add integration to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"create",
			datadogEntity,
			"--apiKey",
			datadogKey,
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
		assertions.Len(objects, 3, "should have 3 objects in the output: project, integration secret, atlas secret")
		integrationSecret := objects[1].(*corev1.Secret)
		password, ok := integrationSecret.Data["password"]
		assertions.True(ok, "should have password field in the integration secret")
		assertions.True(compareStingsWithHiddenPart(datadogKey, string(password), uint8('*')), "should have correct password in the integration secret")
	})
}

func TestProjectWithMaintenanceWindow(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	newMaintenanceWindow := akov2project.MaintenanceWindow{
		DayOfWeek: 1,
		HourOfDay: 1,
	}
	expectedProject.Spec.MaintenanceWindow = newMaintenanceWindow
	expectedProject.Spec.AlertConfigurations = defaultMaintenanceWindowAlertConfigs()

	t.Run("Add integration to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			maintenanceEntity,
			"update",
			"--dayOfWeek",
			strconv.Itoa(newMaintenanceWindow.DayOfWeek),
			"--hourOfDay",
			strconv.Itoa(newMaintenanceWindow.HourOfDay),
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithNetworkPeering(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	atlasCidrBlock := "10.8.0.0/18"
	networkPeer := akov2.NetworkPeer{
		ProviderName: akov2provider.ProviderGCP,
		NetworkName:  "test-network",
		GCPProjectID: "test-project-gcp",
	}
	expectedProject.Spec.NetworkPeers = []akov2.NetworkPeer{
		networkPeer,
	}

	t.Run("Add network peer to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			networkingEntity,
			networkPeeringEntity,
			"create",
			gcpEntity,
			"--atlasCidrBlock",
			atlasCidrBlock,
			"--network",
			networkPeer.NetworkName,
			"--gcpProjectId",
			networkPeer.GCPProjectID,
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		t.Cleanup(func() {
			deleteAllNetworkPeers(t, cliPath, generator.projectID, gcpEntity)
		})
		var createdNetworkPeer atlasv2.BaseNetworkPeeringConnectionSettings
		err = json.Unmarshal(resp, &createdNetworkPeer)
		require.NoError(t, err)
		expectedProject.Spec.NetworkPeers[0].ContainerID = createdNetworkPeer.ContainerId

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err = cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithPrivateEndpoint_Azure(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	const region = "northeurope"
	newPrivateEndpoint := akov2.PrivateEndpoint{
		Provider: akov2provider.ProviderAzure,
		Region:   "EUROPE_NORTH",
	}
	expectedProject.Spec.PrivateEndpoints = []akov2.PrivateEndpoint{
		newPrivateEndpoint,
	}

	t.Run("Add network peer to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"create",
			"--region",
			region,
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err)
		t.Cleanup(func() {
			deleteAllPrivateEndpoints(t, cliPath, generator.projectID, azureEntity)
		})
		var createdNetworkPeer *atlasv2.EndpointService
		err = json.Unmarshal(resp, &createdNetworkPeer)
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"watch",
			createdNetworkPeer.GetId(),
			"--projectId", generator.projectID)
		cmd.Env = os.Environ()
		_, err = cmd.CombinedOutput()
		require.NoError(t, err)

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err = cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectAndTeams(t *testing.T) {
	s := InitialSetupWithTeam(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	teamRole := "GROUP_OWNER"

	t.Run("Add team to project", func(t *testing.T) {
		expectedTeam := referenceTeam(generator.teamName, targetNamespace, []akov2.TeamUser{
			akov2.TeamUser(generator.teamUser),
		}, generator.projectName, expectedLabels)

		expectedProject.Spec.Teams = []akov2.Team{
			{
				TeamRef: akov2common.ResourceRefNamespaced{
					Namespace: targetNamespace,
					Name:      expectedTeam.Name,
				},
				Roles: []akov2.TeamRole{
					akov2.TeamRole(teamRole),
				},
			},
		}

		cmd := exec.Command(cliPath,
			projectsEntity,
			teamsEntity,
			"add",
			generator.teamID,
			"--role",
			teamRole,
			"--projectId",
			generator.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		cmd = exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err = cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		checkProject(t, objects, expectedProject, assertions)
		t.Run("Team is created", func(_ *testing.T) {
			for _, obj := range objects {
				if team, ok := obj.(*akov2.AtlasTeam); ok {
					assertions.Equal(expectedTeam, team)
				}
			}
		})
	})
}

func referenceTeam(name, namespace string, users []akov2.TeamUser, projectName string, labels map[string]string) *akov2.AtlasTeam {
	dictionary := resources.AtlasNameToKubernetesName()

	return &akov2.AtlasTeam{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasTeam",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-team-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: akov2.TeamSpec{
			Name:      name,
			Usernames: users,
		},
		Status: akov2status.TeamStatus{
			Common: akov2status.Common{
				Conditions: []akov2status.Condition{},
			},
		},
	}
}

func checkProject(t *testing.T, output []runtime.Object, expected *akov2.AtlasProject, asserts *assert.Assertions) {
	t.Helper()
	t.Run("Project presents with expected data", func(t *testing.T) {
		found := false
		var p *akov2.AtlasProject
		var ok bool
		for i := range output {
			p, ok = output[i].(*akov2.AtlasProject)
			if ok {
				found = true
				break
			}
		}
		if !found {
			t.Fatal("AtlasProject is not found in results")
		}

		// secretref names are randomly generated so we can't determine those in forehand
		expected.Spec.EncryptionAtRest.AwsKms = p.Spec.EncryptionAtRest.AwsKms
		expected.Spec.EncryptionAtRest.GoogleCloudKms = p.Spec.EncryptionAtRest.GoogleCloudKms
		expected.Spec.EncryptionAtRest.AzureKeyVault = p.Spec.EncryptionAtRest.AzureKeyVault

		for i := range p.Spec.AlertConfigurations {
			alertConfig := &p.Spec.AlertConfigurations[i]
			for j := range alertConfig.Notifications {
				expected.Spec.AlertConfigurations[i].Notifications[j].APITokenRef = p.Spec.AlertConfigurations[i].Notifications[j].APITokenRef
				expected.Spec.AlertConfigurations[i].Notifications[j].DatadogAPIKeyRef = p.Spec.AlertConfigurations[i].Notifications[j].DatadogAPIKeyRef
				expected.Spec.AlertConfigurations[i].Notifications[j].OpsGenieAPIKeyRef = p.Spec.AlertConfigurations[i].Notifications[j].OpsGenieAPIKeyRef
				expected.Spec.AlertConfigurations[i].Notifications[j].ServiceKeyRef = p.Spec.AlertConfigurations[i].Notifications[j].ServiceKeyRef
				expected.Spec.AlertConfigurations[i].Notifications[j].VictorOpsSecretRef = p.Spec.AlertConfigurations[i].Notifications[j].VictorOpsSecretRef
			}
		}

		asserts.Equal(expected, p)
	})
}

func referenceProject(name, namespace string, labels map[string]string) *akov2.AtlasProject {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasProject{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasProject",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(name, dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Status: akov2status.AtlasProjectStatus{
			Common: akov2status.Common{
				Conditions: []akov2status.Condition{},
			},
		},
		Spec: akov2.AtlasProjectSpec{
			Name: name,
			ConnectionSecret: &akov2common.ResourceRefNamespaced{
				Name: resources.NormalizeAtlasName(fmt.Sprintf("%s-credentials", name), dictionary),
			},
			Settings: &akov2.ProjectSettings{
				IsCollectDatabaseSpecificsStatisticsEnabled: pointer.Get(true),
				IsDataExplorerEnabled:                       pointer.Get(true),
				IsPerformanceAdvisorEnabled:                 pointer.Get(true),
				IsRealtimePerformancePanelEnabled:           pointer.Get(true),
				IsSchemaAdvisorEnabled:                      pointer.Get(true),
			},
			Auditing: &akov2.Auditing{
				AuditAuthorizationSuccess: false,
				Enabled:                   false,
			},
			EncryptionAtRest: &akov2.EncryptionAtRest{
				AwsKms: akov2.AwsKms{
					Enabled: pointer.Get(false),
					Valid:   pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-aws-credentials", name), dictionary),
						Namespace: namespace,
					},
				},
				AzureKeyVault: akov2.AzureKeyVault{
					Enabled: pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-azure-credentials", name), dictionary),
						Namespace: namespace,
					},
				},
				GoogleCloudKms: akov2.GoogleCloudKms{
					Enabled: pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-gcp-credentials", name), dictionary),
						Namespace: namespace,
					},
				},
			},
		},
	}
}

func referenceAdvancedCluster(name, region, namespace, projectName string, labels map[string]string) *akov2.AtlasDeployment {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: akov2.AtlasDeploymentSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			BackupScheduleRef: akov2common.ResourceRefNamespaced{
				Namespace: targetNamespace,
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backupschedule", projectName, name), dictionary),
			},
			DeploymentSpec: &akov2.AdvancedDeploymentSpec{
				BackupEnabled: pointer.Get(true),
				BiConnector: &akov2.BiConnectorSpec{
					Enabled:        pointer.Get(false),
					ReadPreference: "secondary",
				},
				ClusterType:              string(akov2.TypeReplicaSet),
				DiskSizeGB:               nil,
				EncryptionAtRestProvider: "NONE",
				Name:                     name,
				Paused:                   pointer.Get(false),
				PitEnabled:               pointer.Get(true),
				ReplicationSpecs: []*akov2.AdvancedReplicationSpec{
					{
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*akov2.AdvancedRegionConfig{
							{
								AnalyticsSpecs: &akov2.Specs{
									DiskIOPS:      pointer.Get(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointer.Get(0),
								},
								ElectableSpecs: &akov2.Specs{
									DiskIOPS:      pointer.Get(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointer.Get(3),
								},
								ReadOnlySpecs: &akov2.Specs{
									DiskIOPS:      pointer.Get(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
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
								ProviderName: string(akov2provider.ProviderAWS),
								RegionName:   region,
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
		Status: akov2status.AtlasDeploymentStatus{
			Common: akov2status.Common{
				Conditions: []akov2status.Condition{},
			},
		},
	}
}

func referenceServerless(name, region, namespace, projectName string, labels map[string]string) *akov2.AtlasDeployment {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: akov2.AtlasDeploymentSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			ServerlessSpec: &akov2.ServerlessSpec{
				Name: name,
				ProviderSettings: &akov2.ProviderSettingsSpec{
					BackingProviderName: string(akov2provider.ProviderAWS),
					ProviderName:        akov2provider.ProviderServerless,
					RegionName:          region,
				},
			},
		},
		Status: akov2status.AtlasDeploymentStatus{
			Common: akov2status.Common{
				Conditions: []akov2status.Condition{},
			},
		},
	}
}

func referenceSharedCluster(name, region, namespace, projectName string, labels map[string]string) *akov2.AtlasDeployment {
	cluster := referenceAdvancedCluster(name, region, namespace, projectName, labels)
	cluster.Spec.DeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs = &akov2.Specs{
		DiskIOPS:     nil,
		InstanceSize: e2eSharedClusterTier,
	}
	cluster.Spec.DeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs = nil
	cluster.Spec.DeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs = nil
	cluster.Spec.DeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].AutoScaling = nil
	cluster.Spec.DeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].BackingProviderName = string(akov2provider.ProviderAWS)
	cluster.Spec.DeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ProviderName = string(akov2provider.ProviderTenant)

	cluster.Spec.DeploymentSpec.BackupEnabled = nil
	cluster.Spec.DeploymentSpec.BiConnector = nil
	cluster.Spec.DeploymentSpec.EncryptionAtRestProvider = ""
	cluster.Spec.DeploymentSpec.PitEnabled = nil
	cluster.Spec.BackupScheduleRef = akov2common.ResourceRefNamespaced{}
	return cluster
}

func defaultMaintenanceWindowAlertConfigs() []akov2.AlertConfiguration {
	ownerNotifications := func() []akov2.Notification {
		return []akov2.Notification{
			{
				EmailEnabled: pointer.Get(true),
				IntervalMin:  60,
				DelayMin:     pointer.Get(0),
				SMSEnabled:   pointer.Get(false),
				TypeName:     "GROUP",
				Roles:        []string{"GROUP_OWNER"},
			},
		}
	}

	return []akov2.AlertConfiguration{
		{
			Enabled:       true,
			EventTypeName: "MAINTENANCE_IN_ADVANCED",
			Threshold:     &akov2.Threshold{},
			Notifications: []akov2.Notification{
				{
					EmailEnabled: pointer.Get(true),
					IntervalMin:  60,
					DelayMin:     pointer.Get(0),
					SMSEnabled:   pointer.Get(false),
					TypeName:     "GROUP",
					Roles:        []string{"GROUP_OWNER"},
				},
			},
			MetricThreshold: &akov2.MetricThreshold{},
		},
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_STARTED",
			Threshold:       &akov2.Threshold{},
			Notifications:   ownerNotifications(),
			MetricThreshold: &akov2.MetricThreshold{},
		},
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_NO_LONGER_NEEDED",
			Threshold:       &akov2.Threshold{},
			Notifications:   ownerNotifications(),
			MetricThreshold: &akov2.MetricThreshold{},
		},
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_AUTO_DEFERRED",
			Threshold:       &akov2.Threshold{},
			Notifications:   ownerNotifications(),
			MetricThreshold: &akov2.MetricThreshold{},
		},
	}
}

func referenceBackupSchedule(namespace, projectName, clusterName string, labels map[string]string) *akov2.AtlasBackupSchedule {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasBackupSchedule{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasBackupSchedule",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: akov2.AtlasBackupScheduleSpec{
			PolicyRef: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName), dictionary),
				Namespace: namespace,
			},
			ReferenceHourOfDay:    1,
			ReferenceMinuteOfHour: 0,
			RestoreWindowDays:     7,
		},
	}
}

func referenceBackupPolicy(namespace, projectName, clusterName string, labels map[string]string) *akov2.AtlasBackupPolicy {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasBackupPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasBackupPolicy",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
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

func checkClustersData(t *testing.T, deployments []*akov2.AtlasDeployment, clusterNames []string, region, namespace, projectName string) {
	t.Helper()
	assert.Len(t, deployments, len(clusterNames))
	var entries []string
	for _, deployment := range deployments {
		if deployment.Spec.ServerlessSpec != nil {
			if ok := slices.Contains(clusterNames, deployment.Spec.ServerlessSpec.Name); ok {
				name := deployment.Spec.ServerlessSpec.Name
				expectedDeployment := referenceServerless(name, region, namespace, projectName, expectedLabels)
				assert.Equal(t, expectedDeployment, deployment)
				entries = append(entries, name)
			}
		} else if deployment.Spec.DeploymentSpec != nil {
			if ok := slices.Contains(clusterNames, deployment.Spec.DeploymentSpec.Name); ok {
				name := deployment.Spec.DeploymentSpec.Name
				expectedDeployment := referenceAdvancedCluster(name, region, namespace, projectName, expectedLabels)
				assert.Equal(t, expectedDeployment, deployment)
				entries = append(entries, name)
			}
		}
	}
	assert.Len(t, entries, len(clusterNames))
	assert.ElementsMatch(t, clusterNames, entries)
}

// TODO: add tests for project auditing and encryption at rest

func TestKubernetesConfigGenerate_ClustersWithBackup(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)
	g := newAtlasE2ETestGenerator(t)
	g.enableBackup = true
	g.generateProject(fmt.Sprintf("kubernetes-%s", n))
	g.generateCluster()
	g.generateServerlessCluster()

	expectedDeployment := referenceAdvancedCluster(g.clusterName, g.clusterRegion, targetNamespace, g.projectName, expectedLabels)
	expectedBackupSchedule := referenceBackupSchedule(targetNamespace, g.projectName, g.clusterName, expectedLabels)
	expectedBackupPolicy := referenceBackupPolicy(targetNamespace, g.projectName, g.clusterName, expectedLabels)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	// always register atlas entities
	require.NoError(t, akov2.AddToScheme(scheme.Scheme))

	t.Run("Update backup schedule", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"update",
			"--referenceHourOfDay",
			strconv.FormatInt(expectedBackupSchedule.Spec.ReferenceHourOfDay, 10),
			"--referenceMinuteOfHour",
			strconv.FormatInt(expectedBackupSchedule.Spec.ReferenceMinuteOfHour, 10),
			"--projectId",
			g.projectID,
			"--clusterName",
			g.clusterName)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))
	})

	t.Run("Generate valid resources of ONE project and ONE cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--clusterName",
			g.clusterName,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects, "result should not be empty")
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			require.True(t, found, "AtlasProject is not found in results")
			assert.Equal(t, targetNamespace, p.Namespace)
		})

		t.Run("Deployment present with valid data", func(t *testing.T) {
			found := false
			var deployment *akov2.AtlasDeployment
			var ok bool
			for i := range objects {
				deployment, ok = objects[i].(*akov2.AtlasDeployment)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasDeployment is not found in results")
			}
			assert.Equal(t, expectedDeployment, deployment)
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			secret, found := findSecret(objects)
			require.True(t, found, "Secret is not found in results")
			assert.Equal(t, targetNamespace, secret.Namespace)
		})

		t.Run("Backup Schedule present with valid data", func(t *testing.T) {
			schedule, found := atlasBackupSchedule(objects)
			require.True(t, found, "AtlasBackupSchedule is not found in results")
			assert.Equal(t, expectedBackupSchedule, schedule)
		})

		t.Run("Backup policy present with valid data", func(t *testing.T) {
			policy, found := atlasBackupPolicy(objects)
			require.True(t, found, "AtlasBackupPolicy is not found in results")
			assert.Equal(t, expectedBackupPolicy, policy)
		})
	})

	t.Run("Generate valid resources of ONE project and TWO clusters", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--clusterName",
			fmt.Sprintf("%s,%s", g.clusterName, g.serverlessName),
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, targetNamespace, p.Namespace)
		})

		t.Run("Deployments present with valid data", func(t *testing.T) {
			ds := atlasDeployments(objects)
			require.Len(t, ds, 2)
			checkClustersData(t, ds, []string{g.clusterName, g.serverlessName}, g.clusterRegion, targetNamespace, g.projectName)
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			secret, found := findSecret(objects)
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, targetNamespace, secret.Namespace)
		})
	})

	t.Run("Generate valid resources of ONE project and TWO clusters without listing clusters", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, targetNamespace, p.Namespace)
		})

		t.Run("Deployments present with valid data", func(t *testing.T) {
			ds := atlasDeployments(objects)
			checkClustersData(t, ds, []string{g.clusterName, g.serverlessName}, g.clusterRegion, targetNamespace, g.projectName)
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			secret, found := findSecret(objects)
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, targetNamespace, secret.Namespace)
		})
	})
}

func atlasBackupPolicy(objects []runtime.Object) (*akov2.AtlasBackupPolicy, bool) {
	for i := range objects {
		if policy, ok := objects[i].(*akov2.AtlasBackupPolicy); ok {
			return policy, ok
		}
	}
	return nil, false
}

func TestKubernetesConfigGenerateSharedCluster(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)
	g := newAtlasE2ETestGenerator(t)
	g.generateProject(fmt.Sprintf("kubernetes-%s", n))
	g.tier = e2eSharedClusterTier
	g.generateCluster()

	expectedDeployment := referenceSharedCluster(g.clusterName, g.clusterRegion, targetNamespace, g.projectName, expectedLabels)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	// always register atlas entities
	require.NoError(t, akov2.AddToScheme(scheme.Scheme))

	t.Run("Generate valid resources of ONE project and TWO clusters without listing clusters", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))
		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects)
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, targetNamespace, p.Namespace)
		})

		t.Run("Deployment present with valid data", func(t *testing.T) {
			ds := atlasDeployments(objects)
			assert.Len(t, ds, 1)
			assert.Equal(t, expectedDeployment, ds[0])
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			secret, found := findSecret(objects)
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, targetNamespace, secret.Namespace)
		})
	})
}

func atlasDeployments(objects []runtime.Object) []*akov2.AtlasDeployment {
	var ds []*akov2.AtlasDeployment
	for i := range objects {
		d, ok := objects[i].(*akov2.AtlasDeployment)
		if ok {
			ds = append(ds, d)
		}
	}
	return ds
}

func findAtlasProject(objects []runtime.Object) (*akov2.AtlasProject, bool) {
	for i := range objects {
		if p, ok := objects[i].(*akov2.AtlasProject); ok {
			return p, ok
		}
	}
	return nil, false
}

func findSecret(objects []runtime.Object) (*corev1.Secret, bool) {
	for i := range objects {
		if secret, ok := objects[i].(*corev1.Secret); ok {
			return secret, ok
		}
	}
	return nil, false
}

func atlasBackupSchedule(objects []runtime.Object) (*akov2.AtlasBackupSchedule, bool) {
	for i := range objects {
		if schedule, ok := objects[i].(*akov2.AtlasBackupSchedule); ok {
			return schedule, ok
		}
	}
	return nil, false
}

func referenceDataFederation(name, namespace, projectName string, labels map[string]string) *akov2.AtlasDataFederation {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasDataFederation{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasDataFederation",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: akov2.DataFederationSpec{
			Project: akov2common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			Name:                name,
			CloudProviderConfig: &akov2.CloudProviderConfig{},
			DataProcessRegion: &akov2.DataProcessRegion{
				CloudProvider: "AWS",
				Region:        "DUBLIN_IRL",
			},
			Storage: &akov2.Storage{
				Databases: nil,
				Stores:    nil,
			},
		},
		Status: akov2status.DataFederationStatus{
			Common: akov2status.Common{
				Conditions: []akov2status.Condition{},
			},
		},
	}
}

func TestKubernetesConfigGenerate_DataFederation(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)
	g := newAtlasE2ETestGenerator(t)
	g.generateProject(fmt.Sprintf("kubernetes-%s", n))
	g.generateDataFederation()
	var storeNames []string
	storeNames = append(storeNames, g.dataFedName)
	g.generateDataFederation()
	storeNames = append(storeNames, g.dataFedName)
	expectedDataFederation := referenceDataFederation(storeNames[0], targetNamespace, g.projectName, expectedLabels)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	// always register atlas entities
	require.NoError(t, akov2.AddToScheme(scheme.Scheme))

	t.Run("Generate valid resources of ONE project and ONE data federation", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--dataFederationName",
			storeNames[0],
			"--targetNamespace",
			targetNamespace)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))

		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects, "result should not be empty")
		})
		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, targetNamespace, p.Namespace)
		})
		t.Run("Deployment present with valid data", func(t *testing.T) {
			found := false
			var datafederation *akov2.AtlasDataFederation
			var ok bool
			for i := range objects {
				datafederation, ok = objects[i].(*akov2.AtlasDataFederation)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasDataFederation is not found in results")
			}
			assert.Equal(t, expectedDataFederation, datafederation)
		})
	})

	t.Run("Generate valid resources of ONE project and TWO data federation", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--dataFederationName",
			fmt.Sprintf("%s,%s", storeNames[0], storeNames[1]),
			"--targetNamespace",
			targetNamespace)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects, "result should not be empty")
		})
		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, targetNamespace, p.Namespace)
		})
		t.Run("Deployments present with valid data", func(t *testing.T) {
			dataFeds := atlasDataFederations(objects)
			require.Len(t, dataFeds, len(storeNames))
			checkDataFederationData(t, dataFeds, storeNames, targetNamespace, g.projectName)
		})
	})

	t.Run("Generate valid resources of ONE project and TWO data federation without listing data federation instances", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--targetNamespace",
			targetNamespace)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))
		require.NoError(t, err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.NotEmpty(t, objects, "result should not be empty")
		})
		t.Run("Project present with valid name", func(t *testing.T) {
			p, found := findAtlasProject(objects)
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, targetNamespace, p.Namespace)
		})
		t.Run("Deployments present with valid data", func(t *testing.T) {
			dataFeds := atlasDataFederations(objects)
			checkDataFederationData(t, dataFeds, storeNames, targetNamespace, g.projectName)
		})
	})
}

func atlasDataFederations(objects []runtime.Object) []*akov2.AtlasDataFederation {
	var df []*akov2.AtlasDataFederation
	for i := range objects {
		d, ok := objects[i].(*akov2.AtlasDataFederation)
		if ok {
			df = append(df, d)
		}
	}
	return df
}

func checkDataFederationData(t *testing.T, dataFederations []*akov2.AtlasDataFederation, dataFedNames []string, namespace, projectName string) {
	t.Helper()
	assert.Len(t, dataFederations, len(dataFedNames))
	var entries []string
	for _, instance := range dataFederations {
		if ok := slices.Contains(dataFedNames, instance.Spec.Name); ok {
			name := instance.Spec.Name
			expectedDeployment := referenceDataFederation(name, namespace, projectName, expectedLabels)
			assert.Equal(t, expectedDeployment, instance)
			entries = append(entries, name)
		}
	}
	assert.Len(t, entries, len(dataFedNames))
	assert.ElementsMatch(t, dataFedNames, entries)
}
