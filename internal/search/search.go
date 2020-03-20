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

package search

import (
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
)

func StringInSlice(a []string, x string) bool {
	for _, b := range a {
		if b == x {
			return true
		}
	}
	return false
}

// Processes return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, Processes returns n and false
func Processes(a []*om.Process, f func(*om.Process) bool) (int, bool) {
	for i, p := range a {
		if f(p) {
			return i, true
		}
	}
	return len(a), false
}

// Members return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, Members returns n and false
func Members(a []om.Member, f func(om.Member) bool) (int, bool) {
	for i, m := range a {
		if f(m) {
			return i, true
		}
	}
	return len(a), false
}

// ReplicaSets return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, ReplicaSets returns n and false
func ReplicaSets(a []*om.ReplicaSet, f func(*om.ReplicaSet) bool) (int, bool) {
	for i, m := range a {
		if f(m) {
			return i, true
		}
	}
	return len(a), false
}

// MongoDBUser return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, MongoDBUser returns n and false
func MongoDBUsers(a []*om.MongoDBUser, f func(*om.MongoDBUser) bool) (int, bool) {
	for i, m := range a {
		if f(m) {
			return i, true
		}
	}
	return len(a), false
}

// ClusterExists return true if a cluster exists for the given name
func ClusterExists(c *om.AutomationConfig, name string) bool {
	_, found := ReplicaSets(c.ReplicaSets, func(r *om.ReplicaSet) bool {
		return r.ID == name
	})

	return found
}
