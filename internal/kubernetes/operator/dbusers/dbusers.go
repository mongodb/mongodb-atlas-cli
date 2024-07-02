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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const timeFormatISO8601 = "2006-01-02T15:04:05.999Z"

func BuildDBUsers(provider store.OperatorDBUsersStore, projectID, projectName, targetNamespace string, dictionary map[string]string, version string) ([]*akov2.AtlasDatabaseUser, []*corev1.Secret, error) {
	users, err := provider.DatabaseUsers(projectID, &store.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	if len(users.GetResults()) == 0 {
		return nil, nil, nil
	}

	mappedUsers := map[string]*akov2.AtlasDatabaseUser{}
	relatedSecrets := make([]*corev1.Secret, 0, len(users.GetResults()))

	for _, u := range users.GetResults() {
		user := pointer.Get(u)
		resourceName := suggestResourceName(projectName, user.Username, mappedUsers, dictionary)
		labels := convertUserLabels(user)
		roles := convertUserRoles(user)
		if len(roles) == 0 {
			continue
		}
		scopes := convertUserScopes(user)

		mappedUsers[resourceName] = &akov2.AtlasDatabaseUser{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasDatabaseUser",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resourceName,
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: version,
				},
			},
			Spec: akov2.AtlasDatabaseUserSpec{
				Project: akov2common.ResourceRefNamespaced{
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
			Status: akov2status.AtlasDatabaseUserStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		if user.GetX509Type() != "MANAGED" {
			secret := buildUserSecret(resourceName, targetNamespace, projectID, projectName, dictionary)
			relatedSecrets = append(relatedSecrets, secret)

			mappedUsers[resourceName].Spec.PasswordSecret = &akov2common.ResourceRef{
				Name: resources.NormalizeAtlasName(secret.Name, dictionary),
			}
		}
	}

	result := make([]*akov2.AtlasDatabaseUser, 0, len(users.GetResults()))
	for _, mappedUser := range mappedUsers {
		result = append(result, mappedUser)
	}

	return result, relatedSecrets, nil
}

func getDeleteAfterDate(user *atlasv2.CloudDatabaseUser) string {
	if user.DeleteAfterDate != nil {
		return user.DeleteAfterDate.Format(timeFormatISO8601)
	}
	return ""
}

func buildUserSecret(resourceName, targetNamespace, projectID, projectName string, dictionary map[string]string) *corev1.Secret {
	secret := secrets.NewAtlasSecretBuilder(resourceName, targetNamespace, dictionary).
		WithData(map[string][]byte{secrets.PasswordField: []byte("")}).
		WithProjectLabels(projectID, projectName).
		Build()
	return secret
}

func convertUserScopes(user *atlasv2.CloudDatabaseUser) []akov2.ScopeSpec {
	result := make([]akov2.ScopeSpec, 0, len(user.GetScopes()))
	for _, scope := range user.GetScopes() {
		result = append(result, akov2.ScopeSpec{
			Name: scope.Name,
			Type: akov2.ScopeType(scope.Type),
		})
	}
	return result
}

func convertUserRoles(user *atlasv2.CloudDatabaseUser) []akov2.RoleSpec {
	if len(user.GetRoles()) == 0 {
		return nil
	}
	result := make([]akov2.RoleSpec, 0, len(user.GetRoles()))
	for _, role := range user.GetRoles() {
		result = append(result, akov2.RoleSpec{
			RoleName:       role.RoleName,
			DatabaseName:   role.DatabaseName,
			CollectionName: role.GetCollectionName(),
		})
	}
	return result
}

func convertUserLabels(user *atlasv2.CloudDatabaseUser) []akov2common.LabelSpec {
	result := make([]akov2common.LabelSpec, 0, len(user.GetLabels()))
	for _, label := range user.GetLabels() {
		result = append(result, akov2common.LabelSpec{
			Key:   *label.Key,
			Value: *label.Value,
		})
	}
	return result
}

func suggestResourceName(
	projectName string,
	username string,
	mappedDatabaseUsers map[string]*akov2.AtlasDatabaseUser,
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
