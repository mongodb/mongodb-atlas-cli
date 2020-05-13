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

package search_test

import (
	"testing"

	"github.com/mongodb/mongocli/internal/fixture"
	"github.com/mongodb/mongocli/internal/search"
)

func TestStringInSlice(t *testing.T) {
	s := []string{"a", "b", "c"}
	t.Run("value exists", func(t *testing.T) {
		if !search.StringInSlice(s, "b") {
			t.Error("StringInSlice() should find the value")
		}
	})

	t.Run("value not exists", func(t *testing.T) {
		if search.StringInSlice(s, "d") {
			t.Error("StringInSlice() should not find the value")
		}
	})
}

func TestClusterExists(t *testing.T) {
	t.Run("value exists", func(t *testing.T) {
		if !search.ClusterExists(fixture.AutomationConfig(), "myReplicaSet") {
			t.Error("ClusterExists() should find the value")
		}
	})

	t.Run("value not exists", func(t *testing.T) {
		if search.ClusterExists(fixture.AutomationConfig(), "X") {
			t.Error("StringInSlice() should not find the value")
		}
	})
}
