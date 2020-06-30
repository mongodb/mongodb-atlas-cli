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
	"testing"

	"github.com/go-test/deep"
	"github.com/mongodb/mongocli/internal/fixture"
)

func TestFromAutomationConfig(t *testing.T) {
	f := false
	buildIndexes := true
	var one float64 = 1
	var zero float64 = 0
	name := "cluster_1"
	t.Run("replica set", func(t *testing.T) {
		config := fixture.AutomationConfigWithOneReplicaSet(name, false)
		expected := []*ClusterConfig{
			{
				RSConfig: RSConfig{
					Name: name,
					ProcessConfigs: []*ProcessConfig{
						{
							ArbiterOnly:  &f,
							BuildIndexes: &buildIndexes,
							DBPath:       "/data/db/",
							Disabled:     false,
							Hidden:       &f,
							Hostname:     "host0",
							LogPath:      "/data/db/mongodb.log",
							Port:         27017,
							Priority:     &one,
							ProcessType:  mongod,
							SlaveDelay:   &zero,
							Votes:        &one,
							FCVersion:    "4.2",
							Version:      "4.2.2",
							Name:         name + "_0",
						},
					},
				},
				MongoURI: "mongodb://host0:27017",
			},
		}

		result := FromAutomationConfig(config)
		if diff := deep.Equal(result, expected); diff != nil {
			t.Error(diff)
		}
	})
	t.Run("sharded cluster", func(t *testing.T) {
		config := fixture.AutomationConfigWithOneShardedCluster(name, false)
		expected := []*ClusterConfig{
			{
				MongoURI: "mongodb://example:3",
				RSConfig: RSConfig{
					Name: name,
				},
				Shards: []*RSConfig{
					{
						Name: "myShard_0",
						ProcessConfigs: []*ProcessConfig{
							{
								ArbiterOnly:  &f,
								BuildIndexes: &buildIndexes,
								DBPath:       "/data/myShard_0",
								Disabled:     false,
								Hidden:       &f,
								Hostname:     "example",
								LogPath:      "/log/myShard_0",
								Port:         1,
								Priority:     &one,
								ProcessType:  mongod,
								SlaveDelay:   &zero,
								Votes:        &one,
								FCVersion:    "4.2",
								Version:      "4.2.2",
								Name:         name + "_myShard_0_0",
							},
						},
					},
				},
				Config: &RSConfig{
					Name: "configRS",
					ProcessConfigs: []*ProcessConfig{
						{
							ArbiterOnly:  &f,
							BuildIndexes: &buildIndexes,
							DBPath:       "/data/configRS",
							Disabled:     false,
							Hidden:       &f,
							Hostname:     "example",
							LogPath:      "/log/configRS",
							Port:         2,
							Priority:     &one,
							ProcessType:  mongod,
							SlaveDelay:   &zero,
							Votes:        &one,
							FCVersion:    "4.2",
							Version:      "4.2.2",
							Name:         name + "_configRS_1",
						},
					},
				},
				Mongos: []*ProcessConfig{
					{
						Disabled:    false,
						Hostname:    "example",
						LogPath:     "/log/mongos",
						Port:        3,
						ProcessType: "mongos",
						FCVersion:   "4.2",
						Version:     "4.2.2",
						Name:        name + "_mongos_2",
					},
				},
			},
		}
		result := FromAutomationConfig(config)
		if diff := deep.Equal(result, expected); diff != nil {
			t.Error(diff)
		}
	})
}
