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

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/pointers"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/auth"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const orgID = "TestOrgID"
const projectID = "TestProjectID"
const teamID = "TestTeamID"

// nolint
type MockAtlasOperatorProjectStore struct {
	publicApiKey           string
	privateApiKey          string
	projects               map[string]*mongodbatlas.Project
	ipAccessLists          map[string]*mongodbatlas.ProjectIPAccessLists
	auditing               map[string]*mongodbatlas.Auditing
	cpas                   map[string]*mongodbatlas.CloudProviderAccessRoles
	encryptionAtRest       map[string]*mongodbatlas.EncryptionAtRest
	thirdPartyIntegrations map[string]*mongodbatlas.ThirdPartyIntegrations
	mw                     map[string]*mongodbatlas.MaintenanceWindow
	peeringConnections     map[string][]mongodbatlas.Peer
	privateEndpoints       map[string]map[string][]mongodbatlas.PrivateEndpointConnection
	projectSettings        map[string]*mongodbatlas.ProjectSettings
	alertConfigs           map[string][]mongodbatlas.AlertConfiguration
	customRoles            map[string]*[]mongodbatlas.CustomDBRole
	teams                  map[string]map[string]*mongodbatlas.Team
	projectTeams           map[string]*mongodbatlas.TeamsAssigned
	teamUsers              map[string]map[string][]mongodbatlas.AtlasUser
}

func (m *MockAtlasOperatorProjectStore) TeamUsers(orgID string, teamID string) (interface{}, error) {
	return m.teamUsers[orgID][teamID], nil
}

func (m *MockAtlasOperatorProjectStore) TeamByID(orgID, teamID string) (*mongodbatlas.Team, error) {
	return m.teams[orgID][teamID], nil
}

// nolint
func (m *MockAtlasOperatorProjectStore) TeamByName(_, _ string) (*mongodbatlas.Team, error) {
	return nil, fmt.Errorf("shoudn't be called")
}

func (m *MockAtlasOperatorProjectStore) ProjectTeams(projectID string) (interface{}, error) {
	return m.projectTeams[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) Project(projectID string) (interface{}, error) {
	return m.projects[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) ProjectIPAccessLists(projectID string, _ *mongodbatlas.ListOptions) (*mongodbatlas.ProjectIPAccessLists, error) {
	return m.ipAccessLists[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) ProjectSettings(projectID string) (*mongodbatlas.ProjectSettings, error) {
	return m.projectSettings[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) Integrations(projectID string) (*mongodbatlas.ThirdPartyIntegrations, error) {
	return m.thirdPartyIntegrations[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) MaintenanceWindow(projectID string) (*mongodbatlas.MaintenanceWindow, error) {
	return m.mw[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) PrivateEndpoints(projectID, providerName string, _ *mongodbatlas.ListOptions) ([]mongodbatlas.PrivateEndpointConnection, error) {
	return m.privateEndpoints[projectID][providerName], nil
}

func (m *MockAtlasOperatorProjectStore) CloudProviderAccessRoles(projectID string) (*mongodbatlas.CloudProviderAccessRoles, error) {
	return m.cpas[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) PeeringConnections(projectID string, _ *mongodbatlas.ContainersListOptions) ([]mongodbatlas.Peer, error) {
	return m.peeringConnections[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) EncryptionAtRest(projectID string) (*mongodbatlas.EncryptionAtRest, error) {
	return m.encryptionAtRest[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) Auditing(projectID string) (*mongodbatlas.Auditing, error) {
	return m.auditing[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) AlertConfigurations(projectID string, _ *mongodbatlas.ListOptions) ([]mongodbatlas.AlertConfiguration, error) {
	return m.alertConfigs[projectID], nil
}

func (m *MockAtlasOperatorProjectStore) DatabaseRoles(projectID string, _ *mongodbatlas.ListOptions) (*[]mongodbatlas.CustomDBRole, error) {
	return m.customRoles[projectID], nil
}

func TestBuildAtlasProject(t *testing.T) {
	t.Run("Can convert Project entity with secrets data", func(t *testing.T) {
		targetNamespace := "test-namespace"
		projectStore := &MockAtlasOperatorProjectStore{
			publicApiKey:  "TestPublicKey",
			privateApiKey: "TestPrivateKey",
			projects: map[string]*mongodbatlas.Project{
				projectID: {
					ID:                        projectID,
					OrgID:                     orgID,
					Name:                      "TestProjectName",
					ClusterCount:              0,
					Created:                   "",
					RegionUsageRestrictions:   "",
					Links:                     nil,
					WithDefaultAlertsSettings: pointers.MakePtr(false),
				},
			},
			ipAccessLists: map[string]*mongodbatlas.ProjectIPAccessLists{
				projectID: {
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
				},
			},
			auditing: map[string]*mongodbatlas.Auditing{
				projectID: {
					AuditAuthorizationSuccess: pointers.MakePtr(true),
					AuditFilter:               "TestFilter",
					ConfigurationType:         "TestConfigType",
					Enabled:                   pointers.MakePtr(true),
				},
			},
			cpas: map[string]*mongodbatlas.CloudProviderAccessRoles{
				projectID: {
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
				},
			},
			encryptionAtRest: map[string]*mongodbatlas.EncryptionAtRest{
				projectID: {
					GroupID:       "TestGroupID",
					AwsKms:        mongodbatlas.AwsKms{},
					AzureKeyVault: mongodbatlas.AzureKeyVault{},
					GoogleCloudKms: mongodbatlas.GoogleCloudKms{
						Enabled:              pointers.MakePtr(true),
						ServiceAccountKey:    "TestServiceAccountKey",
						KeyVersionResourceID: "TestKeyVersionResourceID",
					},
				},
			},
			thirdPartyIntegrations: map[string]*mongodbatlas.ThirdPartyIntegrations{
				projectID: {
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
				},
			},
			mw: map[string]*mongodbatlas.MaintenanceWindow{
				projectID: {
					DayOfWeek:            1,
					HourOfDay:            pointers.MakePtr(10),
					StartASAP:            pointers.MakePtr(false),
					NumberOfDeferrals:    0,
					AutoDeferOnceEnabled: pointers.MakePtr(false),
				},
			},
			peeringConnections: map[string][]mongodbatlas.Peer{
				projectID: {
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
				},
			},
			privateEndpoints: map[string]map[string][]mongodbatlas.PrivateEndpointConnection{
				projectID: {
					string(provider.ProviderAWS): {
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
					},
				},
			},
			alertConfigs: map[string][]mongodbatlas.AlertConfiguration{
				projectID: {
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
				},
			},
			projectSettings: map[string]*mongodbatlas.ProjectSettings{
				projectID: {
					IsCollectDatabaseSpecificsStatisticsEnabled: pointers.MakePtr(true),
					IsDataExplorerEnabled:                       pointers.MakePtr(true),
					IsPerformanceAdvisorEnabled:                 pointers.MakePtr(true),
					IsRealtimePerformancePanelEnabled:           pointers.MakePtr(true),
					IsSchemaAdvisorEnabled:                      pointers.MakePtr(true),
				},
			},
			customRoles: map[string]*[]mongodbatlas.CustomDBRole{
				projectID: {
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
				},
			},
			projectTeams: map[string]*mongodbatlas.TeamsAssigned{
				projectID: {
					Links: nil,
					Results: []*mongodbatlas.Result{
						{
							Links:     nil,
							TeamID:    teamID,
							RoleNames: []string{string(atlasV1.TeamRoleClusterManager)},
						},
					},
					TotalCount: 1,
				},
			},
			teams: map[string]map[string]*mongodbatlas.Team{
				orgID: {
					teamID: {
						ID:        teamID,
						Name:      "TestTeamName",
						Usernames: []string{},
					},
				},
			},
			teamUsers: map[string]map[string][]mongodbatlas.AtlasUser{
				orgID: {
					teamID: {
						{
							EmailAddress: "testuser@mooooongodb.com",
							FirstName:    "TestName",
							ID:           "TestID",
							LastName:     "TestLastName",
						},
					},
				},
			},
		}

		projectResult, err := BuildAtlasProject(projectStore, orgID, projectID, targetNamespace, true)
		if err != nil {
			t.Fatalf("%v", err)
		}
		gotProject := projectResult.Project
		gotTeams := projectResult.Teams

		expectedThreshold := &atlasV1.Threshold{
			Operator:  projectStore.alertConfigs[projectID][0].Threshold.Operator,
			Units:     projectStore.alertConfigs[projectID][0].Threshold.Units,
			Threshold: fmt.Sprintf("%f", projectStore.alertConfigs[projectID][0].Threshold.Threshold),
		}
		expectedMatchers := []atlasV1.Matcher{
			{
				FieldName: projectStore.alertConfigs[projectID][0].Matchers[0].FieldName,
				Operator:  projectStore.alertConfigs[projectID][0].Matchers[0].Operator,
				Value:     projectStore.alertConfigs[projectID][0].Matchers[0].Value,
			},
		}
		expectedNotifications := []atlasV1.Notification{
			{
				APIToken:            projectStore.alertConfigs[projectID][0].Notifications[0].APIToken,
				ChannelName:         projectStore.alertConfigs[projectID][0].Notifications[0].ChannelName,
				DatadogAPIKey:       projectStore.alertConfigs[projectID][0].Notifications[0].DatadogAPIKey,
				DatadogRegion:       projectStore.alertConfigs[projectID][0].Notifications[0].DatadogRegion,
				DelayMin:            projectStore.alertConfigs[projectID][0].Notifications[0].DelayMin,
				EmailAddress:        projectStore.alertConfigs[projectID][0].Notifications[0].EmailAddress,
				EmailEnabled:        projectStore.alertConfigs[projectID][0].Notifications[0].EmailEnabled,
				FlowdockAPIToken:    projectStore.alertConfigs[projectID][0].Notifications[0].FlowdockAPIToken,
				FlowName:            projectStore.alertConfigs[projectID][0].Notifications[0].FlowName,
				IntervalMin:         projectStore.alertConfigs[projectID][0].Notifications[0].IntervalMin,
				MobileNumber:        projectStore.alertConfigs[projectID][0].Notifications[0].MobileNumber,
				OpsGenieAPIKey:      projectStore.alertConfigs[projectID][0].Notifications[0].OpsGenieAPIKey,
				OpsGenieRegion:      projectStore.alertConfigs[projectID][0].Notifications[0].OpsGenieRegion,
				OrgName:             projectStore.alertConfigs[projectID][0].Notifications[0].OrgName,
				ServiceKey:          projectStore.alertConfigs[projectID][0].Notifications[0].ServiceKey,
				SMSEnabled:          projectStore.alertConfigs[projectID][0].Notifications[0].SMSEnabled,
				TeamID:              projectStore.alertConfigs[projectID][0].Notifications[0].TeamID,
				TeamName:            projectStore.alertConfigs[projectID][0].Notifications[0].TeamName,
				TypeName:            projectStore.alertConfigs[projectID][0].Notifications[0].TypeName,
				Username:            projectStore.alertConfigs[projectID][0].Notifications[0].Username,
				VictorOpsAPIKey:     projectStore.alertConfigs[projectID][0].Notifications[0].VictorOpsAPIKey,
				VictorOpsRoutingKey: projectStore.alertConfigs[projectID][0].Notifications[0].VictorOpsRoutingKey,
				Roles:               projectStore.alertConfigs[projectID][0].Notifications[0].Roles,
			},
		}
		expectedMetricThreshold := &atlasV1.MetricThreshold{
			MetricName: projectStore.alertConfigs[projectID][0].MetricThreshold.MetricName,
			Operator:   projectStore.alertConfigs[projectID][0].MetricThreshold.Operator,
			Threshold:  fmt.Sprintf("%f", projectStore.alertConfigs[projectID][0].MetricThreshold.Threshold),
			Units:      projectStore.alertConfigs[projectID][0].MetricThreshold.Units,
			Mode:       projectStore.alertConfigs[projectID][0].MetricThreshold.Mode,
		}
		expectedTeams := []*atlasV1.AtlasTeam{
			{
				TypeMeta: v1.TypeMeta{
					Kind:       "AtlasTeam",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      fmt.Sprintf("%s-team-%s", strings.ToLower(projectStore.projects[projectID].Name), strings.ToLower(projectStore.teams[orgID][teamID].Name)),
					Namespace: targetNamespace,
				},
				Spec: atlasV1.TeamSpec{
					Name:      projectStore.teams[orgID][teamID].Name,
					Usernames: []atlasV1.TeamUser{atlasV1.TeamUser(projectStore.teamUsers[orgID][teamID][0].Username)},
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
				Name:      strings.ToLower(projectStore.projects[projectID].Name),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasProjectSpec{
				Name: projectStore.projects[projectID].Name,
				ConnectionSecret: &common.ResourceRef{
					Name: fmt.Sprintf(credSecretFormat, projectStore.projects[projectID].Name),
				},
				ProjectIPAccessList: []project.IPAccessList{
					{
						AwsSecurityGroup: projectStore.ipAccessLists[projectID].Results[0].AwsSecurityGroup,
						CIDRBlock:        projectStore.ipAccessLists[projectID].Results[0].CIDRBlock,
						Comment:          projectStore.ipAccessLists[projectID].Results[0].Comment,
						DeleteAfterDate:  projectStore.ipAccessLists[projectID].Results[0].DeleteAfterDate,
						IPAddress:        projectStore.ipAccessLists[projectID].Results[0].IPAddress,
					},
				},
				MaintenanceWindow: project.MaintenanceWindow{
					DayOfWeek: projectStore.mw[projectID].DayOfWeek,
					HourOfDay: pointers.PtrValOrDefault(projectStore.mw[projectID].HourOfDay, 0),
					AutoDefer: pointers.PtrValOrDefault(projectStore.mw[projectID].AutoDeferOnceEnabled, false),
					StartASAP: pointers.PtrValOrDefault(projectStore.mw[projectID].StartASAP, false),
					Defer:     false,
				},
				PrivateEndpoints: []atlasV1.PrivateEndpoint{
					{
						Provider:          provider.ProviderAWS,
						Region:            projectStore.privateEndpoints[projectID]["AWS"][0].Region,
						ID:                projectStore.privateEndpoints[projectID]["AWS"][0].ID,
						IP:                "",
						GCPProjectID:      "",
						EndpointGroupName: "",
						Endpoints:         atlasV1.GCPEndpoints{},
					},
				},
				CloudProviderAccessRoles: []atlasV1.CloudProviderAccessRole{
					{
						ProviderName:      projectStore.cpas[projectID].AWSIAMRoles[0].ProviderName,
						IamAssumedRoleArn: projectStore.cpas[projectID].AWSIAMRoles[0].IAMAssumedRoleARN,
					},
				},
				AlertConfigurations: []atlasV1.AlertConfiguration{
					{
						Enabled:         *projectStore.alertConfigs[projectID][0].Enabled,
						EventTypeName:   projectStore.alertConfigs[projectID][0].EventTypeName,
						Matchers:        expectedMatchers,
						Threshold:       expectedThreshold,
						Notifications:   expectedNotifications,
						MetricThreshold: expectedMetricThreshold,
					},
				},
				AlertConfigurationSyncEnabled: false,
				NetworkPeers: []atlasV1.NetworkPeer{
					{
						AccepterRegionName:  projectStore.peeringConnections[projectID][0].AccepterRegionName,
						ContainerRegion:     "",
						AWSAccountID:        projectStore.peeringConnections[projectID][0].AWSAccountID,
						ContainerID:         projectStore.peeringConnections[projectID][0].ContainerID,
						ProviderName:        provider.ProviderName(projectStore.peeringConnections[projectID][0].ProviderName),
						RouteTableCIDRBlock: projectStore.peeringConnections[projectID][0].RouteTableCIDRBlock,
						VpcID:               projectStore.peeringConnections[projectID][0].VpcID,
						AtlasCIDRBlock:      projectStore.peeringConnections[projectID][0].AtlasCIDRBlock,
						AzureDirectoryID:    projectStore.peeringConnections[projectID][0].AzureDirectoryID,
						AzureSubscriptionID: projectStore.peeringConnections[projectID][0].AzureSubscriptionID,
						ResourceGroupName:   projectStore.peeringConnections[projectID][0].ResourceGroupName,
						VNetName:            projectStore.peeringConnections[projectID][0].VNetName,
						GCPProjectID:        projectStore.peeringConnections[projectID][0].GCPProjectID,
						NetworkName:         projectStore.peeringConnections[projectID][0].NetworkName,
					},
				},
				WithDefaultAlertsSettings: false,
				X509CertRef:               nil,
				Integrations: []project.Integration{
					{
						Type:     projectStore.thirdPartyIntegrations[projectID].Results[0].Type,
						UserName: projectStore.thirdPartyIntegrations[projectID].Results[0].UserName,
						PasswordRef: common.ResourceRefNamespaced{
							Name: fmt.Sprintf("%s-integration-%s",
								strings.ToLower(projectID),
								strings.ToLower(projectStore.thirdPartyIntegrations[projectID].Results[0].Type)),
							Namespace: targetNamespace,
						},
						ServiceDiscovery: projectStore.thirdPartyIntegrations[projectID].Results[0].ServiceDiscovery,
					},
				},
				EncryptionAtRest: &atlasV1.EncryptionAtRest{
					AwsKms:        atlasV1.AwsKms{},
					AzureKeyVault: atlasV1.AzureKeyVault{},
					GoogleCloudKms: atlasV1.GoogleCloudKms{
						Enabled:              projectStore.encryptionAtRest[projectID].GoogleCloudKms.Enabled,
						ServiceAccountKey:    projectStore.encryptionAtRest[projectID].GoogleCloudKms.ServiceAccountKey,
						KeyVersionResourceID: projectStore.encryptionAtRest[projectID].GoogleCloudKms.KeyVersionResourceID,
					},
				},
				Auditing: &atlasV1.Auditing{
					AuditAuthorizationSuccess: projectStore.auditing[projectID].AuditAuthorizationSuccess,
					AuditFilter:               projectStore.auditing[projectID].AuditFilter,
					Enabled:                   projectStore.auditing[projectID].Enabled,
				},
				Settings: &atlasV1.ProjectSettings{
					IsCollectDatabaseSpecificsStatisticsEnabled: projectStore.projectSettings[projectID].IsCollectDatabaseSpecificsStatisticsEnabled,
					IsDataExplorerEnabled:                       projectStore.projectSettings[projectID].IsDataExplorerEnabled,
					IsPerformanceAdvisorEnabled:                 projectStore.projectSettings[projectID].IsPerformanceAdvisorEnabled,
					IsRealtimePerformancePanelEnabled:           projectStore.projectSettings[projectID].IsRealtimePerformancePanelEnabled,
					IsSchemaAdvisorEnabled:                      projectStore.projectSettings[projectID].IsSchemaAdvisorEnabled,
				},
				CustomRoles: []atlasV1.CustomRole{
					{
						Name: (*projectStore.customRoles[projectID])[0].RoleName,
						InheritedRoles: []atlasV1.Role{
							{
								Name:     (*projectStore.customRoles[projectID])[0].InheritedRoles[0].Role,
								Database: (*projectStore.customRoles[projectID])[0].InheritedRoles[0].Db,
							},
						},
						Actions: []atlasV1.Action{
							{
								Name: (*projectStore.customRoles[projectID])[0].Actions[0].Action,
								Resources: []atlasV1.Resource{
									{
										Cluster:    (*projectStore.customRoles[projectID])[0].Actions[0].Resources[0].Cluster,
										Database:   (*projectStore.customRoles[projectID])[0].Actions[0].Resources[0].DB,
										Collection: (*projectStore.customRoles[projectID])[0].Actions[0].Resources[0].Collection,
									},
								},
							},
						},
					},
				},
				Teams: []atlasV1.Team{
					{
						TeamRef: common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-team-%s", strings.ToLower(projectStore.projects[projectID].Name), strings.ToLower(projectStore.teams[orgID][teamID].Name)),
							Namespace: targetNamespace,
						},
						Roles: []atlasV1.TeamRole{atlasV1.TeamRole(projectStore.projectTeams[projectID].Results[0].RoleNames[0])},
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

type MockCredsStore struct {
	publicAPIKey  string
	privateAPIKey string
}

func (m *MockCredsStore) PublicAPIKey() string {
	return m.publicAPIKey
}

func (m *MockCredsStore) PrivateAPIKey() string {
	return m.privateAPIKey
}

// nolint
func (m *MockCredsStore) Token() (*auth.Token, error) {
	return nil, nil
}

func TestBuildProjectConnectionSecret(t *testing.T) {
	t.Run("Can generate a valid connection secret WITH data", func(t *testing.T) {
		credsProvider := &MockCredsStore{
			publicAPIKey:  "TestPublicKey",
			privateAPIKey: "TestPrivateKey",
		}
		name := "TestSecret-1"
		namespace := "TestNamespace-1"

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
				secrets.CredPublicAPIKey:  []byte(credsProvider.publicAPIKey),
				secrets.CredPrivateAPIKey: []byte(credsProvider.privateAPIKey),
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Credentials secret mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can generate a valid connection secret WITHOUT data", func(t *testing.T) {
		credsProvider := &MockCredsStore{
			publicAPIKey:  "TestPublicKey",
			privateAPIKey: "TestPrivateKey",
		}
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

type MockAccessListStore struct {
	data map[string]*mongodbatlas.ProjectIPAccessLists
}

func (m *MockAccessListStore) ProjectIPAccessLists(projectID string, _ *mongodbatlas.ListOptions) (*mongodbatlas.ProjectIPAccessLists, error) {
	return m.data[projectID], nil
}

func Test_buildAccessLists(t *testing.T) {
	t.Run("Can convert Access Lists", func(t *testing.T) {
		alProvider := &MockAccessListStore{
			data: map[string]*mongodbatlas.ProjectIPAccessLists{
				projectID: {
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
				},
			},
		}

		got, err := buildAccessLists(alProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []project.IPAccessList{
			{
				AwsSecurityGroup: alProvider.data[projectID].Results[0].AwsSecurityGroup,
				CIDRBlock:        alProvider.data[projectID].Results[0].CIDRBlock,
				Comment:          alProvider.data[projectID].Results[0].Comment,
				DeleteAfterDate:  alProvider.data[projectID].Results[0].DeleteAfterDate,
				IPAddress:        alProvider.data[projectID].Results[0].IPAddress,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("IPAccessList mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockAuditingStore struct {
	data map[string]*mongodbatlas.Auditing
}

func (m *MockAuditingStore) Auditing(projectID string) (*mongodbatlas.Auditing, error) {
	return m.data[projectID], nil
}

func Test_buildAuditing(t *testing.T) {
	t.Run("Can convert Auditing", func(t *testing.T) {
		auditingProvider := &MockAuditingStore{
			data: map[string]*mongodbatlas.Auditing{
				projectID: {
					AuditAuthorizationSuccess: pointers.MakePtr(true),
					AuditFilter:               "TestFilter",
					ConfigurationType:         "TestType",
					Enabled:                   pointers.MakePtr(true),
				},
			},
		}

		got, err := buildAuditing(auditingProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.Auditing{
			AuditAuthorizationSuccess: auditingProvider.data[projectID].AuditAuthorizationSuccess,
			AuditFilter:               auditingProvider.data[projectID].AuditFilter,
			Enabled:                   auditingProvider.data[projectID].Enabled,
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Auditing mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockCPAStore struct {
	data map[string]*mongodbatlas.CloudProviderAccessRoles
}

func (m *MockCPAStore) CloudProviderAccessRoles(projectID string) (*mongodbatlas.CloudProviderAccessRoles, error) {
	return m.data[projectID], nil
}

func Test_buildCloudProviderAccessRoles(t *testing.T) {
	t.Run("Can convert CPA roles", func(t *testing.T) {
		cpaProvider := &MockCPAStore{
			data: map[string]*mongodbatlas.CloudProviderAccessRoles{
				projectID: {
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
				},
			},
		}

		got, err := buildCloudProviderAccessRoles(cpaProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []atlasV1.CloudProviderAccessRole{
			{
				ProviderName:      cpaProvider.data[projectID].AWSIAMRoles[0].ProviderName,
				IamAssumedRoleArn: cpaProvider.data[projectID].AWSIAMRoles[0].IAMAssumedRoleARN,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("CPA mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockEncryptionAtRestStore struct {
	data map[string]*mongodbatlas.EncryptionAtRest
}

func (m *MockEncryptionAtRestStore) EncryptionAtRest(projectID string) (*mongodbatlas.EncryptionAtRest, error) {
	return m.data[projectID], nil
}

func Test_buildEncryptionAtREST(t *testing.T) {
	t.Run("Can convert Encryption at REST AWS", func(t *testing.T) {
		dataProvider := &MockEncryptionAtRestStore{
			data: map[string]*mongodbatlas.EncryptionAtRest{
				projectID: {
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
				},
			},
		}

		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms: atlasV1.AwsKms{
				Enabled:             dataProvider.data[projectID].AwsKms.Enabled,
				AccessKeyID:         dataProvider.data[projectID].AwsKms.AccessKeyID,
				SecretAccessKey:     dataProvider.data[projectID].AwsKms.SecretAccessKey,
				CustomerMasterKeyID: dataProvider.data[projectID].AwsKms.CustomerMasterKeyID,
				Region:              dataProvider.data[projectID].AwsKms.Region,
				RoleID:              dataProvider.data[projectID].AwsKms.RoleID,
				Valid:               dataProvider.data[projectID].AwsKms.Valid,
			},
			AzureKeyVault:  atlasV1.AzureKeyVault{},
			GoogleCloudKms: atlasV1.GoogleCloudKms{},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can convert Encryption at REST GCP", func(t *testing.T) {
		dataProvider := &MockEncryptionAtRestStore{
			data: map[string]*mongodbatlas.EncryptionAtRest{
				projectID: {
					GroupID:       "TestGroupID",
					AwsKms:        mongodbatlas.AwsKms{},
					AzureKeyVault: mongodbatlas.AzureKeyVault{},
					GoogleCloudKms: mongodbatlas.GoogleCloudKms{
						Enabled:              pointers.MakePtr(true),
						ServiceAccountKey:    "TestServiceAccountKey",
						KeyVersionResourceID: "TestVersionResourceID",
					},
				},
			},
		}

		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms:        atlasV1.AwsKms{},
			AzureKeyVault: atlasV1.AzureKeyVault{},
			GoogleCloudKms: atlasV1.GoogleCloudKms{
				Enabled:              dataProvider.data[projectID].GoogleCloudKms.Enabled,
				ServiceAccountKey:    dataProvider.data[projectID].GoogleCloudKms.ServiceAccountKey,
				KeyVersionResourceID: dataProvider.data[projectID].GoogleCloudKms.KeyVersionResourceID,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
	t.Run("Can convert Encryption at REST Azure", func(t *testing.T) {
		dataProvider := &MockEncryptionAtRestStore{
			data: map[string]*mongodbatlas.EncryptionAtRest{
				projectID: {
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
				},
			},
		}

		got, err := buildEncryptionAtRest(dataProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &atlasV1.EncryptionAtRest{
			AwsKms: atlasV1.AwsKms{},
			AzureKeyVault: atlasV1.AzureKeyVault{
				Enabled:           dataProvider.data[projectID].AzureKeyVault.Enabled,
				ClientID:          dataProvider.data[projectID].AzureKeyVault.ClientID,
				AzureEnvironment:  dataProvider.data[projectID].AzureKeyVault.AzureEnvironment,
				SubscriptionID:    dataProvider.data[projectID].AzureKeyVault.SubscriptionID,
				ResourceGroupName: dataProvider.data[projectID].AzureKeyVault.ResourceGroupName,
				KeyVaultName:      dataProvider.data[projectID].AzureKeyVault.KeyVaultName,
				KeyIdentifier:     dataProvider.data[projectID].AzureKeyVault.KeyIdentifier,
				Secret:            dataProvider.data[projectID].AzureKeyVault.Secret,
				TenantID:          dataProvider.data[projectID].AzureKeyVault.TenantID,
			},
			GoogleCloudKms: atlasV1.GoogleCloudKms{},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("EncryptionAtREST mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockIntegrationsStore struct {
	ints map[string]*mongodbatlas.ThirdPartyIntegrations
}

func (m *MockIntegrationsStore) Integrations(projectID string) (*mongodbatlas.ThirdPartyIntegrations, error) {
	return m.ints[projectID], nil
}

func Test_buildIntegrations(t *testing.T) {
	t.Run("Can convert third-party integrations WITH secrets: Prometheus", func(t *testing.T) {
		const targetNamespace = "test-namespace-3"
		const includeSecrets = true
		intProvider := &MockIntegrationsStore{ints: map[string]*mongodbatlas.ThirdPartyIntegrations{
			projectID: {
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
			},
		}}

		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []project.Integration{
			{
				Type:             intProvider.ints[projectID].Results[0].Type,
				ServiceDiscovery: intProvider.ints[projectID].Results[0].ServiceDiscovery,
				UserName:         intProvider.ints[projectID].Results[0].UserName,
				PasswordRef: common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(intProvider.ints[projectID].Results[0].Type)),
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
						strings.ToLower(intProvider.ints[projectID].Results[0].Type)),
					Namespace: targetNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{
					secrets.PasswordField: []byte(intProvider.ints[projectID].Results[0].Password),
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
		intProvider := &MockIntegrationsStore{ints: map[string]*mongodbatlas.ThirdPartyIntegrations{
			projectID: {
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
			},
		}}

		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []project.Integration{
			{
				Type:             intProvider.ints[projectID].Results[0].Type,
				ServiceDiscovery: intProvider.ints[projectID].Results[0].ServiceDiscovery,
				UserName:         intProvider.ints[projectID].Results[0].UserName,
				PasswordRef: common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(intProvider.ints[projectID].Results[0].Type)),
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
						strings.ToLower(intProvider.ints[projectID].Results[0].Type)),
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

type MockMaintenanceWindowStore struct {
	mw map[string]*mongodbatlas.MaintenanceWindow
}

func (m *MockMaintenanceWindowStore) MaintenanceWindow(projectID string) (*mongodbatlas.MaintenanceWindow, error) {
	return m.mw[projectID], nil
}

func Test_buildMaintenanceWindows(t *testing.T) {
	t.Run("Can convert maintenance window", func(t *testing.T) {
		mwProvider := &MockMaintenanceWindowStore{
			mw: map[string]*mongodbatlas.MaintenanceWindow{
				projectID: {
					DayOfWeek:            3,
					HourOfDay:            pointers.MakePtr(10),
					StartASAP:            pointers.MakePtr(false),
					NumberOfDeferrals:    0,
					AutoDeferOnceEnabled: pointers.MakePtr(false),
				},
			},
		}

		got, err := buildMaintenanceWindows(mwProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := project.MaintenanceWindow{
			DayOfWeek: mwProvider.mw[projectID].DayOfWeek,
			HourOfDay: *mwProvider.mw[projectID].HourOfDay,
			AutoDefer: *mwProvider.mw[projectID].AutoDeferOnceEnabled,
			StartASAP: *mwProvider.mw[projectID].StartASAP,
			Defer:     false,
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("MaintenanceWindows mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockNetworkPeeringStore struct {
	peeringConnections map[string][]mongodbatlas.Peer
}

func (m *MockNetworkPeeringStore) PeeringConnections(projectID string, _ *mongodbatlas.ContainersListOptions) ([]mongodbatlas.Peer, error) {
	return m.peeringConnections[projectID], nil
}

func Test_buildNetworkPeering(t *testing.T) {
	t.Run("Can convert Peering connections", func(t *testing.T) {
		peerProvider := &MockNetworkPeeringStore{
			peeringConnections: map[string][]mongodbatlas.Peer{
				projectID: {
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
				},
			},
		}

		got, err := buildNetworkPeering(peerProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.NetworkPeer{
			{
				AccepterRegionName:  peerProvider.peeringConnections[projectID][0].AccepterRegionName,
				ContainerRegion:     "",
				AWSAccountID:        peerProvider.peeringConnections[projectID][0].AWSAccountID,
				ContainerID:         peerProvider.peeringConnections[projectID][0].ContainerID,
				ProviderName:        provider.ProviderName(peerProvider.peeringConnections[projectID][0].ProviderName),
				RouteTableCIDRBlock: peerProvider.peeringConnections[projectID][0].RouteTableCIDRBlock,
				VpcID:               peerProvider.peeringConnections[projectID][0].VpcID,
				AtlasCIDRBlock:      peerProvider.peeringConnections[projectID][0].AtlasCIDRBlock,
				AzureDirectoryID:    peerProvider.peeringConnections[projectID][0].AzureDirectoryID,
				AzureSubscriptionID: peerProvider.peeringConnections[projectID][0].AzureSubscriptionID,
				ResourceGroupName:   peerProvider.peeringConnections[projectID][0].ResourceGroupName,
				VNetName:            peerProvider.peeringConnections[projectID][0].VNetName,
				GCPProjectID:        peerProvider.peeringConnections[projectID][0].GCPProjectID,
				NetworkName:         peerProvider.peeringConnections[projectID][0].NetworkName,
			},
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("NetworkPeerings mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockPrivateEndpointStore struct {
	// Project ID -> Provider name -> [] Private endpoints
	privateEndpoints map[string]map[string][]mongodbatlas.PrivateEndpointConnection
}

func (m *MockPrivateEndpointStore) PrivateEndpoints(projectID, providerName string, _ *mongodbatlas.ListOptions) ([]mongodbatlas.PrivateEndpointConnection, error) {
	return m.privateEndpoints[projectID][providerName], nil
}

func Test_buildPrivateEndpoints(t *testing.T) {
	t.Run("Can convert PrivateEndpointConnection for AWS", func(t *testing.T) {
		providerName := string(provider.ProviderAWS)
		peProvider := &MockPrivateEndpointStore{
			privateEndpoints: map[string]map[string][]mongodbatlas.PrivateEndpointConnection{
				projectID: {
					providerName: {
						{
							ID:                           "1",
							ProviderName:                 providerName,
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
						},
					},
				},
			},
		}

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.PrivateEndpoint{
			{
				Provider:          provider.ProviderName(peProvider.privateEndpoints[projectID][providerName][0].ProviderName),
				Region:            peProvider.privateEndpoints[projectID][providerName][0].Region,
				ID:                peProvider.privateEndpoints[projectID][providerName][0].ID,
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
		providerName := string(provider.ProviderAzure)
		peProvider := &MockPrivateEndpointStore{
			privateEndpoints: map[string]map[string][]mongodbatlas.PrivateEndpointConnection{
				projectID: {
					providerName: {
						{
							ID:                           "1",
							ProviderName:                 providerName,
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
						},
					},
				},
			},
		}

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []atlasV1.PrivateEndpoint{
			{
				Provider:          provider.ProviderName(peProvider.privateEndpoints[projectID][providerName][0].ProviderName),
				Region:            peProvider.privateEndpoints[projectID][providerName][0].Region,
				ID:                peProvider.privateEndpoints[projectID][providerName][0].ID,
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

type MockProjectSettingsStore struct {
	projectSettings map[string]*mongodbatlas.ProjectSettings
}

func (m *MockProjectSettingsStore) ProjectSettings(projectID string) (*mongodbatlas.ProjectSettings, error) {
	return m.projectSettings[projectID], nil
}

func Test_buildProjectSettings(t *testing.T) {
	t.Run("Can convert project settings", func(t *testing.T) {
		settingsProvider := &MockProjectSettingsStore{
			projectSettings: map[string]*mongodbatlas.ProjectSettings{
				projectID: {
					IsCollectDatabaseSpecificsStatisticsEnabled: pointers.MakePtr(true),
					IsDataExplorerEnabled:                       pointers.MakePtr(true),
					IsPerformanceAdvisorEnabled:                 pointers.MakePtr(true),
					IsRealtimePerformancePanelEnabled:           pointers.MakePtr(true),
					IsSchemaAdvisorEnabled:                      pointers.MakePtr(true),
				},
			},
		}

		got, err := buildProjectSettings(settingsProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expected := &atlasV1.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: settingsProvider.projectSettings[projectID].IsCollectDatabaseSpecificsStatisticsEnabled,
			IsDataExplorerEnabled:                       settingsProvider.projectSettings[projectID].IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 settingsProvider.projectSettings[projectID].IsPerformanceAdvisorEnabled,
			IsRealtimePerformancePanelEnabled:           settingsProvider.projectSettings[projectID].IsRealtimePerformancePanelEnabled,
			IsSchemaAdvisorEnabled:                      settingsProvider.projectSettings[projectID].IsSchemaAdvisorEnabled,
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Project settings mismatch. expected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}

type MockCustomRolesStore struct {
	data map[string]*[]mongodbatlas.CustomDBRole
}

func (m *MockCustomRolesStore) DatabaseRoles(projectID string, _ *mongodbatlas.ListOptions) (*[]mongodbatlas.CustomDBRole, error) {
	return m.data[projectID], nil
}

func Test_buildCustomRoles(t *testing.T) {
	t.Run("Can build custom roles", func(t *testing.T) {
		rolesProvider := &MockCustomRolesStore{
			data: map[string]*[]mongodbatlas.CustomDBRole{
				projectID: {
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
				},
			},
		}

		role := (*rolesProvider.data[projectID])[0]
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
