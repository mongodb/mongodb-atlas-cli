// Copyright 2021 MongoDB Inc
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

// +build unit

package customdbroles

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseRoleUpdater(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.CustomDBRole{}

	updateOpts := &UpdateOpts{
		store: mockStore,
	}

	t.Run("default", func(t *testing.T) {
		mockStore.
			EXPECT().
			UpdateDatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName, updateOpts.newCustomDBRole(expected)).Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
	t.Run("with append", func(t *testing.T) {
		updateOpts.append = true
		mockStore.
			EXPECT().
			DatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName).Return(expected, nil).
			Times(1)
		mockStore.
			EXPECT().
			UpdateDatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName, updateOpts.newCustomDBRole(expected)).Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Privilege, flag.InheritedRole, flag.Append},
	)
}

func Test_appendActions(t *testing.T) {
	type args struct {
		existingActions []mongodbatlas.Action
		newActions      []mongodbatlas.Action
	}
	tests := []struct {
		name string
		args args
		want []mongodbatlas.Action
	}{
		{
			name: "empty",
			args: args{
				existingActions: []mongodbatlas.Action{},
				newActions:      []mongodbatlas.Action{},
			},
			want: []mongodbatlas.Action{},
		},
		{
			name: "no new actions",
			args: args{
				existingActions: []mongodbatlas.Action{
					{
						Action: "TEST",
						Resources: []mongodbatlas.Resource{
							{
								Collection: "test",
								Db:         "test",
							},
						},
					},
				},
				newActions: []mongodbatlas.Action{},
			},
			want: []mongodbatlas.Action{
				{
					Action: "TEST",
					Resources: []mongodbatlas.Resource{
						{
							Collection: "test",
							Db:         "test",
						},
					},
				},
			},
		},
		{
			name: "different actions",
			args: args{
				existingActions: []mongodbatlas.Action{
					{
						Action: "TEST",
						Resources: []mongodbatlas.Resource{
							{
								Collection: "test",
								Db:         "test",
							},
						},
					},
				},
				newActions: []mongodbatlas.Action{
					{
						Action: "NEW",
						Resources: []mongodbatlas.Resource{
							{
								Collection: "test",
								Db:         "test",
							},
						},
					},
				},
			},
			want: []mongodbatlas.Action{
				{
					Action: "TEST",
					Resources: []mongodbatlas.Resource{
						{
							Collection: "test",
							Db:         "test",
						},
					},
				},
				{
					Action: "NEW",
					Resources: []mongodbatlas.Resource{
						{
							Collection: "test",
							Db:         "test",
						},
					},
				},
			},
		},
		{
			name: "merge",
			args: args{
				existingActions: []mongodbatlas.Action{
					{
						Action: "TEST",
						Resources: []mongodbatlas.Resource{
							{
								Collection: "test",
								Db:         "test2",
							},
						},
					},
				},
				newActions: []mongodbatlas.Action{
					{
						Action: "TEST",
						Resources: []mongodbatlas.Resource{
							{
								Collection: "test",
								Db:         "test",
							},
						},
					},
				},
			},
			want: []mongodbatlas.Action{
				{
					Action: "TEST",
					Resources: []mongodbatlas.Resource{
						{
							Collection: "test",
							Db:         "test",
						},
						{
							Collection: "test",
							Db:         "test2",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		args := tt.args
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			got := appendActions(args.existingActions, args.newActions)
			assert.ElementsMatch(t, got, want)
		})
	}
}
