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

//go:build unit
// +build unit

package customdbroles

import (
	"testing"

	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		5,
		[]string{},
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

func Test_joinActions(t *testing.T) {
	type args struct {
		newActions []mongodbatlas.Action
	}
	tests := []struct {
		name string
		args args
		want []mongodbatlas.Action
	}{
		{
			name: "empty",
			args: args{
				newActions: []mongodbatlas.Action{},
			},
			want: []mongodbatlas.Action{},
		},
		{
			name: "no duplicate",
			args: args{
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
					{
						Action: "TEST2",
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
					Action: "TEST2",
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
			name: "duplicates",
			args: args{
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
					{
						Action: "TEST",
						Resources: []mongodbatlas.Resource{
							{
								Collection: "test",
								Db:         "test1",
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
							Db:         "test1",
						},
						{
							Collection: "test",
							Db:         "test",
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
			got := joinActions(args.newActions)
			assert.ElementsMatch(t, got, want)
		})
	}
}
