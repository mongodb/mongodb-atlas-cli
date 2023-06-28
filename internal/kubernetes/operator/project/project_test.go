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

package project

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const orgID = "TestOrgID"
const projectID = "TestProjectID"
const teamID = "TestTeamID"
const resourceVersion = "x.y.z"

func TestBuildAtlasProject(t *testing.T) {
	ctl := gomock.NewController(t)
	projectStore := mocks.NewMockAtlasOperatorProjectStore(ctl)
	featureValidator := mocks.NewMockFeatureValidator(ctl)
	t.Run("Can convert Project entity with secrets data", func(t *testing.T) {
		targetNamespace := "test-namespace"

		p := &mongodbatlas.Project{
			ID:                        projectID,
			OrgID:                     orgID,
			Name:                      "TestProjectName",
			ClusterCount:              0,
			Created:                   "",
			RegionUsageRestrictions:   "",
			Links:                     nil,
			WithDefaultAlertsSettings: pointer.Get(false),
		}

		ipAccessLists := &atlasv2.PaginatedNetworkAccess{
			Links: nil,
			Results: []atlasv2.NetworkPermissionEntry{
				{
					AwsSecurityGroup: pointer.Get("TestSecurity group"),
					CidrBlock:        pointer.Get("0.0.0.0/0"),
					Comment:          pointer.Get("Allow everyone"),
					DeleteAfterDate:  pointer.Get(time.Now()),
					GroupId:          pointer.Get("TestGroupID"),
					IpAddress:        pointer.Get("0.0.0.0"),
				},
			},
			TotalCount: pointer.Get(1),
		}

		auditing := &atlasv2.AuditLog{
			AuditAuthorizationSuccess: true,
			AuditFilter:               "TestFilter",
			ConfigurationType:         pointer.Get("TestConfigType"),
			Enabled:                   true,
		}

		authDate, _ := time.Parse(time.RFC3339, "01-01-2001")
		createDate, _ := time.Parse(time.RFC3339, "01-02-2001")

		cpas := &atlasv2.CloudProviderAccessRoles{
			AwsIamRoles: []atlasv2.CloudProviderAccessAWSIAMRole{
				{
					AtlasAWSAccountArn:         pointer.Get("TestARN"),
					AtlasAssumedRoleExternalId: pointer.Get("TestExternalRoleID"),
					AuthorizedDate:             &authDate,
					CreatedDate:                &createDate,
					FeatureUsages:              nil,
					IamAssumedRoleArn:          pointer.Get("TestRoleARN"),
					ProviderName:               string(provider.ProviderAWS),
					RoleId:                     pointer.Get("TestRoleID"),
				},
			},
		}

		encryptionAtRest := &atlasv2.EncryptionAtRest{
			AwsKms:        &atlasv2.AWSKMSConfiguration{},
			AzureKeyVault: &atlasv2.AzureKeyVault{},
			GoogleCloudKms: &atlasv2.GoogleCloudKMS{
				Enabled:              pointer.Get(true),
				ServiceAccountKey:    pointer.Get("TestServiceAccountKey"),
				KeyVersionResourceID: pointer.Get("TestKeyVersionResourceID"),
			},
		}

		thirdPartyIntegrations := &atlasv2.PaginatedIntegration{
			Links: nil,
			Results: []atlasv2.ThridPartyIntegration{
				{
					Prometheus: &atlasv2.Prometheus{
						Type:             pointer.Get("PROMETHEUS"),
						Username:         "TestPrometheusUserName",
						Password:         pointer.Get("TestPrometheusPassword"),
						ServiceDiscovery: "TestPrometheusServiceDiscovery",
					},
				},
			},
			TotalCount: pointer.Get(1),
		}

		mw := &atlasv2.GroupMaintenanceWindow{
			DayOfWeek:            1,
			HourOfDay:            10,
			StartASAP:            pointer.Get(false),
			AutoDeferOnceEnabled: pointer.Get(false),
		}

		peeringConnectionAWS := &atlasv2.AwsNetworkPeeringConnectionSettings{
			AccepterRegionName:  "TestRegionName",
			AwsAccountId:        "TestAWSAccountID",
			ConnectionId:        pointer.Get("TestConnID"),
			ContainerId:         "TestContainerID",
			ErrorStateName:      pointer.Get("TestErrStateName"),
			Id:                  pointer.Get("TestID"),
			ProviderName:        pointer.Get(string(provider.ProviderAWS)),
			RouteTableCidrBlock: "0.0.0.0/0",
			StatusName:          pointer.Get("TestStatusName"),
			VpcId:               "TestVPCID",
		}

		peeringConnections := []interface{}{
			peeringConnectionAWS,
		}

		privateAWSEndpoint := atlasv2.EndpointService{
			Id:                  pointer.Get("TestID"),
			CloudProvider:       string(provider.ProviderAWS),
			RegionName:          pointer.Get("US_WEST_2"),
			EndpointServiceName: nil,
			ErrorMessage:        nil,
			InterfaceEndpoints:  nil,
			Status:              nil,
		}
		privateEndpoints := []atlasv2.EndpointService{privateAWSEndpoint}

		alertConfigs := []mongodbatlas.AlertConfiguration{
			{
				EventTypeName: "TestEventTypeName",
				Enabled:       pointer.Get(true),
				Matchers: []mongodbatlas.Matcher{
					{
						FieldName: "TestFieldName",
						Operator:  "TestOperator",
						Value:     "TestValue",
					},
				},
				MetricThreshold: &mongodbatlas.MetricThreshold{
					MetricName: "TestMetricName",
					Operator:   "TestOperator",
					Threshold:  10,
					Units:      "TestUnits",
					Mode:       "TestMode",
				},
				Threshold: &mongodbatlas.Threshold{
					Operator:  "TestOperator",
					Units:     "TestUnits",
					Threshold: 10,
				},
				Notifications: []mongodbatlas.Notification{
					{
						APIToken:            "TestAPIToken",
						ChannelName:         "TestChannelName",
						DatadogAPIKey:       "TestDatadogAPIKey",
						DatadogRegion:       "TestDatadogRegion",
						DelayMin:            pointer.Get(5),
						EmailAddress:        "TestEmail@mongodb.com",
						EmailEnabled:        pointer.Get(true),
						FlowdockAPIToken:    "TestFlowDockApiToken",
						FlowName:            "TestFlowName",
						IntervalMin:         0,
						MobileNumber:        "+12345678900",
						OpsGenieAPIKey:      "TestGenieAPIKey",
						OpsGenieRegion:      "TestGenieRegion",
						OrgName:             "TestOrgName",
						ServiceKey:          "TestServiceKey",
						SMSEnabled:          pointer.Get(true),
						TeamID:              "TestTeamID",
						TeamName:            "TestTeamName",
						TypeName:            "TestTypeName",
						Username:            "TestUserName",
						VictorOpsAPIKey:     "TestVictorOpsAPIKey",
						VictorOpsRoutingKey: "TestVictorOpsRoutingKey",
						Roles:               []string{"Role1", "Role2"},
					},
				},
			},
		}

		projectSettings := &atlasv2.GroupSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: pointer.Get(true),
			IsDataExplorerEnabled:                       pointer.Get(true),
			IsPerformanceAdvisorEnabled:                 pointer.Get(true),
			IsRealtimePerformancePanelEnabled:           pointer.Get(true),
			IsSchemaAdvisorEnabled:                      pointer.Get(true),
		}

		customRoles := []atlasv2.UserCustomDBRole{
			{
				Actions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "Action-1",
						Resources: []atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: "Collection-1",
								Db:         "DB-1",
								Cluster:    true,
							},
						},
					},
				},
				InheritedRoles: []atlasv2.DatabaseInheritedRole{
					{
						Db:   "Inherited-DB",
						Role: "Inherited-ROLE",
					},
				},
				RoleName: "TestCustomRoleName",
			},
		}

		projectTeams := &mongodbatlas.TeamsAssigned{
			Links: nil,
			Results: []*mongodbatlas.Result{
				{
					Links:     nil,
					TeamID:    teamID,
					RoleNames: []string{string(atlasV1.TeamRoleClusterManager)},
				},
			},
			TotalCount: 1,
		}
		teams := &mongodbatlas.Team{
			ID:        teamID,
			Name:      "TestTeamName",
			Usernames: []string{},
		}

		teamUsers := []mongodbatlas.AtlasUser{
			{
				EmailAddress: "testuser@mooooongodb.com",
				FirstName:    "TestName",
				ID:           "TestID",
				LastName:     "TestLastName",
			},
		}

		listOption := &mongodbatlas.ListOptions{ItemsPerPage: MaxItems}
		containerListOptionAWS := &mongodbatlas.ContainersListOptions{ListOptions: *listOption, ProviderName: string(provider.ProviderAWS)}
		containerListOptionGCP := &mongodbatlas.ContainersListOptions{ListOptions: *listOption, ProviderName: string(provider.ProviderGCP)}
		containerListOptionAzure := &mongodbatlas.ContainersListOptions{ListOptions: *listOption, ProviderName: string(provider.ProviderAzure)}
		projectStore.EXPECT().Project(projectID).Return(p, nil)
		projectStore.EXPECT().ProjectIPAccessLists(projectID, listOption).Return(ipAccessLists, nil)
		projectStore.EXPECT().MaintenanceWindow(projectID).Return(mw, nil)
		projectStore.EXPECT().Integrations(projectID).Return(thirdPartyIntegrations, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOptionAWS).Return(peeringConnections, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOptionGCP).Return(nil, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOptionAzure).Return(nil, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAWS)).Return(privateEndpoints, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderGCP)).Return(nil, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAzure)).Return(nil, nil)
		projectStore.EXPECT().EncryptionAtRest(projectID).Return(encryptionAtRest, nil)
		projectStore.EXPECT().CloudProviderAccessRoles(projectID).Return(cpas, nil)
		projectStore.EXPECT().ProjectSettings(projectID).Return(projectSettings, nil)
		projectStore.EXPECT().Auditing(projectID).Return(auditing, nil)
		projectStore.EXPECT().AlertConfigurations(projectID, listOption).Return(alertConfigs, nil)
		projectStore.EXPECT().DatabaseRoles(projectID).Return(customRoles, nil)
		projectStore.EXPECT().ProjectTeams(projectID).Return(projectTeams, nil)
		projectStore.EXPECT().TeamByID(orgID, teamID).Return(teams, nil)
		projectStore.EXPECT().TeamUsers(orgID, teamID).Return(teamUsers, nil)

		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureAccessLists).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureMaintenanceWindows).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureIntegrations).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureNetworkPeering).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featurePrivateEndpoints).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureEncryptionAtRest).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureCloudProviderAccessRoles).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureProjectSettings).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureAuditing).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureAlertConfiguration).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureCustomRoles).Return(true)
		featureValidator.EXPECT().FeatureExist(features.ResourceAtlasProject, featureTeams).Return(true)

		dictionary := resources.AtlasNameToKubernetesName()
		projectResult, err := BuildAtlasProject(projectStore, featureValidator, orgID, projectID, targetNamespace, true, dictionary, resourceVersion)
		if err != nil {
			t.Fatalf("%v", err)
		}
		gotProject := projectResult.Project
		gotTeams := projectResult.Teams

		expectedThreshold := &atlasV1.Threshold{
			Operator:  alertConfigs[0].Threshold.Operator,
			Units:     alertConfigs[0].Threshold.Units,
			Threshold: fmt.Sprintf("%f", alertConfigs[0].Threshold.Threshold),
		}
		expectedMatchers := []atlasV1.Matcher{
			{
				FieldName: alertConfigs[0].Matchers[0].FieldName,
				Operator:  alertConfigs[0].Matchers[0].Operator,
				Value:     alertConfigs[0].Matchers[0].Value,
			},
		}
		expectedNotifications := []atlasV1.Notification{
			{
				APIToken:            alertConfigs[0].Notifications[0].APIToken,
				ChannelName:         alertConfigs[0].Notifications[0].ChannelName,
				DatadogAPIKey:       alertConfigs[0].Notifications[0].DatadogAPIKey,
				DatadogRegion:       alertConfigs[0].Notifications[0].DatadogRegion,
				DelayMin:            alertConfigs[0].Notifications[0].DelayMin,
				EmailAddress:        alertConfigs[0].Notifications[0].EmailAddress,
				EmailEnabled:        alertConfigs[0].Notifications[0].EmailEnabled,
				FlowdockAPIToken:    alertConfigs[0].Notifications[0].FlowdockAPIToken,
				FlowName:            alertConfigs[0].Notifications[0].FlowName,
				IntervalMin:         alertConfigs[0].Notifications[0].IntervalMin,
				MobileNumber:        alertConfigs[0].Notifications[0].MobileNumber,
				OpsGenieAPIKey:      alertConfigs[0].Notifications[0].OpsGenieAPIKey,
				OpsGenieRegion:      alertConfigs[0].Notifications[0].OpsGenieRegion,
				OrgName:             alertConfigs[0].Notifications[0].OrgName,
				ServiceKey:          alertConfigs[0].Notifications[0].ServiceKey,
				SMSEnabled:          alertConfigs[0].Notifications[0].SMSEnabled,
				TeamID:              alertConfigs[0].Notifications[0].TeamID,
				TeamName:            alertConfigs[0].Notifications[0].TeamName,
				TypeName:            alertConfigs[0].Notifications[0].TypeName,
				Username:            alertConfigs[0].Notifications[0].Username,
				VictorOpsAPIKey:     alertConfigs[0].Notifications[0].VictorOpsAPIKey,
				VictorOpsRoutingKey: alertConfigs[0].Notifications[0].VictorOpsRoutingKey,
				Roles:               alertConfigs[0].Notifications[0].Roles,
			},
		}
		expectedMetricThreshold := &atlasV1.MetricThreshold{
			MetricName: alertConfigs[0].MetricThreshold.MetricName,
			Operator:   alertConfigs[0].MetricThreshold.Operator,
			Threshold:  fmt.Sprintf("%f", alertConfigs[0].MetricThreshold.Threshold),
			Units:      alertConfigs[0].MetricThreshold.Units,
			Mode:       alertConfigs[0].MetricThreshold.Mode,
		}
		expectedTeams := []*atlasV1.AtlasTeam{
			{
				TypeMeta: v1.TypeMeta{
					Kind:       "AtlasTeam",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      fmt.Sprintf("%s-team-%s", strings.ToLower(p.Name), strings.ToLower(teams.Name)),
					Namespace: targetNamespace,
					Labels: map[string]string{
						features.ResourceVersion: resourceVersion,
					},
				},
				Spec: atlasV1.TeamSpec{
					Name:      teams.Name,
					Usernames: []atlasV1.TeamUser{atlasV1.TeamUser(teamUsers[0].Username)},
				},
				Status: status.TeamStatus{
					Common: status.Common{
						Conditions: []status.Condition{},
					},
				},
			},
		}
		expectedProject := &atlasV1.AtlasProject{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasProject",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(p.Name, dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: atlasV1.AtlasProjectSpec{
				Name: p.Name,
				ConnectionSecret: &common.ResourceRefNamespaced{
					Name: resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, p.Name), dictionary),
				},
				ProjectIPAccessList: []project.IPAccessList{
					{
						AwsSecurityGroup: ipAccessLists.Results[0].GetAwsSecurityGroup(),
						CIDRBlock:        ipAccessLists.Results[0].GetCidrBlock(),
						Comment:          ipAccessLists.Results[0].GetComment(),
						DeleteAfterDate:  ipAccessLists.Results[0].GetDeleteAfterDate().String(),
						IPAddress:        ipAccessLists.Results[0].GetIpAddress(),
					},
				},
				MaintenanceWindow: project.MaintenanceWindow{
					DayOfWeek: mw.DayOfWeek,
					HourOfDay: mw.HourOfDay,
					AutoDefer: pointer.GetOrDefault(mw.AutoDeferOnceEnabled, false),
					StartASAP: pointer.GetOrDefault(mw.StartASAP, false),
					Defer:     false,
				},
				PrivateEndpoints: []atlasV1.PrivateEndpoint{
					{
						Provider:          provider.ProviderAWS,
						Region:            *privateAWSEndpoint.RegionName,
						ID:                *privateAWSEndpoint.Id,
						IP:                "",
						GCPProjectID:      "",
						EndpointGroupName: "",
						Endpoints:         atlasV1.GCPEndpoints{},
					},
				},
				CloudProviderAccessRoles: []atlasV1.CloudProviderAccessRole{
					{
						ProviderName:      cpas.AwsIamRoles[0].ProviderName,
						IamAssumedRoleArn: *cpas.AwsIamRoles[0].IamAssumedRoleArn,
					},
				},
				AlertConfigurations: []atlasV1.AlertConfiguration{
					{
						Enabled:         *alertConfigs[0].Enabled,
						EventTypeName:   alertConfigs[0].EventTypeName,
						Matchers:        expectedMatchers,
						Threshold:       expectedThreshold,
						Notifications:   expectedNotifications,
						MetricThreshold: expectedMetricThreshold,
					},
				},
				AlertConfigurationSyncEnabled: false,
				NetworkPeers: []atlasV1.NetworkPeer{
					{
						AccepterRegionName:  peeringConnectionAWS.AccepterRegionName,
						ContainerRegion:     "",
						AWSAccountID:        peeringConnectionAWS.AwsAccountId,
						ContainerID:         peeringConnectionAWS.ContainerId,
						ProviderName:        provider.ProviderName(*peeringConnectionAWS.ProviderName),
						RouteTableCIDRBlock: peeringConnectionAWS.RouteTableCidrBlock,
						VpcID:               peeringConnectionAWS.VpcId,
					},
				},
				WithDefaultAlertsSettings: false,
				X509CertRef:               nil,
				Integrations: []project.Integration{
					{
						Type:     thirdPartyIntegrations.Results[0].Prometheus.GetType(),
						UserName: thirdPartyIntegrations.Results[0].Prometheus.GetUsername(),
						PasswordRef: common.ResourceRefNamespaced{
							Name: fmt.Sprintf("%s-integration-%s",
								strings.ToLower(projectID),
								strings.ToLower(thirdPartyIntegrations.Results[0].Prometheus.GetType())),
							Namespace: targetNamespace,
						},
						ServiceDiscovery: thirdPartyIntegrations.Results[0].Prometheus.ServiceDiscovery,
					},
				},
				EncryptionAtRest: &atlasV1.EncryptionAtRest{
					AwsKms:        atlasV1.AwsKms{},
					AzureKeyVault: atlasV1.AzureKeyVault{},
					GoogleCloudKms: atlasV1.GoogleCloudKms{
						Enabled:              encryptionAtRest.GoogleCloudKms.Enabled,
						ServiceAccountKey:    encryptionAtRest.GoogleCloudKms.GetServiceAccountKey(),
						KeyVersionResourceID: encryptionAtRest.GoogleCloudKms.GetKeyVersionResourceID(),
					},
				},
				Auditing: &atlasV1.Auditing{
					AuditAuthorizationSuccess: &auditing.AuditAuthorizationSuccess,
					AuditFilter:               auditing.AuditFilter,
					Enabled:                   &auditing.Enabled,
				},
				Settings: &atlasV1.ProjectSettings{
					IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
					IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
					IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
					IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
					IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
				},
				CustomRoles: []atlasV1.CustomRole{
					{
						Name: customRoles[0].RoleName,
						InheritedRoles: []atlasV1.Role{
							{
								Name:     customRoles[0].InheritedRoles[0].Role,
								Database: customRoles[0].InheritedRoles[0].Db,
							},
						},
						Actions: []atlasV1.Action{
							{
								Name: customRoles[0].Actions[0].Action,
								Resources: []atlasV1.Resource{
									{
										Cluster:    &customRoles[0].Actions[0].Resources[0].Cluster,
										Database:   &customRoles[0].Actions[0].Resources[0].Db,
										Collection: &customRoles[0].Actions[0].Resources[0].Collection,
									},
								},
							},
						},
					},
				},
				Teams: []atlasV1.Team{
					{
						TeamRef: common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-team-%s", strings.ToLower(p.Name), strings.ToLower(teams.Name)),
							Namespace: targetNamespace,
						},
						Roles: []atlasV1.TeamRole{atlasV1.TeamRole(projectTeams.Results[0].RoleNames[0])},
					},
				},
			},
			Status: status.AtlasProjectStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		if !reflect.DeepEqual(expectedProject, gotProject) {
			t.Fatalf("Project mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expectedProject, gotProject)
		}

		if !reflect.DeepEqual(expectedTeams, gotTeams) {
			t.Fatalf("Teams mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expectedTeams, gotTeams)
		}
	})
}

func TestBuildProjectConnectionSecret(t *testing.T) {
	ctl := gomock.NewController(t)
	dictionary := resources.AtlasNameToKubernetesName()
	credsProvider := mocks.NewMockCredentialsGetter(ctl)
	t.Run("Can generate a valid connection secret WITH data", func(t *testing.T) {
		publicAPIKey := "TestPublicKey"
		privateAPIKey := "TestPrivateKey"

		name := "TestSecret-1"
		namespace := "TestNamespace-1"

		credsProvider.EXPECT().PublicAPIKey().Return(publicAPIKey)
		credsProvider.EXPECT().PrivateAPIKey().Return(privateAPIKey)

		got := BuildProjectConnectionSecret(credsProvider, name, namespace,
			orgID, true, dictionary)

		expected := &corev1.Secret{
			TypeMeta: v1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-credentials", name)),
				Namespace: namespace,
				Labels: map[string]string{
					secrets.TypeLabelKey: secrets.CredLabelVal,
				},
			},
			Data: map[string][]byte{
				secrets.CredOrgID:         []byte(orgID),
				secrets.CredPublicAPIKey:  []byte(publicAPIKey),
				secrets.CredPrivateAPIKey: []byte(privateAPIKey),
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Credentials secret mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can generate a valid connection secret WITHOUT data", func(t *testing.T) {
		name := "TestSecret"
		namespace := "TestNamespace"

		got := BuildProjectConnectionSecret(credsProvider, name, namespace,
			orgID, false, dictionary)

		expected := &corev1.Secret{
			TypeMeta: v1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-credentials", name)),
				Namespace: namespace,
				Labels: map[string]string{
					secrets.TypeLabelKey: secrets.CredLabelVal,
				},
			},
			Data: map[string][]byte{
				secrets.CredOrgID:         []byte(""),
				secrets.CredPublicAPIKey:  []byte(""),
				secrets.CredPrivateAPIKey: []byte(""),
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Credentials secret mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildAccessLists(t *testing.T) {
	ctl := gomock.NewController(t)

	alProvider := mocks.NewMockProjectIPAccessListLister(ctl)
	t.Run("Can convert Access Lists", func(t *testing.T) {
		data := &atlasv2.PaginatedNetworkAccess{
			Links: nil,
			Results: []atlasv2.NetworkPermissionEntry{
				{
					AwsSecurityGroup: pointer.Get("TestSecGroup"),
					CidrBlock:        pointer.Get("0.0.0.0/0"),
					Comment:          pointer.Get("TestComment"),
					DeleteAfterDate:  pointer.Get(time.Now()),
					GroupId:          pointer.Get("TestGroupID"),
					IpAddress:        pointer.Get("0.0.0.0"),
				},
			},
			TotalCount: pointer.Get(1),
		}

		listOptions := &mongodbatlas.ListOptions{ItemsPerPage: MaxItems}

		alProvider.EXPECT().ProjectIPAccessLists(projectID, listOptions).Return(data, nil)

		got, err := buildAccessLists(alProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []project.IPAccessList{
			{
				AwsSecurityGroup: data.Results[0].GetAwsSecurityGroup(),
				CIDRBlock:        data.Results[0].GetCidrBlock(),
				Comment:          data.Results[0].GetComment(),
				DeleteAfterDate:  data.Results[0].GetDeleteAfterDate().String(),
				IPAddress:        data.Results[0].GetIpAddress(),
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("IPAccessList mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildAuditing(t *testing.T) {
	ctl := gomock.NewController(t)

	auditingProvider := mocks.NewMockAuditingDescriber(ctl)
	t.Run("Can convert Auditing", func(t *testing.T) {
		data := &atlasv2.AuditLog{
			AuditAuthorizationSuccess: true,
			AuditFilter:               "TestFilter",
			ConfigurationType:         pointer.Get("TestType"),
			Enabled:                   true,
		}

		auditingProvider.EXPECT().Auditing(projectID).Return(data, nil)

		got, err := buildAuditing(auditingProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.Auditing{
			AuditAuthorizationSuccess: &data.AuditAuthorizationSuccess,
			AuditFilter:               data.AuditFilter,
			Enabled:                   &data.Enabled,
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Auditing mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildCloudProviderAccessRoles(t *testing.T) {
	ctl := gomock.NewController(t)

	cpaProvider := mocks.NewMockCloudProviderAccessRoleLister(ctl)
	t.Run("Can convert CPA roles", func(t *testing.T) {
		data := &atlasv2.CloudProviderAccessRoles{
			AwsIamRoles: []atlasv2.CloudProviderAccessAWSIAMRole{
				{
					AtlasAWSAccountArn:         pointer.Get("TestARN"),
					AtlasAssumedRoleExternalId: pointer.Get("TestRoleID"),
					AuthorizedDate:             &time.Time{},
					CreatedDate:                &time.Time{},
					FeatureUsages:              nil,
					IamAssumedRoleArn:          pointer.Get("TestAssumedRoleARN"),
					ProviderName:               string(provider.ProviderAWS),
					RoleId:                     pointer.Get("TestRoleID"),
				},
			},
		}

		cpaProvider.EXPECT().CloudProviderAccessRoles(projectID).Return(data, nil)

		got, err := buildCloudProviderAccessRoles(cpaProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []atlasV1.CloudProviderAccessRole{
			{
				ProviderName:      data.AwsIamRoles[0].ProviderName,
				IamAssumedRoleArn: *data.AwsIamRoles[0].IamAssumedRoleArn,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("CPA mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildEncryptionAtREST(t *testing.T) {
	ctl := gomock.NewController(t)

	dataProvider := mocks.NewMockEncryptionAtRestDescriber(ctl)
	t.Run("Can convert Encryption at REST AWS", func(t *testing.T) {
		data := &atlasv2.EncryptionAtRest{
			AwsKms: &atlasv2.AWSKMSConfiguration{
				Enabled:             pointer.Get(true),
				AccessKeyID:         pointer.Get("TestAccessKey"),
				SecretAccessKey:     pointer.Get("TestSecretAccessKey"),
				CustomerMasterKeyID: pointer.Get("TestCustomerMasterKeyID"),
				Region:              pointer.Get("US_EAST_1"),
				RoleId:              pointer.Get("TestRoleID"),
				Valid:               pointer.Get(true),
			},
			AzureKeyVault:  &atlasv2.AzureKeyVault{},
			GoogleCloudKms: &atlasv2.GoogleCloudKMS{},
		}

		dataProvider.EXPECT().EncryptionAtRest(projectID).Return(data, nil)

		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms: atlasV1.AwsKms{
				Enabled:             data.AwsKms.Enabled,
				AccessKeyID:         data.AwsKms.GetAccessKeyID(),
				SecretAccessKey:     data.AwsKms.GetSecretAccessKey(),
				CustomerMasterKeyID: data.AwsKms.GetCustomerMasterKeyID(),
				Region:              data.AwsKms.GetRegion(),
				RoleID:              data.AwsKms.GetRoleId(),
				Valid:               data.AwsKms.Valid,
			},
			AzureKeyVault:  atlasV1.AzureKeyVault{},
			GoogleCloudKms: atlasV1.GoogleCloudKms{},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can convert Encryption at REST GCP", func(t *testing.T) {
		data := &atlasv2.EncryptionAtRest{
			AwsKms:        &atlasv2.AWSKMSConfiguration{},
			AzureKeyVault: &atlasv2.AzureKeyVault{},
			GoogleCloudKms: &atlasv2.GoogleCloudKMS{
				Enabled:              pointer.Get(true),
				ServiceAccountKey:    pointer.Get("TestServiceAccountKey"),
				KeyVersionResourceID: pointer.Get("TestVersionResourceID"),
			},
		}

		dataProvider.EXPECT().EncryptionAtRest(projectID).Return(data, nil)
		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms:        atlasV1.AwsKms{},
			AzureKeyVault: atlasV1.AzureKeyVault{},
			GoogleCloudKms: atlasV1.GoogleCloudKms{
				Enabled:              data.GoogleCloudKms.Enabled,
				ServiceAccountKey:    data.GoogleCloudKms.GetServiceAccountKey(),
				KeyVersionResourceID: data.GoogleCloudKms.GetKeyVersionResourceID(),
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can convert Encryption at REST Azure", func(t *testing.T) {
		data := atlasv2.EncryptionAtRest{
			AwsKms: &atlasv2.AWSKMSConfiguration{},
			AzureKeyVault: &atlasv2.AzureKeyVault{
				Enabled:           pointer.Get(true),
				ClientID:          pointer.Get("TestClientID"),
				AzureEnvironment:  pointer.Get("TestAzureEnv"),
				SubscriptionID:    pointer.Get("TestSubID"),
				ResourceGroupName: pointer.Get("TestResourceGroupName"),
				KeyVaultName:      pointer.Get("TestKeyVaultName"),
				KeyIdentifier:     pointer.Get("TestKeyIdentifier"),
				Secret:            pointer.Get("TestSecret"),
				TenantID:          pointer.Get("TestTenantID"),
			},
			GoogleCloudKms: &atlasv2.GoogleCloudKMS{},
		}

		dataProvider.EXPECT().EncryptionAtRest(projectID).Return(&data, nil)
		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms: atlasV1.AwsKms{},
			AzureKeyVault: atlasV1.AzureKeyVault{
				Enabled:           data.AzureKeyVault.Enabled,
				ClientID:          data.AzureKeyVault.GetClientID(),
				AzureEnvironment:  data.AzureKeyVault.GetAzureEnvironment(),
				SubscriptionID:    data.AzureKeyVault.GetSubscriptionID(),
				ResourceGroupName: data.AzureKeyVault.GetResourceGroupName(),
				KeyVaultName:      data.AzureKeyVault.GetKeyVaultName(),
				KeyIdentifier:     data.AzureKeyVault.GetKeyIdentifier(),
				Secret:            data.AzureKeyVault.GetSecret(),
				TenantID:          data.AzureKeyVault.GetTenantID(),
			},
			GoogleCloudKms: atlasV1.GoogleCloudKms{},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildIntegrations(t *testing.T) {
	ctl := gomock.NewController(t)

	intProvider := mocks.NewMockIntegrationLister(ctl)
	dictionary := resources.AtlasNameToKubernetesName()

	t.Run("Can convert third-party integrations WITH secrets: Prometheus", func(t *testing.T) {
		const targetNamespace = "test-namespace-3"
		const includeSecrets = true
		ints := &atlasv2.PaginatedIntegration{
			Links: nil,
			Results: []atlasv2.ThridPartyIntegration{
				{
					Prometheus: &atlasv2.Prometheus{
						Type:             pointer.Get("PROMETHEUS"),
						Password:         pointer.Get("PrometheusTestPassword"),
						Username:         "PrometheusTestUserName",
						ServiceDiscovery: "TestServiceDiscovery",
					},
				},
			},
			TotalCount: pointer.Get(0),
		}

		intProvider.EXPECT().Integrations(projectID).Return(ints, nil)

		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets, dictionary)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []project.Integration{
			{
				Type:             ints.Results[0].Prometheus.GetType(),
				ServiceDiscovery: ints.Results[0].Prometheus.ServiceDiscovery,
				UserName:         ints.Results[0].Prometheus.Username,
				PasswordRef: common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.Results[0].Prometheus.GetType())),
					Namespace: targetNamespace,
				},
			},
		}

		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: v1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: v1.ObjectMeta{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.Results[0].Prometheus.GetType())),
					Namespace: targetNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{
					secrets.PasswordField: []byte(ints.Results[0].Prometheus.GetPassword()),
				},
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Integrations mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}

		if !reflect.DeepEqual(expectedSecrets, intSecrets) {
			t.Fatalf("Integrations secrets mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expectedSecrets, intSecrets)
		}
	})
	t.Run("Can convert third-party integrations WITHOUT secrets: Prometheus", func(t *testing.T) {
		const targetNamespace = "test-namespace-4"
		const includeSecrets = false
		ints := &atlasv2.PaginatedIntegration{
			Links: nil,
			Results: []atlasv2.ThridPartyIntegration{
				{
					Prometheus: &atlasv2.Prometheus{
						Type:             pointer.Get("PROMETHEUS"),
						Password:         pointer.Get("PrometheusTestPassword"),
						Username:         "PrometheusTestUserName",
						ServiceDiscovery: "TestServiceDiscovery",
					},
				},
			},
			TotalCount: pointer.Get(0),
		}
		intProvider.EXPECT().Integrations(projectID).Return(ints, nil)
		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets, dictionary)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []project.Integration{
			{
				Type:             ints.Results[0].Prometheus.GetType(),
				ServiceDiscovery: ints.Results[0].Prometheus.ServiceDiscovery,
				UserName:         ints.Results[0].Prometheus.Username,
				PasswordRef: common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.Results[0].Prometheus.GetType())),
					Namespace: targetNamespace,
				},
			},
		}

		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: v1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: v1.ObjectMeta{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.Results[0].Prometheus.GetType())),
					Namespace: targetNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{
					secrets.PasswordField: []byte(""),
				},
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Integrations mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}

		if !reflect.DeepEqual(expectedSecrets, intSecrets) {
			t.Fatalf("Integrations secrets mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expectedSecrets, intSecrets)
		}
	})
}

func Test_buildMaintenanceWindows(t *testing.T) {
	ctl := gomock.NewController(t)

	mwProvider := mocks.NewMockMaintenanceWindowDescriber(ctl)
	t.Run("Can convert maintenance window", func(t *testing.T) {
		mw := &atlasv2.GroupMaintenanceWindow{
			DayOfWeek:            3,
			HourOfDay:            10,
			StartASAP:            pointer.Get(false),
			AutoDeferOnceEnabled: pointer.Get(false),
		}

		mwProvider.EXPECT().MaintenanceWindow(projectID).Return(mw, nil)

		got, err := buildMaintenanceWindows(mwProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := project.MaintenanceWindow{
			DayOfWeek: mw.DayOfWeek,
			HourOfDay: mw.HourOfDay,
			AutoDefer: *mw.AutoDeferOnceEnabled,
			StartASAP: *mw.StartASAP,
			Defer:     false,
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("MaintenanceWindows mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildNetworkPeering(t *testing.T) {
	ctl := gomock.NewController(t)

	peerProvider := mocks.NewMockPeeringConnectionLister(ctl)
	t.Run("Can convert Peering connections", func(t *testing.T) {
		peeringConnectionAWS := &atlasv2.AwsNetworkPeeringConnectionSettings{
			AccepterRegionName:  "TestRegionName",
			AwsAccountId:        "TestAWSAccountID",
			ConnectionId:        pointer.Get("TestConnID"),
			ContainerId:         "TestContainerID",
			ErrorStateName:      pointer.Get("TestErrStateName"),
			Id:                  pointer.Get("TestID"),
			ProviderName:        pointer.Get(string(provider.ProviderAWS)),
			RouteTableCidrBlock: "0.0.0.0/0",
			StatusName:          pointer.Get("TestStatusName"),
			VpcId:               "TestVPCID",
		}

		peeringConnections := []interface{}{
			peeringConnectionAWS,
		}

		listOptions := mongodbatlas.ListOptions{ItemsPerPage: MaxItems}
		containerListOptionAWS := &mongodbatlas.ContainersListOptions{ListOptions: listOptions, ProviderName: string(provider.ProviderAWS)}
		containerListOptionGCP := &mongodbatlas.ContainersListOptions{ListOptions: listOptions, ProviderName: string(provider.ProviderGCP)}
		containerListOptionAzure := &mongodbatlas.ContainersListOptions{ListOptions: listOptions, ProviderName: string(provider.ProviderAzure)}

		peerProvider.EXPECT().PeeringConnections(projectID, containerListOptionAWS).Return(peeringConnections, nil)
		peerProvider.EXPECT().PeeringConnections(projectID, containerListOptionGCP).Return(nil, nil)
		peerProvider.EXPECT().PeeringConnections(projectID, containerListOptionAzure).Return(nil, nil)

		got, err := buildNetworkPeering(peerProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.NetworkPeer{
			{
				AccepterRegionName:  peeringConnectionAWS.AccepterRegionName,
				ContainerRegion:     "",
				AWSAccountID:        peeringConnectionAWS.AwsAccountId,
				ContainerID:         peeringConnectionAWS.ContainerId,
				ProviderName:        provider.ProviderName(*peeringConnectionAWS.ProviderName),
				RouteTableCIDRBlock: peeringConnectionAWS.RouteTableCidrBlock,
				VpcID:               peeringConnectionAWS.VpcId,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("NetworkPeerings mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildPrivateEndpoints(t *testing.T) {
	ctl := gomock.NewController(t)

	peProvider := mocks.NewMockPrivateEndpointLister(ctl)
	t.Run("Can convert PrivateEndpointConnection for AWS", func(t *testing.T) {
		providerName := provider.ProviderAWS
		privateEndpoint := atlasv2.EndpointService{
			Id:                  pointer.Get("1"),
			CloudProvider:       string(providerName),
			RegionName:          pointer.Get("US_EAST_1"),
			EndpointServiceName: nil,
			ErrorMessage:        nil,
			InterfaceEndpoints:  nil,
			Status:              nil,
		}

		peProvider.EXPECT().PrivateEndpoints(projectID, string(providerName)).Return([]atlasv2.EndpointService{privateEndpoint}, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAzure)).Return(nil, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderGCP)).Return(nil, nil)

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.PrivateEndpoint{
			{
				Provider:          providerName,
				Region:            *privateEndpoint.RegionName,
				ID:                *privateEndpoint.Id,
				IP:                "",
				GCPProjectID:      "",
				EndpointGroupName: "",
				Endpoints:         atlasV1.GCPEndpoints{},
			},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("PrivateEndpoints mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})

	t.Run("Can convert PrivateEndpointConnection for Azure", func(t *testing.T) {
		providerName := provider.ProviderAzure
		privateEndpoint := atlasv2.EndpointService{
			Id:                           pointer.Get("1"),
			CloudProvider:                string(providerName),
			RegionName:                   pointer.Get("uswest3"),
			ErrorMessage:                 pointer.Get(""),
			PrivateEndpoints:             nil,
			PrivateLinkServiceName:       nil,
			PrivateLinkServiceResourceId: nil,
			Status:                       nil,
		}

		peProvider.EXPECT().PrivateEndpoints(projectID, string(providerName)).Return([]atlasv2.EndpointService{privateEndpoint}, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAWS)).Return(nil, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderGCP)).Return(nil, nil)

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.PrivateEndpoint{
			{
				Provider:          providerName,
				Region:            *privateEndpoint.RegionName,
				ID:                *privateEndpoint.Id,
				IP:                "",
				GCPProjectID:      "",
				EndpointGroupName: "",
				Endpoints:         atlasV1.GCPEndpoints{},
			},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("PrivateEndpoints mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildProjectSettings(t *testing.T) {
	ctl := gomock.NewController(t)

	settingsProvider := mocks.NewMockProjectSettingsDescriber(ctl)
	t.Run("Can convert project settings", func(t *testing.T) {
		projectSettings := atlasv2.GroupSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: pointer.Get(true),
			IsDataExplorerEnabled:                       pointer.Get(true),
			IsPerformanceAdvisorEnabled:                 pointer.Get(true),
			IsRealtimePerformancePanelEnabled:           pointer.Get(true),
			IsSchemaAdvisorEnabled:                      pointer.Get(true),
		}
		settingsProvider.EXPECT().ProjectSettings(projectID).Return(&projectSettings, nil)

		got, err := buildProjectSettings(settingsProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expected := &atlasV1.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
			IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
			IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Project settings mismatch. expected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildCustomRoles(t *testing.T) {
	ctl := gomock.NewController(t)

	rolesProvider := mocks.NewMockDatabaseRoleLister(ctl)
	t.Run("Can build custom roles", func(t *testing.T) {
		data := []atlasv2.UserCustomDBRole{
			{
				Actions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TestAction",
						Resources: []atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: "TestCollection",
								Db:         "TestDB",
								Cluster:    true,
							},
						},
					},
				},
				InheritedRoles: []atlasv2.DatabaseInheritedRole{
					{
						Db:   "TestDBMAIN",
						Role: "ADMIN",
					},
				},
				RoleName: "TestRoleName",
			},
		}

		rolesProvider.EXPECT().DatabaseRoles(projectID).Return(data, nil)

		role := data[0]
		expected := []atlasV1.CustomRole{
			{
				Name: role.RoleName,
				InheritedRoles: []atlasV1.Role{
					{
						Name:     role.InheritedRoles[0].Role,
						Database: role.InheritedRoles[0].Db,
					},
				},
				Actions: []atlasV1.Action{
					{
						Name: role.Actions[0].Action,
						Resources: []atlasV1.Resource{
							{
								Cluster:    &role.Actions[0].Resources[0].Cluster,
								Database:   &role.Actions[0].Resources[0].Db,
								Collection: &role.Actions[0].Resources[0].Collection,
							},
						},
					},
				},
			},
		}

		got, err := buildCustomRoles(rolesProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Custom Roles mismatch. expected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}
