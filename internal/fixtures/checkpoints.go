package fixtures

import atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"

func Checkpoint() *atlas.Checkpoint {
	return &atlas.Checkpoint{
		ClusterID: "6b8cd61180eef547110159d9",
		Completed: "2018-02-08T23:20:25Z",
		GroupID:   "6b8cd3c380eef5349ef77gf7",
		ID:        "5a7cdb3980eef53de5bffdcf",
		Links: []*atlas.Link{
			{
				Rel:  "self",
				Href: "https://cloud.mongodb.com/api/public/v1.0/groups/6b8cd3c380eef5349ef77gf7/clusters/Cluster0/checkpoints",
			},
		},
		Parts: []*atlas.Part{

			{
				ReplicaSetName: "Cluster0-shard-1",
				TypeName:       "REPLICA_SET",
				CheckpointPart: atlas.CheckpointPart{
					ShardName:       "Cluster0-shard-1",
					TokenDiscovered: true,
					TokenTimestamp: atlas.SnapshotTimestamp{
						Date:      "2018-02-08T23:20:25Z",
						Increment: 1,
					},
				},
			},
			{
				ReplicaSetName: "Cluster0-shard-0",
				TypeName:       "REPLICA_SET",
				CheckpointPart: atlas.CheckpointPart{
					ShardName:       "Cluster0-shard-0",
					TokenDiscovered: true,
					TokenTimestamp: atlas.SnapshotTimestamp{
						Date:      "2018-02-08T23:20:25Z",
						Increment: 1,
					}},
			},
			{
				ReplicaSetName: "Cluster0-config-0",
				TypeName:       "CONFIG_SERVER_REPLICA_SET",
				CheckpointPart: atlas.CheckpointPart{
					TokenDiscovered: true,
					TokenTimestamp: atlas.SnapshotTimestamp{
						Date:      "2018-02-08T23:20:25Z",
						Increment: 2,
					}},
			},
		},
		Restorable: true,
		Started:    "2018-02-08T23:20:25Z",
		Timestamp:  "2018-02-08T23:19:37Z",
	}
}

func Checkpoints() *atlas.Checkpoints {
	return &atlas.Checkpoints{
		Results:    []*atlas.Checkpoint{Checkpoint()},
		Links:      nil,
		TotalCount: 1,
	}
}
