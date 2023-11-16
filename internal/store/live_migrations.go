// Copyright 2021 MongoDB Inc
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

package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115001/admin"
)

//go:generate mockgen -destination=../mocks/mock_live_migrations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store LiveMigrationCreator,LiveMigrationDescriber

type LiveMigrationCreator interface {
	LiveMigrationCreate(string, *atlasv2.LiveMigrationRequest) (*atlasv2.LiveMigrationResponse, error)
}

type LiveMigrationDescriber interface {
	LiveMigrationDescribe(string, string) (*atlasv2.LiveMigrationResponse, error)
}

// LiveMigrationCreate encapsulates the logic to manage different cloud providers.
func (s *Store) LiveMigrationCreate(groupID string, liveMigrationRequest *atlasv2.LiveMigrationRequest) (*atlasv2.LiveMigrationResponse, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.CloudMigrationServiceApi.CreatePushMigration(context.Background(), groupID, liveMigrationRequest).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// LiveMigrationDescribe encapsulates the logic to manage different cloud providers.
func (s *Store) LiveMigrationDescribe(groupID, migrationID string) (*atlasv2.LiveMigrationResponse, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.CloudMigrationServiceApi.GetPushMigration(context.Background(), groupID, migrationID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
