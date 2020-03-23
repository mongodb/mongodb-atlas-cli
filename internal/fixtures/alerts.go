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
