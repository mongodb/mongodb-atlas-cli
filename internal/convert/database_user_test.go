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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestBuildAtlasRoles(t *testing.T) {
	type test struct {
		name  string
		input []string
		want  []atlasv2.DatabaseUserRole
	}

	tests := []test{
		{
			name:  "one with no db",
			input: []string{"admin"},
			want: []atlasv2.DatabaseUserRole{
				{
					RoleName:     "admin",
					DatabaseName: "admin",
				},
			},
		},
		{
			name:  "one with db",
			input: []string{"admin@test"},
			want: []atlasv2.DatabaseUserRole{
				{
					RoleName:     "admin",
					DatabaseName: "test",
				},
			},
		},
		{
			name:  "one with db one without db",
			input: []string{"admin@test", "something"},
			want: []atlasv2.DatabaseUserRole{
				{
					RoleName:     "admin",
					DatabaseName: "test",
				},
				{
					RoleName:     "something",
					DatabaseName: "admin",
				},
			},
		},
		{
			name:  "one with db and collection",
			input: []string{"admin@db.collection"},
			want: []atlasv2.DatabaseUserRole{
				{
					RoleName:       "admin",
					DatabaseName:   "db",
					CollectionName: pointer.Get("collection"),
				},
			},
		},
		{
			name:  "one with db and collection with multiple dots",
			input: []string{"admin@db.collection.name"},
			want: []atlasv2.DatabaseUserRole{
				{
					RoleName:       "admin",
					DatabaseName:   "db",
					CollectionName: pointer.Get("collection.name"),
				},
			},
		},
	}

	for _, tc := range tests {
		input := tc.input
		want := tc.want
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := BuildAtlasRoles(input)
			if err := deep.Equal(want, got); err != nil {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestBuildAtlasScopes(t *testing.T) {
	type test struct {
		name  string
		input []string
		want  []atlasv2.UserScope
	}

	tests := []test{
		{
			name:  "default to cluster",
			input: []string{"clusterName"},
			want: []atlasv2.UserScope{
				{
					Name: "clusterName",
					Type: "CLUSTER",
				},
			},
		},
		{
			name:  "with cluster type",
			input: []string{"clusterName:CLUSTER"},
			want: []atlasv2.UserScope{
				{
					Name: "clusterName",
					Type: "CLUSTER",
				},
			},
		},
		{
			name:  "default to cluster and a DATA_LAKE",
			input: []string{"clusterName", "name:DATA_LAKE"},
			want: []atlasv2.UserScope{
				{
					Name: "clusterName",
					Type: "CLUSTER",
				},
				{
					Name: "name",
					Type: "DATA_LAKE",
				},
			},
		},
		{
			name:  "data lake",
			input: []string{"name:DATA_LAKE"},
			want: []atlasv2.UserScope{
				{
					Name: "name",
					Type: "DATA_LAKE",
				},
			},
		},
	}

	for _, tc := range tests {
		input := tc.input
		want := tc.want
		name := tc.name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := BuildAtlasScopes(input)
			if err := deep.Equal(want, got); err != nil {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}
