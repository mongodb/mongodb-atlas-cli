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

//go:generate mockgen -destination=../mocks/api_keys_whitelist.go -package=mocks github.com/mongodb/mongocli/internal/store OrganizationAPIKeyWhitelistLister,OrganizationAPIKeyWhitelistCreator,OrganizationAPIKeyWhitelistDeleter

type OrganizationAPIKeyWhitelistLister interface {
	OrganizationAPIKeyWhitelists(string, string, *atlas.ListOptions) (*atlas.WhitelistAPIKeys, error)
}

type OrganizationAPIKeyWhitelistCreator interface {
	CreateOrganizationAPIKeyWhite(string, string, []*atlas.WhitelistAPIKeysReq) (*atlas.WhitelistAPIKeys, error)
}

type OrganizationAPIKeyWhitelistDeleter interface {
	DeleteOrganizationAPIKeyWhitelist(string, string, string) error
}

// OrganizationAPIKeys encapsulates the logic to manage different cloud providers
func (s *Store) OrganizationAPIKeyWhitelists(orgID, apiKeyID string, opts *atlas.ListOptions) (*atlas.WhitelistAPIKeys, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).WhitelistAPIKeys.List(context.Background(), orgID, apiKeyID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).WhitelistAPIKeys.List(context.Background(), orgID, apiKeyID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateOrganizationAPIKeyWhite encapsulates the logic to manage different cloud providers
func (s *Store) CreateOrganizationAPIKeyWhite(orgID, apiKeyID string, opts []*atlas.WhitelistAPIKeysReq) (*atlas.WhitelistAPIKeys, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).WhitelistAPIKeys.Create(context.Background(), orgID, apiKeyID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).WhitelistAPIKeys.Create(context.Background(), orgID, apiKeyID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteOrganizationAPIKeyWhitelist encapsulates the logic to manage different cloud providers
func (s *Store) DeleteOrganizationAPIKeyWhitelist(orgID, apiKeyID, ipAddress string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).WhitelistAPIKeys.Delete(context.Background(), orgID, apiKeyID, ipAddress)
		return err
	case config.OpsManagerService, config.CloudManagerService:
		_, err := s.client.(*opsmngr.Client).WhitelistAPIKeys.Delete(context.Background(), orgID, apiKeyID, ipAddress)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
