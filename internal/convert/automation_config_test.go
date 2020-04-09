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
	"github.com/mongodb/mongocli/internal/fixtures"
)

func TestFromAutomationConfig(t *testing.T) {
	name := "cluster_1"
	config := fixtures.AutomationConfigWithOneReplicaSet(name, false)

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

	result := FromAutomationConfig(config)
	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestEnableMechanism(t *testing.T) {
	config := fixtures.AutomationConfigWithoutMongoDBUsers()

	e := EnableMechanism(config, []string{"SCRAM-SHA-256"})

	if e != nil {
		t.Fatalf("EnableMechanism() unexpected error: %v\n", e)
	}

	if config.Auth.Disabled {
		t.Error("config.Auth.Disabled is true\n")
	}

	if config.Auth.AutoAuthMechanisms[0] != "SCRAM-SHA-256" {
		t.Error("AutoAuthMechanisms not set\n")
	}

	if config.Auth.AutoUser == "" || config.Auth.AutoPwd == "" {
		t.Error("config.Auth.Auto* not set\n")
	}

	if config.Auth.Key == "" || config.Auth.KeyFileWindows == "" || config.Auth.KeyFile == "" {
		t.Error("config.Auth.Key* not set\n")
	}

	if len(config.Auth.Users) != 2 {
		t.Error("automation and monitoring users not set\n")
	}
}
