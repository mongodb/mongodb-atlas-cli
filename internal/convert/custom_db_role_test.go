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

package convert

import (
	"testing"

	"github.com/go-test/deep"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestBuildAtlasInheritedRoles(t *testing.T) {
	tests := []struct {
		input []string
		want  []atlasv2.DatabaseInheritedRole
	}{
		{
			input: []string{"admin"},
			want: []atlasv2.DatabaseInheritedRole{
				{
					Role: "admin",
					Db:   "admin",
				},
			},
		},
		{
			input: []string{"admin@test"},
			want: []atlasv2.DatabaseInheritedRole{
				{
					Role: "admin",
					Db:   "test",
				},
			},
		},
		{
			input: []string{"admin@test", "something"},
			want: []atlasv2.DatabaseInheritedRole{
				{
					Role: "admin",
					Db:   "test",
				},
				{
					Role: "something",
					Db:   "admin",
				},
			},
		},
	}

	for _, tc := range tests {
		input := tc.input
		want := tc.want
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := BuildAtlasInheritedRoles(input)
			if diff := deep.Equal(want, got); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestBuildAtlasActions(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []atlasv2.DatabasePrivilegeAction
	}{
		{
			name:  "role",
			input: []string{"clusterName"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Cluster: true,
						},
					},
				},
			},
		},
		{
			name:  "role and fqn",
			input: []string{"clusterName@testdb.collection"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Db:         "testdb",
							Collection: "collection",
						},
					},
				},
			},
		},
		{
			name:  "role and fqn",
			input: []string{"clusterName@testdb.collection.with.dots"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Db:         "testdb",
							Collection: "collection.with.dots",
						},
					},
				},
			},
		},
		{
			name:  "role and fqn",
			input: []string{"clusterName", "name@DATA_LAKE"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Cluster: true,
						},
					},
				},
				{
					Action: "name",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Db: "DATA_LAKE",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := BuildAtlasActions(tc.input)
			if diff := deep.Equal(tc.want, got); diff != nil {
				t.Error(diff)
			}
		})
	}
}
