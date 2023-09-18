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

package watchers

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const notFoundState = "NOT_FOUND"

var LocalSearchIndexCreated = &StateTransition{
	EndState: pointer.Get("STEADY"),
}

type LocalSearchIndexStateDescriber struct {
	connectionString string
	indexName        string
	db               string
	collection       string
}

func (d *LocalSearchIndexStateDescriber) GetState() (string, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.connectionString))
	if err != nil {
		return "", err
	}
	defer client.Disconnect(ctx)

	db := client.Database(d.db)
	col := db.Collection(d.collection)
	cursor, err := col.Aggregate(ctx, mongo.Pipeline{
		{
			{Key: "$listSearchIndexes", Value: bson.D{}},
		},
	})
	if err != nil {
		return "", err
	}
	var results []bson.M
	err = cursor.All(ctx, &results)
	if err != nil {
		return "", err
	}
	if len(results) == 0 {
		return notFoundState, nil
	}
	status, ok := results[0]["status"].(string)
	if !ok {
		return notFoundState, nil
	}
	return status, nil
}

func NewLocalSearchIndexStateDescriber(connectionString string, indexName string, db string, collection string) *LocalSearchIndexStateDescriber {
	return &LocalSearchIndexStateDescriber{
		connectionString: connectionString,
		indexName:        indexName,
		db:               db,
		collection:       collection,
	}
}
