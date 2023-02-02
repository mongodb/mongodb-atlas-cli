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
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/pointers"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

const targetNamespace = "importer-namespace"

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
	s.generator = newAtlasE2ETestGenerator(s.t)
	s.generator.generateTeam("Kubernetes")
	s.generator.generateEmptyProject(fmt.Sprintf("Kubernetes-%s", s.generator.projectName))
	s.expectedProject = referenceProject(s.generator.projectName, targetNamespace)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(s.t, err)
	s.cliPath = cliPath

	s.assertions = assert.New(t)

	// always register atlas entities
	require.NoError(s.t, atlasV1.AddToScheme(scheme.Scheme))
	return s
}

func InitialSetup(t *testing.T) KubernetesConfigGenerateProjectSuite {
	t.Helper()
	s := KubernetesConfigGenerateProjectSuite{
		t: t,
	}
	s.generator = newAtlasE2ETestGenerator(s.t)
	s.generator.generateEmptyProject(fmt.Sprintf("Kubernetes-%s", s.generator.projectName))
	s.expectedProject = referenceProject(s.generator.projectName, targetNamespace)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(s.t, err)
	s.cliPath = cliPath

	s.assertions = assert.New(t)

	// always register atlas entities
	require.NoError(s.t, atlasV1.AddToScheme(scheme.Scheme))
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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})
}

func TestProjectWithNonDefaultSettings(t *testing.T) {
	s := InitialSetup(t)
	cliPath := s.cliPath
	generator := s.generator
	expectedProject := s.expectedProject
	assertions := s.assertions
	expectedProject.Spec.Settings.IsCollectDatabaseSpecificsStatisticsEnabled = pointers.MakePtr(false)

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
				DelayMin:     pointers.MakePtr(delayMin),
				SMSEnabled:   pointers.MakePtr(false),
				EmailEnabled: pointers.MakePtr(true),
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
						Database: pointers.MakePtr("test-db"),
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

	atlasCidrBlock := "10.8.0.0/18" //
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
		require.NoError(t, err)
		t.Cleanup(func() {
			deleteNetworkPeering(t, generator.projectID, gcpEntity)
		})
		var createdNetworkPeer mongodbatlas.Peer
		err = json.Unmarshal(resp, &createdNetworkPeer)
		require.NoError(t, err)
		expectedProject.Spec.NetworkPeers[0].ContainerID = createdNetworkPeer.ContainerID

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

	region := "northeurope"
	newPrivateEndpoint := atlasV1.PrivateEndpoint{
		Provider: provider.ProviderAzure,
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
			deletePrivateEndpoints(t, generator.projectID, azureEntity)
		})
		var createdNetworkPeer mongodbatlas.PrivateEndpointConnection
		err = json.Unmarshal(resp, &createdNetworkPeer)
		require.NoError(t, err)
		expectedProject.Spec.PrivateEndpoints[0].ID = createdNetworkPeer.ID

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
		}, generator.projectName)

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

func referenceTeam(name, namespace string, users []atlasV1.TeamUser, projectName string) *atlasV1.AtlasTeam {
	dictionary := resources.AtlasNameToKubernetesName()

	return &atlasV1.AtlasTeam{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasTeam",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-team-%s", projectName, name), dictionary),
			Namespace: namespace,
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
		var project *atlasV1.AtlasProject
		var ok bool
		for i := range output {
			project, ok = output[i].(*atlasV1.AtlasProject)
			if ok {
				found = true
				break
			}
		}
		if !found {
			t.Fatal("AtlasProject is not found in results")
		}
		asserts.Equal(expected, project)
	})
}

func referenceProject(name, namespace string) *atlasV1.AtlasProject {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasProject{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasProject",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(name, dictionary),
			Namespace: namespace,
		},
		Status: status.AtlasProjectStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
		Spec: atlasV1.AtlasProjectSpec{
			Name: name,
			ConnectionSecret: &common.ResourceRef{
				Name: resources.NormalizeAtlasName(fmt.Sprintf("%s-credentials", name), dictionary),
			},
			Settings: &atlasV1.ProjectSettings{
				IsCollectDatabaseSpecificsStatisticsEnabled: pointers.MakePtr(true),
				IsDataExplorerEnabled:                       pointers.MakePtr(true),
				IsPerformanceAdvisorEnabled:                 pointers.MakePtr(true),
				IsRealtimePerformancePanelEnabled:           pointers.MakePtr(true),
				IsSchemaAdvisorEnabled:                      pointers.MakePtr(true),
			},
			Auditing: &atlasV1.Auditing{
				AuditAuthorizationSuccess: pointers.MakePtr(false),
				Enabled:                   pointers.MakePtr(false),
			},
			EncryptionAtRest: &atlasV1.EncryptionAtRest{
				AwsKms: atlasV1.AwsKms{
					Enabled: pointers.MakePtr(false),
					Valid:   pointers.MakePtr(false),
				},
				AzureKeyVault: atlasV1.AzureKeyVault{
					Enabled: pointers.MakePtr(false),
				},
				GoogleCloudKms: atlasV1.GoogleCloudKms{
					Enabled: pointers.MakePtr(false),
				},
			},
		},
	}
}

func referenceAdvancedCluster(name, namespace, projectName string) *atlasV1.AtlasDeployment {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(name, dictionary),
			Namespace: namespace,
		},
		Spec: atlasV1.AtlasDeploymentSpec{
			Project: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(projectName, dictionary),
				Namespace: namespace,
			},
			BackupScheduleRef: common.ResourceRefNamespaced{
				Namespace: targetNamespace,
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-backupschedule", name), dictionary),
			},
			AdvancedDeploymentSpec: &atlasV1.AdvancedDeploymentSpec{
				BackupEnabled: pointers.MakePtr(true),
				BiConnector: &atlasV1.BiConnectorSpec{
					Enabled:        pointers.MakePtr(false),
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
				Paused:     pointers.MakePtr(false),
				PitEnabled: pointers.MakePtr(true),
				ReplicationSpecs: []*atlasV1.AdvancedReplicationSpec{
					{
						NumShards: 1,
						ZoneName:  "Zone 1",
						RegionConfigs: []*atlasV1.AdvancedRegionConfig{
							{
								AnalyticsSpecs: &atlasV1.Specs{
									DiskIOPS:      pointers.MakePtr(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointers.MakePtr(0),
								},
								ElectableSpecs: &atlasV1.Specs{
									DiskIOPS:      pointers.MakePtr(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointers.MakePtr(3),
								},
								ReadOnlySpecs: &atlasV1.Specs{
									DiskIOPS:      pointers.MakePtr(int64(3000)),
									EbsVolumeType: "STANDARD",
									InstanceSize:  e2eClusterTier,
									NodeCount:     pointers.MakePtr(0),
								},
								AutoScaling: &atlasV1.AdvancedAutoScalingSpec{
									DiskGB: &atlasV1.DiskGB{
										Enabled: pointers.MakePtr(false),
									},
									Compute: &atlasV1.ComputeSpec{
										Enabled:          pointers.MakePtr(false),
										ScaleDownEnabled: pointers.MakePtr(false),
									},
								},
								Priority:     pointers.MakePtr(7),
								ProviderName: string(provider.ProviderAWS),
								RegionName:   "US_EAST_1",
							},
						},
					},
				},
				RootCertType:         "ISRGROOTX1",
				VersionReleaseSystem: "LTS",
			},
			ProcessArgs: &atlasV1.ProcessArgs{
				MinimumEnabledTLSProtocol: "TLS1_2",
				JavascriptEnabled:         pointers.MakePtr(true),
				NoTableScan:               pointers.MakePtr(false),
			},
		},
		Status: status.AtlasDeploymentStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
}

func referenceAdvancedClusterWithSearch(name, namespace, projectName, collectionName, indexName string) *atlasV1.AtlasDeployment {
	deployment := referenceAdvancedCluster(name, namespace, projectName)
	deployment.Spec.AdvancedDeploymentSpec.AtlasSearch = &atlasV1.AtlasSearch{
		Databases: []atlasV1.AtlasSearchDatabase{
			{
				Database: "test",
				Collections: []atlasV1.AtlasSearchCollection{
					{
						CollectionName: collectionName,
						Indexes: []atlasV1.SearchIndex{
							{
								Name: indexName,
								Mappings: atlasV1.IndexMapping{
									Dynamic: true,
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

func referenceServerless(name, namespace, projectName string) *atlasV1.AtlasDeployment {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasDeployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasDeployment",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(name, dictionary),
			Namespace: namespace,
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
					RegionName:          "US_EAST_1",
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

func referenceSharedCluster(name, namespace, projectName string) *atlasV1.AtlasDeployment {
	cluster := referenceAdvancedCluster(name, namespace, projectName)
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ElectableSpecs = &atlasV1.Specs{
		InstanceSize: e2eSharedClusterTier,
	}
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ReadOnlySpecs = nil
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].AnalyticsSpecs = nil
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].AutoScaling = nil
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].BackingProviderName = string(provider.ProviderAWS)
	cluster.Spec.AdvancedDeploymentSpec.ReplicationSpecs[0].RegionConfigs[0].ProviderName = string(provider.ProviderTenant)

	cluster.Spec.AdvancedDeploymentSpec.PitEnabled = pointers.MakePtr(false)
	cluster.Spec.BackupScheduleRef = common.ResourceRefNamespaced{}
	return cluster
}

func defaultMaintenanceWindowAlertConfigs() []atlasV1.AlertConfiguration {
	ownerNotifications := []atlasV1.Notification{
		{
			EmailEnabled: pointers.MakePtr(true),
			IntervalMin:  60,
			DelayMin:     pointers.MakePtr(0),
			SMSEnabled:   pointers.MakePtr(false),
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

func referenceBackupSchedule(namespace, clusterName string) *atlasV1.AtlasBackupSchedule {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasBackupSchedule{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasBackupSchedule",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-backupschedule", clusterName), dictionary),
			Namespace: namespace,
		},
		Spec: atlasV1.AtlasBackupScheduleSpec{
			PolicyRef: common.ResourceRefNamespaced{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-backuppolicy", clusterName), dictionary),
				Namespace: namespace,
			},
			ReferenceHourOfDay:    1,
			ReferenceMinuteOfHour: 0,
			RestoreWindowDays:     7,
		},
	}
}

func referenceBackupPolicy(namespace, clusterName string) *atlasV1.AtlasBackupPolicy {
	dictionary := resources.AtlasNameToKubernetesName()
	return &atlasV1.AtlasBackupPolicy{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasBackupPolicy",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-backuppolicy", clusterName), dictionary),
			Namespace: namespace,
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

func checkClustersData(t *testing.T, deployments []*atlasV1.AtlasDeployment, clusterNames []string, namespace, projectName string) {
	t.Helper()
	assert.Len(t, deployments, len(clusterNames))
	var entries []string
	for _, deployment := range deployments {
		if deployment.Spec.ServerlessSpec != nil {
			if ok := search.StringInSlice(clusterNames, deployment.Spec.ServerlessSpec.Name); ok {
				name := deployment.Spec.ServerlessSpec.Name
				expectedDeployment := referenceServerless(name, namespace, projectName)
				assert.Equal(t, expectedDeployment, deployment)
				entries = append(entries, name)
			}
		} else if deployment.Spec.AdvancedDeploymentSpec != nil {
			if ok := search.StringInSlice(clusterNames, deployment.Spec.AdvancedDeploymentSpec.Name); ok {
				name := deployment.Spec.AdvancedDeploymentSpec.Name
				expectedDeployment := referenceAdvancedCluster(name, namespace, projectName)
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

	expectedDeployment := referenceAdvancedCluster(g.clusterName, targetNamespace, g.projectName)
	expectedBackupSchedule := referenceBackupSchedule(targetNamespace, g.clusterName)
	expectedBackupPolicy := referenceBackupPolicy(targetNamespace, g.clusterName)

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
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})

		t.Run("Backup Schedule present with valid data", func(t *testing.T) {
			found := false
			var schedule *atlasV1.AtlasBackupSchedule
			var ok bool
			for i := range objects {
				schedule, ok = objects[i].(*atlasV1.AtlasBackupSchedule)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasBackupSchedule is not found in results")
			}
			assert.Equal(t, expectedBackupSchedule, schedule)
		})

		t.Run("Backup policy present with valid data", func(t *testing.T) {
			found := false
			var policy *atlasV1.AtlasBackupPolicy
			var ok bool
			for i := range objects {
				policy, ok = objects[i].(*atlasV1.AtlasBackupPolicy)
				if ok {
					found = true
					break
				}
			}
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
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
		})

		t.Run("Deployments present with valid data", func(t *testing.T) {
			var deployments []*atlasV1.AtlasDeployment
			for i := range objects {
				deployment, ok := objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					deployments = append(deployments, deployment)
				}
			}
			clustersCount := len(deployments)
			require.True(t, clustersCount == 2, "result should contain two clusters. actual: ", clustersCount)
			checkClustersData(t, deployments, []string{g.clusterName, g.serverlessName}, targetNamespace, g.projectName)
		})

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
			assert.Equal(t, secret.Namespace, targetNamespace)
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

		a := assert.New(t)
		a.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
		})

		t.Run("Deployments present with valid data", func(t *testing.T) {
			var deployments []*atlasV1.AtlasDeployment
			for i := range objects {
				deployment, ok := objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					deployments = append(deployments, deployment)
				}
			}
			checkClustersData(t, deployments, []string{g.clusterName, g.serverlessName}, targetNamespace, g.projectName)
		})

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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})
}

func TestKubernetesConfigGenerateSharedCluster(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)
	g := newAtlasE2ETestGenerator(t)
	g.generateProject(fmt.Sprintf("kubernetes-%s", n))
	g.tier = e2eSharedClusterTier
	g.generateCluster()

	expectedDeployment := referenceSharedCluster(g.clusterName, targetNamespace, g.projectName)

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

		a := assert.New(t)
		a.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
		})

		t.Run("Deployment present with valid data", func(t *testing.T) {
			var deployments []*atlasV1.AtlasDeployment
			for i := range objects {
				deployment, ok := objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					deployments = append(deployments, deployment)
				}
			}
			a.Len(deployments, 1)
			a.Equal(expectedDeployment, deployments[0])
		})

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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})
}

func TestKubernetesConfigGenerate_ClustersWithAtlasSearch(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)
	g := newAtlasE2ETestGenerator(t)
	g.generateProjectAndCluster(fmt.Sprintf("k8s-atlas-search-%s", n))
	indexName := fmt.Sprintf("index-%v", n)
	collectionName := fmt.Sprintf("collection-%v", n)

	expectedDeployment := referenceAdvancedClusterWithSearch(g.clusterName, targetNamespace, g.projectName, collectionName, indexName)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	// always register atlas entities
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))

	t.Run("Create index", func(t *testing.T) {
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		}()

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "{{ .collectionName }}",
	"database": "test",
	"name": "{{ .indexName }}",
	"mappings": {
		"dynamic": true
	}
}`))
		err = tpl.Execute(file, map[string]string{
			"collectionName": collectionName,
			"indexName":      indexName,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			"--clusterName", g.clusterName,
			"--file",
			fileName,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
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
			"--includeSecrets",
			"--indexFrom",
			fmt.Sprintf("%s.test.%s", g.clusterName, collectionName))
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
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})

	t.Run("Generate valid resources of ONE project and ONE cluster without listing the cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			g.projectID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets",
			"--indexFrom",
			fmt.Sprintf("%s.test.%s", g.clusterName, collectionName))
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
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
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
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})
}
