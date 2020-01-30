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

	"github.com/10gen/mcli/internal/fixtures"
	"github.com/go-test/deep"
)

func TestFromAutomationConfig(t *testing.T) {
	name := "cluster_1"
	cloud := fixtures.AutomationConfigWithOneReplicaSet(name, false)

	buildIndexes := true
	expected := []ClusterConfig{
		{
			Name:     name,
			MongoURI: "mongodb://host0:27017",
			ProcessConfigs: []ProcessConfig{
				{
					ArbiterOnly:  false,
					BuildIndexes: &buildIndexes,
					DBPath:       "/data/db/",
					Disabled:     false,
					Hidden:       false,
					Hostname:     "host0",
					LogPath:      "/data/db/mongodb.log",
					Port:         27017,
					Priority:     1,
					ProcessType:  mongod,
					SlaveDelay:   0,
					Votes:        1,
					FCVersion:    "4.2",
					Version:      "4.2.2",
					Name:         name + "_0",
				},
			},
		},
	}

	result := FromAutomationConfig(cloud)
	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestShutdown(t *testing.T) {
	name := "cluster_1"
	cloud := fixtures.AutomationConfigWithOneReplicaSet(name, false)

	Shutdown(cloud, name)
	if !cloud.Processes[0].Disabled {
		t.Errorf("TestShutdown\n got=%#v\nwant=%#v\n", cloud.Processes[0].Disabled, true)
	}
}

func TestStartup(t *testing.T) {
	name := "cluster_1"
	cloud := fixtures.AutomationConfigWithOneReplicaSet(name, true)

	Startup(cloud, name)
	if cloud.Processes[0].Disabled {
		t.Errorf("TestStartup\n got=%#v\nwant=%#v\n", cloud.Processes[0].Disabled, false)
	}
}
