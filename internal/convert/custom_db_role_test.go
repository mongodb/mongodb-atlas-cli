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

package convert

import (
	"testing"

	"github.com/go-test/deep"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestBuildAtlasInheritedRoles(t *testing.T) {
	type test struct {
		input []string
		want  []mongodbatlas.InheritedRole
	}

	tests := []test{
		{input: []string{"admin"}, want: []mongodbatlas.InheritedRole{
			{
				Role: "admin",
				Db:   "admin",
			}},
		},
		{input: []string{"admin@test"}, want: []mongodbatlas.InheritedRole{
			{
				Role: "admin",
				Db:   "test",
			}},
		},
		{input: []string{"admin@test", "something"}, want: []mongodbatlas.InheritedRole{
			{
				Role: "admin",
				Db:   "test",
			},
			{
				Role: "something",
				Db:   "admin",
			}},
		},
	}

	for _, tc := range tests {
		got := BuildAtlasInheritedRoles(tc.input)
		if err := deep.Equal(tc.want, got); err != nil {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestBuildAtlasActions(t *testing.T) {
	type test struct {
		input []string
		want  []mongodbatlas.Action
	}

	cluster := true

	tests := []test{
		{input: []string{"clusterName"}, want: []mongodbatlas.Action{
			{
				Action: "clusterName",
				Resources: []mongodbatlas.Resource{
					{
						Cluster: &cluster,
					},
				},
			},
		}},
		{input: []string{"clusterName@testdb.collection"}, want: []mongodbatlas.Action{
			{
				Action: "clusterName",
				Resources: []mongodbatlas.Resource{
					{
						Db:         "testdb",
						Collection: "collection",
					},
				},
			},
		}},
		{input: []string{"clusterName", "name@DATA_LAKE"}, want: []mongodbatlas.Action{
			{
				Action: "clusterName",
				Resources: []mongodbatlas.Resource{
					{
						Cluster: &cluster,
					},
				},
			},
			{
				Action: "name",
				Resources: []mongodbatlas.Resource{
					{
						Db: "DATA_LAKE",
					},
				},
			},
		}},
	}

	for _, tc := range tests {
		got := BuildAtlasActions(tc.input)
		if err := deep.Equal(tc.want, got); err != nil {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
