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
	"github.com/mwielbut/pointy"
)

func Alert() *atlas.Alert {
	return &atlas.Alert{
		ID:            "533dc40ae4b00835ff81eaee",
		GroupID:       "535683b3794d371327b",
		EventTypeName: "OUTSIDE_METRIC_THRESHOLD",
		Created:       "2016-08-23T20:26:50Z",
		Updated:       "2016-08-23T20:26:50Z",
		Enabled:       pointy.Bool(true),
		Status:        "OPEN",
		Matchers: []atlas.Matcher{
			{
				FieldName: "HOSTNAME_AND_PORT",
				Operator:  "EQUALS",
				Value:     "mongo.example.com:27017",
			},
		},
		Notifications: []atlas.Notification{
			{
				TypeName:     "SMS",
				IntervalMin:  5,
				DelayMin:     pointy.Int(0),
				MobileNumber: "2343454567",
			},
		},
		MetricThreshold: &atlas.MetricThreshold{
			MetricName: "ASSERT_REGULAR",
			Operator:   "LESS_THAN",
			Threshold:  99.0,
			Units:      "RAW",
			Mode:       "AVERAGE",
		},
	}
}

func AlertsResponse() *atlas.AlertsResponse {
	return &atlas.AlertsResponse{
		Links:      nil,
		Results:    []atlas.Alert{*Alert()},
		TotalCount: 1,
	}
}
