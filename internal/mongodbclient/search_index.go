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
	"fmt"

	"go.mongodb.org/atlas-sdk/v20230201007/admin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	listSearchIndexes = "$listSearchIndexes"
	addFields         = "$addFields"
	idField           = "id"
	collectionField   = "collection"
	databaseField     = "database"
)

var ErrSearchIndexNotFound = errors.New("search Index not found")

type SearchIndex interface {
	CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) error
	SearchIndex(ctx context.Context, id string) (*SearchIndexDefinition, error)
	SearchIndexes(ctx context.Context, coll string) ([]*SearchIndexDefinition, error)
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
}

func (o *database) CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) error {
	// todo: CLOUDP-199915 Use go-driver search index management helpers instead of createSearchIndex command
	return o.db.RunCommand(ctx, bson.D{
		{
			Key:   "createSearchIndexes",
			Value: collection,
		},
		{
			Key: "indexes",
			Value: []bson.D{
				{
					{
						Key:   "name",
						Value: idx.Name,
					},
					{
						Key: "definition",
						Value: &SearchIndexDefinition{
							Name:      idx.Name,
							Analyzer:  idx.Analyzer,
							Analyzers: idx.Analyzers,
							Mappings:  idx.Mappings,
							Synonyms:  idx.Synonyms,
						},
					},
				},
			},
		},
	}).Err()
}

func (o *database) SearchIndex(ctx context.Context, id string) (*SearchIndexDefinition, error) {
	collectionNames, err := o.db.ListCollectionNames(ctx, bson.D{}, nil)
	if err != nil {
		return nil, err
	}

	// We search the index in all the collections of the database
	for _, coll := range collectionNames {
		cursor, err := o.db.Collection(coll).Aggregate(ctx, newSearchIndexPipeline(id))
		if err != nil || cursor == nil {
			return nil, err
		}
		var results []SearchIndexDefinition
		if err = cursor.All(ctx, &results); err != nil {
			return nil, err
		}
		if len(results) >= 1 {
			return &SearchIndexDefinition{
				ID:         results[0].ID,
				Name:       results[0].Name,
				Collection: coll,
				Database:   o.db.Name(),
			}, nil
		}
	}

	return nil, fmt.Errorf("index `%s` not found: %w", id, ErrSearchIndexNotFound)
}

func (o *database) SearchIndexes(ctx context.Context, coll string) ([]*SearchIndexDefinition, error) {
	cursor, err := o.db.Collection(coll).Aggregate(ctx, newSearchIndexesPipeline(o.db.Name(), coll))
	if err != nil || cursor == nil {
		return nil, err
	}

	var results []*SearchIndexDefinition
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
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
