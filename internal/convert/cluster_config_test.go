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
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/fixtures"
)

func TestClusterConfig_PatchAutomationConfig(t *testing.T) {
	testCases := map[string]struct {
		current  *om.AutomationConfig
		expected *om.AutomationConfig
		changes  ClusterConfig
	}{
		"add a replica set to an empty config": {
			current: fixtures.EmptyAutomationConfig(),
			changes: ClusterConfig{
				FCVersion: "4.2",
				Name:      "test_config",
				Version:   "4.2.2",
				ProcessConfigs: []ProcessConfig{
					{
						DBPath:   "/data",
						Hostname: "example",
						LogPath:  "/log",
						Port:     1,
						Priority: 1,
						Votes:    1,
					},
				},
			},
			expected: &om.AutomationConfig{
				Auth: om.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*om.Process{
					{
						Args26: om.Args26{
							NET: om.Net{Port: 1},
							Replication: &om.Replication{
								ReplSetName: "test_config",
							},
							Storage: &om.Storage{
								DBPath: "/data",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "test_config_0",
						Disabled:                    false,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "example",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
				},
				ReplicaSets: []*om.ReplicaSet{
					{
						ID:              "test_config",
						ProtocolVersion: "1",
						Members: []om.Member{
							{
								ID:           0,
								ArbiterOnly:  false,
								BuildIndexes: true,
								Hidden:       false,
								Host:         "test_config_0",
								Priority:     1,
								SlaveDelay:   0,
								Votes:        1,
							},
						},
					},
				},
			},
		},
		"add a replica set to a config with an existing replica set": {
			current: fixtures.AutomationConfigWithOneReplicaSet("replica_set_1", false),
			changes: ClusterConfig{
				FCVersion: "4.2",
				Name:      "test_config",
				Version:   "4.2.2",
				ProcessConfigs: []ProcessConfig{
					{
						DBPath:   "/data",
						Hostname: "example",
						LogPath:  "/log",
						Port:     1,
						Priority: 1,
						Votes:    1,
					},
				},
			},
			expected: &om.AutomationConfig{
				Auth: om.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*om.Process{
					// Old
					{
						Args26: om.Args26{
							NET: om.Net{Port: 27017},
							Replication: &om.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &om.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "replica_set_1_0",
						Disabled:                    false,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "host0",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
					// New
					{
						Args26: om.Args26{
							NET: om.Net{Port: 1},
							Replication: &om.Replication{
								ReplSetName: "test_config",
							},
							Storage: &om.Storage{
								DBPath: "/data",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "test_config_1",
						Disabled:                    false,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "example",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
				},
				ReplicaSets: []*om.ReplicaSet{
					// Old
					{
						ID:              "replica_set_1",
						ProtocolVersion: "1",
						Members: []om.Member{
							{
								ArbiterOnly:  false,
								BuildIndexes: true,
								Hidden:       false,
								Host:         "replica_set_1_0",
								Priority:     1,
								SlaveDelay:   0,
								Votes:        1,
							},
						},
					},
					// New
					{
						ID:              "test_config",
						ProtocolVersion: "1",
						Members: []om.Member{
							{
								ArbiterOnly:  false,
								BuildIndexes: true,
								Hidden:       false,
								Host:         "test_config_1",
								Priority:     1,
								SlaveDelay:   0,
								Votes:        1,
							},
						},
					},
				},
			},
		},
		"add a process to a config with an existing replica set": {
			current: fixtures.AutomationConfigWithOneReplicaSet("replica_set_1", false),
			changes: ClusterConfig{
				FCVersion: "4.2",
				Name:      "replica_set_1",
				Version:   "4.2.2",
				ProcessConfigs: []ProcessConfig{
					{
						DBPath:   "/data/db/",
						Hostname: "host0",
						LogPath:  "/data/db/mongodb.log",
						Port:     27017,
						Priority: 1,
						Votes:    1,
					}, {
						DBPath:   "/data/db/",
						Hostname: "host1",
						LogPath:  "/data/db/mongodb.log",
						Port:     27017,
						Priority: 1,
						Votes:    1,
					},
				},
			},
			expected: &om.AutomationConfig{
				Auth: om.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*om.Process{
					// Old
					{
						Args26: om.Args26{
							NET: om.Net{Port: 27017},
							Replication: &om.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &om.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "replica_set_1_0",
						Disabled:                    false,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "host0",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
					// New
					{
						Args26: om.Args26{
							NET: om.Net{Port: 27017},
							Replication: &om.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &om.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "replica_set_1_2",
						Disabled:                    false,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "host1",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
				},
				ReplicaSets: []*om.ReplicaSet{
					// Old
					{
						ID:              "replica_set_1",
						ProtocolVersion: "1",
						Members: []om.Member{
							{
								ArbiterOnly:  false,
								BuildIndexes: true,
								Hidden:       false,
								Host:         "replica_set_1_0",
								Priority:     1,
								SlaveDelay:   0,
								Votes:        1,
							},
							{
								ID:           1,
								ArbiterOnly:  false,
								BuildIndexes: true,
								Hidden:       false,
								Host:         "replica_set_1_2",
								Priority:     1,
								SlaveDelay:   0,
								Votes:        1,
							},
						},
					},
				},
			},
		},
		"replace a process to a config with an existing replica set": {
			current: fixtures.AutomationConfigWithOneReplicaSet("replica_set_1", false),
			changes: ClusterConfig{
				FCVersion: "4.2",
				Name:      "replica_set_1",
				Version:   "4.2.2",
				ProcessConfigs: []ProcessConfig{
					{
						DBPath:   "/data/db/",
						Hostname: "host1",
						LogPath:  "/data/db/mongodb.log",
						Port:     27017,
						Priority: 1,
						Votes:    1,
					},
				},
			},
			expected: &om.AutomationConfig{
				Auth: om.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*om.Process{
					// Old
					{
						Args26: om.Args26{
							NET:         om.Net{Port: 27017},
							Replication: &om.Replication{},
							Storage: &om.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "replica_set_1_0",
						Disabled:                    true,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "host0",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
					// New
					{
						Args26: om.Args26{
							NET: om.Net{Port: 27017},
							Replication: &om.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &om.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: om.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &om.LogRotate{
							SizeThresholdMB:  1000,
							TimeThresholdHrs: 24,
						},
						AuthSchemaVersion:           5,
						Name:                        "replica_set_1_1",
						Disabled:                    false,
						FeatureCompatibilityVersion: "4.2",
						Hostname:                    "host1",
						ManualMode:                  false,
						ProcessType:                 "mongod",
						Version:                     "4.2.2",
					},
				},
				ReplicaSets: []*om.ReplicaSet{
					// New
					{
						ID:              "replica_set_1",
						ProtocolVersion: "1",
						Members: []om.Member{
							{
								ID:           1,
								ArbiterOnly:  false,
								BuildIndexes: true,
								Hidden:       false,
								Host:         "replica_set_1_1",
								Priority:     1,
								SlaveDelay:   0,
								Votes:        1,
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.changes.PatchAutomationConfig(tc.current)
			if err != nil {
				t.Fatalf("PatchAutomationConfig() unexpected error: %v", err)
			}
			if diff := deep.Equal(tc.current, tc.expected); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestProtocolVersion(t *testing.T) {
	testCases := map[string]struct {
		mdbVersion      string
		protocolVersion string
	}{
		"post 4.0": {
			mdbVersion:      "4.0",
			protocolVersion: "1",
		},
		"pre 4.0": {
			mdbVersion:      "3.6",
			protocolVersion: "0",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ver, err := protocolVer(tc.mdbVersion)
			if err != nil {
				t.Fatalf("protocolVer() unexpected error: %v", err)
			}
			if ver != tc.protocolVersion {
				t.Errorf("protocolVer() expected: %s but got: %s", tc.protocolVersion, ver)
			}
		})
	}
}
