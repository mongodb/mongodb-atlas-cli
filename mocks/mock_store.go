package mocks

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

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
