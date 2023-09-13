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
	"time"

	"go.mongodb.org/atlas-sdk/v20230201006/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var errConnectFailed = errors.New("failed to connect to mongodb server")

//go:generate mockgen -destination=../mocks/mock_mongodb_client.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient MongoDBClient,Database

type MongoDBClient interface {
	Connect(ctx context.Context, connectionString string, waitSeconds int64) error
	Disconnect(ctx context.Context)
	Database(db string) Database
}

type mongodbClient struct {
	client *mongo.Client
}

func NewClient() MongoDBClient {
	return &mongodbClient{}
}

type Database interface {
	RunCommand(ctx context.Context, runCommand interface{}) (interface{}, error)
	InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	InitiateReplicaSet(ctx context.Context, rsName string, hostname string, internalPort int, externalPort int) error
	CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) error
}

type database struct {
	db *mongo.Database
}

func (o *mongodbClient) Connect(ctx context.Context, connectionString string, waitSeconds int64) error {
	ctxConnect, cancel := context.WithTimeout(ctx, time.Duration(waitSeconds)*time.Second)
	defer cancel()

	client, errConnect := mongo.Connect(ctxConnect, options.Client().ApplyURI(connectionString))
	if errConnect != nil {
		return fmt.Errorf("%w: %w", errConnectFailed, errConnect)
	}
	o.client = client

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("%w: %w", errConnectFailed, err)
	}

	return nil
}

func (o *mongodbClient) Disconnect(ctx context.Context) {
	_ = o.client.Disconnect(ctx)
}

func (o *mongodbClient) Database(name string) Database {
	return &database{db: o.client.Database(name)}
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

func (o *database) CreateSearchIndex(ctx context.Context, collection string, idx *admin.ClusterSearchIndex) error {
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
						Value: &searchIndexDefinition{
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

type searchIndexDefinition struct {
	Name      string                                 `json:"name"`
	Analyzer  *string                                `json:"analyzer,omitempty"`
	Analyzers []admin.ApiAtlasFTSAnalyzers           `json:"analyzers,omitempty"`
	Mappings  *admin.ApiAtlasFTSMappings             `json:"mappings,omitempty"`
	Synonyms  []admin.SearchSynonymMappingDefinition `json:"synonyms,omitempty"`
}
