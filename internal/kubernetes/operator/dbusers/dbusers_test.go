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

//go:build unit

package dbusers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const resourceVersion = "x.y.z"

func Test_convertUserLabels(t *testing.T) {
	t.Run("Can convert user labels from Atlas to the Operator format", func(t *testing.T) {
		atlasUser := &atlasv2.CloudDatabaseUser{
			Labels: &[]atlasv2.ComponentLabel{
				{
					Key:   pointer.Get("TestKey"),
					Value: pointer.Get("TestValue"),
				},
			},
		}

		expectedLabels := []akov2common.LabelSpec{
			{
				Key:   "TestKey",
				Value: "TestValue",
			},
		}

		if got := convertUserLabels(atlasUser); !reflect.DeepEqual(got, expectedLabels) {
			t.Errorf("convertUserLabels() = %v, want %v", got, expectedLabels)
		}
	})
}

func Test_convertUserRoles(t *testing.T) {
	t.Run("Can convert user labels from Atlas to the Operator format", func(t *testing.T) {
		atlasUser := &atlasv2.CloudDatabaseUser{
			Roles: &[]atlasv2.DatabaseUserRole{
				{
					RoleName:       "TestRole",
					DatabaseName:   "TestDB",
					CollectionName: pointer.Get("TestCollection"),
				},
			},
		}

		expectedRoles := []akov2.RoleSpec{
			{
				RoleName:       "TestRole",
				DatabaseName:   "TestDB",
				CollectionName: "TestCollection",
			},
		}
		if got := convertUserRoles(atlasUser); !reflect.DeepEqual(got, expectedRoles) {
			t.Errorf("convertUserRoles() = %v, want %v", got, expectedRoles)
		}
	})
}

func Test_buildUserSecret(t *testing.T) {
	dictionary := resources.AtlasNameToKubernetesName()
	t.Run("Can build user secret WITHOUT credentials", func(t *testing.T) {
		projectName := "TestProject-1"
		projectID := "123"
		atlasUser := &atlasv2.CloudDatabaseUser{
			Password: pointer.Get("TestPassword"),
			Username: "TestName",
		}

		expectedSecret := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s", projectName, atlasUser.Username)),
				Namespace: "TestNamespace",
				Labels: map[string]string{
					secrets.TypeLabelKey:        secrets.CredLabelVal,
					secrets.ProjectIDLabelKey:   resources.NormalizeAtlasName(projectID, dictionary),
					secrets.ProjectNameLabelKey: resources.NormalizeAtlasName(projectName, dictionary),
				},
			},
			Data: map[string][]byte{
				secrets.PasswordField: []byte(""),
			},
		}

		got := buildUserSecret(resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, atlasUser.Username), dictionary), "TestNamespace", projectID, projectName, dictionary)
		if diff := deep.Equal(expectedSecret, got); diff != nil {
			t.Fatalf("buildUserSecret() mismatch: %v", diff)
		}
	})
}

func TestBuildDBUsers(t *testing.T) {
	ctl := gomock.NewController(t)
	mockUserStore := mocks.NewMockDatabaseUserLister(ctl)
	dictionary := resources.AtlasNameToKubernetesName()

	projectID := "0"
	projectName := "projectName-1"
	targetNamespace := "TestNamespace-1"

	t.Run("Can build AtlasDatabaseUser from AtlasUser WITHOUT credentials", func(t *testing.T) {
		user := atlasv2.CloudDatabaseUser{
			DatabaseName:    "TestDB",
			DeleteAfterDate: pointer.Get(time.Now()),
			Labels: &[]atlasv2.ComponentLabel{
				{
					Key:   pointer.Get("TestLabelKey"),
					Value: pointer.Get("TestLabelValue"),
				},
			},
			LdapAuthType: pointer.Get("TestType"),
			X509Type:     pointer.Get("TestX509"),
			AwsIAMType:   pointer.Get("TestAWSIAMType"),
			GroupId:      "0",
			Roles: &[]atlasv2.DatabaseUserRole{
				{
					RoleName:       "TestRoleName",
					DatabaseName:   "TestRoleDatabaseName",
					CollectionName: pointer.Get("TestCollectionName"),
				},
			},
			Scopes: &[]atlasv2.UserScope{
				{
					Name: "TestScopeName",
					Type: "CLUSTER",
				},
			},
			Password: pointer.Get("TestPassword"),
			Username: "TestUsername",
		}

		listOptions := &store.ListOptions{}
		mockUserStore.EXPECT().DatabaseUsers(projectID, listOptions).Return(&atlasv2.PaginatedApiAtlasDatabaseUser{
			Results: &[]atlasv2.CloudDatabaseUser{
				user,
			},
		}, nil)

		users, relatedSecrets, err := BuildDBUsers(mockUserStore, projectID, projectName, targetNamespace, dictionary, resourceVersion)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expectedUser := &akov2.AtlasDatabaseUser{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasDatabaseUser",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, user.Username), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: akov2.AtlasDatabaseUserSpec{
				Project: akov2common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(projectName, dictionary),
					Namespace: targetNamespace,
				},
				DatabaseName:    user.DatabaseName,
				DeleteAfterDate: user.DeleteAfterDate.Format(timeFormatISO8601),
				Labels: []akov2common.LabelSpec{
					{
						Key:   *user.GetLabels()[0].Key,
						Value: *user.GetLabels()[0].Value,
					},
				},
				Roles: []akov2.RoleSpec{
					{
						RoleName:       user.GetRoles()[0].RoleName,
						DatabaseName:   user.GetRoles()[0].DatabaseName,
						CollectionName: *user.GetRoles()[0].CollectionName,
					},
				},
				Scopes: []akov2.ScopeSpec{
					{
						Name: user.GetScopes()[0].Name,
						Type: akov2.ScopeType(user.GetScopes()[0].Type),
					},
				},
				PasswordSecret: &akov2common.ResourceRef{
					Name: relatedSecrets[0].Name,
				},
				Username: user.Username,
				X509Type: *user.X509Type,
			},
			Status: akov2status.AtlasDatabaseUserStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		if !reflect.DeepEqual(users[0], expectedUser) {
			ed, err := json.MarshalIndent(expectedUser, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			gd, err := json.MarshalIndent(users[0], "", " ")
			if err != nil {
				t.Fatal(err)
			}
			t.Fatalf("User result doesn't match.\r\nexpected: %v,\r\ngot: %v\r\n", string(ed), string(gd))
		}

		expectedSecret := secrets.NewAtlasSecretBuilder(
			fmt.Sprintf("%v-%v", projectName, user.Username),
			targetNamespace,
			dictionary,
		).WithData(map[string][]byte{
			secrets.PasswordField: []byte(""),
		}).WithProjectLabels(projectID, projectName).Build()

		if diff := deep.Equal(relatedSecrets[0], expectedSecret); diff != nil {
			t.Fatalf("Secret mismatch: %v", diff)
		}
	})

	t.Run("Can build AtlasDatabaseUser when k8s resource name conflicts", func(t *testing.T) {
		atlasUsers := atlasv2.PaginatedApiAtlasDatabaseUser{
			Results: &[]atlasv2.CloudDatabaseUser{
				{
					DatabaseName:    "TestDB",
					DeleteAfterDate: pointer.Get(time.Now()),
					Labels: &[]atlasv2.ComponentLabel{
						{
							Key:   pointer.Get("TestLabelKey"),
							Value: pointer.Get("TestLabelValue"),
						},
					},
					LdapAuthType: pointer.Get("TestType"),
					X509Type:     pointer.Get("TestX509"),
					AwsIAMType:   pointer.Get("TestAWSIAMType"),
					GroupId:      "0",
					Roles: &[]atlasv2.DatabaseUserRole{
						{
							RoleName:       "TestRoleName",
							DatabaseName:   "TestRoleDatabaseName",
							CollectionName: pointer.Get("TestCollectionName"),
						},
					},
					Scopes: &[]atlasv2.UserScope{
						{
							Name: "TestScopeName",
							Type: "CLUSTER",
						},
					},
					Password: pointer.Get("TestPassword"),
					Username: "TestUsername",
				},
				{
					DatabaseName:    "TestDB",
					DeleteAfterDate: pointer.Get(time.Now()),
					Labels: &[]atlasv2.ComponentLabel{
						{
							Key:   pointer.Get("TestLabelKey"),
							Value: pointer.Get("TestLabelValue"),
						},
					},
					LdapAuthType: pointer.Get("TestType"),
					X509Type:     pointer.Get("TestX509"),
					AwsIAMType:   pointer.Get("TestAWSIAMType"),
					GroupId:      "0",
					Roles: &[]atlasv2.DatabaseUserRole{
						{
							RoleName:       "TestRoleName",
							DatabaseName:   "TestRoleDatabaseName",
							CollectionName: pointer.Get("TestCollectionName"),
						},
					},
					Scopes: &[]atlasv2.UserScope{
						{
							Name: "TestScopeName",
							Type: "CLUSTER",
						},
					},
					Password: pointer.Get("TestPassword"),
					Username: "testUsername",
				},
			},
		}

		listOptions := &store.ListOptions{}
		mockUserStore.EXPECT().DatabaseUsers(projectID, listOptions).Return(&atlasUsers, nil)

		users, relatedSecrets, err := BuildDBUsers(mockUserStore, projectID, projectName, targetNamespace, dictionary, resourceVersion)
		if err != nil {
			t.Fatalf("%v", err)
		}

		assert.NotEqual(t, users[0].Name, users[1].Name)
		assert.NotEqual(t, relatedSecrets[0].Name, relatedSecrets[1].Name)
	})
}
