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

//go:generate mockgen -destination=../mocks/mock_api_keys_access_list.go -package=mocks github.com/mongodb/mongocli/internal/store OrganizationAPIKeyAccessListCreator,OrganizationAPIKeyAccessListDeleter,OrganizationAPIKeyAccessListLister

type OrganizationAPIKeyAccessListLister interface {
	OrganizationAPIKeyAccessLists(string, string, *atlas.ListOptions) (*atlas.AccessListAPIKeys, error)
}

type OrganizationAPIKeyAccessListDeleter interface {
	DeleteOrganizationAPIKeyAccessList(string, string, string) error
}

type OrganizationAPIKeyAccessListCreator interface {
	CreateOrganizationAPIKeyAccessList(string, string, []*atlas.AccessListAPIKeysReq) (*atlas.AccessListAPIKeys, error)
}

// CreateOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOrganizationAPIKeyAccessList(orgID, apiKeyID string, opts []*atlas.AccessListAPIKeysReq) (*atlas.AccessListAPIKeys, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).AccessListAPIKeys.Create(context.Background(), orgID, apiKeyID, opts)
		return result, err
	case config.CloudManagerService:
		return createViaAccessList(context.Background(), s, orgID, apiKeyID, opts)
	case config.OpsManagerService:
		omIsAtLeastFive, err := checkOMIsAtLeastFive(s)
		if err != nil {
			return nil, err
		}
		if omIsAtLeastFive {
			return createViaAccessList(context.Background(), s, orgID, apiKeyID, opts)
		}
		return createViaWhitelist(context.Background(), s, orgID, apiKeyID, opts)
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func createViaAccessList(ctx context.Context, s *Store, orgID, apiKeyID string, opts []*atlas.AccessListAPIKeysReq) (*atlas.AccessListAPIKeys, error) {
	result, _, err := s.client.(*opsmngr.Client).AccessListAPIKeys.Create(ctx, orgID, apiKeyID, opts)
	return result, err
}

func createViaWhitelist(ctx context.Context, s *Store, orgID, apiKeyID string, opts []*atlas.AccessListAPIKeysReq) (*atlas.AccessListAPIKeys, error) {
	result, _, err := s.client.(*opsmngr.Client).WhitelistAPIKeys.Create(ctx, orgID, apiKeyID, fromAccessListAPIKeysReqToWhitelistAPIKeysReq(opts))
	return fromWhitelistAPIKeysToAccessListAPIKeys(result), err
}

// DeleteOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOrganizationAPIKeyAccessList(orgID, apiKeyID, ipAddress string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.client.(*atlas.Client).AccessListAPIKeys.Delete(context.Background(), orgID, apiKeyID, ipAddress)
		return err
	case config.CloudManagerService:
		return deleteViaAccessList(context.Background(), s, orgID, apiKeyID, ipAddress)
	case config.OpsManagerService:
		omIsAtLeastFive, err := checkOMIsAtLeastFive(s)
		if err != nil {
			return err
		}
		if omIsAtLeastFive {
			return deleteViaAccessList(context.Background(), s, orgID, apiKeyID, ipAddress)
		}

		return deleteViaWhitelist(context.Background(), s, orgID, apiKeyID, ipAddress)
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func deleteViaAccessList(ctx context.Context, s *Store, orgID, apiKeyID, ipAddress string) error {
	_, err := s.client.(*opsmngr.Client).AccessListAPIKeys.Delete(ctx, orgID, apiKeyID, ipAddress)
	return err
}

func deleteViaWhitelist(ctx context.Context, s *Store, orgID, apiKeyID, ipAddress string) error {
	_, err := s.client.(*opsmngr.Client).AccessListAPIKeys.Delete(ctx, orgID, apiKeyID, ipAddress)
	return err
}

// OrganizationAPIKeyAccessLists encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKeyAccessLists(orgID, apiKeyID string, opts *atlas.ListOptions) (*atlas.AccessListAPIKeys, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).AccessListAPIKeys.List(context.Background(), orgID, apiKeyID, opts)
		return result, err
	case config.CloudManagerService:
		return listViaAccessList(context.Background(), s, orgID, apiKeyID, opts)
	case config.OpsManagerService:
		omIsAtLeastFive, err := checkOMIsAtLeastFive(s)
		if err != nil {
			return nil, err
		}
		if omIsAtLeastFive {
			return listViaAccessList(context.Background(), s, orgID, apiKeyID, opts)
		}

		return listViaWhitelist(context.Background(), s, orgID, apiKeyID, opts)
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func listViaAccessList(ctx context.Context, s *Store, orgID, apiKeyID string, opts *atlas.ListOptions) (*atlas.AccessListAPIKeys, error) {
	result, _, err := s.client.(*opsmngr.Client).AccessListAPIKeys.List(ctx, orgID, apiKeyID, opts)
	return result, err
}

func listViaWhitelist(ctx context.Context, s *Store, orgID, apiKeyID string, opts *atlas.ListOptions) (*atlas.AccessListAPIKeys, error) {
	result, _, e := s.client.(*opsmngr.Client).WhitelistAPIKeys.List(ctx, orgID, apiKeyID, opts)
	return fromWhitelistAPIKeysToAccessListAPIKeys(result), e
}

// fromWhitelistAPIKeysToAccessListAPIKeys convert from atlas.WhitelistAPIKeys format to atlas.AccessListAPIKeys
// We use this function with whitelist endpoints to keep supporting OM 4.2 and OM 4.4.
func fromWhitelistAPIKeysToAccessListAPIKeys(in *atlas.WhitelistAPIKeys) *atlas.AccessListAPIKeys {
	if in == nil {
		return nil
	}

	out := &atlas.AccessListAPIKeys{
		TotalCount: in.TotalCount,
		Links:      in.Links,
	}

	results := make([]*atlas.AccessListAPIKey, len(in.Results))
	for i, element := range in.Results {
		results[i] = fromWhitelistAPIKeyToAccessListAPIKey(element)
	}

	out.Results = results
	return out
}

// fromWhitelistAPIKeyToAccessListAPIKey convert from atlas.WhitelistAPIKey format to atlas.AccessListAPIKey
// We use this function with whitelist endpoints to keep supporting OM 4.2 and OM 4.4.
func fromWhitelistAPIKeyToAccessListAPIKey(in *atlas.WhitelistAPIKey) *atlas.AccessListAPIKey {
	return &atlas.AccessListAPIKey{
		CidrBlock:       in.CidrBlock,
		Count:           in.Count,
		Created:         in.Created,
		IPAddress:       in.IPAddress,
		LastUsed:        in.LastUsed,
		LastUsedAddress: in.LastUsedAddress,
		Links:           in.Links,
	}
}

// fromAccessListAPIKeysReqToWhitelistAPIKeysReq convert from atlas.AccessListAPIKeysReq format to atlas.WhitelistAPIKeysReq
// We use this function with whitelist endpoints to keep supporting OM 4.2 and OM 4.4.
func fromAccessListAPIKeysReqToWhitelistAPIKeysReq(in []*atlas.AccessListAPIKeysReq) []*atlas.WhitelistAPIKeysReq {
	if in == nil {
		return nil
	}

	out := make([]*atlas.WhitelistAPIKeysReq, len(in))

	for i, element := range in {
		accessListElement := &atlas.WhitelistAPIKeysReq{
			IPAddress: element.IPAddress,
			CidrBlock: element.CidrBlock,
		}
		out[i] = accessListElement
	}
	return out
}
