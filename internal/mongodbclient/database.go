// Copyright 2023 MongoDB Inc
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

package mongodbclient

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	RunCommand(ctx context.Context, runCommand interface{}) (interface{}, error)
	InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	InitiateReplicaSet(ctx context.Context, rsName string, hostname string, internalPort int, externalPort int) error
	SearchIndex
	User
	Collection(string) Collection
}

type database struct {
	db *mongo.Database
}

func (o *database) Collection(name string) Collection {
	return &collection{
		collection: o.db.Collection(name),
	}
}

func (o *database) RunCommand(ctx context.Context, runCmd interface{}) (interface{}, error) {
	r := o.db.RunCommand(ctx, runCmd)
	if err := r.Err(); err != nil {
		return nil, err
	}

	var cmdResult bson.M
	if err := r.Decode(&cmdResult); err != nil {
		return nil, err
	}
	return cmdResult, nil
}

func (o *database) InsertOne(ctx context.Context, col string, doc interface{}) (interface{}, error) {
	return o.db.Collection(col).InsertOne(ctx, doc)
}

func (o *database) InitiateReplicaSet(ctx context.Context, rsName string, hostname string, internalPort int, externalPort int) error {
	return o.db.RunCommand(ctx, bson.D{{Key: "replSetInitiate", Value: bson.M{
		"_id":       rsName,
		"version":   1,
		"configsvr": false,
		"members": []bson.M{
			{
				"_id":  0,
				"host": fmt.Sprintf("%s:%d", hostname, internalPort),
				"horizons": bson.M{
					"external": fmt.Sprintf("localhost:%d", externalPort),
				},
			},
		},
	}}}).Err()
}
