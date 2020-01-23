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
