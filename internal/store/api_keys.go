// Copyright 2020 MongoDB Inc
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

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/api_keys.go -package=mocks github.com/mongodb/mongocli/internal/store APIKeyLister,APIKeyDeleter

type APIKeyLister interface {
	APIKeys(string, *atlas.ListOptions) ([]atlas.APIKey, error)
}

type APIKeyDeleter interface {
	DeleteAPIKeys(string, string) error
}

// APIKeys encapsulate the logic to manage different cloud providers
func (s *Store) APIKeys(orgID string, opts *atlas.ListOptions) ([]atlas.APIKey, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).APIKeys.List(context.Background(), orgID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).OrganizationAPIKeys.List(context.Background(), orgID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteAPIKeys encapsulate the logic to manage different cloud providers
func (s *Store) DeleteAPIKeys(orgID, id string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).APIKeys.Delete(context.Background(), orgID, id)
		return err
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).OrganizationAPIKeys.Delete(context.Background(), orgID, id)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
