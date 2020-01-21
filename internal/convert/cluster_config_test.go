// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package convert

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/go-test/deep"
	"github.com/mongodb-labs/pcgc/cloudmanager"
)

func TestClusterConfig_PatchAutomationConfig(t *testing.T) {
	testCases := map[string]struct {
		current  *cloudmanager.AutomationConfig
		expected *cloudmanager.AutomationConfig
		changes  ClusterConfig
	}{
		"add a replica set to an empty config": {
			current: mocks.EmptyAutomationConfig(),
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
			expected: &cloudmanager.AutomationConfig{
				Auth: cloudmanager.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*cloudmanager.Process{
					{
						Args26: cloudmanager.Args26{
							NET: cloudmanager.Net{Port: 1},
							Replication: &cloudmanager.Replication{
								ReplSetName: "test_config",
							},
							Storage: &cloudmanager.Storage{
								DBPath: "/data",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
				ReplicaSets: []*cloudmanager.ReplicaSet{
					{
						ID:              "test_config",
						ProtocolVersion: "1",
						Members: []cloudmanager.Member{
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
			current: mocks.AutomationConfigWithOneReplicaSet("replica_set_1", false),
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
			expected: &cloudmanager.AutomationConfig{
				Auth: cloudmanager.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*cloudmanager.Process{
					// Old
					{
						Args26: cloudmanager.Args26{
							NET: cloudmanager.Net{Port: 27017},
							Replication: &cloudmanager.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &cloudmanager.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
						Args26: cloudmanager.Args26{
							NET: cloudmanager.Net{Port: 1},
							Replication: &cloudmanager.Replication{
								ReplSetName: "test_config",
							},
							Storage: &cloudmanager.Storage{
								DBPath: "/data",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
				ReplicaSets: []*cloudmanager.ReplicaSet{
					// Old
					{
						ID:              "replica_set_1",
						ProtocolVersion: "1",
						Members: []cloudmanager.Member{
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
						Members: []cloudmanager.Member{
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
			current: mocks.AutomationConfigWithOneReplicaSet("replica_set_1", false),
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
			expected: &cloudmanager.AutomationConfig{
				Auth: cloudmanager.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*cloudmanager.Process{
					// Old
					{
						Args26: cloudmanager.Args26{
							NET: cloudmanager.Net{Port: 27017},
							Replication: &cloudmanager.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &cloudmanager.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
						Args26: cloudmanager.Args26{
							NET: cloudmanager.Net{Port: 27017},
							Replication: &cloudmanager.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &cloudmanager.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
				ReplicaSets: []*cloudmanager.ReplicaSet{
					// Old
					{
						ID:              "replica_set_1",
						ProtocolVersion: "1",
						Members: []cloudmanager.Member{
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
			current: mocks.AutomationConfigWithOneReplicaSet("replica_set_1", false),
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
			expected: &cloudmanager.AutomationConfig{
				Auth: cloudmanager.Auth{
					DeploymentAuthMechanisms: []string{},
				},
				Processes: []*cloudmanager.Process{
					// Old
					{
						Args26: cloudmanager.Args26{
							NET:         cloudmanager.Net{Port: 27017},
							Replication: &cloudmanager.Replication{},
							Storage: &cloudmanager.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
						Args26: cloudmanager.Args26{
							NET: cloudmanager.Net{Port: 27017},
							Replication: &cloudmanager.Replication{
								ReplSetName: "replica_set_1",
							},
							Storage: &cloudmanager.Storage{
								DBPath: "/data/db/",
							},
							SystemLog: cloudmanager.SystemLog{
								Destination: "file",
								Path:        "/data/db/mongodb.log",
							},
						},
						LogRotate: &cloudmanager.LogRotate{
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
				ReplicaSets: []*cloudmanager.ReplicaSet{
					// New
					{
						ID:              "replica_set_1",
						ProtocolVersion: "1",
						Members: []cloudmanager.Member{
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
