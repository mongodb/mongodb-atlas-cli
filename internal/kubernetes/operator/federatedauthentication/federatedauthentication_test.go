// Copyright 2024 MongoDB Inc
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

///go:build unit

package federatedauthentication

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20241113004/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	testOrganizationID           = "TestOrgID"
	legacyTestIdentityProviderID = "LegacyTestIdentityProviderID"
	testIdentityProviderID       = "TestIdentityProviderID"
	testFederationSettingID      = "TestFederationSettingID"
	testProjectID                = "TestProjectID"
	targetNamespace              = "test"
	version                      = "2.3.1"
	projectName                  = "my-project"
)

var (
	testProjectIDs         = []string{"test-project-1", "test-project-2"}
	secondTestProjectIDs   = []string{"test-project-3", "test-project-4"}
	testRoleProjects       = []string{"GROUP_OWNER", "GROUP_OWNER"}
	testRoleOrganizations  = []string{"ORG_OWNER", "ORG_OWNER"}
	testExternalGroupNames = []string{"org-admin", "dev-team"}
)
var (
	ErrGetProvidersList       = errors.New("problem when fetching the list of the identity providers")
	ErrGetProjectUnauthorized = errors.New("you are not authorized to get this project details")
)

func Test_BuildAtlasFederatedAuth(t *testing.T) {
	testCases := []struct {
		name               string
		federationSettings *admin.OrgFederationSettings
		setupMocks         func(*mocks.MockOperatorGenericStore)
		expected           *akov2.AtlasFederatedAuth
		expectedError      error
	}{
		{
			name: "should build the atlas federation auth custom resource",
			federationSettings: &admin.OrgFederationSettings{
				Id:                     pointer.Get(testFederationSettingID),
				IdentityProviderId:     pointer.Get(legacyTestIdentityProviderID),
				IdentityProviderStatus: pointer.Get("ACTIVE"),
				HasRoleMappings:        pointer.Get(true),
			},
			setupMocks: func(store *mocks.MockOperatorGenericStore) {
				authRoleMappings := setupAuthRoleMappings(testProjectIDs, secondTestProjectIDs, testRoleProjects, testRoleOrganizations, testExternalGroupNames)

				orgConfig := &admin.ConnectedOrgConfig{
					DomainAllowList:          &[]string{"example.com"},
					PostAuthRoleGrants:       &[]string{"ORG_OWNER"},
					DomainRestrictionEnabled: true,
					RoleMappings:             &authRoleMappings,
					IdentityProviderId:       pointer.Get(legacyTestIdentityProviderID),
				}

				store.EXPECT().
					GetConnectedOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID), OrgId: testOrganizationID}).
					Return(orgConfig, nil)

				identityProvider := &admin.FederationIdentityProvider{
					SsoDebugEnabled: pointer.Get(true),
					OktaIdpId:       *pointer.Get(legacyTestIdentityProviderID),
					Id:              testIdentityProviderID,
				}
				paginatedResult := &admin.PaginatedFederationIdentityProvider{
					Results:    &[]admin.FederationIdentityProvider{*identityProvider},
					TotalCount: pointer.Get(1),
				}

				store.EXPECT().
					IdentityProviders(&admin.ListIdentityProvidersApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID)}).
					Return(paginatedResult, nil)

				// Mocking projects
				firstProject := &admin.Group{Id: pointer.Get("test-project-1"), Name: "test-project-name-1", OrgId: testOrganizationID}
				secondProject := &admin.Group{Id: pointer.Get("test-project-2"), Name: "test-project-name-2", OrgId: testOrganizationID}

				store.EXPECT().Project("test-project-1").Return(firstProject, nil)
				store.EXPECT().Project("test-project-2").Return(secondProject, nil)
				store.EXPECT().Project("test-project-3").Return(firstProject, nil)
				store.EXPECT().Project("test-project-4").Return(secondProject, nil)
			},
			expected: &akov2.AtlasFederatedAuth{
				TypeMeta: metav1.TypeMeta{
					Kind:       "AtlasFederatedAuth",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-project-testfederationsettingid",
					Namespace: targetNamespace,
				},
				Spec: akov2.AtlasFederatedAuthSpec{
					Enabled:                  true,
					DomainAllowList:          []string{"example.com"},
					PostAuthRoleGrants:       []string{"ORG_OWNER"},
					DomainRestrictionEnabled: pointer.Get(true),
					SSODebugEnabled:          pointer.Get(true),
					RoleMappings: []akov2.RoleMapping{
						{
							ExternalGroupName: "org-admin",
							RoleAssignments: []akov2.RoleAssignment{
								{
									ProjectName: "test-project-name-1",
									Role:        "GROUP_OWNER",
								},
							},
						},
						{
							ExternalGroupName: "dev-team",
							RoleAssignments: []akov2.RoleAssignment{
								{
									ProjectName: "test-project-name-2",
									Role:        "GROUP_OWNER",
								},
							},
						},
						{
							ExternalGroupName: "org-admin",
							RoleAssignments: []akov2.RoleAssignment{
								{
									Role: "ORG_OWNER",
								},
								{
									ProjectName: "test-project-name-1",
									Role:        "GROUP_OWNER",
								},
							},
						},
						{
							ExternalGroupName: "dev-team",
							RoleAssignments: []akov2.RoleAssignment{
								{
									Role: "ORG_OWNER",
								},
								{
									ProjectName: "test-project-name-2",
									Role:        "GROUP_OWNER",
								},
							},
						},
					},
				},
				Status: akov2status.AtlasFederatedAuthStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "should return error because lack of project permissions",
			federationSettings: &admin.OrgFederationSettings{
				Id:                     pointer.Get(testFederationSettingID),
				IdentityProviderId:     pointer.Get(legacyTestIdentityProviderID),
				IdentityProviderStatus: pointer.Get("ACTIVE"),
				HasRoleMappings:        pointer.Get(true),
			},
			setupMocks: func(store *mocks.MockOperatorGenericStore) {
				authRoleMappings := setupAuthRoleMappings(testProjectIDs, secondTestProjectIDs, testRoleProjects, testRoleOrganizations, testExternalGroupNames)

				orgConfig := &admin.ConnectedOrgConfig{
					DomainAllowList:          &[]string{"example.com"},
					PostAuthRoleGrants:       &[]string{"ORG_OWNER"},
					DomainRestrictionEnabled: true,
					RoleMappings:             &authRoleMappings,
					IdentityProviderId:       pointer.Get(legacyTestIdentityProviderID),
				}

				store.EXPECT().
					GetConnectedOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID), OrgId: testOrganizationID}).
					Return(orgConfig, nil)

				identityProvider := &admin.FederationIdentityProvider{
					SsoDebugEnabled: pointer.Get(true),
					OktaIdpId:       *pointer.Get(legacyTestIdentityProviderID),
					Id:              testIdentityProviderID,
				}
				paginatedResult := &admin.PaginatedFederationIdentityProvider{
					Results:    &[]admin.FederationIdentityProvider{*identityProvider},
					TotalCount: pointer.Get(1),
				}

				store.EXPECT().
					IdentityProviders(&admin.ListIdentityProvidersApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID)}).
					Return(paginatedResult, nil)

				store.EXPECT().Project("test-project-1").Return(nil, ErrGetProjectUnauthorized)
			},
			expected:      nil,
			expectedError: ErrGetProjectUnauthorized,
		},
		{
			name: "should return error because an error where fetching the identity providers of the fedetaion",
			federationSettings: &admin.OrgFederationSettings{
				Id:                     pointer.Get(testFederationSettingID),
				IdentityProviderId:     pointer.Get(legacyTestIdentityProviderID),
				IdentityProviderStatus: pointer.Get("ACTIVE"),
				HasRoleMappings:        pointer.Get(true),
			},
			setupMocks: func(store *mocks.MockOperatorGenericStore) {
				authRoleMappings := setupAuthRoleMappings(testProjectIDs, secondTestProjectIDs, testRoleProjects, testRoleOrganizations, testExternalGroupNames)

				orgConfig := &admin.ConnectedOrgConfig{
					DomainAllowList:          &[]string{"example.com"},
					PostAuthRoleGrants:       &[]string{"ORG_OWNER"},
					DomainRestrictionEnabled: true,
					RoleMappings:             &authRoleMappings,
					IdentityProviderId:       pointer.Get(legacyTestIdentityProviderID),
				}

				store.EXPECT().
					GetConnectedOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID), OrgId: testOrganizationID}).
					Return(orgConfig, nil)

				store.EXPECT().
					IdentityProviders(&admin.ListIdentityProvidersApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID)}).
					Return(nil, ErrGetProvidersList)
			},
			expected:      nil,
			expectedError: ErrGetProvidersList,
		},
		{
			name: "no identity provider present matching the legacy identityproviderID",
			federationSettings: &admin.OrgFederationSettings{
				Id:                     pointer.Get(testFederationSettingID),
				IdentityProviderId:     pointer.Get(legacyTestIdentityProviderID),
				IdentityProviderStatus: pointer.Get("ACTIVE"),
				HasRoleMappings:        pointer.Get(true),
			},
			setupMocks: func(store *mocks.MockOperatorGenericStore) {
				authRoleMappings := setupAuthRoleMappings(testProjectIDs, secondTestProjectIDs, testRoleProjects, testRoleOrganizations, testExternalGroupNames)

				orgConfig := &admin.ConnectedOrgConfig{
					DomainAllowList:          &[]string{"example.com"},
					PostAuthRoleGrants:       &[]string{"ORG_OWNER"},
					DomainRestrictionEnabled: true,
					RoleMappings:             &authRoleMappings,
					IdentityProviderId:       pointer.Get(legacyTestIdentityProviderID),
				}

				store.EXPECT().
					GetConnectedOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID), OrgId: testOrganizationID}).
					Return(orgConfig, nil)

				paginatedResult := &admin.PaginatedFederationIdentityProvider{
					Results:    &[]admin.FederationIdentityProvider{},
					TotalCount: pointer.Get(0),
				}

				store.EXPECT().
					IdentityProviders(&admin.ListIdentityProvidersApiParams{FederationSettingsId: *pointer.Get(testFederationSettingID)}).
					Return(paginatedResult, nil)
			},
			expected:      nil,
			expectedError: ErrNoMatchingSAMLProvider,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			store := mocks.NewMockOperatorGenericStore(ctl)
			tc.setupMocks(store)

			resources, err := BuildAtlasFederatedAuth(&AtlasFederatedAuthBuildRequest{
				IncludeSecret:                false,
				ProjectStore:                 store,
				ConnectedOrgConfigsDescriber: store,
				IdentityProviderLister:       store,
				IdentityProviderDescriber:    store,
				ProjectID:                    "TestProjectID",
				OrgID:                        testOrganizationID,
				TargetNamespace:              "test",
				Version:                      "2.3.1",
				Dictionary:                   resources.AtlasNameToKubernetesName(),
				ProjectName:                  "my-project",
				FederatedSettings:            tc.federationSettings,
			})

			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, resources)
		})
	}
}
func setupAuthRoleMappings(testProjectID, secondTestProjectID, testRoleProject, testRoleOrganization, testExternalGroupName []string) []admin.AuthFederationRoleMapping {
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
					OrgId: pointer.Get(testOrganizationID),
					Role:  &testRoleOrganization[i],
				},
				{
					GroupId: &secondTestProjectID[i],
					Role:    &testRoleProject[i],
				},
			},
		}
	}
	return AuthRoleMappings
}
