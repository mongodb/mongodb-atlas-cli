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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	listSearchIndexes      = "$listSearchIndexes"
	addFields              = "$addFields"
	idField                = "id"
	collectionField        = "collection"
	databaseField          = "database"
	defaultSearchIndexType = "search"
)

var ErrSearchIndexNotFound = errors.New("search Index not found")

type SearchIndex interface {
	CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) (*admin.ClusterSearchIndex, error)
	SearchIndex(ctx context.Context, id string) (*admin.ClusterSearchIndex, error)
	SearchIndexes(ctx context.Context, coll string) ([]*admin.ClusterSearchIndex, error)
	SearchIndexByName(ctx context.Context, name string, collection string) (*admin.ClusterSearchIndex, error)
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

// todo: CLOUDP-199915 Use go-driver search index management helpers instead of createSearchIndex command
func (o *database) CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) (*admin.ClusterSearchIndex, error) {
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

	indexCommand := bson.D{
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
						Key:   "type",
						Value: pointer.GetOrDefault(idx.Type, defaultSearchIndexType),
					},
					{
						Key:   "definition",
						Value: index,
					},
				},
			},
		},
	}

	_, _ = log.Debugln("Creating search index with definition: ", index)
	if result := o.db.RunCommand(ctx, indexCommand); result.Err() != nil {
		return nil, result.Err()
	}

	return o.SearchIndexByName(ctx, idx.Name, collection)
}

func removeFields(doc bson.D, fields ...string) bson.D {
	cleanedDoc := bson.D{}

	for _, elem := range doc {
		if search.StringInSlice(fields, elem.Key) {
			continue
		}

		cleanedDoc = append(cleanedDoc, elem)
	}

	return cleanedDoc
}

func (o *database) SearchIndex(ctx context.Context, id string) (*admin.ClusterSearchIndex, error) {
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
			searchIndexDef := &SearchIndexDefinition{
				ID:         results[0].ID,
				Name:       results[0].Name,
				Collection: coll,
				Database:   o.db.Name(),
				Status:     results[0].Status,
			}
			return newClusterSearchIndex(searchIndexDef), nil
		}
	}

	return nil, fmt.Errorf("index `%s` not found: %w", id, ErrSearchIndexNotFound)
}

func (o *database) SearchIndexByName(ctx context.Context, name string, collection string) (*admin.ClusterSearchIndex, error) {
	indexes, err := o.SearchIndexes(ctx, collection)
	if err != nil {
		return nil, err
	}

	for _, index := range indexes {
		if index.Name == name && index.Database == o.db.Name() {
			return index, nil
		}
	}

	return nil, ErrSearchIndexNotFound
}

func (o *database) SearchIndexes(ctx context.Context, coll string) ([]*admin.ClusterSearchIndex, error) {
	cursor, err := o.db.Collection(coll).Aggregate(ctx, newSearchIndexesPipeline(o.db.Name(), coll))
	if err != nil || cursor == nil {
		return nil, err
	}

	var results []*SearchIndexDefinition
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return newClusterSearchIndexes(results), nil
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
