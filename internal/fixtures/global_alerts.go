package fixtures

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
)

func GlobalAlert() *om.GlobalAlert {
	return &om.GlobalAlert{
		Alert: atlas.Alert{
			ID:                    "3b7d2de0a4b02fd2c98146de",
			GroupID:               "1",
			AlertConfigID:         "5730f5e1e4b030a9634a3f69",
			EventTypeName:         "OPLOG_BEHIND",
			Created:               "2016-10-09T06:16:36Z",
			Updated:               "2016-10-10T22:03:11Z",
			Status:                "OPEN",
			LastNotified:          "2016-10-10T20:42:32Z",
			ReplicaSetName:        "shardedCluster-shard-0",
			ClusterName:           "shardedCluster",
			AcknowledgedUntil:     "2016-11-01T00:00:00Z",
			AcknowledgingUsername: "admin@example.com",
		},
		Tags:           []string{},
		Links:          []*atlas.Link{},
		SourceTypeName: "REPLICA_SET",
		ClusterID:      "572a00f2e4b051814b144e90",
	}
}

func GlobalAlerts() *om.GlobalAlerts {
	return &om.GlobalAlerts{
		Links:      []*atlas.Link{},
		Results:    []*om.GlobalAlert{GlobalAlert()},
		TotalCount: 1,
	}
}
