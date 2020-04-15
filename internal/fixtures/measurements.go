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

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

func ProcessDisks() *atlas.ProcessDisksResponse {
	return &atlas.ProcessDisksResponse{
		Links: []*atlas.Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks",
			},
		},
		Results: []*atlas.ProcessDisk{
			{
				Links: []*atlas.Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/disks/test",
					},
				},
				PartitionName: "test",
			},
		},
		TotalCount: 1,
	}
}

func ProcessDatabases() *atlas.ProcessDatabasesResponse {
	return &atlas.ProcessDatabasesResponse{
		Links: []*atlas.Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases",
			},
		},
		Results: []*atlas.ProcessDatabase{
			{
				Links: []*atlas.Link{
					{
						Rel:  "self",
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/12345678/processes/shard-00-00.mongodb.net:27017/databases/test",
					},
				},
				DatabaseName: "test",
			},
		},
		TotalCount: 1,
	}
}

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
