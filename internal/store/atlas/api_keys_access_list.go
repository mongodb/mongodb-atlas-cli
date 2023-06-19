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

package atlas

import (
	"go.mongodb.org/atlas-sdk/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_api_keys_access_list.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas OrganizationAPIKeyAccessListCreator,OrganizationAPIKeyAccessListDeleter,OrganizationAPIKeyAccessListLister

type OrganizationAPIKeyAccessListLister interface {
	OrganizationAPIKeyAccessLists(*admin.ListApiKeyAccessListsEntriesApiParams) (*admin.PaginatedApiUserAccessList, error)
}

type OrganizationAPIKeyAccessListDeleter interface {
	DeleteOrganizationAPIKeyAccessList(string, string, string) error
}

type OrganizationAPIKeyAccessListCreator interface {
	CreateOrganizationAPIKeyAccessList(*admin.CreateApiKeyAccessListApiParams) (*admin.PaginatedApiUserAccessList, error)
}

// CreateOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOrganizationAPIKeyAccessList(params *admin.CreateApiKeyAccessListApiParams) (*admin.PaginatedApiUserAccessList, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.CreateApiKeyAccessListWithParams(s.ctx, params).Execute()
	return result, err
}

// DeleteOrganizationAPIKeyAccessList encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOrganizationAPIKeyAccessList(orgID, apiKeyID, ipAddress string) error {
	_, _, err := s.clientv2.ProgrammaticAPIKeysApi.DeleteApiKeyAccessListEntry(s.ctx, orgID, apiKeyID, ipAddress).Execute()
	return err
}

// OrganizationAPIKeyAccessLists encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKeyAccessLists(params *admin.ListApiKeyAccessListsEntriesApiParams) (*admin.PaginatedApiUserAccessList, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.ListApiKeyAccessListsEntriesWithParams(s.ctx, params).Execute()
	return result, err
}
