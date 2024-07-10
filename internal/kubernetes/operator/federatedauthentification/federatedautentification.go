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

package federatedautentification

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115014/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BuildAtlasFederatedAuth(FederationAuthentificationStore store.FederationAuthentificationStore, dataProvider store.OperatorGenericStore, projectName, orgID string, federatedSettings admin.OrgFederationSettings, projectID, operatorVersion, targetNamespace string, dictionary map[string]string) (*akov2.AtlasFederatedAuth, error) {
	orgConfig, err := FederationAuthentificationStore.AtlasFederatedAuthOrgConfig(&admin.GetConnectedOrgConfigApiParams{FederationSettingsId: *federatedSettings.Id, OrgId: orgID})
	if err != nil {
		return nil, err
	}
	federationAuthentification := &akov2.AtlasFederatedAuth{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasFederatedAuth",
		},
		ObjectMeta: metav1.ObjectMeta{
			//Not sure about this one
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, *federatedSettings.Id), dictionary),
			Namespace: targetNamespace,
		},
		Spec: getAtlasFederatedAuthSpec(FederationAuthentificationStore, dataProvider, targetNamespace, orgConfig, projectName, federatedSettings),
		Status: akov2status.AtlasFederatedAuthStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}
	return federationAuthentification, nil
}
func getMetadataName(FederationAuthentificationStore store.FederationAuthentificationStore, federatedSettingsId string, identityProviderId string) string {
	name, err := FederationAuthentificationStore.AtlasIdentityProviderMetadata(&atlasv2.GetIdentityProviderMetadataApiParams{FederationSettingsId: federatedSettingsId, IdentityProviderId: identityProviderId})
	if err != nil {
		return ""
	}
	return name
}
func getAtlasFederatedAuthSpec(FederationAuthentificationStore store.FederationAuthentificationStore, dataProvider store.OperatorGenericStore, targetNamespace string, orgConfig *atlasv2.ConnectedOrgConfig, projectName string, federatedSettings admin.OrgFederationSettings) akov2.AtlasFederatedAuthSpec {
	// Convert DomainAllowList and PostAuthRoleGrants from pointers to values
	var domainAllowList []string
	if orgConfig.DomainAllowList != nil {
		domainAllowList = *orgConfig.DomainAllowList
	}

	var postAuthRoleGrants []string
	if orgConfig.PostAuthRoleGrants != nil {
		postAuthRoleGrants = *orgConfig.PostAuthRoleGrants
	}

	domainRestrictionEnabled := false
	if orgConfig.DomainRestrictionEnabled {
		domainRestrictionEnabled = orgConfig.DomainRestrictionEnabled
	}
	idp := getIdentityProvider(FederationAuthentificationStore, *federatedSettings.Id, *federatedSettings.IdentityProviderId)

	// Initialize the AtlasFederatedAuthSpec
	authSpec := akov2.AtlasFederatedAuthSpec{
		Enabled:                  true,
		ConnectionSecretRef:      akov2common.ResourceRefNamespaced{Name: projectName, Namespace: targetNamespace},
		DomainAllowList:          domainAllowList,
		DomainRestrictionEnabled: &domainRestrictionEnabled,
		PostAuthRoleGrants:       postAuthRoleGrants,
		SSODebugEnabled:          idp.SsoDebugEnabled,
		RoleMappings:             nil,
	}
	if orgConfig.RoleMappings != nil {
		// Convert slice of AuthFederationRoleMapping to RoleMapping slice
		var roleMappings []akov2.RoleMapping
		for _, mapping := range *orgConfig.RoleMappings {
			roleMappings = append(roleMappings, getRoleMappings(mapping, projectName, dataProvider)...)
		}
		authSpec.RoleMappings = roleMappings
	}
	return authSpec
}

func getRoleMappings(mapping admin.AuthFederationRoleMapping, projectName string, dataProvider store.OperatorGenericStore) []akov2.RoleMapping {
	var roleAssignments []akov2.RoleAssignment
	if mapping.RoleAssignments != nil {
		for _, ra := range *mapping.RoleAssignments {
			if ra.GroupId != nil && *ra.GroupId != "" {
				pjname, _ := dataProvider.Project(*ra.GroupId)
				roleAssignments = append(roleAssignments, akov2.RoleAssignment{
					Role:        *ra.Role,
					ProjectName: pjname.Name,
				})
			} else {
				roleAssignments = append(roleAssignments, akov2.RoleAssignment{
					Role: *ra.Role,
				})
			}
		}
	}

	roleMappingList := []akov2.RoleMapping{
		{
			ExternalGroupName: mapping.ExternalGroupName,
			RoleAssignments:   roleAssignments,
		},
	}

	return roleMappingList
}
func getIdentityProvider(FederationAuthentificationStore store.FederationAuthentificationStore, FederationSettingsId string, IdentityProviderId string) *admin.FederationIdentityProvider {
	idp, _ := FederationAuthentificationStore.AtlasIdentityProvider(&atlasv2.GetIdentityProviderApiParams{FederationSettingsId: FederationSettingsId, IdentityProviderId: IdentityProviderId})
	return idp
}
