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
	"reflect"
	"testing"

	"github.com/mongodb-labs/pcgc/cloudmanager"
)

func TestFromAutomationConfig(t *testing.T) {
	name := "cluster_1"
	cloud := automationConfig(name, false)

	buildIndexes := true
	expected := []ClusterConfig{
		{
			Name:     "cluster_1",
			MongoURI: "mongodb://host0:27017",
			Processes: []ProcessConfig{
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
				},
			},
		},
	}

	result := FromAutomationConfig(cloud)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromAutomationConfig\n got=%#v\nwant=%#v", result, expected)
	}
}

func TestShutdown(t *testing.T) {
	name := "cluster_1"
	cloud := automationConfig(name, false)

	Shutdown(cloud, name)
	if !cloud.Processes[0].Disabled {
		t.Errorf("TestShutdown\n got=%#v\nwant=%#v", cloud.Processes[0].Disabled, true)
	}
}

func TestStartup(t *testing.T) {
	name := "cluster_1"
	cloud := automationConfig(name, true)

	Startup(cloud, name)
	if cloud.Processes[0].Disabled {
		t.Errorf("TestStartup\n got=%#v\nwant=%#v", cloud.Processes[0].Disabled, false)
	}
}

func automationConfig(name string, disabled bool) *cloudmanager.AutomationConfig {
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
					Storage: cloudmanager.Storage{
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
				ProcessType: mongod,
				Version:     "4.2.2",
			},
		},
		ReplicaSets: []*cloudmanager.ReplicaSet{
			{
				ID: name,
				Members: []cloudmanager.Member{
					{
						ID:           0,
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
