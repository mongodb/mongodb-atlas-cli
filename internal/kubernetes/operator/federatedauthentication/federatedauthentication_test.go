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

///go:build unit

package federatedauthentication

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExportFederatedAuth(t *testing.T) {
	t.Run("should return exported resources", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)
		testIdentityProviderID := "TestIdentityProviderID"
		testOrganizationID := "TestOrgID"
		testProjectID := []string{"test-project-1", "test-project-2"}
		secondTestProjectID := []string{"test-project-3", "test-project-4"}
		testProjectName := []string{"test-project-name-1", "test-project-name-2"}
		testRoleProject := []string{"GROUP_OWNER", "GROUP_OWNER"}
		testRoleOrganization := []string{"ORG_OWNER", "ORG_OWNER"}
		testExternalGroupName := []string{"org-admin", "dev-team"}

		federationSettings := &admin.OrgFederationSettings{
			Id:                 pointer.Get("TestFederationSettingID"),
			IdentityProviderId: &testIdentityProviderID,
		}
		input := &AtlasFederatedAuthBuildRequest{
			IncludeSecret:                 false,
			FederationAuthenticationStore: atlasOperatorGenericStore,
			ProjectStore:                  atlasOperatorGenericStore,
			ProjectID:                     "TestProjectID",
			OrgID:                         testOrganizationID,
			TargetNamespace:               "test",
			Version:                       "2.3.1",
			Dictionary:                    resources.AtlasNameToKubernetesName(),
			ProjectName:                   "my-project",
			FederatedSettings:             federationSettings,
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
		atlasOperatorGenericStore.EXPECT().AtlasFederatedAuthOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *federationSettings.Id, OrgId: testOrganizationID}).
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
		resources, err := BuildAtlasFederatedAuth(input)
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
		assert.Equal(
			t,
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
			resources,
		)
	})
}
