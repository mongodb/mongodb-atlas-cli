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
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_convertUserLabels(t *testing.T) {
	t.Run("Can convert user labels from Atlas to the Operator format", func(t *testing.T) {
		atlasUser := &mongodbatlas.DatabaseUser{
			Labels: []mongodbatlas.Label{
				{
					Key:   "TestKey",
					Value: "TestValue",
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
		atlasUser := &mongodbatlas.DatabaseUser{
			Roles: []mongodbatlas.Role{
				{
					RoleName:       "TestRole",
					DatabaseName:   "TestDB",
					CollectionName: "TestCollection",
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
	t.Run("Can build user secret WITHOUT credentials", func(t *testing.T) {
		projectName := "TestProject-1"
		atlasUser := &mongodbatlas.DatabaseUser{
			Password: "TestPassword",
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
		if got := buildUserSecret(projectName, atlasUser, "TestNamespace", false); !reflect.DeepEqual(got, expectedSecret) {
			t.Errorf("buildUserSecret(); \r\n got:%v;s\r\n want:%v", got, expectedSecret)
		}
	})

	t.Run("Can build user secret WITH credentials", func(t *testing.T) {
		projectName := "TestProject-2"
		atlasUser := &mongodbatlas.DatabaseUser{
			Password: "TestPassword",
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
				secrets.PasswordField: []byte(atlasUser.Password),
			},
		}
		if got := buildUserSecret(projectName, atlasUser, "TestNamespace", true); !reflect.DeepEqual(got, expectedSecret) {
			t.Errorf("buildUserSecret(); \r\n got:%v;s\r\n want:%v", got, expectedSecret)
		}
	})
}

func TestBuildDBUsers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockUserStore := mocks.NewMockDatabaseUserLister(ctl)

	t.Run("Can build AtlasDatabaseUser from AtlasUser WITHOUT credentials", func(t *testing.T) {
		user := mongodbatlas.DatabaseUser{
			DatabaseName:    "TestDB",
			DeleteAfterDate: "2022",
			Labels: []mongodbatlas.Label{
				{
					Key:   "TestLabelKey",
					Value: "TestLabelValue",
				},
			},
			LDAPAuthType: "TestType",
			X509Type:     "TestX509",
			AWSIAMType:   "TestAWSIAMType",
			GroupID:      "0",
			Roles: []mongodbatlas.Role{
				{
					RoleName:       "TestRoleName",
					DatabaseName:   "TestRoleDatabaseName",
					CollectionName: "TestCollectionName",
				},
			},
			Scopes: []mongodbatlas.Scope{
				{
					Name: "TestScopeName",
					Type: "CLUSTER",
				},
			},
			Password: "TestPassword",
			Username: "TestUsername",
		}

		projectID := "0"
		projectName := "projectName-1"
		targetNamespace := "TestNamespace-1"

		listOptions := &mongodbatlas.ListOptions{}

		mockUserStore.EXPECT().DatabaseUsers(projectID, listOptions).Return([]mongodbatlas.DatabaseUser{
			user,
		}, nil)

		users, relatedSecrets, err := BuildDBUsers(mockUserStore, projectID, projectName, targetNamespace, false)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expectedUser := &atlasV1.AtlasDatabaseUser{
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
				Labels: []common.LabelSpec{
					{
						Key:   user.Labels[0].Key,
						Value: user.Labels[0].Value,
					},
				},
				Roles: []atlasV1.RoleSpec{
					{
						RoleName:       user.Roles[0].RoleName,
						DatabaseName:   user.Roles[0].DatabaseName,
						CollectionName: user.Roles[0].CollectionName,
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
				X509Type: user.X509Type,
			},
			Status: status.AtlasDatabaseUserStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		if !reflect.DeepEqual(users[0], expectedUser) {
			t.Fatalf("User result doesn't match.\r\nexpected: %v,\r\ngot: %v\r\n", expectedUser, users[0])
		}

		expectedSecret := secrets.NewAtlasSecret(
			fmt.Sprintf("%v-%v", projectName, user.Username),
			targetNamespace,
			map[string][]byte{
				secrets.PasswordField: []byte(""),
			})
		if !reflect.DeepEqual(relatedSecrets[0], expectedSecret) {
			t.Fatalf("Secret result doesn't match.\r\nexpected: %v\r\ngot %v\r\n", expectedSecret, relatedSecrets[0])
		}
	})

	t.Run("Can build AtlasDatabaseUser from AtlasUser WITH credentials", func(t *testing.T) {
		user := mongodbatlas.DatabaseUser{
			DatabaseName:    "TestDB",
			DeleteAfterDate: "2022",
			Labels: []mongodbatlas.Label{
				{
					Key:   "TestLabelKey",
					Value: "TestLabelValue",
				},
			},
			LDAPAuthType: "TestType",
			X509Type:     "TestX509",
			AWSIAMType:   "TestAWSIAMType",
			GroupID:      "0",
			Roles: []mongodbatlas.Role{
				{
					RoleName:       "TestRoleName",
					DatabaseName:   "TestRoleDatabaseName",
					CollectionName: "TestCollectionName",
				},
			},
			Scopes: []mongodbatlas.Scope{
				{
					Name: "TestScopeName",
					Type: "CLUSTER",
				},
			},
			Password: "TestPassword",
			Username: "TestUsername",
		}

		projectID := "0"
		projectName := "projectName-2"
		targetNamespace := "TestNamespace-2"

		listOptions := &mongodbatlas.ListOptions{}
		mockUserStore.EXPECT().DatabaseUsers(projectID, listOptions).Return([]mongodbatlas.DatabaseUser{
			user,
		}, nil)

		users, relatedSecrets, err := BuildDBUsers(mockUserStore, projectID, projectName, targetNamespace, true)
		if err != nil {
			t.Fatalf("%v", err)
		}

		expectedUser := &atlasV1.AtlasDatabaseUser{
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
				Labels: []common.LabelSpec{
					{
						Key:   user.Labels[0].Key,
						Value: user.Labels[0].Value,
					},
				},
				Roles: []atlasV1.RoleSpec{
					{
						RoleName:       user.Roles[0].RoleName,
						DatabaseName:   user.Roles[0].DatabaseName,
						CollectionName: user.Roles[0].CollectionName,
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
				X509Type: user.X509Type,
			},
			Status: status.AtlasDatabaseUserStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		if !reflect.DeepEqual(users[0], expectedUser) {
			t.Fatalf("User result doesn't match.\r\nexpected: %v,\r\ngot: %v\r\n", expectedUser, users[0])
		}

		expectedSecret := secrets.NewAtlasSecret(
			fmt.Sprintf("%v-%v", projectName, user.Username),
			targetNamespace,
			map[string][]byte{
				secrets.PasswordField: []byte(user.Password),
			})
		if !reflect.DeepEqual(relatedSecrets[0], expectedSecret) {
			t.Fatalf("Secret result doesn't match.\r\nexpected: %v\r\ngot %v\r\n", expectedSecret, relatedSecrets[0])
		}
	})
}
