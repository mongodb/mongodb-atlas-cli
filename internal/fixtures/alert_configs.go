package fixtures

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

func AlertConfig() *atlas.AlertConfiguration{
	var enabled = true
	var delayMin int = 0
	return &atlas.AlertConfiguration{
		ID:                     "533dc40ae4b00835ff81",
		GroupID:                "535683b3794d371327b",
		AlertConfigID:          "22",
		EventTypeName:          "OUTSIDE_METRIC_THRESHOLD",
		Created:                "2016-08-23T20:26:50Z",
		Updated:                "2016-08-23T20:26:50Z",
		Enabled:                &enabled,
		ClusterID:              "1",
		ClusterName:            "REPLICASET",
		Matchers:               []atlas.Matcher{{
			FieldName: "HOSTNAME_AND_PORT",
			Operator:  "EQUALS",
			Value:     "mongo.example.com:27017",
		}},
		MetricThreshold:		&atlas.MetricThreshold{
			MetricName: "ASSERT_REGULAR",
			Operator:   "LESS_THAN",
			Threshold:  99.0,
			Units:      "RAW",
			Mode:       "RAW",
		},
		Notifications:      []atlas.Notification{{
			DelayMin:            &delayMin,
			IntervalMin:         5,
			MobileNumber:        "2343454567",
			TypeName:            "SMS",
		}},
	}
}

func AlertConfigs() []atlas.AlertConfiguration{
	return []atlas.AlertConfiguration{*AlertConfig()}
}
