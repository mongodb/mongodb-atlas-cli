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

package customdbroles

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
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
		existingActions []atlasv2.DatabasePrivilegeAction
		newActions      []atlasv2.DatabasePrivilegeAction
	}

	test1 := "test1"
	test2 := "test2"

	tests := []struct {
		name string
		args args
		want []atlasv2.DatabasePrivilegeAction
	}{
		{
			name: "empty",
			args: args{
				existingActions: []atlasv2.DatabasePrivilegeAction{},
				newActions:      []atlasv2.DatabasePrivilegeAction{},
			},
			want: []atlasv2.DatabasePrivilegeAction{},
		},
		{
			name: "no new actions",
			args: args{
				existingActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test1,
								Db:         test1,
							},
						},
					},
				},
				newActions: []atlasv2.DatabasePrivilegeAction{},
			},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "TEST",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test1,
							Db:         test1,
						},
					},
				},
			},
		},
		{
			name: "different actions",
			args: args{
				existingActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test1,
								Db:         test1,
							},
						},
					},
				},
				newActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "NEW",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test1,
								Db:         test1,
							},
						},
					},
				},
			},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "TEST",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test1,
							Db:         test1,
						},
					},
				},
				{
					Action: "NEW",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test1,
							Db:         test1,
						},
					},
				},
			},
		},
		{
			name: "merge",
			args: args{
				existingActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test1,
								Db:         test2,
							},
						},
					},
				},
				newActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test1,
								Db:         test1,
							},
						},
					},
				},
			},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "TEST",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test1,
							Db:         test1,
						},
						{
							Collection: test1,
							Db:         test2,
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
		newActions []atlasv2.DatabasePrivilegeAction
	}

	test3 := "test3"
	test4 := "test4"

	tests := []struct {
		name string
		args args
		want []atlasv2.DatabasePrivilegeAction
	}{
		{
			name: "empty",
			args: args{
				newActions: []atlasv2.DatabasePrivilegeAction{},
			},
			want: []atlasv2.DatabasePrivilegeAction{},
		},
		{
			name: "no duplicate",
			args: args{
				newActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test3,
								Db:         test3,
							},
						},
					},
					{
						Action: "TEST2",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test3,
								Db:         test3,
							},
						},
					},
				},
			},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "TEST",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test3,
							Db:         test3,
						},
					},
				},
				{
					Action: "TEST2",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test3,
							Db:         test3,
						},
					},
				},
			},
		},
		{
			name: "duplicates",
			args: args{
				newActions: []atlasv2.DatabasePrivilegeAction{
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test3,
								Db:         test3,
							},
						},
					},
					{
						Action: "TEST",
						Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
							{
								Collection: test3,
								Db:         test4,
							},
						},
					},
				},
			},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "TEST",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: test3,
							Db:         test4,
						},
						{
							Collection: test3,
							Db:         test3,
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
