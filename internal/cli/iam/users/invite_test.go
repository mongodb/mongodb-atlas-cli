// Copyright 2020 MongoDB Inc
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

package users

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestInvite_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockUserCreator(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.AtlasUser{
		Username: "testUser",
	}
	opts := &InviteOpts{
		store:    mockStore,
		username: "testUser",
	}

	user, err := opts.createUserView()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	mockStore.
		EXPECT().
		CreateUser(user).
		Return(expected, nil).
		Times(1)

	err = opts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateAtlasRole(t *testing.T) {
	type test struct {
		input InviteOpts
		want  []mongodbatlas.AtlasRole
	}

	tests := []test{
		{
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER"},
			},
			want: []mongodbatlas.AtlasRole{{
				OrgID:    "5e4e593f70dfbf1010295836",
				RoleName: "ORG_OWNER",
			}},
		},
		{
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
			},
			want: []mongodbatlas.AtlasRole{{
				OrgID:    "5e4e593f70dfbf1010295836",
				RoleName: "ORG_OWNER",
			},
				{
					OrgID:    "5e4e593f70dfbf1010295836",
					RoleName: "ORG_GROUP_CREATOR",
				},
			},
		},
		{
			input: InviteOpts{
				orgRoles:     []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
				projectRoles: []string{"5e4e593f70dfbf1010295836:GROUP_OWNER", "5e4e593f70dfbf1010295836:GROUP_CLUSTER_MANAGER"},
			},
			want: []mongodbatlas.AtlasRole{
				{
					OrgID:    "5e4e593f70dfbf1010295836",
					RoleName: "ORG_OWNER",
				},
				{
					OrgID:    "5e4e593f70dfbf1010295836",
					RoleName: "ORG_GROUP_CREATOR",
				},
				{
					GroupID:  "5e4e593f70dfbf1010295836",
					RoleName: "GROUP_OWNER",
				},
				{
					GroupID:  "5e4e593f70dfbf1010295836",
					RoleName: "GROUP_CLUSTER_MANAGER",
				},
			},
		},
		{
			input: InviteOpts{},
			want:  []mongodbatlas.AtlasRole{},
		},
	}

	for _, tc := range tests {
		got, err := tc.input.createAtlasRole()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCreateUserRole(t *testing.T) {
	type test struct {
		input InviteOpts
		want  []*opsmngr.UserRole
	}

	config.SetService(config.OpsManagerService)

	tests := []test{
		{
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER"},
			},
			want: []*opsmngr.UserRole{{
				OrgID:    "5e4e593f70dfbf1010295836",
				RoleName: "ORG_OWNER",
			}},
		},
		{
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
			},
			want: []*opsmngr.UserRole{{
				OrgID:    "5e4e593f70dfbf1010295836",
				RoleName: "ORG_OWNER",
			},
				{
					OrgID:    "5e4e593f70dfbf1010295836",
					RoleName: "ORG_GROUP_CREATOR",
				},
			},
		},
		{
			input: InviteOpts{
				orgRoles:     []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
				projectRoles: []string{"5e4e593f70dfbf1010295836:GROUP_OWNER", "5e4e593f70dfbf1010295836:GROUP_CLUSTER_MANAGER"},
			},
			want: []*opsmngr.UserRole{
				{
					OrgID:    "5e4e593f70dfbf1010295836",
					RoleName: "ORG_OWNER",
				},
				{
					OrgID:    "5e4e593f70dfbf1010295836",
					RoleName: "ORG_GROUP_CREATOR",
				},
				{
					GroupID:  "5e4e593f70dfbf1010295836",
					RoleName: "GROUP_OWNER",
				},
				{
					GroupID:  "5e4e593f70dfbf1010295836",
					RoleName: "GROUP_CLUSTER_MANAGER",
				},
			},
		},
		{
			input: InviteOpts{},
			want:  []*opsmngr.UserRole{},
		},
	}

	for _, tc := range tests {
		got, err := tc.input.createUserRole()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
