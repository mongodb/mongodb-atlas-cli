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

	"github.com/google/uuid"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BuildDBUsers(provider atlas.OperatorDBUsersStore, projectID, projectName, targetNamespace string, dictionary map[string]string, version string) ([]*atlasV1.AtlasDatabaseUser, []*corev1.Secret, error) {
	users, err := provider.DatabaseUsers(projectID, &atlas.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	if len(users.Results) == 0 {
		return nil, nil, nil
	}

	mappedUsers := map[string]*atlasV1.AtlasDatabaseUser{}
	relatedSecrets := make([]*corev1.Secret, 0, len(users.Results))

	for i := range users.Results {
		user := &users.Results[i]

		resourceName := suggestResourceName(projectName, user.Username, mappedUsers, dictionary)
		labels := convertUserLabels(user)
		roles := convertUserRoles(user)
		if len(roles) == 0 {
			continue
		}
		scopes := convertUserScopes(user)

		mappedUsers[resourceName] = &atlasV1.AtlasDatabaseUser{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDatabaseUser",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      resourceName,
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: version,
				},
			},
			Spec: atlasV1.AtlasDatabaseUserSpec{
				Project: common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(projectName, dictionary),
					Namespace: targetNamespace,
				},
				DatabaseName:    user.DatabaseName,
				DeleteAfterDate: getDeleteAfterDate(user),
				Labels:          labels,
				Roles:           roles,
				Scopes:          scopes,
				Username:        user.Username,
				X509Type:        *user.X509Type,
			},
			Status: status.AtlasDatabaseUserStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		if user.GetX509Type() != "MANAGED" {
			secret := buildUserSecret(resourceName, targetNamespace, dictionary)
			relatedSecrets = append(relatedSecrets, secret)

			mappedUsers[resourceName].Spec.PasswordSecret = &common.ResourceRef{
				Name: resources.NormalizeAtlasName(secret.Name, dictionary),
			}
		}
	}

	result := make([]*atlasV1.AtlasDatabaseUser, 0, len(users.Results))
	for _, mappedUser := range mappedUsers {
		result = append(result, mappedUser)
	}

	return result, relatedSecrets, nil
}

func getDeleteAfterDate(user *atlasv2.CloudDatabaseUser) string {
	if user.DeleteAfterDate != nil {
		return user.DeleteAfterDate.String()
	}
	return ""
}

func buildUserSecret(resourceName string, targetNamespace string, dictionary map[string]string) *corev1.Secret {
	secret := secrets.NewAtlasSecret(resourceName, targetNamespace, map[string][]byte{
		secrets.PasswordField: []byte(""),
	}, dictionary)
	return secret
}

func convertUserScopes(user *atlasv2.CloudDatabaseUser) []atlasV1.ScopeSpec {
	result := make([]atlasV1.ScopeSpec, 0, len(user.Scopes))
	for _, scope := range user.Scopes {
		result = append(result, atlasV1.ScopeSpec{
			Name: scope.Name,
			Type: atlasV1.ScopeType(scope.Type),
		})
	}
	return result
}

func convertUserRoles(user *atlasv2.CloudDatabaseUser) []atlasV1.RoleSpec {
	if len(user.Roles) == 0 {
		return nil
	}
	result := make([]atlasV1.RoleSpec, 0, len(user.Roles))
	for _, role := range user.Roles {
		result = append(result, atlasV1.RoleSpec{
			RoleName:       role.RoleName,
			DatabaseName:   role.DatabaseName,
			CollectionName: pointer.GetOrDefault(role.CollectionName, ""),
		})
	}
	return result
}

func convertUserLabels(user *atlasv2.CloudDatabaseUser) []common.LabelSpec {
	result := make([]common.LabelSpec, 0, len(user.Labels))
	for _, label := range user.Labels {
		result = append(result, common.LabelSpec{
			Key:   *label.Key,
			Value: *label.Value,
		})
	}
	return result
}

func suggestResourceName(
	projectName string,
	username string,
	mappedDatabaseUsers map[string]*atlasV1.AtlasDatabaseUser,
	dictionary map[string]string,
) string {
	resourceName := resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, username), dictionary)
	_, ok := mappedDatabaseUsers[resourceName]

	for ok {
		suffix := uuid.NewString()[:5]
		resourceName = fmt.Sprintf("%s-%s", resourceName, suffix)

		_, ok = mappedDatabaseUsers[resourceName]
	}

	return resourceName
}
