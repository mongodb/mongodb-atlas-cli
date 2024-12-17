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

package federatedauthentication

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113003/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrNoMatchingSAMLProvider = errors.New("failed to retrieve the SAML identity provider matching the legacy ID")
)

type AtlasFederatedAuthBuildRequest struct {
	IncludeSecret                bool
	IdentityProviderLister       store.IdentityProviderLister
	ConnectedOrgConfigsDescriber store.ConnectedOrgConfigsDescriber
	ProjectStore                 store.OperatorProjectStore
	IdentityProviderDescriber    store.IdentityProviderDescriber
	ProjectName                  string
	OrgID                        string
	ProjectID                    string
	FederatedSettings            *atlasv2.OrgFederationSettings
	Version                      string
	TargetNamespace              string
	Dictionary                   map[string]string
}

const credSecretFormat = "%s-credentials"

// BuildAtlasFederatedAuth builds an AtlasFederatedAuth resource.
func BuildAtlasFederatedAuth(br *AtlasFederatedAuthBuildRequest) (*akov2.AtlasFederatedAuth, error) {
	orgConfig, err := getOrgConfig(br)
	if err != nil {
		return nil, err
	}

	spec, err := atlasFederatedAuthSpec(*br, orgConfig)
	if err != nil {
		return nil, err
	}
	return &akov2.AtlasFederatedAuth{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasFederatedAuth",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", br.ProjectName, br.FederatedSettings.GetId()), br.Dictionary),
			Namespace: br.TargetNamespace,
		},
		Spec: *spec,
		Status: akov2status.AtlasFederatedAuthStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}, nil
}

// getOrgConfig retrieves the organization configuration for the AtlasFederatedAuth resource.
func getOrgConfig(br *AtlasFederatedAuthBuildRequest) (*atlasv2.ConnectedOrgConfig, error) {
	return br.ConnectedOrgConfigsDescriber.GetConnectedOrgConfig(&atlasv2.GetConnectedOrgConfigApiParams{
		FederationSettingsId: br.FederatedSettings.GetId(),
		OrgId:                br.OrgID,
	})
}

// atlasFederatedAuthSpec returns the spec for AtlasFederatedAuth.
func atlasFederatedAuthSpec(br AtlasFederatedAuthBuildRequest, orgConfig *atlasv2.ConnectedOrgConfig) (*akov2.AtlasFederatedAuthSpec, error) {
	idp, err := GetIdentityProviderForFederatedSettings(br.IdentityProviderLister, br.FederatedSettings.GetId(), br.FederatedSettings.GetIdentityProviderId())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the federated authentication spec: %w", err)
	}
	authSpec := akov2.AtlasFederatedAuthSpec{
		Enabled:                  true,
		DomainAllowList:          orgConfig.GetDomainAllowList(),
		ConnectionSecretRef:      getSecretRef(br),
		DomainRestrictionEnabled: pointer.Get(orgConfig.GetDomainRestrictionEnabled()),
		PostAuthRoleGrants:       orgConfig.GetPostAuthRoleGrants(),
		SSODebugEnabled:          pointer.Get(idp.GetSsoDebugEnabled()),
	}
	if br.FederatedSettings.HasHasRoleMappings() && orgConfig.HasRoleMappings() {
		authSpec.RoleMappings, err = getRoleMappings(orgConfig.GetRoleMappings(), br.ProjectStore)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve the role mappings: %w", err)
		}
	}

	return &authSpec, nil
}

// getSecretRef generates a secret reference for the AtlasFederatedAuthSpec.
func getSecretRef(br AtlasFederatedAuthBuildRequest) akov2common.ResourceRefNamespaced {
	secretRef := &akov2common.ResourceRefNamespaced{}
	if br.IncludeSecret {
		secretRef.Name = resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, br.ProjectName), br.Dictionary)
		secretRef.Namespace = br.TargetNamespace
	}
	return *secretRef
}

// getRoleMappings converts AuthFederationRoleMapping to RoleMapping.
func getRoleMappings(mappings []atlasv2.AuthFederationRoleMapping, projectStore store.OperatorProjectStore) ([]akov2.RoleMapping, error) {
	roleMappings := make([]akov2.RoleMapping, 0, len(mappings))
	for _, mapping := range mappings {
		if mapping.HasRoleAssignments() {
			roleAssignemnts, err := getRoleAssignments(mapping.GetRoleAssignments(), projectStore)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve the role assignments: %w", err)
			}
			roleMappings = append(roleMappings, akov2.RoleMapping{
				ExternalGroupName: mapping.GetExternalGroupName(),
				RoleAssignments:   roleAssignemnts,
			})
		}
	}
	return roleMappings, nil
}

// getRoleAssignments converts RoleAssignments from AuthFederationRoleMapping.
func getRoleAssignments(assignments []atlasv2.RoleAssignment, projectStore store.OperatorProjectStore) ([]akov2.RoleAssignment, error) {
	roleAssignments := make([]akov2.RoleAssignment, 0, len(assignments))
	for _, ra := range assignments {
		roleAssignment := akov2.RoleAssignment{Role: ra.GetRole()}
		if ra.HasGroupId() {
			project, err := projectStore.Project(ra.GetGroupId())
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve the project: %w", err)
			}
			roleAssignment.ProjectName = project.GetName()
		}
		roleAssignments = append(roleAssignments, roleAssignment)
	}
	return roleAssignments, nil
}

// GetIdentityProviderForFederatedSettings retrieves the requested identityprovider from a list of the identity provider for the given federation settings.
func GetIdentityProviderForFederatedSettings(st store.IdentityProviderLister, federationSettingsID string, identityProviderID string) (*atlasv2.FederationIdentityProvider, error) {
	identityProviders, err := st.IdentityProviders(&atlasv2.ListIdentityProvidersApiParams{
		FederationSettingsId: federationSettingsID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the federation setting's identity providers: %w", err)
	}

	for _, identityProvider := range identityProviders.GetResults() {
		if identityProvider.GetOktaIdpId() == identityProviderID {
			return &identityProvider, nil
		}
	}
	return nil, ErrNoMatchingSAMLProvider
}
