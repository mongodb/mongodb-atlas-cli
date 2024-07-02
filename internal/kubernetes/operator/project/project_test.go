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
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2project "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/project"
	akov2provider "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/provider"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	orgID           = "TestOrgID"
	projectID       = "TestProjectID"
	teamID          = "TestTeamID"
	resourceVersion = "x.y.z"
	targetNamespace = "test-namespace"
)

func TestBuildAtlasProject(t *testing.T) {
	ctl := gomock.NewController(t)
	projectStore := mocks.NewMockOperatorProjectStore(ctl)
	featureValidator := mocks.NewMockFeatureValidator(ctl)
	t.Run("Can convert Project entity with secrets data", func(t *testing.T) {
		p := &atlasv2.Group{
			Id:                        pointer.Get(projectID),
			OrgId:                     orgID,
			Name:                      "TestProjectName",
			ClusterCount:              0,
			WithDefaultAlertsSettings: pointer.Get(false),
		}

		ipAccessLists := &atlasv2.PaginatedNetworkAccess{
			Links: nil,
			Results: &[]atlasv2.NetworkPermissionEntry{
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
			AuditAuthorizationSuccess: pointer.Get(true),
			AuditFilter:               pointer.Get("TestFilter"),
			ConfigurationType:         pointer.Get("TestConfigType"),
			Enabled:                   pointer.Get(true),
		}

		authDate, _ := time.Parse(time.RFC3339, "01-01-2001")
		createDate, _ := time.Parse(time.RFC3339, "01-02-2001")

		cpas := &atlasv2.CloudProviderAccessRoles{
			AwsIamRoles: &[]atlasv2.CloudProviderAccessAWSIAMRole{
				{
					AtlasAWSAccountArn:         pointer.Get("TestARN"),
					AtlasAssumedRoleExternalId: pointer.Get("TestExternalRoleID"),
					AuthorizedDate:             &authDate,
					CreatedDate:                &createDate,
					FeatureUsages:              nil,
					IamAssumedRoleArn:          pointer.Get("TestRoleARN"),
					ProviderName:               string(akov2provider.ProviderAWS),
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
			Results: &[]atlasv2.ThirdPartyIntegration{
				{
					Type:             pointer.Get("PROMETHEUS"),
					Username:         pointer.Get("TestPrometheusUserName"),
					Password:         pointer.Get("TestPrometheusPassword"),
					ServiceDiscovery: pointer.Get("TestPrometheusServiceDiscovery"),
				},
			},
			TotalCount: pointer.Get(1),
		}

		mw := &atlasv2.GroupMaintenanceWindow{
			DayOfWeek:            1,
			HourOfDay:            pointer.Get(10),
			StartASAP:            pointer.Get(false),
			AutoDeferOnceEnabled: pointer.Get(false),
		}

		peeringConnectionAWS := &atlasv2.BaseNetworkPeeringConnectionSettings{
			AccepterRegionName:  pointer.Get("TestRegionName"),
			AwsAccountId:        pointer.Get("TestAWSAccountID"),
			ConnectionId:        pointer.Get("TestConnID"),
			ContainerId:         "TestContainerID",
			ErrorStateName:      pointer.Get("TestErrStateName"),
			Id:                  pointer.Get("TestID"),
			ProviderName:        pointer.Get(string(akov2provider.ProviderAWS)),
			RouteTableCidrBlock: pointer.Get("0.0.0.0/0"),
			StatusName:          pointer.Get("TestStatusName"),
			VpcId:               pointer.Get("TestVPCID"),
		}

		peeringConnections := []atlasv2.BaseNetworkPeeringConnectionSettings{*peeringConnectionAWS}

		privateAWSEndpoint := atlasv2.EndpointService{
			Id:                  pointer.Get("TestID"),
			CloudProvider:       string(akov2provider.ProviderAWS),
			RegionName:          pointer.Get("US_WEST_2"),
			EndpointServiceName: nil,
			ErrorMessage:        nil,
			InterfaceEndpoints:  nil,
			Status:              nil,
		}
		privateEndpoints := []atlasv2.EndpointService{privateAWSEndpoint}

		alertConfigResult := &atlasv2.PaginatedAlertConfig{
			Results: &[]atlasv2.GroupAlertsConfig{
				{
					Enabled:       pointer.Get(true),
					EventTypeName: pointer.Get("TestEventTypeName"),
					Matchers: &[]atlasv2.StreamsMatcher{
						{
							FieldName: "TestFieldName",
							Operator:  "TestOperator",
							Value:     "TestValue",
						},
					},
					MetricThreshold: &atlasv2.ServerlessMetricThreshold{
						MetricName: "TestMetricName",
						Operator:   pointer.Get("TestOperator"),
						Threshold:  pointer.Get(10.0),
						Units:      pointer.Get("TestUnits"),
						Mode:       pointer.Get("TestMode"),
					},
					Threshold: &atlasv2.GreaterThanRawThreshold{
						Operator:  pointer.Get("TestOperator"),
						Units:     pointer.Get("TestUnits"),
						Threshold: pointer.Get(10),
					},
					Notifications: &[]atlasv2.AlertsNotificationRootForGroup{
						{
							ChannelName:         pointer.Get("TestChannelName"),
							DatadogApiKey:       pointer.Get("TestDatadogAPIKey"),
							DatadogRegion:       pointer.Get("TestDatadogRegion"),
							DelayMin:            pointer.Get(5),
							EmailAddress:        pointer.Get("TestEmail@mongodb.com"),
							EmailEnabled:        pointer.Get(true),
							IntervalMin:         pointer.Get(0),
							MobileNumber:        pointer.Get("+12345678900"),
							OpsGenieApiKey:      pointer.Get("TestGenieAPIKey"),
							OpsGenieRegion:      pointer.Get("TestGenieRegion"),
							ServiceKey:          pointer.Get("TestServiceKey"),
							SmsEnabled:          pointer.Get(true),
							TeamId:              pointer.Get("TestTeamID"),
							TeamName:            pointer.Get("TestTeamName"),
							TypeName:            pointer.Get("TestTypeName"),
							Username:            pointer.Get("TestUserName"),
							VictorOpsApiKey:     pointer.Get("TestVictorOpsAPIKey"),
							VictorOpsRoutingKey: pointer.Get("TestVictorOpsRoutingKey"),
							Roles:               &[]string{"Role1", "Role2"},
						},
					},
				},
			},
			TotalCount: pointer.GetNonZeroValue(1),
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
				Actions: &[]atlasv2.DatabasePrivilegeAction{
					{
						Action: "Action-1",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: "Collection-1",
								Db:         "DB-1",
								Cluster:    true,
							},
						},
					},
				},
				InheritedRoles: &[]atlasv2.DatabaseInheritedRole{
					{
						Db:   "Inherited-DB",
						Role: "Inherited-ROLE",
					},
				},
				RoleName: "TestCustomRoleName",
			},
		}

		projectTeams := &atlasv2.PaginatedTeamRole{
			Links: nil,
			Results: &[]atlasv2.TeamRole{
				{
					TeamId:    pointer.Get(teamID),
					RoleNames: &[]string{string(akov2.TeamRoleClusterManager)},
				},
			},
			TotalCount: pointer.Get(1),
		}
		teams := &atlasv2.TeamResponse{
			Id:   pointer.Get(teamID),
			Name: pointer.Get("TestTeamName"),
		}

		teamUsers := &atlasv2.PaginatedApiAppUser{
			Results: &[]atlasv2.CloudAppUser{
				{
					EmailAddress: "testuser@mooooongodb.com",
					FirstName:    "TestName",
					Id:           pointer.Get("TestID"),
					LastName:     "TestLastName",
				},
			},
			TotalCount: pointer.Get(1),
		}

		listOption := &store.ListOptions{ItemsPerPage: MaxItems}
		listAlterOpt := &atlasv2.ListAlertConfigurationsApiParams{
			GroupId:      projectID,
			ItemsPerPage: &listOption.ItemsPerPage,
		}
		containerListOptionAWS := &store.ContainersListOptions{ListOptions: *listOption, ProviderName: string(akov2provider.ProviderAWS)}
		containerListOptionGCP := &store.ContainersListOptions{ListOptions: *listOption, ProviderName: string(akov2provider.ProviderGCP)}
		containerListOptionAzure := &store.ContainersListOptions{ListOptions: *listOption, ProviderName: string(akov2provider.ProviderAzure)}
		projectStore.EXPECT().ProjectIPAccessLists(projectID, listOption).Return(ipAccessLists, nil)
		projectStore.EXPECT().MaintenanceWindow(projectID).Return(mw, nil)
		projectStore.EXPECT().Integrations(projectID).Return(thirdPartyIntegrations, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOptionAWS).Return(peeringConnections, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOptionGCP).Return(nil, nil)
		projectStore.EXPECT().PeeringConnections(projectID, containerListOptionAzure).Return(nil, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderAWS)).Return(privateEndpoints, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderGCP)).Return(nil, nil)
		projectStore.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderAzure)).Return(nil, nil)
		projectStore.EXPECT().EncryptionAtRest(projectID).Return(encryptionAtRest, nil)
		projectStore.EXPECT().CloudProviderAccessRoles(projectID).Return(cpas, nil)
		projectStore.EXPECT().ProjectSettings(projectID).Return(projectSettings, nil)
		projectStore.EXPECT().Auditing(projectID).Return(auditing, nil)
		projectStore.EXPECT().AlertConfigurations(listAlterOpt).Return(alertConfigResult, nil)
		projectStore.EXPECT().DatabaseRoles(projectID).Return(customRoles, nil)
		projectStore.EXPECT().ProjectTeams(projectID, nil).Return(projectTeams, nil)
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
		projectResult, err := BuildAtlasProject(&AtlasProjectBuildRequest{
			ProjectStore:    projectStore,
			Project:         p,
			Validator:       featureValidator,
			OrgID:           orgID,
			ProjectID:       projectID,
			TargetNamespace: targetNamespace,
			IncludeSecret:   true,
			Dictionary:      dictionary,
			Version:         resourceVersion,
		})
		if err != nil {
			t.Fatalf("%v", err)
		}
		gotProject := projectResult.Project
		gotTeams := projectResult.Teams

		alertConfigs := alertConfigResult.GetResults()
		expectedThreshold := &akov2.Threshold{
			Operator:  alertConfigs[0].Threshold.GetOperator(),
			Units:     alertConfigs[0].Threshold.GetUnits(),
			Threshold: strconv.Itoa(alertConfigs[0].Threshold.GetThreshold()),
		}
		expectedMatchers := []akov2.Matcher{
			{
				FieldName: alertConfigs[0].GetMatchers()[0].GetFieldName(),
				Operator:  alertConfigs[0].GetMatchers()[0].GetOperator(),
				Value:     alertConfigs[0].GetMatchers()[0].GetValue(),
			},
		}
		expectedNotifications := []akov2.Notification{
			{
				APITokenRef: akov2common.ResourceRefNamespaced{
					Name:      gotProject.Spec.AlertConfigurations[0].Notifications[0].APITokenRef.Name,
					Namespace: gotProject.Spec.AlertConfigurations[0].Notifications[0].APITokenRef.Namespace,
				},
				ChannelName:   alertConfigs[0].GetNotifications()[0].GetChannelName(),
				DatadogRegion: alertConfigs[0].GetNotifications()[0].GetDatadogRegion(),
				DatadogAPIKeyRef: akov2common.ResourceRefNamespaced{
					Name:      gotProject.Spec.AlertConfigurations[0].Notifications[0].DatadogAPIKeyRef.Name,
					Namespace: gotProject.Spec.AlertConfigurations[0].Notifications[0].DatadogAPIKeyRef.Namespace,
				},
				DelayMin:       alertConfigs[0].GetNotifications()[0].DelayMin,
				EmailAddress:   alertConfigs[0].GetNotifications()[0].GetEmailAddress(),
				EmailEnabled:   alertConfigs[0].GetNotifications()[0].EmailEnabled,
				IntervalMin:    alertConfigs[0].GetNotifications()[0].GetIntervalMin(),
				MobileNumber:   alertConfigs[0].GetNotifications()[0].GetMobileNumber(),
				OpsGenieRegion: alertConfigs[0].GetNotifications()[0].GetOpsGenieRegion(),
				OpsGenieAPIKeyRef: akov2common.ResourceRefNamespaced{
					Name:      gotProject.Spec.AlertConfigurations[0].Notifications[0].OpsGenieAPIKeyRef.Name,
					Namespace: gotProject.Spec.AlertConfigurations[0].Notifications[0].OpsGenieAPIKeyRef.Namespace,
				},
				ServiceKeyRef: akov2common.ResourceRefNamespaced{
					Name:      gotProject.Spec.AlertConfigurations[0].Notifications[0].ServiceKeyRef.Name,
					Namespace: gotProject.Spec.AlertConfigurations[0].Notifications[0].ServiceKeyRef.Namespace,
				},
				SMSEnabled: alertConfigs[0].GetNotifications()[0].SmsEnabled,
				TeamID:     alertConfigs[0].GetNotifications()[0].GetTeamId(),
				TeamName:   alertConfigs[0].GetNotifications()[0].GetTeamName(),
				TypeName:   alertConfigs[0].GetNotifications()[0].GetTypeName(),
				Username:   alertConfigs[0].GetNotifications()[0].GetUsername(),
				Roles:      alertConfigs[0].GetNotifications()[0].GetRoles(),
				VictorOpsSecretRef: akov2common.ResourceRefNamespaced{
					Name:      gotProject.Spec.AlertConfigurations[0].Notifications[0].VictorOpsSecretRef.Name,
					Namespace: gotProject.Spec.AlertConfigurations[0].Notifications[0].VictorOpsSecretRef.Namespace,
				},
			},
		}
		expectedMetricThreshold := &akov2.MetricThreshold{
			MetricName: alertConfigs[0].MetricThreshold.MetricName,
			Operator:   alertConfigs[0].MetricThreshold.GetOperator(),
			Threshold:  fmt.Sprintf("%f", alertConfigs[0].MetricThreshold.GetThreshold()),
			Units:      alertConfigs[0].MetricThreshold.GetUnits(),
			Mode:       alertConfigs[0].MetricThreshold.GetMode(),
		}
		teamsName := teams.GetName()
		expectedTeams := []*akov2.AtlasTeam{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "AtlasTeam",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-team-%s", strings.ToLower(p.Name), strings.ToLower(teamsName)),
					Namespace: targetNamespace,
					Labels: map[string]string{
						features.ResourceVersion: resourceVersion,
					},
				},
				Spec: akov2.TeamSpec{
					Name:      teamsName,
					Usernames: []akov2.TeamUser{akov2.TeamUser(teamUsers.GetResults()[0].Username)},
				},
				Status: akov2status.TeamStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
		}
		expectedProject := &akov2.AtlasProject{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasProject",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(p.Name, dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: akov2.AtlasProjectSpec{
				Name: p.Name,
				ConnectionSecret: &akov2common.ResourceRefNamespaced{
					Name: resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, p.Name), dictionary),
				},
				ProjectIPAccessList: []akov2project.IPAccessList{
					{
						AwsSecurityGroup: ipAccessLists.GetResults()[0].GetAwsSecurityGroup(),
						CIDRBlock:        ipAccessLists.GetResults()[0].GetCidrBlock(),
						Comment:          ipAccessLists.GetResults()[0].GetComment(),
						DeleteAfterDate:  ipAccessLists.GetResults()[0].GetDeleteAfterDate().String(),
						IPAddress:        ipAccessLists.GetResults()[0].GetIpAddress(),
					},
				},
				MaintenanceWindow: akov2project.MaintenanceWindow{
					DayOfWeek: mw.DayOfWeek,
					HourOfDay: mw.GetHourOfDay(),
					AutoDefer: mw.GetAutoDeferOnceEnabled(),
					StartASAP: mw.GetStartASAP(),
					Defer:     false,
				},
				PrivateEndpoints: []akov2.PrivateEndpoint{
					{
						Provider:          akov2provider.ProviderAWS,
						Region:            *privateAWSEndpoint.RegionName,
						ID:                firstElementOrZeroValue(privateAWSEndpoint.GetInterfaceEndpoints()),
						IP:                "",
						GCPProjectID:      "",
						EndpointGroupName: "",
						Endpoints:         akov2.GCPEndpoints{},
					},
				},
				CloudProviderAccessRoles: []akov2.CloudProviderAccessRole{
					{
						ProviderName:      cpas.GetAwsIamRoles()[0].ProviderName,
						IamAssumedRoleArn: *cpas.GetAwsIamRoles()[0].IamAssumedRoleArn,
					},
				},
				AlertConfigurations: []akov2.AlertConfiguration{
					{
						Enabled:         alertConfigs[0].GetEnabled(),
						EventTypeName:   alertConfigs[0].GetEventTypeName(),
						Matchers:        expectedMatchers,
						Threshold:       expectedThreshold,
						Notifications:   expectedNotifications,
						MetricThreshold: expectedMetricThreshold,
					},
				},
				AlertConfigurationSyncEnabled: false,
				NetworkPeers: []akov2.NetworkPeer{
					{
						AccepterRegionName:  peeringConnectionAWS.GetAccepterRegionName(),
						ContainerRegion:     "",
						AWSAccountID:        peeringConnectionAWS.GetAwsAccountId(),
						ContainerID:         peeringConnectionAWS.ContainerId,
						ProviderName:        akov2provider.ProviderName(peeringConnectionAWS.GetProviderName()),
						RouteTableCIDRBlock: peeringConnectionAWS.GetRouteTableCidrBlock(),
						VpcID:               peeringConnectionAWS.GetVpcId(),
					},
				},
				WithDefaultAlertsSettings: false,
				X509CertRef:               nil,
				Integrations: []akov2project.Integration{
					{
						Type:     thirdPartyIntegrations.GetResults()[0].GetType(),
						UserName: thirdPartyIntegrations.GetResults()[0].GetUsername(),
						PasswordRef: akov2common.ResourceRefNamespaced{
							Name: fmt.Sprintf("%s-integration-%s",
								strings.ToLower(projectID),
								strings.ToLower(thirdPartyIntegrations.GetResults()[0].GetType())),
							Namespace: targetNamespace,
						},
						ServiceDiscovery: thirdPartyIntegrations.GetResults()[0].GetServiceDiscovery(),
					},
				},
				EncryptionAtRest: &akov2.EncryptionAtRest{
					AwsKms: akov2.AwsKms{
						SecretRef: akov2common.ResourceRefNamespaced{
							Name:      gotProject.Spec.EncryptionAtRest.AwsKms.SecretRef.Name,
							Namespace: gotProject.Spec.EncryptionAtRest.AwsKms.SecretRef.Namespace,
						},
					},
					AzureKeyVault: akov2.AzureKeyVault{
						SecretRef: akov2common.ResourceRefNamespaced{
							Name:      gotProject.Spec.EncryptionAtRest.AzureKeyVault.SecretRef.Name,
							Namespace: gotProject.Spec.EncryptionAtRest.AzureKeyVault.SecretRef.Namespace,
						},
					},
					GoogleCloudKms: akov2.GoogleCloudKms{
						Enabled: encryptionAtRest.GoogleCloudKms.Enabled,
						SecretRef: akov2common.ResourceRefNamespaced{
							Name:      gotProject.Spec.EncryptionAtRest.GoogleCloudKms.SecretRef.Name,
							Namespace: gotProject.Spec.EncryptionAtRest.GoogleCloudKms.SecretRef.Namespace,
						},
					},
				},
				Auditing: &akov2.Auditing{
					AuditAuthorizationSuccess: auditing.GetAuditAuthorizationSuccess(),
					AuditFilter:               auditing.GetAuditFilter(),
					Enabled:                   auditing.GetEnabled(),
				},
				Settings: &akov2.ProjectSettings{
					IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
					IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
					IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
					IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
					IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
				},
				CustomRoles: []akov2.CustomRole{
					{
						Name: customRoles[0].RoleName,
						InheritedRoles: []akov2.Role{
							{
								Name:     customRoles[0].GetInheritedRoles()[0].Role,
								Database: customRoles[0].GetInheritedRoles()[0].Db,
							},
						},
						Actions: []akov2.Action{
							{
								Name: customRoles[0].GetActions()[0].Action,
								Resources: []akov2.Resource{
									{
										Cluster:    &customRoles[0].GetActions()[0].GetResources()[0].Cluster,
										Database:   &customRoles[0].GetActions()[0].GetResources()[0].Db,
										Collection: &customRoles[0].GetActions()[0].GetResources()[0].Collection,
									},
								},
							},
						},
					},
				},
				Teams: []akov2.Team{
					{
						TeamRef: akov2common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-team-%s", strings.ToLower(p.Name), strings.ToLower(teamsName)),
							Namespace: targetNamespace,
						},
						Roles: []akov2.TeamRole{akov2.TeamRole(projectTeams.GetResults()[0].GetRoleNames()[0])},
					},
				},
			},
			Status: akov2status.AtlasProjectStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		assert.Equal(t, expectedProject, gotProject)
		assert.Equal(t, expectedTeams, gotTeams)
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
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(name + credentialSuffix),
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
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(name + credentialSuffix),
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
			Results: &[]atlasv2.NetworkPermissionEntry{
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

		listOptions := &store.ListOptions{ItemsPerPage: MaxItems}

		alProvider.EXPECT().ProjectIPAccessLists(projectID, listOptions).Return(data, nil)

		got, err := buildAccessLists(alProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []akov2project.IPAccessList{
			{
				AwsSecurityGroup: data.GetResults()[0].GetAwsSecurityGroup(),
				CIDRBlock:        data.GetResults()[0].GetCidrBlock(),
				Comment:          data.GetResults()[0].GetComment(),
				DeleteAfterDate:  data.GetResults()[0].GetDeleteAfterDate().String(),
				IPAddress:        data.GetResults()[0].GetIpAddress(),
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
			AuditAuthorizationSuccess: pointer.Get(true),
			AuditFilter:               pointer.Get("TestFilter"),
			ConfigurationType:         pointer.Get("TestType"),
			Enabled:                   pointer.Get(true),
		}

		auditingProvider.EXPECT().Auditing(projectID).Return(data, nil)

		got, err := buildAuditing(auditingProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &akov2.Auditing{
			AuditAuthorizationSuccess: data.GetAuditAuthorizationSuccess(),
			AuditFilter:               data.GetAuditFilter(),
			Enabled:                   data.GetEnabled(),
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
			AwsIamRoles: &[]atlasv2.CloudProviderAccessAWSIAMRole{
				{
					AtlasAWSAccountArn:         pointer.Get("TestARN"),
					AtlasAssumedRoleExternalId: pointer.Get("TestRoleID"),
					AuthorizedDate:             &time.Time{},
					CreatedDate:                &time.Time{},
					FeatureUsages:              nil,
					IamAssumedRoleArn:          pointer.Get("TestAssumedRoleARN"),
					ProviderName:               string(akov2provider.ProviderAWS),
					RoleId:                     pointer.Get("TestRoleID"),
				},
			},
		}

		cpaProvider.EXPECT().CloudProviderAccessRoles(projectID).Return(data, nil)

		got, err := buildCloudProviderAccessRoles(cpaProvider, projectID)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := []akov2.CloudProviderAccessRole{
			{
				ProviderName:      data.GetAwsIamRoles()[0].ProviderName,
				IamAssumedRoleArn: *data.GetAwsIamRoles()[0].IamAssumedRoleArn,
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
	dictionary := resources.AtlasNameToKubernetesName()
	testProjectName := "test-project"
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

		got, _, err := buildEncryptionAtRest(dataProvider, projectID, testProjectName, targetNamespace, dictionary)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &akov2.EncryptionAtRest{
			AwsKms: akov2.AwsKms{
				Enabled: data.AwsKms.Enabled,
				Region:  data.AwsKms.GetRegion(),
				Valid:   data.AwsKms.Valid,
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.AwsKms.SecretRef.Name,
					Namespace: got.AwsKms.SecretRef.Namespace,
				},
			},
			AzureKeyVault: akov2.AzureKeyVault{
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.AzureKeyVault.SecretRef.Name,
					Namespace: got.AzureKeyVault.SecretRef.Namespace,
				},
			},
			GoogleCloudKms: akov2.GoogleCloudKms{
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.GoogleCloudKms.SecretRef.Name,
					Namespace: got.GoogleCloudKms.SecretRef.Namespace,
				},
			},
		}
		if diff := deep.Equal(expected, got); diff != nil {
			t.Fatalf("EncryptionAtREST mismatch: %v", diff)
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
		got, _, err := buildEncryptionAtRest(dataProvider, projectID, testProjectName, targetNamespace, dictionary)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &akov2.EncryptionAtRest{
			AwsKms: akov2.AwsKms{
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.AwsKms.SecretRef.Name,
					Namespace: got.AwsKms.SecretRef.Namespace,
				},
			},
			AzureKeyVault: akov2.AzureKeyVault{
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.AzureKeyVault.SecretRef.Name,
					Namespace: got.AzureKeyVault.SecretRef.Namespace,
				},
			},
			GoogleCloudKms: akov2.GoogleCloudKms{
				Enabled: data.GoogleCloudKms.Enabled,
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.GoogleCloudKms.SecretRef.Name,
					Namespace: got.GoogleCloudKms.SecretRef.Namespace,
				},
			},
		}

		if diff := deep.Equal(expected, got); diff != nil {
			t.Fatalf("EncryptionAtREST mismatch: %v", diff)
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
		got, _, err := buildEncryptionAtRest(dataProvider, projectID, testProjectName, targetNamespace, dictionary)
		if err != nil {
			t.Errorf("%v", err)
		}

		expected := &akov2.EncryptionAtRest{
			AwsKms: akov2.AwsKms{
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.AwsKms.SecretRef.Name,
					Namespace: got.AwsKms.SecretRef.Namespace,
				},
			},
			AzureKeyVault: akov2.AzureKeyVault{
				Enabled:           data.AzureKeyVault.Enabled,
				ClientID:          data.AzureKeyVault.GetClientID(),
				AzureEnvironment:  data.AzureKeyVault.GetAzureEnvironment(),
				ResourceGroupName: data.AzureKeyVault.GetResourceGroupName(),
				TenantID:          data.AzureKeyVault.GetTenantID(),
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.AzureKeyVault.SecretRef.Name,
					Namespace: got.AzureKeyVault.SecretRef.Namespace,
				},
			},
			GoogleCloudKms: akov2.GoogleCloudKms{
				SecretRef: akov2common.ResourceRefNamespaced{
					Name:      got.GoogleCloudKms.SecretRef.Name,
					Namespace: got.GoogleCloudKms.SecretRef.Namespace,
				},
			},
		}

		if diff := deep.Equal(expected, got); diff != nil {
			t.Fatalf("EncryptionAtREST mismatch: %v", diff)
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
			Results: &[]atlasv2.ThirdPartyIntegration{
				{
					Type:             pointer.Get("PROMETHEUS"),
					Password:         pointer.Get("PrometheusTestPassword"),
					Username:         pointer.Get("PrometheusTestUserName"),
					ServiceDiscovery: pointer.Get("TestServiceDiscovery"),
				},
			},
			TotalCount: pointer.Get(0),
		}

		intProvider.EXPECT().Integrations(projectID).Return(ints, nil)

		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets, dictionary)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []akov2project.Integration{
			{
				Type:             ints.GetResults()[0].GetType(),
				ServiceDiscovery: ints.GetResults()[0].GetServiceDiscovery(),
				UserName:         ints.GetResults()[0].GetUsername(),
				PasswordRef: akov2common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.GetResults()[0].GetType())),
					Namespace: targetNamespace,
				},
			},
		}

		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.GetResults()[0].GetType())),
					Namespace: targetNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{
					secrets.PasswordField: []byte(ints.GetResults()[0].GetPassword()),
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
			Results: &[]atlasv2.ThirdPartyIntegration{
				{

					Type:             pointer.Get("PROMETHEUS"),
					Password:         pointer.Get("PrometheusTestPassword"),
					Username:         pointer.Get("PrometheusTestUserName"),
					ServiceDiscovery: pointer.Get("TestServiceDiscovery"),
				},
			},
			TotalCount: pointer.Get(0),
		}
		intProvider.EXPECT().Integrations(projectID).Return(ints, nil)
		got, intSecrets, err := buildIntegrations(intProvider, projectID, targetNamespace, includeSecrets, dictionary)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []akov2project.Integration{
			{
				Type:             ints.GetResults()[0].GetType(),
				ServiceDiscovery: ints.GetResults()[0].GetServiceDiscovery(),
				UserName:         ints.GetResults()[0].GetUsername(),
				PasswordRef: akov2common.ResourceRefNamespaced{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.GetResults()[0].GetType())),
					Namespace: targetNamespace,
				},
			},
		}

		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("%s-integration-%s",
						strings.ToLower(projectID),
						strings.ToLower(ints.GetResults()[0].GetType())),
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
			HourOfDay:            pointer.Get(10),
			StartASAP:            pointer.Get(false),
			AutoDeferOnceEnabled: pointer.Get(false),
		}

		mwProvider.EXPECT().MaintenanceWindow(projectID).Return(mw, nil)

		got, err := buildMaintenanceWindows(mwProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := akov2project.MaintenanceWindow{
			DayOfWeek: mw.DayOfWeek,
			HourOfDay: mw.GetHourOfDay(),
			AutoDefer: mw.GetAutoDeferOnceEnabled(),
			StartASAP: mw.GetStartASAP(),
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
		peeringConnectionAWS := &atlasv2.BaseNetworkPeeringConnectionSettings{
			AccepterRegionName:  pointer.Get("TestRegionName"),
			AwsAccountId:        pointer.Get("TestAWSAccountID"),
			ConnectionId:        pointer.Get("TestConnID"),
			ContainerId:         "TestContainerID",
			ErrorStateName:      pointer.Get("TestErrStateName"),
			Id:                  pointer.Get("TestID"),
			ProviderName:        pointer.Get(string(akov2provider.ProviderAWS)),
			RouteTableCidrBlock: pointer.Get("0.0.0.0/0"),
			StatusName:          pointer.Get("TestStatusName"),
			VpcId:               pointer.Get("TestVPCID"),
		}

		peeringConnections := []atlasv2.BaseNetworkPeeringConnectionSettings{
			*peeringConnectionAWS,
		}

		listOptions := store.ListOptions{ItemsPerPage: MaxItems}
		containerListOptionAWS := &store.ContainersListOptions{ListOptions: listOptions, ProviderName: string(akov2provider.ProviderAWS)}
		containerListOptionGCP := &store.ContainersListOptions{ListOptions: listOptions, ProviderName: string(akov2provider.ProviderGCP)}
		containerListOptionAzure := &store.ContainersListOptions{ListOptions: listOptions, ProviderName: string(akov2provider.ProviderAzure)}

		peerProvider.EXPECT().PeeringConnections(projectID, containerListOptionAWS).Return(peeringConnections, nil)
		peerProvider.EXPECT().PeeringConnections(projectID, containerListOptionGCP).Return(nil, nil)
		peerProvider.EXPECT().PeeringConnections(projectID, containerListOptionAzure).Return(nil, nil)

		got, err := buildNetworkPeering(peerProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []akov2.NetworkPeer{
			{
				AccepterRegionName:  peeringConnectionAWS.GetAccepterRegionName(),
				ContainerRegion:     "",
				AWSAccountID:        peeringConnectionAWS.GetAwsAccountId(),
				ContainerID:         peeringConnectionAWS.ContainerId,
				ProviderName:        akov2provider.ProviderName(*peeringConnectionAWS.ProviderName),
				RouteTableCIDRBlock: peeringConnectionAWS.GetRouteTableCidrBlock(),
				VpcID:               peeringConnectionAWS.GetVpcId(),
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
		providerName := akov2provider.ProviderAWS
		privateEndpoint := atlasv2.EndpointService{
			Id:                  pointer.Get("1"),
			CloudProvider:       string(providerName),
			RegionName:          pointer.Get("US_EAST_1"),
			EndpointServiceName: nil,
			ErrorMessage:        nil,
			InterfaceEndpoints:  &[]string{"vpce-123456"},
			Status:              nil,
		}

		peProvider.EXPECT().PrivateEndpoints(projectID, string(providerName)).Return([]atlasv2.EndpointService{privateEndpoint}, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderAzure)).Return(nil, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderGCP)).Return(nil, nil)

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []akov2.PrivateEndpoint{
			{
				Provider:          providerName,
				Region:            *privateEndpoint.RegionName,
				ID:                "vpce-123456",
				IP:                "",
				GCPProjectID:      "",
				EndpointGroupName: "",
				Endpoints:         akov2.GCPEndpoints{},
			},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("PrivateEndpoints mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})

	t.Run("Can convert PrivateEndpointConnection for Azure", func(t *testing.T) {
		providerName := akov2provider.ProviderAzure
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
		peProvider.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderAWS)).Return(nil, nil)
		peProvider.EXPECT().PrivateEndpoints(projectID, string(akov2provider.ProviderGCP)).Return(nil, nil)

		got, err := buildPrivateEndpoints(peProvider, projectID)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expected := []akov2.PrivateEndpoint{
			{
				Provider:          providerName,
				Region:            *privateEndpoint.RegionName,
				IP:                "",
				GCPProjectID:      "",
				EndpointGroupName: "",
				Endpoints:         akov2.GCPEndpoints{},
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
		expected := &akov2.ProjectSettings{
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
				Actions: &[]atlasv2.DatabasePrivilegeAction{
					{
						Action: "TestAction",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: "TestCollection",
								Db:         "TestDB",
								Cluster:    true,
							},
						},
					},
				},
				InheritedRoles: &[]atlasv2.DatabaseInheritedRole{
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
		expected := []akov2.CustomRole{
			{
				Name: role.RoleName,
				InheritedRoles: []akov2.Role{
					{
						Name:     role.GetInheritedRoles()[0].Role,
						Database: role.GetInheritedRoles()[0].Db,
					},
				},
				Actions: []akov2.Action{
					{
						Name: role.GetActions()[0].Action,
						Resources: []akov2.Resource{
							{
								Cluster:    &role.GetActions()[0].GetResources()[0].Cluster,
								Database:   &role.GetActions()[0].GetResources()[0].Db,
								Collection: &role.GetActions()[0].GetResources()[0].Collection,
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

func Test_firstElementOrEmpty(t *testing.T) {
	t.Run("should return zero value when slice is empty", func(t *testing.T) {
		assert.Empty(t, firstElementOrZeroValue([]string{}))
	})

	t.Run("should return first item when slice has a single item", func(t *testing.T) {
		assert.Equal(t, "1", firstElementOrZeroValue([]string{"1"}))
	})

	t.Run("should return first item when slice has multiple items", func(t *testing.T) {
		assert.Equal(t, "1", firstElementOrZeroValue([]string{"1", "2", "3"}))
	})
}

func TestToMatcherErrors(t *testing.T) {
	testCases := []struct {
		title            string
		m                atlasv2.StreamsMatcher
		expectedErrorMsg string
	}{
		{
			title:            "Empty map renders nil map error",
			m:                atlasv2.StreamsMatcher{},
			expectedErrorMsg: "matcher is empty",
		},
		{
			title:            "Missing fieldName renders key not set error",
			m:                atlasv2.StreamsMatcher{Operator: "op", Value: "value"},
			expectedErrorMsg: "fieldName is not set",
		},
		{
			title:            "Misnamed fieldName renders key not found error",
			m:                atlasv2.StreamsMatcher{Operator: "op"},
			expectedErrorMsg: "fieldName is not set",
		},
		{
			title:            "Missing operator renders key not found error",
			m:                atlasv2.StreamsMatcher{FieldName: "name"},
			expectedErrorMsg: "operator is not set",
		},
		{
			title:            "Missing value renders key not found error",
			m:                atlasv2.StreamsMatcher{Operator: "op", FieldName: "fieldName"},
			expectedErrorMsg: "value is not set",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			_, err := toMatcher(tc.m)
			log.Printf("err=%v", err)
			assert.ErrorContains(t, err, tc.expectedErrorMsg)
		})
	}
}

func TestConvertMatchers(t *testing.T) {
	configs := []atlasv2.StreamsMatcher{
		{},
		{FieldName: "field"},
		{FieldName: "field", Operator: "op"},
		{FieldName: "field", Operator: "op", Value: "value"},
	}
	expected := []akov2.Matcher{
		{FieldName: "field", Operator: "op", Value: "value"},
	}
	matchers := convertMatchers(configs)
	assert.Equal(t, expected, matchers)
}
