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
