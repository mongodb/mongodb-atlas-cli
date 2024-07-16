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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrNoIndentityProvider = errors.New("could not find the identity provider")
	ErrProjectName         = errors.New("could not get the project name")
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

	spec, err := getAtlasFederatedAuthSpec(*br, orgConfig)
	if err != nil {
		return nil, err
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
		Spec: *spec,
		Status: akov2status.AtlasFederatedAuthStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}

	return federatedAuth, nil
}

// getOrgConfig retrieves the organization configuration for the AtlasFederatedAuth resource.
func getOrgConfig(br *AtlasFederatedAuthBuildRequest) (*atlasv2.ConnectedOrgConfig, error) {
	return br.ConnectedOrgConfigsDescriber.GetConnectedOrgConfig(&atlasv2.GetConnectedOrgConfigApiParams{
		FederationSettingsId: *br.FederatedSettings.Id,
		OrgId:                br.OrgID,
	})
}

// getAtlasFederatedAuthSpec returns the spec for AtlasFederatedAuth.
func getAtlasFederatedAuthSpec(br AtlasFederatedAuthBuildRequest, orgConfig *atlasv2.ConnectedOrgConfig) (*akov2.AtlasFederatedAuthSpec, error) {
	domainAllowList := getDomainAllowList(orgConfig)
	postAuthRoleGrants := getPostAuthRoleGrants(orgConfig)

	idp, err := GetIdentityProviderForFederatedSettings(br.IdentityProviderLister, *br.FederatedSettings.Id, *br.FederatedSettings.IdentityProviderId)
	if err != nil {
		return nil, err
	}
	secretRef := getSecretRef(br)

	authSpec := akov2.AtlasFederatedAuthSpec{
		Enabled:                  true,
		DomainAllowList:          domainAllowList,
		ConnectionSecretRef:      *secretRef,
		DomainRestrictionEnabled: &orgConfig.DomainRestrictionEnabled,
		PostAuthRoleGrants:       postAuthRoleGrants,
		SSODebugEnabled:          idp.SsoDebugEnabled,
	}
	if br.FederatedSettings.HasRoleMappings != nil && *br.FederatedSettings.HasRoleMappings && orgConfig.RoleMappings != nil {
		authSpec.RoleMappings, err = getRoleMappings(orgConfig.RoleMappings, br.ProjectStore)
		if err != nil {
			return nil, err
		}
	}

	return &authSpec, nil
}

// getDomainAllowList retrieves the domain allow list from the organization configuration.
func getDomainAllowList(orgConfig *atlasv2.ConnectedOrgConfig) []string {
	if orgConfig.DomainAllowList != nil {
		return *orgConfig.DomainAllowList
	}
	return nil
}

// getPostAuthRoleGrants retrieves the post-auth role grants from the organization configuration.
func getPostAuthRoleGrants(orgConfig *atlasv2.ConnectedOrgConfig) []string {
	if orgConfig.PostAuthRoleGrants != nil {
		return *orgConfig.PostAuthRoleGrants
	}
	return nil
}

// getSecretRef generates a secret reference for the AtlasFederatedAuthSpec.
func getSecretRef(br AtlasFederatedAuthBuildRequest) *akov2common.ResourceRefNamespaced {
	secretRef := &akov2common.ResourceRefNamespaced{}
	if br.IncludeSecret {
		secretRef.Name = resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, br.ProjectName), br.Dictionary)
		secretRef.Namespace = br.TargetNamespace
	}
	return secretRef
}

// getRoleMappings converts AuthFederationRoleMapping to RoleMapping.
func getRoleMappings(mappings *[]atlasv2.AuthFederationRoleMapping, projectStore store.OperatorProjectStore) ([]akov2.RoleMapping, error) {
	roleMappings := make([]akov2.RoleMapping, 0, len(*mappings))
	for _, mapping := range *mappings {
		roleAssignemnts, err := getRoleAssignments(mapping.RoleAssignments, projectStore)
		if err != nil {
			return nil, err
		}
		roleMappings = append(roleMappings, akov2.RoleMapping{
			ExternalGroupName: mapping.ExternalGroupName,
			RoleAssignments:   roleAssignemnts,
		})
	}
	return roleMappings, nil
}

// getRoleAssignments converts RoleAssignments from AuthFederationRoleMapping.
func getRoleAssignments(assignments *[]atlasv2.RoleAssignment, projectStore store.OperatorProjectStore) ([]akov2.RoleAssignment, error) {
	var roleAssignments []akov2.RoleAssignment
	if assignments != nil {
		for _, ra := range *assignments {
			roleAssignment := akov2.RoleAssignment{Role: *ra.Role}
			if ra.GroupId != nil && *ra.GroupId != "" {
				project, err := projectStore.Project(*ra.GroupId)
				if err != nil {
					return nil, ErrProjectName
				}
				roleAssignment.ProjectName = project.Name
			}
			roleAssignments = append(roleAssignments, roleAssignment)
		}
	}
	return roleAssignments, nil
}

// GetIdentityProviderForFederatedSettings retrieves the list of the identity provider for the given federation settings.
func GetIdentityProviderForFederatedSettings(store store.IdentityProviderLister, federationSettingsID string, identityProviderID string) (*atlasv2.FederationIdentityProvider, error) {
	identityProviders, err := store.IdentityProviders(&atlasv2.ListIdentityProvidersApiParams{
		FederationSettingsId: federationSettingsID,
	})
	if err != nil {
		return nil, err
	}

	for _, identityProvider := range identityProviders.GetResults() {
		if identityProvider.GetOktaIdpId() == identityProviderID {
			return &identityProvider, nil
		}
	}
	return nil, ErrNoIndentityProvider
}
