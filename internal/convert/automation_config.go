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

package convert

import (
	"go.mongodb.org/ops-manager/opsmngr"
)

// FromAutomationConfig convert from opsmngr.AutomationConfig format to []*ClusterConfig
// the given opsmngr.AutomationConfig will be modified
func FromAutomationConfig(c *opsmngr.AutomationConfig) []*ClusterConfig {
	out := make([]*ClusterConfig, 0, len(c.ReplicaSets))

	for i, s := range c.Sharding {
		out = append(out, newShardedCluster(s))
		for j, ss := range s.Shards {
			id := ss.ID
			out[i].Shards[j] = newRSConfig(c, id)
		}

		out[i].Config = newRSConfig(c, s.ConfigServerReplica)
		for j, p := range c.Processes {
			if p.Cluster == s.Name {
				out[i].Mongos = append(out[i].Mongos, newMongosProcessConfig(p))
				out[i].addToMongoURI(p)
				c.Processes = removeProcess(c.Processes, j)
				break
			}
		}
	}
	for i, rs := range c.ReplicaSets {
		out = append(out, newReplicaSetCluster(rs.ID, len(rs.Members)))
		for j, m := range rs.Members {
			for k, p := range c.Processes {
				if p.Name == m.Host {
					out[i].ProcessConfigs[j] = newReplicaSetProcessConfig(m, p)
					out[i].addToMongoURI(p)
					c.Processes = removeProcess(c.Processes, k)
					break
				}
			}
		}
	}

	return out
}

func removeProcess(in []*opsmngr.Process, i int) []*opsmngr.Process {
	return append(in[:i], in[i+1:]...)
}
