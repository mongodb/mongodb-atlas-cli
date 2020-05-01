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

package store

import (
	"strings"
	"testing"

	"go.mongodb.org/ops-manager/opsmngr"
)

func TestStore_apiPath(t *testing.T) {
	t.Run("ops manager", func(t *testing.T) {
		s := &Store{
			service: "ops-manager",
		}
		result := s.apiPath("localhost")
		if !strings.Contains(result, opsmngr.APIPublicV1Path) {
			t.Errorf("apiPath() = %s; want '%s'", result, opsmngr.APIPublicV1Path)
		}
	})
	t.Run("atlas", func(t *testing.T) {
		s := &Store{
			service: "cloud",
		}
		result := s.apiPath("localhost")
		if !strings.Contains(result, atlasAPIPath) {
			t.Errorf("apiPath() = %s; want '%s'", result, atlasAPIPath)
		}
	})
}
