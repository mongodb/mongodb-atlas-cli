// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package project

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CustomRolesRequest struct {
	ProjectName     string
	ProjectID       string
	TargetNamespace string
	Credentials     string
	Version         string
	IsIndependent   bool
	Dict            map[string]string
}

func BuildCustomRoles(provider store.DatabaseRoleLister, request CustomRolesRequest) ([]akov2.AtlasCustomRole, error) {
	roles, err := provider.DatabaseRoles(request.ProjectID)
	if err != nil {
		return nil, err
	}
	if roles == nil {
		return nil, nil
	}

	result := make([]akov2.AtlasCustomRole, 0, len(roles))

	for rIdx := range roles {
		role := &roles[rIdx]

		inhRoles := make([]akov2.Role, 0, len(role.GetInheritedRoles()))
		for _, rl := range role.GetInheritedRoles() {
			inhRoles = append(inhRoles, akov2.Role{
				Name:     rl.Role,
				Database: rl.Db,
			})
		}

		actions := make([]akov2.Action, 0, len(role.GetActions()))
		for _, action := range role.GetActions() {
			r := make([]akov2.Resource, 0, len(action.GetResources()))
			for _, res := range action.GetResources() {
				r = append(r, akov2.Resource{
					Cluster:    pointer.Get(res.Cluster),
					Database:   pointer.Get(res.Db),
					Collection: pointer.Get(res.Collection),
				})
			}
			actions = append(actions, akov2.Action{
				Name:      action.Action,
				Resources: r,
			})
		}

		akoRole := akov2.AtlasCustomRole{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasCustomRole",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: resources.NormalizeAtlasName(
					fmt.Sprintf("%s-custom-role-%s",
						request.ProjectName,
						role.RoleName),
					request.Dict),
				Namespace: request.TargetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: request.Version,
				},
			},
			Spec: akov2.AtlasCustomRoleSpec{
				Role: akov2.CustomRole{
					Name:           role.RoleName,
					InheritedRoles: inhRoles,
					Actions:        actions,
				},
			},
			Status: akov2status.AtlasCustomRoleStatus{
				Common: akoapi.Common{Conditions: []akoapi.Condition{}},
			},
		}
		if request.IsIndependent {
			akoRole.Spec.ExternalProjectIDRef = &akov2.ExternalProjectReference{
				ID: request.ProjectID,
			}
			akoRole.Spec.LocalCredentialHolder = akoapi.LocalCredentialHolder{
				ConnectionSecret: &akoapi.LocalObjectReference{
					Name: resources.NormalizeAtlasName(request.Credentials, request.Dict),
				},
			}
		} else {
			akoRole.Spec.ProjectRef = &akov2common.ResourceRefNamespaced{
				Name:      request.ProjectName,
				Namespace: request.TargetNamespace,
			}
		}
		result = append(result, akoRole)
	}
	return result, nil
}
