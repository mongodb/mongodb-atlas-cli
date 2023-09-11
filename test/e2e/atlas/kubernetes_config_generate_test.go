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
	"strconv"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201006/admin"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

const targetNamespace = "importer-namespace"

var expectedLabels = map[string]string{
	features.ResourceVersion: features.LatestOperatorMajorVersion,
}

func getK8SEntities(data []byte) ([]runtime.Object, error) {
	b := bufio.NewReader(bytes.NewReader(data))
	r := k8syaml.NewYAMLReader(b)

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
	expectedProject *atlasV1.AtlasProject
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
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))
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
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))
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

		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		checkProject(t, objects, expectedProject, assertions)
	})
}

func TestProjectWithNonDefaultAlertConf(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions

	newAlertConfig := atlasV1.AlertConfiguration{
		Threshold:       &atlasV1.Threshold{},
		MetricThreshold: &atlasV1.MetricThreshold{},
		EventTypeName:   eventTypeName,
		Enabled:         true,
		Notifications: []atlasV1.Notification{
			{
				TypeName:     group,
				IntervalMin:  intervalMin,
				DelayMin:     pointer.Get(delayMin),
				SMSEnabled:   pointer.Get(false),
				EmailEnabled: pointer.Get(true),
			},
		},
	}
	expectedProject.Spec.AlertConfigurations = []atlasV1.AlertConfiguration{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
	newIPAccess := project.IPAccessList{
		IPAddress: entry,
		Comment:   "test",
	}
	expectedProject.Spec.ProjectIPAccessList = []project.IPAccessList{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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

	newIPAccess := atlasV1.CloudProviderAccessRole{
		ProviderName: string(provider.ProviderAWS),
	}
	expectedProject.Spec.CloudProviderAccessRoles = []atlasV1.CloudProviderAccessRole{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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

	newCustomRole := atlasV1.CustomRole{
		Name: "test-role",
		Actions: []atlasV1.Action{
			{
				Name: "FIND",
				Resources: []atlasV1.Resource{
					{
						Database:   pointer.Get("test-db	"),
						Collection: pointer.Get(""),
						Cluster:    pointer.Get(false),
					},
				},
			},
		},
	}
	expectedProject.Spec.CustomRoles = []atlasV1.CustomRole{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
	newIntegration := project.Integration{
		Type:   datadogEntity,
		Region: "US", // it's a default value
		APIKeyRef: common.ResourceRefNamespaced{
			Namespace: targetNamespace,
			Name:      fmt.Sprintf("%s-integration-%s", generator.projectID, strings.ToLower(datadogEntity)),
		},
	}
	expectedProject.Spec.Integrations = []project.Integration{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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

	newMaintenanceWindow := project.MaintenanceWindow{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
	networkPeer := atlasV1.NetworkPeer{
		ProviderName: provider.ProviderGCP,
		NetworkName:  "test-network",
		GCPProjectID: "test-project-gcp",
	}
	expectedProject.Spec.NetworkPeers = []atlasV1.NetworkPeer{
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
	newPrivateEndpoint := atlasV1.PrivateEndpoint{
		Provider: provider.ProviderAzure,
		Region:   "EUROPE_NORTH",
	}
	expectedProject.Spec.PrivateEndpoints = []atlasV1.PrivateEndpoint{
		newPrivateEndpoint,
	}

	t.Run("Add network peer to the project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"create",
			"--region="+region,
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
		expectedProject.Spec.PrivateEndpoints[0].ID = createdNetworkPeer.GetId()

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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
		expectedTeam := referenceTeam(generator.teamName, targetNamespace, []atlasV1.TeamUser{
			atlasV1.TeamUser(generator.teamUser),
		}, generator.projectName, expectedLabels)

		expectedProject.Spec.Teams = []atlasV1.Team{
			{
				TeamRef: common.ResourceRefNamespaced{
					Namespace: targetNamespace,
					Name:      expectedTeam.Name,
				},
				Roles: []atlasV1.TeamRole{
					atlasV1.TeamRole(teamRole),
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
		assertions.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		checkProject(t, objects, expectedProject, assertions)
		t.Run("Team is created", func(t *testing.T) {
			for _, obj := range objects {
				if team, ok := obj.(*atlasV1.AtlasTeam); ok {
					assertions.Equal(expectedTeam, team)
				}
			}
		})
	})
}

func referenceTeam(name, namespace string, users []atlasV1.TeamUser, projectName string, labels map[string]string) *atlasV1.AtlasTeam {
	dictionary := resources.AtlasNameToKubernetesName()

	return &atlasV1.AtlasTeam{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasTeam",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-team-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: atlasV1.TeamSpec{
			Name:      name,
			Usernames: users,
		},
		Status: status.TeamStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
}

func checkProject(t *testing.T, output []runtime.Object, expected *atlasV1.AtlasProject, asserts *assert.Assertions) {
	t.Helper()
	t.Run("Project presents with expected data", func(t *testing.T) {
		found := false
		var p *atlasV1.AtlasProject
		var ok bool
		for i := range output {
			p, ok = output[i].(*atlasV1.AtlasProject)
			if ok {
				found = true
				break
			}
		}
		if !found {
			t.Fatal("AtlasProject is not found in results")
		}
		asserts.Equal(expected, p)
	})
}

func referenceProject(name, namespace string, labels map[string]string) *atlasV1.AtlasProject {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasProject{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasProject",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(name, dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Status: status.AtlasProjectStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
		Spec: atlasV1.AtlasProjectSpec{
			Name: name,
			ConnectionSecret: &common.ResourceRefNamespaced{
				Name: resources.NormalizeAtlasName(fmt.Sprintf("%s-credentials", name), dictionary),
			},
			Settings: &atlasV1.ProjectSettings{
				IsCollectDatabaseSpecificsStatisticsEnabled: pointer.Get(true),
				IsDataExplorerEnabled:                       pointer.Get(true),
				IsPerformanceAdvisorEnabled:                 pointer.Get(true),
				IsRealtimePerformancePanelEnabled:           pointer.Get(true),
				IsSchemaAdvisorEnabled:                      pointer.Get(true),
			},
			Auditing: &atlasV1.Auditing{
				AuditAuthorizationSuccess: pointer.Get(false),
				Enabled:                   pointer.Get(false),
			},
			EncryptionAtRest: &atlasV1.EncryptionAtRest{
				AwsKms: atlasV1.AwsKms{
					Enabled: pointer.Get(false),
					Valid:   pointer.Get(false),
				},
				AzureKeyVault: atlasV1.AzureKeyVault{
					Enabled: pointer.Get(false),
				},
				GoogleCloudKms: atlasV1.GoogleCloudKms{
					Enabled: pointer.Get(false),
				},
			},
		},
	}
}

func referenceAdvancedCluster(name, region, namespace, projectName string, labels map[string]string) *atlasV1.AtlasDeployment {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			BackupScheduleRef: common.ResourceRefNamespaced{
				Namespace: targetNamespace,
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backupschedule", projectName, name), dictionary),
			},
			AdvancedDeploymentSpec: &atlasV1.AdvancedDeploymentSpec{
				BackupEnabled: pointer.Get(true),
				BiConnector: &atlasV1.BiConnectorSpec{
					Enabled:        pointer.Get(false),
					ReadPreference: "secondary",
				},
				ClusterType:              string(atlasV1.TypeReplicaSet),
				DiskSizeGB:               nil,
				EncryptionAtRestProvider: "NONE",
				Labels: []common.LabelSpec{
					{
						Key:   "Infrastructure Tool",
						Value: "Atlas CLI",
					},
				},
				Name:       name,
				Paused:     pointer.Get(false),
				PitEnabled: pointer.Get(true),
				ReplicationSpecs: []*atlasV1.AdvancedReplicationSpec{
					{
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*atlasV1.AdvancedRegionConfig{
							{
								AnalyticsSpecs: &atlasV1.Specs{
									DiskIOPS:      pointer.Get(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointer.Get(0),
								},
								ElectableSpecs: &atlasV1.Specs{
									DiskIOPS:      pointer.Get(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointer.Get(3),
								},
								ReadOnlySpecs: &atlasV1.Specs{
									DiskIOPS:      pointer.Get(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointer.Get(0),
								},
								AutoScaling: &atlasV1.AdvancedAutoScalingSpec{
									DiskGB: &atlasV1.DiskGB{
										Enabled: pointer.Get(false),
									},
									Compute: &atlasV1.ComputeSpec{
										Enabled:          pointer.Get(false),
										ScaleDownEnabled: pointer.Get(false),
									},
								},
								Priority:     pointer.Get(7),
								ProviderName: string(provider.ProviderAWS),
								RegionName:   region,
							},
						},
					},
				},
				RootCertType:         "ISRGROOTX1",
				VersionReleaseSystem: "LTS",
			},
			ProcessArgs: &atlasV1.ProcessArgs{
				MinimumEnabledTLSProtocol: "TLS1_2",
				JavascriptEnabled:         pointer.Get(true),
				NoTableScan:               pointer.Get(false),
			},
		},
		Status: status.AtlasDeploymentStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
}

func referenceServerless(name, region, namespace, projectName string, labels map[string]string) *atlasV1.AtlasDeployment {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			ServerlessSpec: &atlasV1.ServerlessSpec{
				Name: name,
				ProviderSettings: &atlasV1.ProviderSettingsSpec{
					BackingProviderName: string(provider.ProviderAWS),
					ProviderName:        provider.ProviderServerless,
					RegionName:          region,
				},
			},
		},
		Status: status.AtlasDeploymentStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
}

func referenceSharedCluster(name, region, namespace, projectName string, labels map[string]string) *atlasV1.AtlasDeployment {
	cluster := referenceAdvancedCluster(name, region, namespace, projectName, labels)
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs = &atlasV1.Specs{
		DiskIOPS:     pointer.Get(int64(0)),
		InstanceSize: e2eSharedClusterTier,
	}
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs = nil
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs = nil
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].AutoScaling = nil
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].BackingProviderName = string(provider.ProviderAWS)
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ProviderName = string(provider.ProviderTenant)

	cluster.Spec.AdvancedDeploymentSpec.PitEnabled = pointer.Get(false)
	cluster.Spec.BackupScheduleRef = common.ResourceRefNamespaced{}
	return cluster
}

func defaultMaintenanceWindowAlertConfigs() []atlasV1.AlertConfiguration {
	ownerNotifications := []atlasV1.Notification{
		{
			EmailEnabled: pointer.Get(true),
			IntervalMin:  60,
			DelayMin:     pointer.Get(0),
			SMSEnabled:   pointer.Get(false),
			TypeName:     "GROUP",
			Roles:        []string{"GROUP_OWNER"},
		},
	}
	return []atlasV1.AlertConfiguration{
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_IN_ADVANCED",
			Threshold:       &atlasV1.Threshold{},
			Notifications:   ownerNotifications,
			MetricThreshold: &atlasV1.MetricThreshold{},
		},
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_STARTED",
			Threshold:       &atlasV1.Threshold{},
			Notifications:   ownerNotifications,
			MetricThreshold: &atlasV1.MetricThreshold{},
		},
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_NO_LONGER_NEEDED",
			Threshold:       &atlasV1.Threshold{},
			Notifications:   ownerNotifications,
			MetricThreshold: &atlasV1.MetricThreshold{},
		},
		{
			Enabled:         true,
			EventTypeName:   "MAINTENANCE_AUTO_DEFERRED",
			Threshold:       &atlasV1.Threshold{},
			Notifications:   ownerNotifications,
			MetricThreshold: &atlasV1.MetricThreshold{},
		},
	}
}

func referenceBackupSchedule(namespace, projectName, clusterName string, labels map[string]string) *atlasV1.AtlasBackupSchedule {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasBackupSchedule{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasBackupSchedule",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backupschedule", projectName, clusterName), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: atlasV1.AtlasBackupScheduleSpec{
			PolicyRef: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName), dictionary),
				Namespace: namespace,
			},
			ReferenceHourOfDay:    1,
			ReferenceMinuteOfHour: 0,
			RestoreWindowDays:     7,
		},
	}
}

func referenceBackupPolicy(namespace, projectName, clusterName string, labels map[string]string) *atlasV1.AtlasBackupPolicy {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasBackupPolicy{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasBackupPolicy",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-backuppolicy", projectName, clusterName), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: atlasV1.AtlasBackupPolicySpec{
			Items: []atlasV1.AtlasBackupPolicyItem{
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

func checkClustersData(t *testing.T, deployments []*atlasV1.AtlasDeployment, clusterNames []string, region, namespace, projectName string) {
	t.Helper()
	assert.Len(t, deployments, len(clusterNames))
	var entries []string
	for _, deployment := range deployments {
		if deployment.Spec.ServerlessSpec != nil {
			if ok := search.StringInSlice(clusterNames, deployment.Spec.ServerlessSpec.Name); ok {
				name := deployment.Spec.ServerlessSpec.Name
				expectedDeployment := referenceServerless(name, region, namespace, projectName, expectedLabels)
				assert.Equal(t, expectedDeployment, deployment)
				entries = append(entries, name)
			}
		} else if deployment.Spec.AdvancedDeploymentSpec != nil {
			if ok := search.StringInSlice(clusterNames, deployment.Spec.AdvancedDeploymentSpec.Name); ok {
				name := deployment.Spec.AdvancedDeploymentSpec.Name
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
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))

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

		a := assert.New(t)
		a.NoError(err, string(resp))
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

		a := assert.New(t)
		a.NoError(err, string(resp))

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
			assert.Equal(t, p.Namespace, targetNamespace)
		})

		t.Run("Deployment present with valid data", func(t *testing.T) {
			found := false
			var deployment *atlasV1.AtlasDeployment
			var ok bool
			for i := range objects {
				deployment, ok = objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasDeployment is not found in results")
			}
			a.Equal(expectedDeployment, deployment)
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			secret, found := findSecret(objects)
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, secret.Namespace, targetNamespace)
		})

		t.Run("Backup Schedule present with valid data", func(t *testing.T) {
			schedule, found := atlasBackupSchedule(objects)
			if !found {
				t.Fatal("AtlasBackupSchedule is not found in results")
			}
			assert.Equal(t, expectedBackupSchedule, schedule)
		})

		t.Run("Backup policy present with valid data", func(t *testing.T) {
			policy, found := atlasBackupPolicy(objects)
			if !found {
				t.Fatal("AtlasBackupSchedule is not found in results")
			}
			a.Equal(expectedBackupPolicy, policy)
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

		a := assert.New(t)
		a.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})
}

func atlasBackupPolicy(objects []runtime.Object) (*atlasV1.AtlasBackupPolicy, bool) {
	for i := range objects {
		if policy, ok := objects[i].(*atlasV1.AtlasBackupPolicy); ok {
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
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))

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
			assert.Equal(t, p.Namespace, targetNamespace)
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

func atlasDeployments(objects []runtime.Object) []*atlasV1.AtlasDeployment {
	var ds []*atlasV1.AtlasDeployment
	for i := range objects {
		d, ok := objects[i].(*atlasV1.AtlasDeployment)
		if ok {
			ds = append(ds, d)
		}
	}
	return ds
}

func findAtlasProject(objects []runtime.Object) (*atlasV1.AtlasProject, bool) {
	for i := range objects {
		if p, ok := objects[i].(*atlasV1.AtlasProject); ok {
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

func atlasBackupSchedule(objects []runtime.Object) (*atlasV1.AtlasBackupSchedule, bool) {
	for i := range objects {
		if schedule, ok := objects[i].(*atlasV1.AtlasBackupSchedule); ok {
			return schedule, ok
		}
	}
	return nil, false
}

func referenceDataFederation(name, namespace, projectName string, labels map[string]string) *atlasV1.AtlasDataFederation {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasDataFederation{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDataFederation",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, name), dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: atlasV1.DataFederationSpec{
			Project: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			Name:                name,
			CloudProviderConfig: &atlasV1.CloudProviderConfig{},
			DataProcessRegion: &atlasV1.DataProcessRegion{
				CloudProvider: "AWS",
				Region:        "DUBLIN_IRL",
			},
			Storage: &atlasV1.Storage{
				Databases: nil,
				Stores:    nil,
			},
		},
		Status: status.DataFederationStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
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
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))

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

		a := assert.New(t)
		a.NoError(err, string(resp))

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
			var datafederation *atlasV1.AtlasDataFederation
			var ok bool
			for i := range objects {
				datafederation, ok = objects[i].(*atlasV1.AtlasDataFederation)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasDataFederation is not found in results")
			}
			a.Equal(expectedDataFederation, datafederation)
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

		a := assert.New(t)
		a.NoError(err, string(resp))

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

		a := assert.New(t)
		a.NoError(err, string(resp))

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

func atlasDataFederations(objects []runtime.Object) []*atlasV1.AtlasDataFederation {
	var df []*atlasV1.AtlasDataFederation
	for i := range objects {
		d, ok := objects[i].(*atlasV1.AtlasDataFederation)
		if ok {
			df = append(df, d)
		}
	}
	return df
}

func checkDataFederationData(t *testing.T, dataFederations []*atlasV1.AtlasDataFederation, dataFedNames []string, namespace, projectName string) {
	t.Helper()
	assert.Len(t, dataFederations, len(dataFedNames))
	var entries []string
	for _, instance := range dataFederations {
		if ok := search.StringInSlice(dataFedNames, instance.Spec.Name); ok {
			name := instance.Spec.Name
			expectedDeployment := referenceDataFederation(name, namespace, projectName, expectedLabels)
			assert.Equal(t, expectedDeployment, instance)
			entries = append(entries, name)
		}
	}
	assert.Len(t, entries, len(dataFedNames))
	assert.ElementsMatch(t, dataFedNames, entries)
}
