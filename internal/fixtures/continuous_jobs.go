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

// TODO: https://github.com/mongodb/go-client-mongodb-atlas/pull/64
func ContinuousJob() *atlas.ContinuousJob {
	return &atlas.ContinuousJob{}
}

func ContinuousJobs() *atlas.ContinuousJobs {
	pointInTime := false

	return &atlas.ContinuousJobs{
		Links: []*atlas.Link{
			{
				Href: "http://mms:9080/backup/restore/v2/pull/5e6a4f56917b225d8c10e708/ODgwODQzZDZmNjk3NGE2ZGExMDM2M2U1YmU1MjgwMGM=/myReplicaSet4-1584012528-5e6a4f56917b225d8c10e708.tar.gz",
				Rel:  "self",
			},
		},
		Results: []*atlas.ContinuousJob{
			{
				ClusterID: "5e662732917b220fbd8be844",
				Created:   "2020-03-12T15:03:50Z",
				Delivery: &atlas.Delivery{
					Expires:         "2020-03-12T17:03:50Z",
					ExpirationHours: 2,
					MaxDownloads:    1,
					MethodName:      "HTTP",
					StatusName:      "READY",
					URL:             "http://mms:9080/backup/restore/v2/pull/5e6a4f56917b225d8c10e708/ODgwODQzZDZmNjk3NGE2ZGExMDM2M2U1YmU1MjgwMGM=/myReplicaSet4-1584012528-5e6a4f56917b225d8c10e708.tar.gz",
				},
				EncryptionEnabled: false,
				GroupID:           "5e66185d917b220fbd8bb4d1",
				ID:                "5e6a4f56917b225d8c10e708",
				Links: []*atlas.Link{
					{
						Href: "http://mms:9080/api/public/v1.0/groups/5e66185d917b220fbd8bb4d1/clusters/5e662732917b220fbd8be844/restoreJobs/5e6a4f56917b225d8c10e708",
						Rel:  "self",
					},
				},
				SnapshotID:  "5e6a1d3f917b22609860fd74",
				StatusName:  "FINISHED",
				PointInTime: &pointInTime,
				Timestamp: atlas.SnapshotTimestamp{
					Date:      "2020-03-12T11:28:48Z",
					Increment: 1,
				},
			},
		},
		TotalCount: 1,
	}
}
