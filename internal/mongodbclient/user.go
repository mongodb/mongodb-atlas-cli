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

	"go.mongodb.org/mongo-driver/bson"
)

type User interface {
	CreateUser(ctx context.Context, username, password string, roles []string) error
	DropUser(ctx context.Context, username string) error
}

func (d *database) CreateUser(ctx context.Context, username, password string, roles []string) error {
	rolesDoc := bson.A{}
	for _, r := range roles {
		rolesDoc = append(rolesDoc, bson.D{
			{Key: "role", Value: r},
			{Key: "db", Value: d.db.Name()},
		})
	}

	return d.db.Client().
		Database("admin").
		RunCommand(ctx, bson.D{
			{Key: "createUser", Value: username},
			{Key: "pwd", Value: password},
			{Key: "roles", Value: rolesDoc},
		}).Err()
}

func (d *database) DropUser(ctx context.Context, username string) error {
	return d.db.Client().
		Database("admin").
		RunCommand(ctx, bson.D{
			{Key: "dropUser", Value: username},
		}).Err()
}
