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

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201006/admin"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const resourceVersion = "x.y.z"

func Test_convertUserLabels(t *testing.T) {
	t.Run("Can convert user labels from Atlas to the Operator format", func(t *testing.T) {
		atlasUser := &atlasv2.CloudDatabaseUser{
			Labels: []atlasv2.ComponentLabel{
				{
					Key:   pointer.Get("TestKey"),
					Value: pointer.Get("TestValue"),
				},
			},
		}

		expectedLabels := []common.LabelSpec{
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
			Roles: []atlasv2.DatabaseUserRole{
				{
					RoleName:       "TestRole",
					DatabaseName:   "TestDB",
					CollectionName: pointer.Get("TestCollection"),
				},
			},
		}

		expectedRoles := []atlasV1.RoleSpec{
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
		atlasUser := &atlasv2.CloudDatabaseUser{
			Password: pointer.Get("TestPassword"),
			Username: "TestName",
		}

		expectedSecret := &corev1.Secret{
			TypeMeta: v1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s", projectName, atlasUser.Username)),
				Namespace: "TestNamespace",
				Labels: map[string]string{
					secrets.TypeLabelKey: secrets.CredLabelVal,
				},
			},
			Data: map[string][]byte{
				secrets.PasswordField: []byte(""),
			},
		}

		got := buildUserSecret(resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, atlasUser.Username), dictionary), "TestNamespace", dictionary)
		if !reflect.DeepEqual(got, expectedSecret) {
			t.Errorf("buildUserSecret(); \r\n got:%v;s\r\n want:%v", got, expectedSecret)
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
			Labels: []atlasv2.ComponentLabel{
				{
					Key:   pointer.Get("TestLabelKey"),
					Value: pointer.Get("TestLabelValue"),
				},
			},
			LdapAuthType: pointer.Get("TestType"),
			X509Type:     pointer.Get("TestX509"),
			AwsIAMType:   pointer.Get("TestAWSIAMType"),
			GroupId:      "0",
			Roles: []atlasv2.DatabaseUserRole{
				{
					RoleName:       "TestRoleName",
					DatabaseName:   "TestRoleDatabaseName",
					CollectionName: pointer.Get("TestCollectionName"),
				},
			},
			Scopes: []atlasv2.UserScope{
				{
					Name: "TestScopeName",
					Type: "CLUSTER",
				},
			},
			Password: pointer.Get("TestPassword"),
			Username: "TestUsername",
		}

		listOptions := &mongodbatlas.ListOptions{}
		mockUserStore.EXPECT().DatabaseUsers(projectID, listOptions).Return(&atlasv2.PaginatedApiAtlasDatabaseUser{
			Results: []atlasv2.CloudDatabaseUser{
				user,
			},
		}, nil)

		users, relatedSecrets, err := BuildDBUsers(mockUserStore, projectID, projectName, targetNamespace, dictionary, resourceVersion)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expectedUser := &atlasV1.AtlasDatabaseUser{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDatabaseUser",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, user.Username), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: atlasV1.AtlasDatabaseUserSpec{
				Project: common.ResourceRefNamespaced{
					Name:      resources.NormalizeAtlasName(projectName, dictionary),
					Namespace: targetNamespace,
				},
				DatabaseName:    user.DatabaseName,
				DeleteAfterDate: user.DeleteAfterDate.String(),
				Labels: []common.LabelSpec{
					{
						Key:   *user.Labels[0].Key,
						Value: *user.Labels[0].Value,
					},
				},
				Roles: []atlasV1.RoleSpec{
					{
						RoleName:       user.Roles[0].RoleName,
						DatabaseName:   user.Roles[0].DatabaseName,
						CollectionName: *user.Roles[0].CollectionName,
					},
				},
				Scopes: []atlasV1.ScopeSpec{
					{
						Name: user.Scopes[0].Name,
						Type: atlasV1.ScopeType(user.Scopes[0].Type),
					},
				},
				PasswordSecret: &common.ResourceRef{
					Name: relatedSecrets[0].Name,
				},
				Username: user.Username,
				X509Type: *user.X509Type,
			},
			Status: status.AtlasDatabaseUserStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
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

		expectedSecret := secrets.NewAtlasSecret(
			fmt.Sprintf("%v-%v", projectName, user.Username),
			targetNamespace,
			map[string][]byte{
				secrets.PasswordField: []byte(""),
			}, dictionary)
		if !reflect.DeepEqual(relatedSecrets[0], expectedSecret) {
			t.Fatalf("Secret result doesn't match.\r\nexpected: %v\r\ngot %v\r\n", expectedSecret, relatedSecrets[0])
		}
	})

	t.Run("Can build AtlasDatabaseUser when k8s resource name conflicts", func(t *testing.T) {
		atlasUsers := atlasv2.PaginatedApiAtlasDatabaseUser{
			Results: []atlasv2.CloudDatabaseUser{
				{
					DatabaseName:    "TestDB",
					DeleteAfterDate: pointer.Get(time.Now()),
					Labels: []atlasv2.ComponentLabel{
						{
							Key:   pointer.Get("TestLabelKey"),
							Value: pointer.Get("TestLabelValue"),
						},
					},
					LdapAuthType: pointer.Get("TestType"),
					X509Type:     pointer.Get("TestX509"),
					AwsIAMType:   pointer.Get("TestAWSIAMType"),
					GroupId:      "0",
					Roles: []atlasv2.DatabaseUserRole{
						{
							RoleName:       "TestRoleName",
							DatabaseName:   "TestRoleDatabaseName",
							CollectionName: pointer.Get("TestCollectionName"),
						},
					},
					Scopes: []atlasv2.UserScope{
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
					Labels: []atlasv2.ComponentLabel{
						{
							Key:   pointer.Get("TestLabelKey"),
							Value: pointer.Get("TestLabelValue"),
						},
					},
					LdapAuthType: pointer.Get("TestType"),
					X509Type:     pointer.Get("TestX509"),
					AwsIAMType:   pointer.Get("TestAWSIAMType"),
					GroupId:      "0",
					Roles: []atlasv2.DatabaseUserRole{
						{
							RoleName:       "TestRoleName",
							DatabaseName:   "TestRoleDatabaseName",
							CollectionName: pointer.Get("TestCollectionName"),
						},
					},
					Scopes: []atlasv2.UserScope{
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

		listOptions := &mongodbatlas.ListOptions{}

		mockUserStore.EXPECT().DatabaseUsers(projectID, listOptions).Return(&atlasUsers, nil)

		users, relatedSecrets, err := BuildDBUsers(mockUserStore, projectID, projectName, targetNamespace, dictionary, resourceVersion)
		if err != nil {
			t.Fatalf("%v", err)
		}

		assert.NotEqual(t, users[0].Name, users[1].Name)
		assert.NotEqual(t, relatedSecrets[0].Name, relatedSecrets[1].Name)
	})
}
