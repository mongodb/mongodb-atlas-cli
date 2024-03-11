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
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestBuildOMRoles(t *testing.T) {
	type test struct {
		input []string
		want  []*opsmngr.Role
	}

	tests := []test{
		{
			input: []string{"admin"},
			want: []*opsmngr.Role{
				{
					Role:     "admin",
					Database: "admin",
				},
			},
		},
		{
			input: []string{"admin@test"},
			want: []*opsmngr.Role{
				{
					Role:     "admin",
					Database: "test",
				},
			},
		},
		{
			input: []string{"admin@test", "something"},
			want: []*opsmngr.Role{
				{
					Role:     "admin",
					Database: "test",
				},
				{
					Role:     "something",
					Database: "admin",
				},
			},
		},
	}

	for _, tc := range tests {
		input := tc.input
		want := tc.want
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := BuildOMRoles(input)
			if err := deep.Equal(want, got); err != nil {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}
