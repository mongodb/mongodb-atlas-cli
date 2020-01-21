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

package mocks

import "github.com/mongodb-labs/pcgc/cloudmanager"

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
