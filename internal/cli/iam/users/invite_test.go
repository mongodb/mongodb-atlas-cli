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
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/mocks"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestInvite_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockUserCreator(ctrl)

	expected := &opsmngr.User{
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

func TestCreateUserRole(t *testing.T) {
	type test struct {
		name  string
		input InviteOpts
		want  []*opsmngr.UserRole
	}
	tests := []test{
		{
			name: "one role",
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER"},
			},
			want: []*opsmngr.UserRole{{
				OrgID:    "5e4e593f70dfbf1010295836",
				RoleName: "ORG_OWNER",
			}},
		},
		{
			name: "multiple org roles",
			input: InviteOpts{
				orgRoles: []string{"5e4e593f70dfbf1010295836:ORG_OWNER", "5e4e593f70dfbf1010295836:ORG_GROUP_CREATOR"},
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
			},
		},
		{
			name: "projects and orgs",
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
			name:  "empty",
			input: InviteOpts{},
			want:  []*opsmngr.UserRole{},
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.createUserRole()
			require.NoError(t, err)
			if diff := deep.Equal(tt.want, got); diff != nil {
				t.Fatalf("diff: %v", diff)
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
