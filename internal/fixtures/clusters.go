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

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func Cluster() *atlas.Cluster {
	var falseValue = true
	var one float64 = 1
	var two int64 = 1
	return &atlas.Cluster{
		ID:                       "1",
		AutoScaling:              atlas.AutoScaling{DiskGBEnabled: &falseValue},
		BackupEnabled:            &falseValue,
		BiConnector:              atlas.BiConnector{Enabled: &falseValue, ReadPreference: "secondary"},
		ClusterType:              "REPLICASET",
		DiskSizeGB:               &one,
		EncryptionAtRestProvider: "AWS",
		GroupID:                  "asdasdads",
		MongoDBVersion:           "3.4.9",
		MongoURI:                 "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		MongoURIUpdated:          "2017-10-23T21:26:17Z",
		MongoURIWithOptions:      "mongodb://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=mongo-shard-0",
		Name:                     "AppData",
		NumShards:                &two,
		Paused:                   &falseValue,
		ProviderSettings: &atlas.ProviderSettings{
			ProviderName:     "AWS",
			DiskIOPS:         &two,
			EncryptEBSVolume: &falseValue,
			InstanceSizeName: "M40",
			RegionName:       "US_WEST_2",
		},
		ReplicationFactor: &two,
		ReplicationSpec: map[string]atlas.RegionsConfig{
			"US_EAST_1": {
				ElectableNodes: &two,
				Priority:       &two,
				ReadOnlyNodes:  &two,
			},
		},
		SrvAddress: "mongodb+srv://mongo-shard-00-00.mongodb.net:27017,mongo-shard-00-01.mongodb.net:27017,mongo-shard-00-02.mongodb.net:27017",
		StateName:  "CREATING",
	}
}

func Clusters() []atlas.Cluster {
	return []atlas.Cluster{*Cluster()}
}
