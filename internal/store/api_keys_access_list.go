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

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_api_keys_access_list.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store OrganizationAPIKeyAccessListCreator,OrganizationAPIKeyAccessListDeleter,OrganizationAPIKeyAccessListLister

type OrganizationAPIKeyAccessListLister interface {
	OrganizationAPIKeyAccessLists(string, string, *opsmngr.ListOptions) (*opsmngr.AccessListAPIKeys, error)
}

type OrganizationAPIKeyAccessListDeleter interface {
	DeleteOrganizationAPIKeyAccessList(string, string, string) error
}

type OrganizationAPIKeyAccessListCreator interface {
	CreateOrganizationAPIKeyAccessList(string, string, []*opsmngr.AccessListAPIKeysReq) (*opsmngr.AccessListAPIKeys, error)
}

// CreateOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOrganizationAPIKeyAccessList(orgID, apiKeyID string, opts []*opsmngr.AccessListAPIKeysReq) (*opsmngr.AccessListAPIKeys, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.AccessListAPIKeys.Create(s.ctx, orgID, apiKeyID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOrganizationAPIKeyAccessList(orgID, apiKeyID, ipAddress string) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.AccessListAPIKeys.Delete(s.ctx, orgID, apiKeyID, ipAddress)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OrganizationAPIKeyAccessLists encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKeyAccessLists(orgID, apiKeyID string, opts *opsmngr.ListOptions) (*opsmngr.AccessListAPIKeys, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.AccessListAPIKeys.List(s.ctx, orgID, apiKeyID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
