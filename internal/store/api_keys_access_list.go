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
	"errors"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_api_keys_access_list.go -package=mocks github.com/mongodb/mongocli/internal/store OrganizationAPIKeyAccessListCreator,OrganizationAPIKeyAccessListDeleter,OrganizationAPIKeyAccessListLister

const resourceNotFound = 404

type OrganizationAPIKeyAccessListLister interface {
	OrganizationAPIKeyAccessLists(string, string, *atlas.ListOptions) (*atlas.AccessListAPIKeys, error)
}

type OrganizationAPIKeyAccessListDeleter interface {
	DeleteOrganizationAPIKeyAccessList(string, string, string) error
}

type OrganizationAPIKeyAccessListCreator interface {
	CreateOrganizationAPIKeyAccessList(string, string, []*atlas.AccessListAPIKeysReq) (*atlas.AccessListAPIKeys, error)
}

// CreateOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers
func (s *Store) CreateOrganizationAPIKeyAccessList(orgID, apiKeyID string, opts []*atlas.AccessListAPIKeysReq) (*atlas.AccessListAPIKeys, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).AccessListAPIKeys.Create(context.Background(), orgID, apiKeyID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).AccessListAPIKeys.Create(context.Background(), orgID, apiKeyID, opts)
		var target *atlas.ErrorResponse
		if err != nil && errors.As(err, &target) {
			// We keep supporting OM 4.2 and OM 4.4
			if target.HTTPCode == resourceNotFound {
				result, _, e := s.client.(*opsmngr.Client).WhitelistAPIKeys.Create(context.Background(), orgID, apiKeyID, convert.FromAccessListAPIKeysReqToWhitelistAPIKeysReq(opts))
				return convert.FromWhitelistAPIKeysToAccessListAPIKeys(result), e
			}
		}

		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers
func (s *Store) DeleteOrganizationAPIKeyAccessList(orgID, apiKeyID, ipAddress string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).AccessListAPIKeys.Delete(context.Background(), orgID, apiKeyID, ipAddress)
		return err
	case config.OpsManagerService, config.CloudManagerService:
		_, err := s.client.(*opsmngr.Client).AccessListAPIKeys.Delete(context.Background(), orgID, apiKeyID, ipAddress)
		var target *atlas.ErrorResponse
		if err != nil && errors.As(err, &target) {
			// We keep supporting OM 4.2 and OM 4.4
			if target.HTTPCode == resourceNotFound {
				_, e := s.client.(*opsmngr.Client).WhitelistAPIKeys.Delete(context.Background(), orgID, apiKeyID, ipAddress)
				return e
			}
		}
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// OrganizationAPIKeyAccessLists encapsulates the logic to manage different cloud providers
func (s *Store) OrganizationAPIKeyAccessLists(orgID, apiKeyID string, opts *atlas.ListOptions) (*atlas.AccessListAPIKeys, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).AccessListAPIKeys.List(context.Background(), orgID, apiKeyID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).AccessListAPIKeys.List(context.Background(), orgID, apiKeyID, opts)
		var target *atlas.ErrorResponse
		if err != nil && errors.As(err, &target) {
			// We keep supporting OM 4.2 and OM 4.4
			if target.HTTPCode == resourceNotFound {
				result, _, e := s.client.(*opsmngr.Client).WhitelistAPIKeys.List(context.Background(), orgID, apiKeyID, opts)
				return convert.FromWhitelistAPIKeysToAccessListAPIKeys(result), e
			}
		}
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
