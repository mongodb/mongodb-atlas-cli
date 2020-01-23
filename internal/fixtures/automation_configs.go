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

package fixtures

import "github.com/mongodb-labs/pcgc/cloudmanager"

func AutomationConfig() *cloudmanager.AutomationConfig {
	return &cloudmanager.AutomationConfig{
		Auth: cloudmanager.Auth{
			AutoAuthMechanism: "MONGODB-CR",
			Disabled:          true,
			AuthoritativeSet:  false,
		},
		Processes: []*cloudmanager.Process{
			{
				Name:                        "myReplicaSet_1",
				ProcessType:                 "mongod",
				Version:                     "4.2.2",
				AuthSchemaVersion:           5,
				FeatureCompatibilityVersion: "4.2",
				Disabled:                    false,
				ManualMode:                  false,
				Hostname:                    "host0",
				Args26: cloudmanager.Args26{
					NET: cloudmanager.Net{
						Port: 27000,
					},
					Storage: &cloudmanager.Storage{
						DBPath: "/data/rs1",
					},
					SystemLog: cloudmanager.SystemLog{
						Destination: "file",
						Path:        "/data/rs1/mongodb.log",
					},
					Replication: &cloudmanager.Replication{
						ReplSetName: "myReplicaSet",
					},
				},
				LogRotate: &cloudmanager.LogRotate{
					SizeThresholdMB:  1000.0,
					TimeThresholdHrs: 24,
				},
				LastGoalVersionAchieved: 0,
				Cluster:                 "",
			},
			{
				Name:                        "myReplicaSet_2",
				ProcessType:                 "mongod",
				Version:                     "4.2.2",
				AuthSchemaVersion:           5,
				FeatureCompatibilityVersion: "4.2",
				Disabled:                    false,
				ManualMode:                  false,
				Hostname:                    "host1",
				Args26: cloudmanager.Args26{
					NET: cloudmanager.Net{
						Port: 27010,
					},
					Storage: &cloudmanager.Storage{
						DBPath: "/data/rs2",
					},
					SystemLog: cloudmanager.SystemLog{
						Destination: "file",
						Path:        "/data/rs2/mongodb.log",
					},
					Replication: &cloudmanager.Replication{
						ReplSetName: "myReplicaSet",
					},
				},
				LogRotate: &cloudmanager.LogRotate{
					SizeThresholdMB:  1000.0,
					TimeThresholdHrs: 24,
				},
				LastGoalVersionAchieved: 0,
				Cluster:                 "",
			},
			{
				Name:                        "myReplicaSet_3",
				ProcessType:                 "mongod",
				Version:                     "4.2.2",
				AuthSchemaVersion:           5,
				FeatureCompatibilityVersion: "4.2",
				Disabled:                    false,
				ManualMode:                  false,
				Hostname:                    "host0",
				Args26: cloudmanager.Args26{
					NET: cloudmanager.Net{
						Port: 27020,
					},
					Storage: &cloudmanager.Storage{
						DBPath: "/data/rs3",
					},
					SystemLog: cloudmanager.SystemLog{
						Destination: "file",
						Path:        "/data/rs3/mongodb.log",
					},
					Replication: &cloudmanager.Replication{
						ReplSetName: "myReplicaSet",
					},
				},
				LogRotate: &cloudmanager.LogRotate{
					SizeThresholdMB:  1000.0,
					TimeThresholdHrs: 24,
				},
				LastGoalVersionAchieved: 0,
				Cluster:                 "",
			},
		},
		ReplicaSets: []*cloudmanager.ReplicaSet{
			{
				ID:              "myReplicaSet",
				ProtocolVersion: "1",
				Members: []cloudmanager.Member{
					{
						ID:           0,
						ArbiterOnly:  false,
						BuildIndexes: true,
						Hidden:       false,
						Host:         "myReplicaSet_1",
						Priority:     1,
						SlaveDelay:   0,
						Votes:        1,
					},
					{
						ID:           1,
						ArbiterOnly:  false,
						BuildIndexes: true,
						Hidden:       false,
						Host:         "myReplicaSet_2",
						Priority:     1,
						SlaveDelay:   0,
						Votes:        1,
					},
					{
						ID:           2,
						ArbiterOnly:  false,
						BuildIndexes: true,
						Hidden:       false,
						Host:         "myReplicaSet_3",
						Priority:     1,
						SlaveDelay:   0,
						Votes:        1,
					},
				},
			},
		},
		Version:   1,
		UIBaseURL: "",
	}
}

func AutomationConfigWithOneReplicaSet(name string, disabled bool) *cloudmanager.AutomationConfig {
	return &cloudmanager.AutomationConfig{
		Processes: []*cloudmanager.Process{
			{
				Args26: cloudmanager.Args26{
					NET: cloudmanager.Net{
						Port: 27017,
					},
					Replication: &cloudmanager.Replication{
						ReplSetName: name,
					},
					Sharding: nil,
					Storage: &cloudmanager.Storage{
						DBPath: "/data/db/",
					},
					SystemLog: cloudmanager.SystemLog{
						Destination: "file",
						Path:        "/data/db/mongodb.log",
					},
				},
				AuthSchemaVersion:           5,
				Name:                        name + "_0",
				Disabled:                    disabled,
				FeatureCompatibilityVersion: "4.2",
				Hostname:                    "host0",
				LogRotate: &cloudmanager.LogRotate{
					SizeThresholdMB:  1000,
					TimeThresholdHrs: 24,
				},
				ProcessType: "mongod",
				Version:     "4.2.2",
			},
		},
		ReplicaSets: []*cloudmanager.ReplicaSet{
			{
				ID:              name,
				ProtocolVersion: "1",
				Members: []cloudmanager.Member{
					{
						ArbiterOnly:  false,
						BuildIndexes: true,
						Hidden:       false,
						Host:         name + "_0",
						Priority:     1,
						SlaveDelay:   0,
						Votes:        1,
					},
				},
			},
		},
	}
}

func EmptyAutomationConfig() *cloudmanager.AutomationConfig {
	return &cloudmanager.AutomationConfig{
		Processes:   make([]*cloudmanager.Process, 0),
		ReplicaSets: make([]*cloudmanager.ReplicaSet, 0),
	}
}
