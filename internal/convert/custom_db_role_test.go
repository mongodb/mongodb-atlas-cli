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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestBuildAtlasInheritedRoles(t *testing.T) {
	type test struct {
		input []string
		want  []atlasv2.DatabaseInheritedRole
	}

	tests := []test{
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
			if err := deep.Equal(want, got); err != nil {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestBuildAtlasActions(t *testing.T) {
	type test struct {
		input []string
		want  []atlasv2.DatabasePrivilegeAction
	}

	cluster := true

	testdb := "testdb"
	collection := "collection"
	datalake := "DATA_LAKE"

	tests := []test{
		{
			input: []string{"clusterName"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: []atlasv2.DatabasePermittedNamespaceResource{
						{
							Cluster: cluster,
						},
					},
				},
			},
		},
		{
			input: []string{"clusterName@testdb.collection"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: []atlasv2.DatabasePermittedNamespaceResource{
						{
							Db:         testdb,
							Collection: collection,
						},
					},
				},
			},
		},
		{
			input: []string{"clusterName", "name@DATA_LAKE"},
			want: []atlasv2.DatabasePrivilegeAction{
				{
					Action: "clusterName",
					Resources: []atlasv2.DatabasePermittedNamespaceResource{
						{
							Cluster: cluster,
						},
					},
				},
				{
					Action: "name",
					Resources: []atlasv2.DatabasePermittedNamespaceResource{
						{
							Db: datalake,
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		input := tc.input
		want := tc.want
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := BuildAtlasActions(input)
			if err := deep.Equal(want, got); err != nil {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}
