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

//go:build unit

package users

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestInvite_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockUserCreator(ctrl)

	expected := &atlasv2.CloudAppUser{
		Username: "testUser",
	}
	opts := &InviteOpts{
		store:    mockStore,
		username: "testUser",
	}

	user, err := opts.newUserRequest()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	mockStore.
		EXPECT().
		CreateUser(user).
		Return(expected, nil).
		Times(1)

	if err = opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateRoles(t *testing.T) {
	type test struct {
		name  string
		input InviteOpts
		want  []atlasv2.CloudAccessRoleAssignment
	}

	tests := []test{
		{
			name: "one role",
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER"},
			},
			want: []atlasv2.CloudAccessRoleAssignment{{
				OrgId:    pointer.Get("5e4e593f70dfbf1010295836"),
				RoleName: pointer.Get("ORG_OWNER"),
			}},
		},
		{
			name: "multple roles",
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
			},
			want: []atlasv2.CloudAccessRoleAssignment{
				{
					OrgId:    pointer.Get("5e4e593f70dfbf1010295836"),
					RoleName: pointer.Get("ORG_OWNER"),
				},
				{
					OrgId:    pointer.Get("5e4e593f70dfbf1010295836"),
					RoleName: pointer.Get("ORG_GROUP_CREATOR"),
				},
			},
		},
		{
			name: "org and project roles",
			input: InviteOpts{
				orgRoles:     []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
				projectRoles: []string{"5e4e593f70dfbf1010295836:GROUP_OWNER", "5e4e593f70dfbf1010295836:GROUP_CLUSTER_MANAGER"},
			},
			want: []atlasv2.CloudAccessRoleAssignment{
				{
					OrgId:    pointer.Get("5e4e593f70dfbf1010295836"),
					RoleName: pointer.Get("ORG_OWNER"),
				},
				{
					OrgId:    pointer.Get("5e4e593f70dfbf1010295836"),
					RoleName: pointer.Get("ORG_GROUP_CREATOR"),
				},
				{
					GroupId:  pointer.Get("5e4e593f70dfbf1010295836"),
					RoleName: pointer.Get("GROUP_OWNER"),
				},
				{
					GroupId:  pointer.Get("5e4e593f70dfbf1010295836"),
					RoleName: pointer.Get("GROUP_CLUSTER_MANAGER"),
				},
			},
		},
		{
			name:  "empty",
			input: InviteOpts{},
			want:  []atlasv2.CloudAccessRoleAssignment{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.createRoles()
			if err != nil {
				t.Fatalf("Run() unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestNewUserRequestWithPasswordStin(t *testing.T) {
	password := "p@ssw0rd"
	opts := &InviteOpts{
		username: "testUser",
	}

	require.NoError(t, opts.InitInput(bytes.NewReader([]byte(password)))())
	require.NoError(t, opts.Prompt())

	user, err := opts.newUserRequest()
	if err != nil {
		t.Fatalf("newUserRequest() unexpected error: %v", err)
	}

	if user.Password != password {
		t.Fatalf("failed to read password from stream")
	}
}
