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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_live_migration_validations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store LiveMigrationValidationsCreator,LiveMigrationCutoverCreator,LiveMigrationValidationsDescriber

type LiveMigrationValidationsCreator interface {
	CreateValidation(string, *atlasv2.LiveMigrationRequest20240530) (*atlasv2.LiveImportValidation, error)
}

type LiveMigrationCutoverCreator interface {
	CreateLiveMigrationCutover(string, string) error
}

type LiveMigrationValidationsDescriber interface {
	GetValidationStatus(string, string) (*atlasv2.LiveImportValidation, error)
}

// CreateValidation encapsulate the logic to manage different cloud providers.
func (s *Store) CreateValidation(groupID string, liveMigration *atlasv2.LiveMigrationRequest20240530) (*atlasv2.LiveImportValidation, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.CloudMigrationServiceApi.ValidateMigration(s.ctx, groupID, liveMigration).Execute()
	return result, err
}

// CreateLiveMigrationCutover encapsulate the logic to manage different cloud providers.
func (s *Store) CreateLiveMigrationCutover(groupID, liveMigrationID string) error {
	if s.service == config.CloudGovService {
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	_, err := s.clientv2.CloudMigrationServiceApi.CutoverMigration(s.ctx, groupID, liveMigrationID).Execute()
	return err
}

// GetValidationStatus encapsulate the logic to manage different cloud providers.
func (s *Store) GetValidationStatus(groupID, liveMigrationID string) (*atlasv2.LiveImportValidation, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.CloudMigrationServiceApi.GetValidationStatus(context.Background(), groupID, liveMigrationID).Execute()
	return result, err
}
