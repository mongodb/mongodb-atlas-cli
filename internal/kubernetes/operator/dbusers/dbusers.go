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

package dbusers

import (
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BuildDBUsers(provider store.AtlasOperatorDBUsersStore, projectID, projectName, targetNamespace string, includeSecrets bool) ([]*atlasV1.AtlasDatabaseUser, []*corev1.Secret, error) {
	users, err := provider.DatabaseUsers(projectID, &mongodbatlas.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	if len(users) == 0 {
		return nil, nil, nil
	}

	result := make([]*atlasV1.AtlasDatabaseUser, 0, len(users))
	relatedSecrets := make([]*corev1.Secret, 0, len(users))

	for i := range users {
		user := &users[i]
		labels := convertUserLabels(user)
		roles := convertUserRoles(user)
		if len(roles) == 0 {
			continue
		}
		scopes := convertUserScopes(user)

		secret := buildUserSecret(projectName, user, targetNamespace, includeSecrets)
		relatedSecrets = append(relatedSecrets, secret)

		result = append(result, &atlasV1.AtlasDatabaseUser{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDatabaseUser",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s", projectName, user.Username)),
				Namespace: targetNamespace,
			},
			Spec: atlasV1.AtlasDatabaseUserSpec{
				Project: common.ResourceRefNamespaced{
					Name:      projectName,
					Namespace: targetNamespace,
				},
				DatabaseName:    user.DatabaseName,
				DeleteAfterDate: user.DeleteAfterDate,
				Labels:          labels,
				Roles:           roles,
				Scopes:          scopes,
				PasswordSecret: &common.ResourceRef{
					Name: secret.Name,
				},
				Username: user.Username,
				X509Type: user.X509Type,
			},
			Status: status.AtlasDatabaseUserStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		})
	}

	return result, relatedSecrets, nil
}

func buildUserSecret(projectName string, user *mongodbatlas.DatabaseUser, targetNamespace string, includeSecrets bool) *corev1.Secret {
	secret := secrets.NewAtlasSecret(fmt.Sprintf("%s-%s", projectName, user.Username), targetNamespace, map[string][]byte{
		secrets.PasswordField: []byte(""),
	})
	if includeSecrets {
		secret.Data[secrets.PasswordField] = []byte(user.Password)
	}
	return secret
}

func convertUserScopes(user *mongodbatlas.DatabaseUser) []atlasV1.ScopeSpec {
	result := make([]atlasV1.ScopeSpec, 0, len(user.Scopes))
	for _, scope := range user.Scopes {
		result = append(result, atlasV1.ScopeSpec{
			Name: scope.Name,
			Type: atlasV1.ScopeType(scope.Type),
		})
	}
	return result
}

func convertUserRoles(user *mongodbatlas.DatabaseUser) []atlasV1.RoleSpec {
	if len(user.Roles) == 0 {
		return nil
	}
	result := make([]atlasV1.RoleSpec, 0, len(user.Roles))
	for _, role := range user.Roles {
		result = append(result, atlasV1.RoleSpec{
			RoleName:       role.RoleName,
			DatabaseName:   role.DatabaseName,
			CollectionName: role.CollectionName,
		})
	}
	return result
}

func convertUserLabels(user *mongodbatlas.DatabaseUser) []common.LabelSpec {
	result := make([]common.LabelSpec, 0, len(user.Labels))
	for _, label := range user.Labels {
		result = append(result, common.LabelSpec{
			Key:   label.Key,
			Value: label.Value,
		})
	}
	return result
}
