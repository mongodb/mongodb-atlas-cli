// Copyright 2023 MongoDB Inc
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

package operator

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530005/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const orgID = "TestOrgID"
const projectID = "TestProjectID"

func Test_fetchDataFederationNames(t *testing.T) {
	ctl := gomock.NewController(t)
	atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)

	t.Run("Can fetch Data Federation Instance names", func(t *testing.T) {
		dataFederations := []admin.DataLakeTenant{
			{
				DataProcessRegion: &admin.DataLakeDataProcessRegion{
					CloudProvider: "TestProvider",
					Region:        "TestRegion",
				},
				Name:  pointer.Get("DataFederationInstance0"),
				State: pointer.Get("TestState"),
				Storage: &admin.DataLakeStorage{
					Databases: nil,
					Stores:    nil,
				},
			},
			{
				DataProcessRegion: &admin.DataLakeDataProcessRegion{
					CloudProvider: "TestProvider",
					Region:        "TestRegion",
				},
				Name:  pointer.Get("DataFederationInstance1"),
				State: pointer.Get("TestState"),
				Storage: &admin.DataLakeStorage{
					Databases: nil,
					Stores:    nil,
				},
			},
			{
				DataProcessRegion: &admin.DataLakeDataProcessRegion{
					CloudProvider: "TestProvider",
					Region:        "TestRegion",
				},
				Name:  pointer.Get("DataFederationInstance2"),
				State: pointer.Get("TestState"),
				Storage: &admin.DataLakeStorage{
					Databases: nil,
					Stores:    nil,
				},
			},
		}

		atlasOperatorGenericStore.EXPECT().DataFederationList(projectID).Return(dataFederations, nil)
		expected := []string{"DataFederationInstance0", "DataFederationInstance1", "DataFederationInstance2"}
		ce := NewConfigExporter(
			atlasOperatorGenericStore,
			nil,           // credsProvider (not used)
			projectID, "", // orgID (not used)
		)
		got, err := ce.fetchDataFederationNames()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if diff := deep.Equal(got, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func TestProjectWithWrongOrgID(t *testing.T) {
	ctl := gomock.NewController(t)
	atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)

	t.Run("should fail flagging when the org id does not match the owner of the project", func(t *testing.T) {
		project := &admin.Group{
			Id:    pointer.Get("project-id"),
			Name:  "test-project",
			OrgId: "right-org-id",
		}

		atlasOperatorGenericStore.EXPECT().Project(projectID).Return(project, nil)
		ce := NewConfigExporter(
			atlasOperatorGenericStore,
			nil, // credsProvider (not used)
			projectID, "wrong-org-id",
		)
		_, got := ce.Run()
		expected := errors.New("the project test-project (project-id) is not part of the " +
			"organization \"wrong-org-id\", please confirm the arguments provided " +
			"to the command or you are using the correct profile")
		if diff := deep.Equal(got, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func TestExportAtlasStreamProcessing(t *testing.T) {
	t.Run("should return nil when resource is not supported", func(t *testing.T) {
		ctl := gomock.NewController(t)
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		featureValidator := mocks.NewMockFeatureValidator(ctl)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamInstance).
			Return(false)

		ce := NewConfigExporter(atlasOperatorGenericStore, nil, projectID, orgID).
			WithFeatureValidator(featureValidator)

		resources, err := ce.exportAtlasStreamProcessing("my-project")
		require.NoError(t, err)
		assert.Nil(t, resources)
	})

	t.Run("should return error when fail to list streams instances", func(t *testing.T) {
		ctl := gomock.NewController(t)
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		atlasOperatorGenericStore.EXPECT().
			ProjectStreams(&admin.ListStreamInstancesApiParams{GroupId: projectID}).
			Return(nil, errors.New("failed to list streams instances"))

		featureValidator := mocks.NewMockFeatureValidator(ctl)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamInstance).
			Return(true)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamConnection).
			Return(true)

		ce := NewConfigExporter(atlasOperatorGenericStore, nil, projectID, orgID).
			WithFeatureValidator(featureValidator)

		resources, err := ce.exportAtlasStreamProcessing("my-project")
		require.ErrorContains(t, err, "failed to list streams instances")
		assert.Nil(t, resources)
	})

	t.Run("should return error when fail to list streams connections", func(t *testing.T) {
		ctl := gomock.NewController(t)
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		atlasOperatorGenericStore.EXPECT().
			ProjectStreams(&admin.ListStreamInstancesApiParams{GroupId: projectID}).
			Return(&admin.PaginatedApiStreamsTenant{Results: &[]admin.StreamsTenant{{Name: pointer.Get("instance-0")}}}, nil)
		atlasOperatorGenericStore.EXPECT().
			StreamsConnections(projectID, "instance-0").
			Return(nil, errors.New("failed to list streams connections"))

		featureValidator := mocks.NewMockFeatureValidator(ctl)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamInstance).
			Return(true)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamConnection).
			Return(true)

		ce := NewConfigExporter(atlasOperatorGenericStore, nil, projectID, orgID).
			WithFeatureValidator(featureValidator)

		resources, err := ce.exportAtlasStreamProcessing("my-project")
		require.ErrorContains(t, err, "failed to list streams connections")
		assert.Nil(t, resources)
	})

	t.Run("should return error when fail to build resources", func(t *testing.T) {
		ctl := gomock.NewController(t)
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		atlasOperatorGenericStore.EXPECT().
			ProjectStreams(&admin.ListStreamInstancesApiParams{GroupId: projectID}).
			Return(&admin.PaginatedApiStreamsTenant{Results: &[]admin.StreamsTenant{{Name: pointer.Get("instance-0")}}}, nil)
		atlasOperatorGenericStore.EXPECT().
			StreamsConnections(projectID, "instance-0").
			Return(
				&admin.PaginatedApiStreamsConnection{
					Results: &[]admin.StreamsConnection{
						{Name: pointer.Get("unknown"), Type: pointer.Get("RabbitMQ")},
					},
				},
				nil,
			)

		featureValidator := mocks.NewMockFeatureValidator(ctl)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamInstance).
			Return(true)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamConnection).
			Return(true)

		ce := NewConfigExporter(atlasOperatorGenericStore, nil, projectID, orgID).
			WithFeatureValidator(featureValidator)

		resources, err := ce.exportAtlasStreamProcessing("my-project")
		require.ErrorContains(t, err, "trying to generate an unsupported connection type")
		assert.Nil(t, resources)
	})

	t.Run("should return exported resources", func(t *testing.T) {
		ctl := gomock.NewController(t)
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		atlasOperatorGenericStore.EXPECT().
			ProjectStreams(&admin.ListStreamInstancesApiParams{GroupId: projectID}).
			Return(
				&admin.PaginatedApiStreamsTenant{
					Results: &[]admin.StreamsTenant{
						{
							Id:   pointer.Get("instance-0-id"),
							Name: pointer.Get("instance-0"),
							DataProcessRegion: &admin.StreamsDataProcessRegion{
								CloudProvider: "AWS",
								Region:        "VIRGINIA_USA",
							},
							StreamConfig: &admin.StreamConfig{
								Tier: pointer.Get("SP30"),
							},
							Hostnames: &[]string{"https://server1", "https://server2"},
							GroupId:   pointer.Get(projectID),
						},
					},
				},
				nil,
			)
		atlasOperatorGenericStore.EXPECT().
			StreamsConnections(projectID, "instance-0").
			Return(
				&admin.PaginatedApiStreamsConnection{
					Results: &[]admin.StreamsConnection{
						{
							Name: pointer.Get("sample_stream_solar"),
							Type: pointer.Get("Sample"),
						},
						{
							Name: pointer.Get("kafka-config"),
							Type: pointer.Get("Kafka"),
							Authentication: &admin.StreamsKafkaAuthentication{
								Mechanism: pointer.Get("SCRAM-SHA512"),
								Username:  pointer.Get("kafka-user"),
							},
							BootstrapServers: pointer.Get("kafka://server1:9001,kafka://server:9002"),
							Config:           pointer.Get(map[string]string{"config": "value"}),
							Security: &admin.StreamsKafkaSecurity{
								Protocol: pointer.Get("PLAINTEXT"),
							},
						},
					},
				},
				nil,
			)

		featureValidator := mocks.NewMockFeatureValidator(ctl)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamInstance).
			Return(true)
		featureValidator.EXPECT().
			IsResourceSupported(features.ResourceAtlasStreamConnection).
			Return(true)

		ce := NewConfigExporter(atlasOperatorGenericStore, nil, projectID, orgID).
			WithFeatureValidator(featureValidator).
			WithTargetNamespace("test").
			WithTargetOperatorVersion("2.4.0")

		resources, err := ce.exportAtlasStreamProcessing("my-project")
		require.NoError(t, err)
		assert.Equal(
			t,
			[]runtime.Object{
				&akov2.AtlasStreamInstance{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasStreamInstance",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-project-instance-0",
						Namespace: "test",
						Labels: map[string]string{
							"mongodb.com/atlas-resource-version": "2.4.0",
						},
					},
					Spec: akov2.AtlasStreamInstanceSpec{
						Name: "instance-0",
						Config: akov2.Config{
							Provider: "AWS",
							Region:   "VIRGINIA_USA",
							Tier:     "SP30",
						},
						Project: akov2common.ResourceRefNamespaced{
							Name:      "my-project",
							Namespace: "test",
						},
						ConnectionRegistry: []akov2common.ResourceRefNamespaced{
							{
								Name:      "my-project-instance-0-samplelowlinestreamlowlinesolar",
								Namespace: "test",
							},
							{
								Name:      "my-project-instance-0-kafka-config",
								Namespace: "test",
							},
						},
					},
					Status: akov2status.AtlasStreamInstanceStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				&akov2.AtlasStreamConnection{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasStreamConnection",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-project-instance-0-samplelowlinestreamlowlinesolar",
						Namespace: "test",
						Labels: map[string]string{
							"mongodb.com/atlas-resource-version": "2.4.0",
						},
					},
					Spec: akov2.AtlasStreamConnectionSpec{
						Name:           "sample_stream_solar",
						ConnectionType: "Sample",
					},
					Status: akov2status.AtlasStreamConnectionStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				&akov2.AtlasStreamConnection{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasStreamConnection",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-project-instance-0-kafka-config",
						Namespace: "test",
						Labels: map[string]string{
							"mongodb.com/atlas-resource-version": "2.4.0",
						},
					},
					Spec: akov2.AtlasStreamConnectionSpec{
						Name:           "kafka-config",
						ConnectionType: "Kafka",
						KafkaConfig: &akov2.StreamsKafkaConnection{
							Authentication: akov2.StreamsKafkaAuthentication{
								Mechanism: "SCRAM-SHA512",
								Credentials: akov2common.ResourceRefNamespaced{
									Name:      "my-project-instance-0-kafka-config-userpass",
									Namespace: "test",
								},
							},
							BootstrapServers: "kafka://server1:9001,kafka://server:9002",
							Security: akov2.StreamsKafkaSecurity{
								Protocol: "PLAINTEXT",
							},
							Config: map[string]string{"config": "value"},
						},
					},
					Status: akov2status.AtlasStreamConnectionStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				&corev1.Secret{
					TypeMeta: metav1.TypeMeta{
						Kind:       "Secret",
						APIVersion: "v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-project-instance-0-kafka-config-userpass",
						Namespace: "test",
						Labels: map[string]string{
							secrets.TypeLabelKey: secrets.CredLabelVal,
						},
					},
					Data: map[string][]byte{secrets.UsernameField: []byte("kafka-user"), secrets.PasswordField: []byte("")},
				},
			},
			resources,
		)
	})
}
func TestExportFederatedAuth(t *testing.T) {
	t.Run("should return exported resources", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		ce := NewConfigExporter(atlasOperatorGenericStore, nil, projectID, orgID).
			WithTargetNamespace("test").
			WithTargetOperatorVersion("2.3.0")

		testIdentityProviderID := "TestIdentityProviderID"
		testOrganizationID := "test-org"

		testProjectID := []string{"test-project-1", "test-project-2"}
		secondTestProjectID := []string{"test-project-3", "test-project-4"}
		testProjectName := []string{"test-project-name-1", "test-project-name-2"}
		testRoleProject := []string{"GROUP_OWNER", "GROUP_OWNER"}
		testRoleOrganization := []string{"ORG_OWNER", "ORG_OWNER"}
		testExternalGroupName := []string{"org-admin", "dev-team"}

		// Constructing federationSettings
		federationSettings := &admin.OrgFederationSettings{
			Id:                 pointer.Get("TestFederationSettingID"),
			IdentityProviderId: &testIdentityProviderID,
		}

		// Constructing AuthRoleMappings
		AuthRoleMappings := make([]admin.AuthFederationRoleMapping, len(testRoleProject)+len(testRoleOrganization))
		for i := range testProjectID {
			AuthRoleMappings[i] = admin.AuthFederationRoleMapping{
				ExternalGroupName: testExternalGroupName[i],
				RoleAssignments: &[]admin.RoleAssignment{
					{
						GroupId: &testProjectID[i],
						Role:    &testRoleProject[i],
					},
				},
			}
		}
		for i := range testRoleOrganization {
			AuthRoleMappings[len(testProjectID)+i] = admin.AuthFederationRoleMapping{
				ExternalGroupName: testExternalGroupName[i],
				RoleAssignments: &[]admin.RoleAssignment{
					{
						OrgId: &testOrganizationID,
						Role:  &testRoleOrganization[i],
					},
					{
						GroupId: &secondTestProjectID[i],
						Role:    &testRoleProject[i],
					},
				},
			}
		}

		orgConfig := &admin.ConnectedOrgConfig{
			DomainAllowList:          &[]string{"example.com"},
			PostAuthRoleGrants:       &[]string{"role1"},
			DomainRestrictionEnabled: true,
			RoleMappings:             &AuthRoleMappings,
		}
		identityProvider := &admin.FederationIdentityProvider{
			SsoDebugEnabled: pointer.Get(true),
		}
		atlasOperatorGenericStore.EXPECT().FederationSetting(&admin.GetFederationSettingsApiParams{OrgId: orgID}).
			Return(federationSettings, nil)

		atlasOperatorGenericStore.EXPECT().AtlasFederatedAuthOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *federationSettings.Id, OrgId: orgID}).
			Return(orgConfig, nil)

		atlasOperatorGenericStore.EXPECT().AtlasIdentityProvider(&admin.GetIdentityProviderApiParams{FederationSettingsId: *federationSettings.Id, IdentityProviderId: testIdentityProviderID}).
			Return(identityProvider, nil)

		firstProject := &admin.Group{
			Id:    pointer.Get("test-project-1"),
			Name:  "test-project-name-1",
			OrgId: "right-org-id",
		}
		secondProject := &admin.Group{
			Id:    pointer.Get("test-project-1"),
			Name:  "test-project-name-2",
			OrgId: "right-org-id",
		}
		atlasOperatorGenericStore.EXPECT().Project("test-project-1").
			Return(firstProject, nil)

		atlasOperatorGenericStore.EXPECT().Project("test-project-2").
			Return(secondProject, nil)

		atlasOperatorGenericStore.EXPECT().Project("test-project-3").
			Return(firstProject, nil)

		atlasOperatorGenericStore.EXPECT().Project("test-project-4").
			Return(secondProject, nil)
		resources, err := ce.exportAtlasFederatedAuth("my-project")
		require.NoError(t, err)

		// Constructing roleMappings
		roleMappings := make([]akov2.RoleMapping, 0, len(testRoleProject)+len(testRoleOrganization))

		for i := range testProjectID {
			roleMapping := akov2.RoleMapping{
				ExternalGroupName: testExternalGroupName[i],
				RoleAssignments: []akov2.RoleAssignment{
					{
						ProjectName: testProjectName[i],
						Role:        testRoleProject[i],
					},
				},
			}
			roleMappings = append(roleMappings, roleMapping)
		}

		for i := range testRoleOrganization {
			roleMapping := akov2.RoleMapping{
				ExternalGroupName: testExternalGroupName[i],
				RoleAssignments: []akov2.RoleAssignment{
					{
						Role: testRoleOrganization[i],
					},
					{
						ProjectName: testProjectName[i],
						Role:        testRoleProject[i],
					},
				},
			}
			roleMappings = append(roleMappings, roleMapping)
		}
		expected := []runtime.Object{
			&akov2.AtlasFederatedAuth{
				TypeMeta: metav1.TypeMeta{
					Kind:       "AtlasFederatedAuth",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-project-testfederationsettingid",
					Namespace: "test",
				},
				Spec: akov2.AtlasFederatedAuthSpec{
					// ConnectionSecretRef: akov2common.ResourceRefNamespaced{
					// 	Name:      "my-project",
					// 	Namespace: "test",
					// },
					Enabled:                  true,
					DomainAllowList:          []string{"example.com"},
					PostAuthRoleGrants:       []string{"role1"},
					DomainRestrictionEnabled: pointer.Get(true),
					SSODebugEnabled:          pointer.Get(true),
					RoleMappings:             roleMappings,
				},
				Status: akov2status.AtlasFederatedAuthStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
		}
		assert.Equal(
			t,
			expected,
			resources,
		)
	})
}
