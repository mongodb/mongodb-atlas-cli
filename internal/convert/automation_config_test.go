package convert

import (
	"reflect"
	"testing"

	"github.com/mongodb-labs/pcgc/cloudmanager"
)

func TestFromAutomationConfig(t *testing.T) {
	cloud := &cloudmanager.AutomationConfig{
		Processes: []*cloudmanager.Process{
			{
				Args26: cloudmanager.Args26{
					NET: cloudmanager.Net{
						Port: 27017,
					},
					Replication: &cloudmanager.Replication{
						ReplSetName: "cluster_1",
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
				Name:                        "cluster_1_0",
				Disabled:                    false,
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
				ID: "cluster_1",
				Members: []cloudmanager.Member{
					{
						ID:           0,
						ArbiterOnly:  false,
						BuildIndexes: true,
						Hidden:       false,
						Host:         "cluster_1_0",
						Priority:     1,
						SlaveDelay:   0,
						Votes:        1,
					},
				},
			},
		},
	}

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
