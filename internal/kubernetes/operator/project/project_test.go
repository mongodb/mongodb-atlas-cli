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

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/pointers"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const orgID = "TestOrgID"
const projectID = "TestProjectID"
const teamID = "TestTeamID"

func TestBuildAtlasProject(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	projectStore := mocks.NewMockAtlasOperatorProjectStore(ctl)

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
			WithDefaultAlertsSettings: pointers.MakePtr(false),
		}

		ipAccessLists := &mongodbatlas.ProjectIPAccessLists{
			Links: nil,
			Results: []mongodbatlas.ProjectIPAccessList{
				{
					AwsSecurityGroup: "TestSecurity group",
					CIDRBlock:        "0.0.0.0/0",
					Comment:          "Allow everyone",
					DeleteAfterDate:  "",
					GroupID:          "TestGroupID",
					IPAddress:        "0.0.0.0",
				},
			},
			TotalCount: 1,
		}

		auditing := &mongodbatlas.Auditing{
			AuditAuthorizationSuccess: pointers.MakePtr(true),
			AuditFilter:               "TestFilter",
			ConfigurationType:         "TestConfigType",
			Enabled:                   pointers.MakePtr(true),
		}

		cpas := &mongodbatlas.CloudProviderAccessRoles{
			AWSIAMRoles: []mongodbatlas.AWSIAMRole{
				{
					AtlasAWSAccountARN:         "TestARN",
					AtlasAssumedRoleExternalID: "TestExternalRoleID",
					AuthorizedDate:             "01-01-2001",
					CreatedDate:                "01-02-2001",
					FeatureUsages:              nil,
					IAMAssumedRoleARN:          "TestRoleARN",
					ProviderName:               string(provider.ProviderAWS),
					RoleID:                     "TestRoleID",
				},
			},
		}

		encryptionAtRest := &mongodbatlas.EncryptionAtRest{
			GroupID:       "TestGroupID",
			AwsKms:        mongodbatlas.AwsKms{},
			AzureKeyVault: mongodbatlas.AzureKeyVault{},
			GoogleCloudKms: mongodbatlas.GoogleCloudKms{
				Enabled:              pointers.MakePtr(true),
				ServiceAccountKey:    "TestServiceAccountKey",
				KeyVersionResourceID: "TestKeyVersionResourceID",
			},
		}

		thirdPartyIntegrations := &mongodbatlas.ThirdPartyIntegrations{
			Links: nil,
			Results: []*mongodbatlas.ThirdPartyIntegration{
				{
					Type:             "PROMETHEUS",
					UserName:         "TestPrometheusUserName",
					Password:         "TestPrometheusPassword",
					ServiceDiscovery: "TestPrometheusServiceDiscovery",
				},
			},
			TotalCount: 1,
		}

		mw := &mongodbatlas.MaintenanceWindow{
			DayOfWeek:            1,
			HourOfDay:            pointers.MakePtr(10),
			StartASAP:            pointers.MakePtr(false),
			NumberOfDeferrals:    0,
			AutoDeferOnceEnabled: pointers.MakePtr(false),
		}

		peeringConnections := []mongodbatlas.Peer{
			{
				AccepterRegionName:  "US_EAST_1",
				AWSAccountID:        "TestAwsAccountID",
				ConnectionID:        "TestConnectionID",
				ContainerID:         "TestContainerID",
				ErrorStateName:      "TestErrorStateName",
				ID:                  "TestID",
				ProviderName:        string(provider.ProviderAWS),
				RouteTableCIDRBlock: "0.0.0.0/0",
				StatusName:          "",
				VpcID:               "TestVPCID",
				AtlasCIDRBlock:      "0.0.0.0/0",
				AzureDirectoryID:    "TestDirectoryID",
				AzureSubscriptionID: "TestAzureSubID",
				ResourceGroupName:   "TestResourceGroupName",
				VNetName:            "TestVNetName",
				ErrorState:          "TestErrorState",
				Status:              "TestStatus",
				GCPProjectID:        "TestGCPProjectID",
				NetworkName:         "TestNetworkName",
				ErrorMessage:        "TestErrorMessage",
			},
		}

		privateEndpoints := []mongodbatlas.PrivateEndpointConnection{

			{
				ID:                           "TestID",
				ProviderName:                 string(provider.ProviderAWS),
				Region:                       "US_WEST_2",
				EndpointServiceName:          "",
				ErrorMessage:                 "",
				InterfaceEndpoints:           nil,
				PrivateEndpoints:             nil,
				PrivateLinkServiceName:       "",
				PrivateLinkServiceResourceID: "",
				Status:                       "",
				EndpointGroupNames:           nil,
				RegionName:                   "",
				ServiceAttachmentNames:       nil,
			},
		}

		alertConfigs := []mongodbatlas.AlertConfiguration{
			{
				EventTypeName: "TestEventTypeName",
				Enabled:       pointers.MakePtr(true),
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
						DelayMin:            pointers.MakePtr(5),
						EmailAddress:        "TestEmail@mongodb.com",
						EmailEnabled:        pointers.MakePtr(true),
						FlowdockAPIToken:    "TestFlowDockApiToken",
						FlowName:            "TestFlowName",
						IntervalMin:         0,
						MobileNumber:        "+12345678900",
						OpsGenieAPIKey:      "TestGenieAPIKey",
						OpsGenieRegion:      "TestGenieRegion",
						OrgName:             "TestOrgName",
						ServiceKey:          "TestServiceKey",
						SMSEnabled:          pointers.MakePtr(true),
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

		projectSettings := &mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: pointers.MakePtr(true),
			IsDataExplorerEnabled:                       pointers.MakePtr(true),
			IsPerformanceAdvisorEnabled:                 pointers.MakePtr(true),
			IsRealtimePerformancePanelEnabled:           pointers.MakePtr(true),
			IsSchemaAdvisorEnabled:                      pointers.MakePtr(true),
		}

		customRoles := []mongodbatlas.CustomDBRole{
			{
				Actions: []mongodbatlas.Action{
					{
						Action: "Action-1",
						Resources: []mongodbatlas.Resource{
							{
								Collection: pointers.MakePtr("Collection-1"),
								DB:         pointers.MakePtr("DB-1"),
								Cluster:    pointers.MakePtr(true),
							},
						},
					},
				},
				InheritedRoles: []mongodbatlas.InheritedRole{
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
		containerListOption := &mongodbatlas.ContainersListOptions{ListOptions: *listOption}
		projectStore.EXPECT().Project(projectID).Return(p, nil)
		projectStore.EXPECT().ProjectIPAccessLists(projectID, listOption).Return(ipAccessLists, nil)
		projectStore.EXPECT().MaintenanceWindow(projectID).Return(mw, nil)
		projectStore.EXPECT().Integrations(projectID).Return(thirdPartyIntegrations, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOption).Return(peeringConnections, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAWS), listOption).Return(privateEndpoints, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderGCP), listOption).Return(nil, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAzure), listOption).Return(nil, nil)
		projectStore.EXPECT().EncryptionAtRest(projectID).Return(encryptionAtRest, nil)
		projectStore.EXPECT().CloudProviderAccessRoles(projectID).Return(cpas, nil)
		projectStore.EXPECT().ProjectSettings(projectID).Return(projectSettings, nil)
		projectStore.EXPECT().Auditing(projectID).Return(auditing, nil)
		projectStore.EXPECT().AlertConfigurations(projectID, listOption).Return(alertConfigs, nil)
		projectStore.EXPECT().DatabaseRoles(projectID, listOption).Return(&customRoles, nil)
		projectStore.EXPECT().ProjectTeams(projectID).Return(projectTeams, nil)
		projectStore.EXPECT().TeamByID(orgID, teamID).Return(teams, nil)
		projectStore.EXPECT().TeamUsers(orgID, teamID).Return(teamUsers, nil)

		projectResult, err := BuildAtlasProject(projectStore, orgID, projectID, targetNamespace, true)
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
				Name:      strings.ToLower(p.Name),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasProjectSpec{
				Name: p.Name,
				ConnectionSecret: &common.ResourceRef{
					Name: fmt.Sprintf(credSecretFormat, p.Name),
				},
				ProjectIPAccessList: []project.IPAccessList{
					{
						AwsSecurityGroup: ipAccessLists.Results[0].AwsSecurityGroup,
						CIDRBlock:        ipAccessLists.Results[0].CIDRBlock,
						Comment:          ipAccessLists.Results[0].Comment,
						DeleteAfterDate:  ipAccessLists.Results[0].DeleteAfterDate,
						IPAddress:        ipAccessLists.Results[0].IPAddress,
					},
				},
				MaintenanceWindow: project.MaintenanceWindow{
					DayOfWeek: mw.DayOfWeek,
					HourOfDay: pointers.PtrValOrDefault(mw.HourOfDay, 0),
					AutoDefer: pointers.PtrValOrDefault(mw.AutoDeferOnceEnabled, false),
					StartASAP: pointers.PtrValOrDefault(mw.StartASAP, false),
					Defer:     false,
				},
				PrivateEndpoints: []atlasV1.PrivateEndpoint{
					{
						Provider:          provider.ProviderAWS,
						Region:            privateEndpoints[0].Region,
						ID:                privateEndpoints[0].ID,
						IP:                "",
						GCPProjectID:      "",
						EndpointGroupName: "",
						Endpoints:         atlasV1.GCPEndpoints{},
					},
				},
				CloudProviderAccessRoles: []atlasV1.CloudProviderAccessRole{
					{
						ProviderName:      cpas.AWSIAMRoles[0].ProviderName,
						IamAssumedRoleArn: cpas.AWSIAMRoles[0].IAMAssumedRoleARN,
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
						AccepterRegionName:  peeringConnections[0].AccepterRegionName,
						ContainerRegion:     "",
						AWSAccountID:        peeringConnections[0].AWSAccountID,
						ContainerID:         peeringConnections[0].ContainerID,
						ProviderName:        provider.ProviderName(peeringConnections[0].ProviderName),
						RouteTableCIDRBlock: peeringConnections[0].RouteTableCIDRBlock,
						VpcID:               peeringConnections[0].VpcID,
						AtlasCIDRBlock:      peeringConnections[0].AtlasCIDRBlock,
						AzureDirectoryID:    peeringConnections[0].AzureDirectoryID,
						AzureSubscriptionID: peeringConnections[0].AzureSubscriptionID,
						ResourceGroupName:   peeringConnections[0].ResourceGroupName,
						VNetName:            peeringConnections[0].VNetName,
						GCPProjectID:        peeringConnections[0].GCPProjectID,
						NetworkName:         peeringConnections[0].NetworkName,
					},
				},
				WithDefaultAlertsSettings: false,
				X509CertRef:               nil,
				Integrations: []project.Integration{
					{
						Type:     thirdPartyIntegrations.Results[0].Type,
						UserName: thirdPartyIntegrations.Results[0].UserName,
						PasswordRef: common.ResourceRefNamespaced{
							Name: fmt.Sprintf("%s-integration-%s",
								strings.ToLower(projectID),
								strings.ToLower(thirdPartyIntegrations.Results[0].Type)),
							Namespace: targetNamespace,
						},
						ServiceDiscovery: thirdPartyIntegrations.Results[0].ServiceDiscovery,
					},
				},
				EncryptionAtRest: &atlasV1.EncryptionAtRest{
					AwsKms:        atlasV1.AwsKms{},
					AzureKeyVault: atlasV1.AzureKeyVault{},
					GoogleCloudKms: atlasV1.GoogleCloudKms{
						Enabled:              encryptionAtRest.GoogleCloudKms.Enabled,
						ServiceAccountKey:    encryptionAtRest.GoogleCloudKms.ServiceAccountKey,
						KeyVersionResourceID: encryptionAtRest.GoogleCloudKms.KeyVersionResourceID,
					},
				},
				Auditing: &atlasV1.Auditing{
					AuditAuthorizationSuccess: auditing.AuditAuthorizationSuccess,
					AuditFilter:               auditing.AuditFilter,
					Enabled:                   auditing.Enabled,
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
										Cluster:    customRoles[0].Actions[0].Resources[0].Cluster,
										Database:   customRoles[0].Actions[0].Resources[0].DB,
										Collection: customRoles[0].Actions[0].Resources[0].Collection,
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
	defer ctl.Finish()
	credsProvider := mocks.NewMockCredentialsGetter(ctl)
	t.Run("Can generate a valid connection secret WITH data", func(t *testing.T) {
		publicAPIKey := "TestPublicKey"
		privateAPIKey := "TestPrivateKey"

		name := "TestSecret-1"
		namespace := "TestNamespace-1"

		credsProvider.EXPECT().PublicAPIKey().Return(publicAPIKey)
		credsProvider.EXPECT().PrivateAPIKey().Return(privateAPIKey)

		got := BuildProjectConnectionSecret(credsProvider, name, namespace,
			orgID, true)

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
			orgID, false)

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
	defer ctl.Finish()
	alProvider := mocks.NewMockProjectIPAccessListLister(ctl)
	t.Run("Can convert Access Lists", func(t *testing.T) {
		data := &mongodbatlas.ProjectIPAccessLists{
			Links: nil,
			Results: []mongodbatlas.ProjectIPAccessList{
				{
					AwsSecurityGroup: "TestSecGroup",
					CIDRBlock:        "0.0.0.0/0",
					Comment:          "TestComment",
					DeleteAfterDate:  "TestDate",
					GroupID:          "TestGroupID",
					IPAddress:        "0.0.0.0",
				},
			},
			TotalCount: 1,
		}

		listOptions := &mongodbatlas.ListOptions{ItemsPerPage: MaxItems}

		alProvider.EXPECT().ProjectIPAccessLists(projectID, listOptions).Return(data, nil)

		got, err := buildAccessLists(alProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []project.IPAccessList{
			{
				AwsSecurityGroup: data.Results[0].AwsSecurityGroup,
				CIDRBlock:        data.Results[0].CIDRBlock,
				Comment:          data.Results[0].Comment,
				DeleteAfterDate:  data.Results[0].DeleteAfterDate,
				IPAddress:        data.Results[0].IPAddress,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("IPAccessList mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildAuditing(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	auditingProvider := mocks.NewMockAuditingDescriber(ctl)
	t.Run("Can convert Auditing", func(t *testing.T) {
		data := &mongodbatlas.Auditing{
			AuditAuthorizationSuccess: pointers.MakePtr(true),
			AuditFilter:               "TestFilter",
			ConfigurationType:         "TestType",
			Enabled:                   pointers.MakePtr(true),
		}

		auditingProvider.EXPECT().Auditing(projectID).Return(data, nil)

		got, err := buildAuditing(auditingProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.Auditing{
			AuditAuthorizationSuccess: data.AuditAuthorizationSuccess,
			AuditFilter:               data.AuditFilter,
			Enabled:                   data.Enabled,
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Auditing mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildCloudProviderAccessRoles(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	cpaProvider := mocks.NewMockCloudProviderAccessRoleLister(ctl)
	t.Run("Can convert CPA roles", func(t *testing.T) {
		data := &mongodbatlas.CloudProviderAccessRoles{
			AWSIAMRoles: []mongodbatlas.AWSIAMRole{
				{
					AtlasAWSAccountARN:         "TestARN",
					AtlasAssumedRoleExternalID: "TestRoleID",
					AuthorizedDate:             "TestAuthDate",
					CreatedDate:                "TestCreatedDate",
					FeatureUsages:              nil,
					IAMAssumedRoleARN:          "TestAssumedRoleARN",
					ProviderName:               string(provider.ProviderAWS),
					RoleID:                     "TestRoleID",
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
				ProviderName:      data.AWSIAMRoles[0].ProviderName,
				IamAssumedRoleArn: data.AWSIAMRoles[0].IAMAssumedRoleARN,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("CPA mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildEncryptionAtREST(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	dataProvider := mocks.NewMockEncryptionAtRestDescriber(ctl)
	t.Run("Can convert Encryption at REST AWS", func(t *testing.T) {
		data := &mongodbatlas.EncryptionAtRest{
			GroupID: "TestGroupID",
			AwsKms: mongodbatlas.AwsKms{
				Enabled:             pointers.MakePtr(true),
				AccessKeyID:         "TestAccessKey",
				SecretAccessKey:     "TestSecretAccessKey",
				CustomerMasterKeyID: "TestCustomerMasterKeyID",
				Region:              "US_EAST_1",
				RoleID:              "TestRoleID",
				Valid:               pointers.MakePtr(true),
			},
			AzureKeyVault:  mongodbatlas.AzureKeyVault{},
			GoogleCloudKms: mongodbatlas.GoogleCloudKms{},
		}

		dataProvider.EXPECT().EncryptionAtRest(projectID).Return(data, nil)

		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms: atlasV1.AwsKms{
				Enabled:             data.AwsKms.Enabled,
				AccessKeyID:         data.AwsKms.AccessKeyID,
				SecretAccessKey:     data.AwsKms.SecretAccessKey,
				CustomerMasterKeyID: data.AwsKms.CustomerMasterKeyID,
				Region:              data.AwsKms.Region,
				RoleID:              data.AwsKms.RoleID,
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
		data := &mongodbatlas.EncryptionAtRest{
			GroupID:       "TestGroupID",
			AwsKms:        mongodbatlas.AwsKms{},
			AzureKeyVault: mongodbatlas.AzureKeyVault{},
			GoogleCloudKms: mongodbatlas.GoogleCloudKms{
				Enabled:              pointers.MakePtr(true),
				ServiceAccountKey:    "TestServiceAccountKey",
				KeyVersionResourceID: "TestVersionResourceID",
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
				ServiceAccountKey:    data.GoogleCloudKms.ServiceAccountKey,
				KeyVersionResourceID: data.GoogleCloudKms.KeyVersionResourceID,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can convert Encryption at REST Azure", func(t *testing.T) {
		data := mongodbatlas.EncryptionAtRest{
			GroupID: "TestGroupID",
			AwsKms:  mongodbatlas.AwsKms{},
			AzureKeyVault: mongodbatlas.AzureKeyVault{
				Enabled:           pointers.MakePtr(true),
				ClientID:          "TestClientID",
				AzureEnvironment:  "TestAzureEnv",
				SubscriptionID:    "TestSubID",
				ResourceGroupName: "TestResourceGroupName",
				KeyVaultName:      "TestKeyVaultName",
				KeyIdentifier:     "TestKeyIdentifier",
				Secret:            "TestSecret",
				TenantID:          "TestTenantID",
			},
			GoogleCloudKms: mongodbatlas.GoogleCloudKms{},
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
				ClientID:          data.AzureKeyVault.ClientID,
				AzureEnvironment:  data.AzureKeyVault.AzureEnvironment,
				SubscriptionID:    data.AzureKeyVault.SubscriptionID,
				ResourceGroupName: data.AzureKeyVault.ResourceGroupName,
				KeyVaultName:      data.AzureKeyVault.KeyVaultName,
				KeyIdentifier:     data.AzureKeyVault.KeyIdentifier,
				Secret:            data.AzureKeyVault.Secret,
				TenantID:          data.AzureKeyVault.TenantID,
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
	defer ctl.Finish()
	intProvider := mocks.NewMockIntegrationLister(ctl)

	t.Run("Can convert third-party integrations WITH secrets: Prometheus", func(t *testing.T) {
		const targetNamespace = "test-namespace-3"
		const includeSecrets = true
		ints := &mongodbatlas.ThirdPartyIntegrations{
			Links: nil,
			Results: []*mongodbatlas.ThirdPartyIntegration{
				{
					Type:             "PROMETHEUS",
					Password:         "PrometheusTestPassword",
					UserName:         "PrometheusTestUserName",
					ServiceDiscovery: "TestServiceDiscovery",
				},
			},
			TotalCount: 0,
		}

		intProvider.EXPECT().Integrations(projectID).Return(ints, nil)

		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []project.Integration{
			{
				Type:             ints.Results[0].Type,
				ServiceDiscovery: ints.Results[0].ServiceDiscovery,
				UserName:         ints.Results[0].UserName,
				PasswordRef: common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.Results[0].Type)),
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
						strings.ToLower(ints.Results[0].Type)),
					Namespace: targetNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{
					secrets.PasswordField: []byte(ints.Results[0].Password),
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
		ints := &mongodbatlas.ThirdPartyIntegrations{
			Links: nil,
			Results: []*mongodbatlas.ThirdPartyIntegration{
				{
					Type:             "PROMETHEUS",
					Password:         "PrometheusTestPassword",
					UserName:         "PrometheusTestUserName",
					ServiceDiscovery: "TestServiceDiscovery",
				},
			},
			TotalCount: 0,
		}
		intProvider.EXPECT().Integrations(projectID).Return(ints, nil)
		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []project.Integration{
			{
				Type:             ints.Results[0].Type,
				ServiceDiscovery: ints.Results[0].ServiceDiscovery,
				UserName:         ints.Results[0].UserName,
				PasswordRef: common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.Results[0].Type)),
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
						strings.ToLower(ints.Results[0].Type)),
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
	defer ctl.Finish()
	mwProvider := mocks.NewMockMaintenanceWindowDescriber(ctl)
	t.Run("Can convert maintenance window", func(t *testing.T) {
		mw := &mongodbatlas.MaintenanceWindow{
			DayOfWeek:            3,
			HourOfDay:            pointers.MakePtr(10),
			StartASAP:            pointers.MakePtr(false),
			NumberOfDeferrals:    0,
			AutoDeferOnceEnabled: pointers.MakePtr(false),
		}

		mwProvider.EXPECT().MaintenanceWindow(projectID).Return(mw, nil)

		got, err := buildMaintenanceWindows(mwProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := project.MaintenanceWindow{
			DayOfWeek: mw.DayOfWeek,
			HourOfDay: *mw.HourOfDay,
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
	defer ctl.Finish()
	peerProvider := mocks.NewMockPeeringConnectionLister(ctl)
	t.Run("Can convert Peering connections", func(t *testing.T) {
		peeringConnections := []mongodbatlas.Peer{

			{
				AccepterRegionName:  "TestRegionName",
				AWSAccountID:        "TestAWSAccountID",
				ConnectionID:        "TestConnID",
				ContainerID:         "TestContainerID",
				ErrorStateName:      "TestErrStateName",
				ID:                  "TestID",
				ProviderName:        string(provider.ProviderAWS),
				RouteTableCIDRBlock: "0.0.0.0/0",
				StatusName:          "TestStatusName",
				VpcID:               "TestVPCID",
				AtlasCIDRBlock:      "0.0.0.0/0",
				AzureDirectoryID:    "TestDir",
				AzureSubscriptionID: "TestSub",
				ResourceGroupName:   "TestResourceName",
				VNetName:            "TestNETName",
				ErrorState:          "TestErrState",
				Status:              "TestStatus",
				GCPProjectID:        "TestProjectID",
				NetworkName:         "TestNetworkName",
				ErrorMessage:        "TestErrMessage",
			},
		}
		listOptions := &mongodbatlas.ContainersListOptions{ListOptions: mongodbatlas.ListOptions{ItemsPerPage: MaxItems}}
		peerProvider.EXPECT().PeeringConnections(projectID, listOptions).Return(peeringConnections, nil)

		got, err := buildNetworkPeering(peerProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.NetworkPeer{
			{
				AccepterRegionName:  peeringConnections[0].AccepterRegionName,
				ContainerRegion:     "",
				AWSAccountID:        peeringConnections[0].AWSAccountID,
				ContainerID:         peeringConnections[0].ContainerID,
				ProviderName:        provider.ProviderName(peeringConnections[0].ProviderName),
				RouteTableCIDRBlock: peeringConnections[0].RouteTableCIDRBlock,
				VpcID:               peeringConnections[0].VpcID,
				AtlasCIDRBlock:      peeringConnections[0].AtlasCIDRBlock,
				AzureDirectoryID:    peeringConnections[0].AzureDirectoryID,
				AzureSubscriptionID: peeringConnections[0].AzureSubscriptionID,
				ResourceGroupName:   peeringConnections[0].ResourceGroupName,
				VNetName:            peeringConnections[0].VNetName,
				GCPProjectID:        peeringConnections[0].GCPProjectID,
				NetworkName:         peeringConnections[0].NetworkName,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("NetworkPeerings mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

func Test_buildPrivateEndpoints(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	peProvider := mocks.NewMockPrivateEndpointLister(ctl)
	t.Run("Can convert PrivateEndpointConnection for AWS", func(t *testing.T) {
		providerName := provider.ProviderAWS
		privateEndpoint := mongodbatlas.PrivateEndpointConnection{
			ID:                           "1",
			ProviderName:                 string(providerName),
			Region:                       "US_EAST_1",
			EndpointServiceName:          "",
			ErrorMessage:                 "",
			InterfaceEndpoints:           nil,
			PrivateEndpoints:             nil,
			PrivateLinkServiceName:       "",
			PrivateLinkServiceResourceID: "",
			Status:                       "",
			EndpointGroupNames:           nil,
			RegionName:                   "",
			ServiceAttachmentNames:       nil,
		}

		listOptions := &mongodbatlas.ListOptions{ItemsPerPage: MaxItems}
		peProvider.EXPECT().PrivateEndpoints(projectID, string(providerName), listOptions).Return([]mongodbatlas.PrivateEndpointConnection{privateEndpoint}, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAzure), listOptions).Return(nil, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderGCP), listOptions).Return(nil, nil)

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.PrivateEndpoint{
			{
				Provider:          providerName,
				Region:            privateEndpoint.Region,
				ID:                privateEndpoint.ID,
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
		privateEndpoint := mongodbatlas.PrivateEndpointConnection{

			ID:                           "1",
			ProviderName:                 string(providerName),
			Region:                       "uswest3",
			EndpointServiceName:          "",
			ErrorMessage:                 "",
			InterfaceEndpoints:           nil,
			PrivateEndpoints:             nil,
			PrivateLinkServiceName:       "",
			PrivateLinkServiceResourceID: "",
			Status:                       "",
			EndpointGroupNames:           nil,
			RegionName:                   "",
			ServiceAttachmentNames:       nil,
		}

		listOptions := &mongodbatlas.ListOptions{ItemsPerPage: MaxItems}
		peProvider.EXPECT().PrivateEndpoints(projectID, string(providerName), listOptions).Return([]mongodbatlas.PrivateEndpointConnection{privateEndpoint}, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderAWS), listOptions).Return(nil, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(provider.ProviderGCP), listOptions).Return(nil, nil)

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.PrivateEndpoint{
			{
				Provider:          providerName,
				Region:            privateEndpoint.Region,
				ID:                privateEndpoint.ID,
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
	defer ctl.Finish()
	settingsProvider := mocks.NewMockProjectSettingsDescriber(ctl)
	t.Run("Can convert project settings", func(t *testing.T) {
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: pointers.MakePtr(true),
			IsDataExplorerEnabled:                       pointers.MakePtr(true),
			IsPerformanceAdvisorEnabled:                 pointers.MakePtr(true),
			IsRealtimePerformancePanelEnabled:           pointers.MakePtr(true),
			IsSchemaAdvisorEnabled:                      pointers.MakePtr(true),
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
	defer ctl.Finish()
	rolesProvider := mocks.NewMockDatabaseRoleLister(ctl)
	t.Run("Can build custom roles", func(t *testing.T) {
		data := []mongodbatlas.CustomDBRole{
			{
				Actions: []mongodbatlas.Action{
					{
						Action: "TestAction",
						Resources: []mongodbatlas.Resource{
							{
								Collection: pointers.MakePtr("TestCollection"),
								DB:         pointers.MakePtr("TestDB"),
								Cluster:    pointers.MakePtr(true),
							},
						},
					},
				},
				InheritedRoles: []mongodbatlas.InheritedRole{
					{
						Db:   "TestDBMAIN",
						Role: "ADMIN",
					},
				},
				RoleName: "TestRoleName",
			},
		}

		listOptions := &mongodbatlas.ListOptions{ItemsPerPage: MaxItems}
		rolesProvider.EXPECT().DatabaseRoles(projectID, listOptions).Return(&data, nil)

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
								Cluster:    role.Actions[0].Resources[0].Cluster,
								Database:   role.Actions[0].Resources[0].DB,
								Collection: role.Actions[0].Resources[0].Collection,
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
