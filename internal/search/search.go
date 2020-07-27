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
	"go.mongodb.org/ops-manager/opsmngr"
	"go.mongodb.org/ops-manager/search"
)

func StringInSlice(a []string, x string) bool {
	for _, b := range a {
		if b == x {
			return true
		}
	}
	return false
}

// ClusterExists return true if a cluster exists for the given name
func ClusterExists(c *opsmngr.AutomationConfig, name string) bool {
	_, rsFound := search.ReplicaSets(c.ReplicaSets, func(r *opsmngr.ReplicaSet) bool {
		return r.ID == name
	})

	_, shardedFound := search.ShardingConfig(c.Sharding, func(r *opsmngr.ShardingConfig) bool {
		return r.Name == name
	})

	return rsFound || shardedFound
}
