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

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	RunCommand(ctx context.Context, runCommand any) (any, error)
	SearchIndex(ctx context.Context, id string) (*admin.ClusterSearchIndex, error)
	Collection(string) Collection
}

type database struct {
	db *mongo.Database
}

func (d *database) Collection(name string) Collection {
	return &collection{
		collection: d.db.Collection(name),
	}
}

func (d *database) RunCommand(ctx context.Context, runCmd any) (any, error) {
	r := d.db.RunCommand(ctx, runCmd)
	if err := r.Err(); err != nil {
		return nil, err
	}

	var cmdResult bson.M
	if err := r.Decode(&cmdResult); err != nil {
		return nil, err
	}
	return cmdResult, nil
}

func newSearchIndexPipeline(id string) []*bson.D {
	return []*bson.D{
		{
			{
				Key: listSearchIndexes, Value: []bson.E{
					{
						Key: idField, Value: id,
					},
				},
			},
		},
	}
}

func (d *database) SearchIndex(ctx context.Context, id string) (*admin.ClusterSearchIndex, error) {
	collectionNames, err := d.db.ListCollectionNames(ctx, bson.D{}, nil)
	if err != nil {
		return nil, err
	}

	// We search the index in all the collections of the database
	for _, coll := range collectionNames {
		cursor, err := d.db.Collection(coll).Aggregate(ctx, newSearchIndexPipeline(id))
		if err != nil || cursor == nil {
			return nil, err
		}
		var results []SearchIndexDefinition
		if err = cursor.All(ctx, &results); err != nil {
			return nil, err
		}
		if len(results) >= 1 {
			searchIndexDef := &SearchIndexDefinition{
				ID:         results[0].ID,
				Name:       results[0].Name,
				Collection: coll,
				Database:   d.db.Name(),
				Status:     results[0].Status,
			}
			return newClusterSearchIndex(searchIndexDef), nil
		}
	}

	return nil, fmt.Errorf("index `%s` not found: %w", id, ErrSearchIndexNotFound)
}
