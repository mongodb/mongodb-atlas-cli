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
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrSearchIndexNotFound = errors.New("search index not found")

type Collection interface {
	Aggregate(ctx context.Context, pipeline any) (*mongo.Cursor, error)
	CreateSearchIndex(ctx context.Context, name, indexType string, definition any) (*SearchIndexDefinition, error)
	SearchIndexes(ctx context.Context) ([]*SearchIndexDefinition, error)
	SearchIndexByName(ctx context.Context, name string) (*SearchIndexDefinition, error)
	DropSearchIndex(ctx context.Context, name string) error
}

type collection struct {
	collection *mongo.Collection
}

func (c *collection) Aggregate(ctx context.Context, pipeline any) (*mongo.Cursor, error) {
	return c.collection.Aggregate(ctx, pipeline)
}

type SearchIndexDefinition struct {
	IndexID          *string `bson:"id,omitempty"`
	Database         *string
	CollectionName   *string
	Name             *string `bson:"name,omitempty"`
	Type             *string `bson:"type,omitempty"`
	Status           *string `bson:"status,omitempty"`
	Queryable        *bool   `bson:"queryable,omitempty"`
	LatestDefinition any     `bson:"latestDefinition,omitempty"`
	LatestVersion    *int    `bson:"latestVersion,omitempty"`
}

func (c *collection) CreateSearchIndex(ctx context.Context, name, indexType string, definition any) (*SearchIndexDefinition, error) {
	o := options.SearchIndexes().
		SetName(name).
		SetType(indexType)

	model := mongo.SearchIndexModel{
		Definition: definition,
		Options:    o,
	}

	_, err := c.collection.SearchIndexes().CreateOne(ctx, model)

	_, _ = log.Debugln("Creating search index with definition: ", model)
	if err != nil {
		return nil, err
	}

	return c.SearchIndexByName(ctx, name)
}

func (c *collection) SearchIndexByName(ctx context.Context, name string) (*SearchIndexDefinition, error) {
	cursor, err := c.collection.SearchIndexes().List(ctx, &options.SearchIndexesOptions{
		Name: &name,
	})
	if err != nil || cursor == nil {
		return nil, err
	}

	var results []*SearchIndexDefinition
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, ErrSearchIndexNotFound
	}

	result := results[0]

	coll := c.collection.Name()
	db := c.collection.Database().Name()
	result.CollectionName = &coll
	result.Database = &db

	return result, nil
}

func (c *collection) SearchIndexes(ctx context.Context) ([]*SearchIndexDefinition, error) {
	cursor, err := c.collection.SearchIndexes().List(ctx, nil)
	if err != nil || cursor == nil {
		return nil, err
	}

	var results []*SearchIndexDefinition
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	for _, result := range results {
		coll := c.collection.Name()
		db := c.collection.Database().Name()
		result.CollectionName = &coll
		result.Database = &db
	}

	return results, nil
}

func (c *collection) DropSearchIndex(ctx context.Context, name string) error {
	return c.collection.SearchIndexes().DropOne(ctx, name)
}
