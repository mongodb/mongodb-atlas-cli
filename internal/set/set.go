// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

// Define a generic Set type.
type Set[T comparable] map[T]struct{}

// NewSet creates and returns a new Set.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add adds an element to the Set.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Remove removes an element from the Set.
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Contains checks if an element is in the Set.
func (s Set[T]) Contains(elem T) bool {
	_, exists := s[elem]
	return exists
}
