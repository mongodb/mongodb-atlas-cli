package fixtures

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

func ContinuousSnapshots() *atlas.ContinuousSnapshots {
	doNotDelete := false

	return &atlas.ContinuousSnapshots{
		Links: []*atlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots?pageNum=1&itemsPerPage=100",
				Rel:  "self",
			},
		},
		Results: []*atlas.ContinuousSnapshot{
			{
				ClusterID: "7c2487d833e9e75286093696",
				Complete:  true,
				Created: &atlas.SnapshotTimestamp{
					Date:      "2017-12-26T16:32:16Z",
					Increment: 1,
				},
				DoNotDelete: &doNotDelete,
				Expires:     "2018-12-25T16:32:16Z",
				GroupID:     "6c7498dg87d9e6526801572b",
				ID:          "5a4279d4fcc178500596745a",
				LastOplogAppliedTimestamp: &atlas.SnapshotTimestamp{
					Date:      "2017-12-26T16:32:15Z",
					Increment: 1,
				},
				Links: []*atlas.Link{
					{
						Href: "https://cloud.mongodb.com/api/atlas/v1.0/groups/6c7498dg87d9e6526801572b/clusters/Cluster0/snapshots/5a4279d4fcc178500596745a",
						Rel:  "self",
					},
				},
				Parts: []*atlas.Part{
					{
						ClusterID:          "7c2487d833e9e75286093696",
						CompressionSetting: "GZIP",
						DataSizeBytes:      4502,
						EncryptionEnabled:  false,
						FileSizeBytes:      324760,
						MongodVersion:      "3.6.10",
						ReplicaSetName:     "Cluster0-shard-0",
						StorageSizeBytes:   53248,
						TypeName:           "REPLICA_SET",
					},
				},
			},
		},
		TotalCount: 1,
	}
}
