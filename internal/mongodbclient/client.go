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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var errConnectFailed = errors.New("failed to connect to mongodb server")
var errPingFailed = errors.New("failed to ping mongodb server")

//go:generate go tool go.uber.org/mock/mockgen -destination=../mocks/mock_mongodb_client.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongodbclient MongoDBClient,Database,Collection

type MongoDBClient interface {
	Connect(ctx context.Context, connectionString string, waitSeconds int64) error
	Disconnect(ctx context.Context) error
	Database(db string) Database
	SearchIndex(ctx context.Context, id string) (*SearchIndexDefinition, error)
}

type mongodbClient struct {
	client *mongo.Client
}

func NewClient() MongoDBClient {
	return &mongodbClient{}
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
		_ = client.Disconnect(ctx)
		o.client = nil
		return fmt.Errorf("%w: %w", errPingFailed, err)
	}

	return nil
}

func (o *mongodbClient) SearchIndex(ctx context.Context, id string) (*SearchIndexDefinition, error) {
	dbs, err := o.client.ListDatabaseNames(ctx, bson.D{}, nil)
	if err != nil {
		return nil, err
	}

	// We search the index in all the databases
	for _, db := range dbs {
		if index, err := o.Database(db).SearchIndex(ctx, id); err == nil && index != nil {
			return index, nil
		}
	}

	return nil, ErrSearchIndexNotFound
}

func (o *mongodbClient) Disconnect(ctx context.Context) error {
	return o.client.Disconnect(ctx)
}

func (o *mongodbClient) Database(name string) Database {
	return &database{db: o.client.Database(name)}
}
