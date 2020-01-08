package mocks

import (
	"github.com/mongodb-labs/pcgc/cloudmanager"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

var Project1 = &atlas.Project{
	ClusterCount: 2,
	Created:      "2016-07-14T14:19:33Z",
	ID:           "5a0a1e7e0f2912c554080ae6",
	Links: []*atlas.Link{
		{
			Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae6",
			Rel:  "self",
		},
	},
	Name:  "ProjectBar",
	OrgID: "5a0a1e7e0f2912c554080adc",
}

var Project2 = &atlas.Project{
	ClusterCount: 0,
	Created:      "2017-10-16T15:24:01Z",
	ID:           "5a0a1e7e0f2912c554080ae7",
	Links: []*atlas.Link{
		{
			Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/5a0a1e7e0f2912c554080ae7",
			Rel:  "self",
		},
	},
	Name:  "Project Foo",
	OrgID: "5a0a1e7e0f2912c554080adc",
}

func ProjectsMock() *atlas.Projects {
	return &atlas.Projects{
		Links: []*atlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups",
				Rel:  "self",
			},
		},
		Results:    []*atlas.Project{Project1, Project2},
		TotalCount: 2,
	}
}

func ClusterMock() *atlas.Cluster {
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

func ClustersMock() []atlas.Cluster {
	return []atlas.Cluster{*ClusterMock()}
}

func DatabaseUserMock() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		Roles: []atlas.Role{
			{
				RoleName:     "admin",
				DatabaseName: "admin",
			},
		},
		GroupID:      "5def8d5dce4bd936ac99ae9c",
		Username:     "test4",
		DatabaseName: "admin",
		LDAPAuthType: "NONE",
	}
}

func ProjectIPWhitelistMock() []atlas.ProjectIPWhitelist {
	return []atlas.ProjectIPWhitelist{
		{
			Comment:   "test",
			GroupID:   "5def8d5dce4bd936ac99ae9c",
			CIDRBlock: "37.228.254.100/32",
			IPAddress: "37.228.254.100",
		},
	}
}

func AutomationMock() *cloudmanager.AutomationConfig {
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
					Storage: cloudmanager.Storage{
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
					Storage: cloudmanager.Storage{
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
					Storage: cloudmanager.Storage{
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
