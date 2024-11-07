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
	"encoding/json"
	"errors"
	"slices"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"go.mongodb.org/atlas-sdk/v20240805005/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	listSearchIndexes = "$listSearchIndexes"
	addFields         = "$addFields"
	idField           = "id"
	collectionField   = "collection"
	databaseField     = "database"
)

var ErrSearchIndexNotFound = errors.New("search Index not found")

type Collection interface {
	Aggregate(context.Context, any) (*mongo.Cursor, error)
	CreateSearchIndex(ctx context.Context, idx *admin.ClusterSearchIndex) (*admin.ClusterSearchIndex, error)
	SearchIndexes(ctx context.Context) ([]*admin.ClusterSearchIndex, error)
	SearchIndexByName(ctx context.Context, name string) (*admin.ClusterSearchIndex, error)
}

type collection struct {
	collection *mongo.Collection
}

func (c *collection) Aggregate(ctx context.Context, pipeline any) (*mongo.Cursor, error) {
	return c.collection.Aggregate(ctx, pipeline)
}

type SearchIndexDefinition struct {
	ID         string                                 `bson:"id,omitempty"`
	Name       string                                 `bson:"name,omitempty"`
	Collection string                                 `bson:"collection,omitempty"`
	Database   string                                 `bson:"database,omitempty"`
	Analyzer   *string                                `bson:"analyzer,omitempty"`
	Analyzers  []admin.ApiAtlasFTSAnalyzers           `bson:"analyzers,omitempty"`
	Synonyms   []admin.SearchSynonymMappingDefinition `bson:"synonyms,omitempty"`
	Mappings   *admin.ApiAtlasFTSMappings             `bson:"mappings,omitempty"`
	Status     *string                                `bson:"status,omitempty"`
}

func (c *collection) CreateSearchIndex(ctx context.Context, idx *admin.ClusterSearchIndex) (*admin.ClusterSearchIndex, error) {
	// To maintain formatting of the SDK, marshal object into JSON and then unmarshal into BSON
	jsonIndex, err := json.Marshal(idx)
	if err != nil {
		return nil, err
	}

	var index bson.D
	err = bson.UnmarshalExtJSON(jsonIndex, true, &index)
	if err != nil {
		return nil, err
	}

	// Empty these fields so that they are not included into the index definition for the MongoDB command
	index = removeFields(index, "id", "collectionName", "database", "type")

	o := options.SearchIndexes().SetName(idx.Name)
	if idx.Type != nil {
		o.SetType(*idx.Type)
	}

	_, err = c.collection.SearchIndexes().CreateOne(ctx, mongo.SearchIndexModel{
		Definition: index,
		Options:    o,
	})

	_, _ = log.Debugln("Creating search index with definition: ", index)
	if err != nil {
		return nil, err
	}

	return c.SearchIndexByName(ctx, idx.Name)
}

func removeFields(doc bson.D, fields ...string) bson.D {
	cleanedDoc := bson.D{}

	for _, elem := range doc {
		if slices.Contains(fields, elem.Key) {
			continue
		}

		cleanedDoc = append(cleanedDoc, elem)
	}

	return cleanedDoc
}

func (c *collection) SearchIndexByName(ctx context.Context, name string) (*admin.ClusterSearchIndex, error) {
	indexes, err := c.SearchIndexes(ctx)
	if err != nil {
		return nil, err
	}

	for _, index := range indexes {
		if index.Name == name && index.Database == c.collection.Database().Name() {
			return index, nil
		}
	}

	return nil, ErrSearchIndexNotFound
}

func (c *collection) SearchIndexes(ctx context.Context) ([]*admin.ClusterSearchIndex, error) {
	cursor, err := c.Aggregate(ctx, newSearchIndexesPipeline(c.collection.Database().Name(), c.collection.Name()))
	if err != nil || cursor == nil {
		return nil, err
	}

	var results []*SearchIndexDefinition
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return newClusterSearchIndexes(results), nil
}

func newSearchIndexesPipeline(db, coll string) []*bson.D {
	return []*bson.D{
		{
			{
				Key: listSearchIndexes, Value: bson.D{},
			},
		},
		{
			{
				Key: addFields, Value: bson.D{
					{
						Key: collectionField, Value: coll,
					},
					{
						Key: databaseField, Value: db,
					},
				},
			},
		},
	}
}

func newClusterSearchIndex(index *SearchIndexDefinition) *admin.ClusterSearchIndex {
	return &admin.ClusterSearchIndex{
		Name:           index.Name,
		IndexID:        &index.ID,
		CollectionName: index.Collection,
		Database:       index.Database,
		Status:         index.Status,
	}
}

func newClusterSearchIndexes(indexes []*SearchIndexDefinition) []*admin.ClusterSearchIndex {
	out := make([]*admin.ClusterSearchIndex, len(indexes))

	for i, v := range indexes {
		out[i] = newClusterSearchIndex(v)
	}

	return out
}
