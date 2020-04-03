package fixtures

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

func ProcessMeasurements() *atlas.ProcessMeasurements {
	return &atlas.ProcessMeasurements{
		End:         "2017-08-22T20:31:14Z",
		Granularity: "PT1M",
		GroupID:     "12345678",
		HostID:      "shard-00-00.mongodb.net:27017",
		Links: []*atlas.Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/measurements?granularity=PT1M&period=PT1M",
			},
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017",
				Rel:  "http://mms.mongodb.com/host",
			},
		},
		Measurements: []*atlas.Measurements{
			{
				DataPoints: []*atlas.DataPoints{
					{
						Timestamp: "2017-08-22T20:31:12Z",
						Value:     nil,
					},
					{
						Timestamp: "2017-08-22T20:31:14Z",
						Value:     nil,
					},
				},
				Name:  "ASSERT_REGULAR",
				Units: "SCALAR_PER_SECOND",
			},
		},
		ProcessID: "shard-00-00.mongodb.net:27017",
		Start:     "2017-08-22T20:30:45Z",
	}
}
