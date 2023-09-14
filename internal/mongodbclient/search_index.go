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

	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	listSearchIndexes = "$listSearchIndexes"
	idField           = "id"
)

var ErrSearchIndexNotFound = errors.New("search Index not found")

type SearchIndex interface {
	CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) error
	SearchIndex(ctx context.Context, id string) (*SearchIndexDefinition, error)
}

type SearchIndexDefinition struct {
	ID         string                                 `json:"id"`
	Name       string                                 `json:"name"`
	Collection string                                 `json:"collection"`
	Database   string                                 `json:"database"`
	Analyzer   *string                                `json:"analyzer,omitempty"`
	Analyzers  []admin.ApiAtlasFTSAnalyzers           `json:"analyzers,omitempty"`
	Synonyms   []admin.SearchSynonymMappingDefinition `json:"synonyms,omitempty"`
	Mappings   *admin.ApiAtlasFTSMappings             `json:"mappings,omitempty"`
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
		if err != nil {
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
