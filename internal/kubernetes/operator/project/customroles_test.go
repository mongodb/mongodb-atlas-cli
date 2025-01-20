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

//nolint:all
package project

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/status"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBuildCustomRoles(t *testing.T) {
	projectID := "pid-1"
	projectName := "p-1"
	targetNamespace := "n-1"
	credentialName := "creds"
	type args struct {
		provider store.DatabaseRoleLister
		request  CustomRolesRequest
	}
	tests := []struct {
		name         string
		args         args
		rolesInAtlas []atlasv2.UserCustomDBRole
		errInAtlas   error
		want         []akov2.AtlasCustomRole
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name: "Should return err if provider return error",
			args: args{
				request: CustomRolesRequest{
					ProjectName:     projectName,
					ProjectID:       projectID,
					TargetNamespace: targetNamespace,
					Credentials:     credentialName,
					Version:         "v2.6.0",
					IsIndependent:   false,
					Dict:            resources.AtlasNameToKubernetesName(),
				},
			},
			errInAtlas: errors.New("no roles found"),
			want:       nil,
			wantErr: func(_ assert.TestingT, err error, i ...any) bool {
				return true
			},
		},
		{
			name: "Should return nil if provider returned empty list of custom roles",
			args: args{
				request: CustomRolesRequest{
					ProjectName:     projectName,
					ProjectID:       projectID,
					TargetNamespace: targetNamespace,
					Credentials:     credentialName,
					Version:         "v2.6.0",
					IsIndependent:   false,
					Dict:            resources.AtlasNameToKubernetesName(),
				},
			},
			rolesInAtlas: nil,
			errInAtlas:   nil,
			want:         nil,
			wantErr: func(_ assert.TestingT, err error, i ...any) bool {
				return false
			},
		},
		{
			name: "Should return AKO custom roles if provider returned custom roles",
			args: args{
				request: CustomRolesRequest{
					ProjectName:     projectName,
					ProjectID:       projectID,
					TargetNamespace: targetNamespace,
					Credentials:     credentialName,
					Version:         "v2.6.0",
					IsIndependent:   false,
					Dict:            resources.AtlasNameToKubernetesName(),
				},
			},
			rolesInAtlas: []atlasv2.UserCustomDBRole{
				{
					Actions: &([]atlasv2.DatabasePrivilegeAction{
						{
							Action: "test",
							Resources: &([]atlasv2.DatabasePermittedNamespaceResource{
								{
									Cluster:    false,
									Collection: "c-1",
									Db:         "d-1",
								},
							}),
						},
					}),
					InheritedRoles: &([]atlasv2.DatabaseInheritedRole{
						{
							Db:   "d-1",
							Role: "ADMIN",
						},
					}),
					RoleName: "r-1",
				},
			},
			errInAtlas: nil,
			want: []akov2.AtlasCustomRole{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasCustomRole",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      projectName + "%-custom-role-r-1",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "v2.6.0",
						},
					},
					Spec: akov2.AtlasCustomRoleSpec{
						Role: akov2.CustomRole{
							Name: "r-1",
							InheritedRoles: []akov2.Role{
								{
									Name:     "ADMIN",
									Database: "d-1",
								},
							},
							Actions: []akov2.Action{
								{
									Name: "test",
									Resources: []akov2.Resource{
										{
											Cluster:    pointer.Get(false),
											Database:   pointer.Get("d-1"),
											Collection: pointer.Get("c-1"),
										},
									},
								},
							},
						},
						ProjectRef: &akov2common.ResourceRefNamespaced{
							Name:      projectName,
							Namespace: targetNamespace,
						},
					},
					Status: akov2status.AtlasCustomRoleStatus{},
				},
			},
			wantErr: func(_ assert.TestingT, err error, i ...any) bool {
				return false
			},
		},
		{
			name: "Should return AKO custom roles if provider returned custom roles, as independent",
			args: args{
				request: CustomRolesRequest{
					ProjectName:     projectName,
					ProjectID:       projectID,
					TargetNamespace: targetNamespace,
					Credentials:     credentialName,
					Version:         "v2.6.0",
					IsIndependent:   true,
					Dict:            resources.AtlasNameToKubernetesName(),
				},
			},
			rolesInAtlas: []atlasv2.UserCustomDBRole{
				{
					Actions: &([]atlasv2.DatabasePrivilegeAction{
						{
							Action: "test",
							Resources: &([]atlasv2.DatabasePermittedNamespaceResource{
								{
									Cluster:    false,
									Collection: "c-1",
									Db:         "d-1",
								},
							}),
						},
					}),
					InheritedRoles: &([]atlasv2.DatabaseInheritedRole{
						{
							Db:   "d-1",
							Role: "ADMIN",
						},
					}),
					RoleName: "r-1",
				},
			},
			errInAtlas: nil,
			want: []akov2.AtlasCustomRole{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasCustomRole",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      projectName + "%-custom-role-r-1",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "v2.6.0",
						},
					},
					Spec: akov2.AtlasCustomRoleSpec{
						Role: akov2.CustomRole{
							Name: "r-1",
							InheritedRoles: []akov2.Role{
								{
									Name:     "ADMIN",
									Database: "d-1",
								},
							},
							Actions: []akov2.Action{
								{
									Name: "test",
									Resources: []akov2.Resource{
										{
											Cluster:    pointer.Get(false),
											Database:   pointer.Get("d-1"),
											Collection: pointer.Get("c-1"),
										},
									},
								},
							},
						},
						ExternalProjectIDRef: &akov2.ExternalProjectReference{ID: projectID},
					},
					Status: akov2status.AtlasCustomRoleStatus{},
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...any) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			crStore := mocks.NewMockOperatorProjectStore(c)
			crStore.EXPECT().DatabaseRoles(projectID).Return(tt.rolesInAtlas, tt.errInAtlas)

			got, err := BuildCustomRoles(crStore, tt.args.request)
			if !tt.wantErr(t, err, fmt.Sprintf("BuildCustomRoles(%v, %v)", tt.args.provider, tt.args.request)) {
				return
			}
			assert.Equalf(t, tt.want, got, "BuildCustomRoles(%v, %v)", tt.args.provider, tt.args.request)
		})
	}
}
