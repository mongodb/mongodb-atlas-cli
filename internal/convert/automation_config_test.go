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

//go:build unit

package convert

import (
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/pointer"
	"github.com/andreangiolillo/mongocli-test/internal/test/fixture"
	"github.com/go-test/deep"
)

func TestFromAutomationConfig(t *testing.T) {
	name := "cluster_1"
	fipsMode := true
	t.Run("replica set", func(t *testing.T) {
		t.Parallel()
		config := fixture.AutomationConfigWithOneReplicaSet(name, false)
		expected := []*ClusterConfig{
			{
				RSConfig: RSConfig{
					Name: name,
					Processes: []*ProcessConfig{
						{
							ArbiterOnly:                 pointer.Get(false),
							BuildIndexes:                pointer.Get(true),
							DBPath:                      "/data/db/",
							Disabled:                    false,
							Hidden:                      pointer.Get(false),
							Hostname:                    "host0",
							LogPath:                     "/data/db/mongodb.log",
							LogDestination:              file,
							AuditLogDestination:         file,
							AuditLogPath:                "/data/db/audit.log",
							Port:                        27017,
							Priority:                    pointer.Get[float64](1),
							ProcessType:                 mongod,
							SlaveDelay:                  pointer.Get[float64](1),
							SecondaryDelaySecs:          pointer.Get[float64](1),
							Votes:                       pointer.Get[float64](1),
							FeatureCompatibilityVersion: "4.2",
							Version:                     "4.2.2",
							Name:                        name + "_0",
							OplogSizeMB:                 pointer.Get(10),
							TLS: &TLS{
								CAFile:                     "CAFile",
								CertificateKeyFile:         "CertificateKeyFile",
								CertificateKeyFilePassword: "CertificateKeyFilePassword",
								CertificateSelector:        "CertificateSelector",
								ClusterCertificateSelector: "ClusterCertificateSelector",
								ClusterFile:                "ClusterFile",
								ClusterPassword:            "ClusterPassword",
								CRLFile:                    "CRLFile",
								DisabledProtocols:          "DisabledProtocols",
								FIPSMode:                   &fipsMode,
								Mode:                       "Mode",
								PEMKeyFile:                 "PEMKeyFile",
							},
							Security: &map[string]interface{}{
								"test": "test",
							},
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
		t.Parallel()
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
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data/myShard_0",
								Disabled:                    false,
								Hidden:                      pointer.Get(false),
								Hostname:                    "example",
								LogPath:                     "/log/myShard_0",
								LogDestination:              file,
								Port:                        1,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SlaveDelay:                  pointer.Get[float64](1),
								SecondaryDelaySecs:          pointer.Get[float64](1),
								Votes:                       pointer.Get[float64](1),
								FeatureCompatibilityVersion: "4.2",
								Version:                     "4.2.2",
								Name:                        name + "_myShard_0_0",
							},
						},
					},
				},
				Config: &RSConfig{
					Name: "configRS",
					Processes: []*ProcessConfig{
						{
							ArbiterOnly:                 pointer.Get(false),
							BuildIndexes:                pointer.Get(true),
							DBPath:                      "/data/configRS",
							Disabled:                    false,
							Hidden:                      pointer.Get(false),
							Hostname:                    "example",
							LogPath:                     "/log/configRS",
							LogDestination:              file,
							Port:                        2,
							Priority:                    pointer.Get[float64](1),
							ProcessType:                 mongod,
							SlaveDelay:                  pointer.Get[float64](1),
							SecondaryDelaySecs:          pointer.Get[float64](1),
							Votes:                       pointer.Get[float64](1),
							FeatureCompatibilityVersion: "4.2",
							Version:                     "4.2.2",
							Name:                        name + "_configRS_1",
						},
					},
				},
				Mongos: []*ProcessConfig{
					{
						Disabled:                    false,
						Hostname:                    "example",
						LogPath:                     "/log/mongos",
						LogDestination:              file,
						Port:                        3,
						ProcessType:                 "mongos",
						FeatureCompatibilityVersion: "4.2",
						Version:                     "4.2.2",
						Name:                        name + "_mongos_2",
					},
				},
			},
		}
		result := FromAutomationConfig(config)
		if diff := deep.Equal(result, expected); diff != nil {
			t.Error(diff)
		}
	})
	t.Run("Sharded multi mongos cluster", func(t *testing.T) {
		t.Parallel()
		config := fixture.MultiMongosAutomationConfig()

		engineConfig := make(map[string]interface{})
		engineConfig["cacheSizeGB"] = 0.5
		expected := []*ClusterConfig{
			{
				MongoURI: "mongodb://ip-172-31-43-144.eu-west-1.compute.internal:27021,ip-172-31-39-241.eu-west-1.compute.internal:27021,ip-172-31-37-180.eu-west-1.compute.internal:27021,ip-172-31-35-62.eu-west-1.compute.internal:27021",
				RSConfig: RSConfig{
					Name: "myCluster",
				},
				Shards: []*RSConfig{
					{
						Name: "myShard_0",
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data",
								Disabled:                    false,
								FeatureCompatibilityVersion: "6.0",
								Hidden:                      pointer.Get(false),
								Hostname:                    "ip-172-31-33-34.eu-west-1.compute.internal",
								LogPath:                     "/data/mongodb.log",
								Name:                        "myCluster_myShard_0_1",
								LogDestination:              "file",
								Port:                        27017,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SecondaryDelaySecs:          pointer.Get[float64](0),
								Votes:                       pointer.Get[float64](1),
								Version:                     "6.0.6-ent",
								WiredTiger:                  &engineConfig,
							},
						},
					},
					{
						Name: "myShard_1",
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data",
								Disabled:                    false,
								FeatureCompatibilityVersion: "6.0",
								Hidden:                      pointer.Get(false),
								Hostname:                    "ip-172-31-35-62.eu-west-1.compute.internal",
								LogPath:                     "/data/mongodb.log",
								LogDestination:              "file",
								Name:                        "myCluster_myShard_1_2",
								Port:                        27017,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SecondaryDelaySecs:          pointer.Get[float64](0),
								Votes:                       pointer.Get[float64](1),
								Version:                     "6.0.6-ent",
								WiredTiger:                  &engineConfig,
							},
						},
					},
					{
						Name: "myShard_2",
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data",
								Disabled:                    false,
								FeatureCompatibilityVersion: "6.0",
								Hidden:                      pointer.Get(false),
								Hostname:                    "ip-172-31-37-180.eu-west-1.compute.internal",
								LogPath:                     "/data/mongodb.log",
								LogDestination:              "file",
								Name:                        "myCluster_myShard_2_3",
								Port:                        27017,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SecondaryDelaySecs:          pointer.Get[float64](0),
								Votes:                       pointer.Get[float64](1),
								Version:                     "6.0.6-ent",
								WiredTiger:                  &engineConfig,
							},
						},
					},
					{
						Name: "myShard_3",
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data",
								Disabled:                    false,
								FeatureCompatibilityVersion: "6.0",
								Hidden:                      pointer.Get(false),
								Hostname:                    "ip-172-31-39-241.eu-west-1.compute.internal",
								LogPath:                     "/data/mongodb.log",
								LogDestination:              "file",
								Name:                        "myCluster_myShard_3_4",
								Port:                        27017,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SecondaryDelaySecs:          pointer.Get[float64](0),
								Votes:                       pointer.Get[float64](1),
								Version:                     "6.0.6-ent",
								WiredTiger:                  &engineConfig,
							},
						},
					},
					{
						Name: "myShard_4",
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data",
								Disabled:                    false,
								FeatureCompatibilityVersion: "6.0",
								Hidden:                      pointer.Get(false),
								Hostname:                    "ip-172-31-43-144.eu-west-1.compute.internal",
								LogPath:                     "/data/mongodb.log",
								LogDestination:              "file",
								Name:                        "myCluster_myShard_4_5",
								Port:                        27017,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SecondaryDelaySecs:          pointer.Get[float64](0),
								Votes:                       pointer.Get[float64](1),
								Version:                     "6.0.6-ent",
								WiredTiger:                  &engineConfig,
							},
						},
					},
					{
						Name: "myShard_5",
						Processes: []*ProcessConfig{
							{
								ArbiterOnly:                 pointer.Get(false),
								BuildIndexes:                pointer.Get(true),
								DBPath:                      "/data",
								Disabled:                    false,
								FeatureCompatibilityVersion: "6.0",
								Hidden:                      pointer.Get(false),
								Hostname:                    "ip-172-31-43-246.eu-west-1.compute.internal",
								LogPath:                     "/data/mongodb.log",
								LogDestination:              "file",
								Name:                        "myCluster_myShard_5_6",
								Port:                        27017,
								Priority:                    pointer.Get[float64](1),
								ProcessType:                 mongod,
								SecondaryDelaySecs:          pointer.Get[float64](0),
								Votes:                       pointer.Get[float64](1),
								Version:                     "6.0.6-ent",
								WiredTiger:                  &engineConfig,
							},
						},
					},
				},
				Config: &RSConfig{
					Name: "configRS",
					Processes: []*ProcessConfig{
						{
							ArbiterOnly:                 pointer.Get(false),
							BuildIndexes:                pointer.Get(true),
							DBPath:                      "/data/n12",
							Disabled:                    false,
							Hidden:                      pointer.Get(false),
							Hostname:                    "ip-172-31-33-34.eu-west-1.compute.internal",
							LogPath:                     "/data/n12/mongodb.log",
							LogDestination:              "file",
							Port:                        27020,
							Priority:                    pointer.Get[float64](1),
							ProcessType:                 mongod,
							SecondaryDelaySecs:          pointer.Get[float64](0),
							Votes:                       pointer.Get[float64](1),
							FeatureCompatibilityVersion: "6.0",
							Version:                     "6.0.6-ent",
							Name:                        "myCluster_configRS_7",
							WiredTiger:                  &engineConfig,
						},
					},
				},
				Mongos: []*ProcessConfig{
					{
						Disabled:                    false,
						FeatureCompatibilityVersion: "6.0",
						Hostname:                    "ip-172-31-43-144.eu-west-1.compute.internal",
						LogDestination:              "file",
						LogPath:                     "/data/n1/mongodb.log",
						Name:                        "myCluster_mongos_11",
						Port:                        27021,
						ProcessType:                 "mongos",
						Version:                     "6.0.6-ent",
					},
					{
						Disabled:                    false,
						FeatureCompatibilityVersion: "6.0",
						Hostname:                    "ip-172-31-39-241.eu-west-1.compute.internal",
						LogDestination:              "file",
						LogPath:                     "/data/n1/mongodb.log",
						Name:                        "myCluster_mongos_10",
						Port:                        27021,
						ProcessType:                 "mongos",
						Version:                     "6.0.6-ent",
					},
					{
						Disabled:                    false,
						FeatureCompatibilityVersion: "6.0",
						Hostname:                    "ip-172-31-37-180.eu-west-1.compute.internal",
						LogDestination:              "file",
						LogPath:                     "/data/n1/mongodb.log",
						Name:                        "myCluster_mongos_9",
						Port:                        27021,
						ProcessType:                 "mongos",
						Version:                     "6.0.6-ent",
					},
					{
						Disabled:                    false,
						FeatureCompatibilityVersion: "6.0",
						Hostname:                    "ip-172-31-35-62.eu-west-1.compute.internal",
						LogDestination:              "file",
						LogPath:                     "/data/n1/mongodb.log",
						Name:                        "myCluster_mongos_8",
						Port:                        27021,
						ProcessType:                 "mongos",
						Version:                     "6.0.6-ent",
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
