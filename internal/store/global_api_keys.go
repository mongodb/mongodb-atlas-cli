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
	"fmt"

	"github.com/andreangiolillo/mongocli-test/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_global_api_keys.go -package=mocks github.com/andreangiolillo/mongocli-test/internal/store GlobalAPIKeyLister,GlobalAPIKeyDescriber,GlobalAPIKeyUpdater,GlobalAPIKeyCreator,GlobalAPIKeyDeleter

type GlobalAPIKeyLister interface {
	GlobalAPIKeys(*opsmngr.ListOptions) ([]opsmngr.APIKey, error)
}

type GlobalAPIKeyDescriber interface {
	GlobalAPIKey(string) (*opsmngr.APIKey, error)
}

type GlobalAPIKeyUpdater interface {
	UpdateGlobalAPIKey(string, *opsmngr.APIKeyInput) (*opsmngr.APIKey, error)
}

type GlobalAPIKeyCreator interface {
	CreateGlobalAPIKey(*opsmngr.APIKeyInput) (*opsmngr.APIKey, error)
}

type GlobalAPIKeyDeleter interface {
	DeleteGlobalAPIKey(string) error
}

// GlobalAPIKeys encapsulates the logic to manage different cloud providers.
func (s *Store) GlobalAPIKeys(opts *opsmngr.ListOptions) ([]opsmngr.APIKey, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).GlobalAPIKeys.List(s.ctx, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// GlobalAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) GlobalAPIKey(apiKeyID string) (*opsmngr.APIKey, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).GlobalAPIKeys.Get(s.ctx, apiKeyID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UpdateGlobalAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateGlobalAPIKey(apiKeyID string, input *atlas.APIKeyInput) (*opsmngr.APIKey, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).GlobalAPIKeys.Update(s.ctx, apiKeyID, input)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateGlobalAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) CreateGlobalAPIKey(input *opsmngr.APIKeyInput) (*opsmngr.APIKey, error) {
	switch s.service {
	case config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).GlobalAPIKeys.Create(s.ctx, input)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteGlobalAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteGlobalAPIKey(id string) error {
	switch s.service {
	case config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).GlobalAPIKeys.Delete(s.ctx, id)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
