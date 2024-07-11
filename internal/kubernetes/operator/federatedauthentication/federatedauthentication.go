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

package federatedauthentication

import (
	"fmt"
	"log"

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

type AtlasFederatedAuthBuildRequest struct {
	IncludeSecret                 bool
	FederationAuthenticationStore store.FederationAuthenticationStore
	ProjectStore                  store.OperatorProjectStore
	ProjectName                   string
	OrgID                         string
	ProjectID                     string
	FederatedSettings             *admin.OrgFederationSettings
	Version                       string
	TargetNamespace               string
	Dictionary                    map[string]string
}

const credSecretFormat = "%s-credentials"

// BuildAtlasFederatedAuth builds an AtlasFederatedAuth resource based on the provided build request.
func BuildAtlasFederatedAuth(br *AtlasFederatedAuthBuildRequest) (*akov2.AtlasFederatedAuth, error) {
	orgConfig, err := br.FederationAuthenticationStore.AtlasFederatedAuthOrgConfig(&admin.GetConnectedOrgConfigApiParams{
		FederationSettingsId: *br.FederatedSettings.Id,
		OrgId:                br.OrgID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get org config: %w", err)
	}

	federatedAuth := &akov2.AtlasFederatedAuth{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasFederatedAuth",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", br.ProjectName, *br.FederatedSettings.Id), br.Dictionary),
			Namespace: br.TargetNamespace,
		},
		Spec: getAtlasFederatedAuthSpec(*br, orgConfig),
		Status: akov2status.AtlasFederatedAuthStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}

	return federatedAuth, nil
}

// getAtlasFederatedAuthSpec returns the spec for AtlasFederatedAuth based on the provided build request and org config.
func getAtlasFederatedAuthSpec(br AtlasFederatedAuthBuildRequest, orgConfig *atlasv2.ConnectedOrgConfig) akov2.AtlasFederatedAuthSpec {
	var domainAllowList []string
	if orgConfig.DomainAllowList != nil {
		domainAllowList = *orgConfig.DomainAllowList
	}

	var postAuthRoleGrants []string
	if orgConfig.PostAuthRoleGrants != nil {
		postAuthRoleGrants = *orgConfig.PostAuthRoleGrants
	}

	idp := getIdentityProvider(br.FederationAuthenticationStore, *br.FederatedSettings.Id, *br.FederatedSettings.IdentityProviderId)

	secretRef := &akov2common.ResourceRefNamespaced{}
	if br.IncludeSecret {
		secretRef.Name = resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, br.ProjectName), br.Dictionary)
	}

	authSpec := akov2.AtlasFederatedAuthSpec{
		Enabled:                  true,
		DomainAllowList:          domainAllowList,
		ConnectionSecretRef:      *secretRef,
		DomainRestrictionEnabled: &orgConfig.DomainRestrictionEnabled,
		PostAuthRoleGrants:       postAuthRoleGrants,
		SSODebugEnabled:          idp.SsoDebugEnabled,
	}

	if orgConfig.RoleMappings != nil {
		var roleMappings []akov2.RoleMapping
		for _, mapping := range *orgConfig.RoleMappings {
			roleMappings = append(roleMappings, getRoleMappings(mapping, br.ProjectStore)...)
		}
		authSpec.RoleMappings = roleMappings
	}

	return authSpec
}

// getRoleMappings converts AuthFederationRoleMapping to a slice of RoleMapping.
func getRoleMappings(mapping admin.AuthFederationRoleMapping, projectStore store.OperatorProjectStore) []akov2.RoleMapping {
	var roleAssignments []akov2.RoleAssignment
	if mapping.RoleAssignments != nil {
		for _, ra := range *mapping.RoleAssignments {
			if ra.GroupId != nil && *ra.GroupId != "" {
				project, err := projectStore.Project(*ra.GroupId)
				if err != nil {
					log.Printf("failed to get project name for GroupId %s: %v", *ra.GroupId, err)
					continue
				}
				roleAssignments = append(roleAssignments, akov2.RoleAssignment{
					Role:        *ra.Role,
					ProjectName: project.Name,
				})
			} else {
				roleAssignments = append(roleAssignments, akov2.RoleAssignment{
					Role: *ra.Role,
				})
			}
		}
	}

	return []akov2.RoleMapping{
		{
			ExternalGroupName: mapping.ExternalGroupName,
			RoleAssignments:   roleAssignments,
		},
	}
}

// getIdentityProvider retrieves the identity provider for the given federation settings and identity provider ID.
func getIdentityProvider(store store.FederationAuthenticationStore, federationSettingsID, identityProviderID string) *admin.FederationIdentityProvider {
	idp, err := store.AtlasIdentityProvider(&atlasv2.GetIdentityProviderApiParams{
		FederationSettingsId: federationSettingsID,
		IdentityProviderId:   identityProviderID,
	})
	if err != nil {
		log.Printf("failed to get identity provider for FederationSettingsId %s and IdentityProviderId %s: %v", federationSettingsID, identityProviderID, err)
		return nil
	}
	return idp
}
