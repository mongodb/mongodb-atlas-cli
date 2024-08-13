// Copyright 2024 MongoDB Inc
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

package set

import "testing"

func Test_Add(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)

	if len(s) != 1 {
		t.Errorf("Expected set length to be 1, got %d", len(s))
	}
	if !s.Contains(1) {
		t.Errorf("Set should contain element 1 after adding it")
	}

	s.Add(1)
	if len(s) != 1 {
		t.Errorf("Expected set length to be 1, got %d", len(s))
	}
}

func Test_Remove(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Remove(1)

	if len(s) != 0 {
		t.Errorf("Expected set length to be 0, got %d", len(s))
	}
	if s.Contains(1) {
		t.Errorf("Set should not contain element 1 after removing it")
	}
}

func Test_Contains(t *testing.T) {
	s := NewSet[int]()

	if s.Contains(1) {
		t.Errorf("Set should not contain element 1")
	}

	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Set should contain element 1 after adding it")
	}

	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Set should not contain element 1 after removing it")
	}
}
