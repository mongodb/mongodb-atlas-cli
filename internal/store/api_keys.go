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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

// OrganizationAPIKeys encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKeys(orgID string, opts *ListOptions) (*atlasv2.PaginatedApiApiUser, error) {
	res := s.clientv2.ProgrammaticAPIKeysApi.ListOrgApiKeys(s.ctx, orgID)
	if opts != nil {
		res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// OrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKey(orgID, apiKeyID string) (*atlasv2.ApiKeyUserDetails, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.GetOrgApiKey(s.ctx, orgID, apiKeyID).Execute()
	return result, err
}

// UpdateOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateOrganizationAPIKey(orgID, apiKeyID string, input *atlasv2.UpdateAtlasOrganizationApiKey) (*atlasv2.ApiKeyUserDetails, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.UpdateOrgApiKey(s.ctx, orgID, apiKeyID, input).Execute()
	return result, err
}

// CreateOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOrganizationAPIKey(orgID string, input *atlasv2.CreateAtlasOrganizationApiKey) (*atlasv2.ApiKeyUserDetails, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.CreateOrgApiKey(s.ctx, orgID, input).Execute()
	return result, err
}

// DeleteOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOrganizationAPIKey(orgID, id string) error {
	_, err := s.clientv2.ProgrammaticAPIKeysApi.DeleteOrgApiKey(s.ctx, orgID, id).Execute()
	return err
}

// ProjectAPIKeys returns the API Keys for a specific project.
func (s *Store) ProjectAPIKeys(projectID string, opts *ListOptions) (*atlasv2.PaginatedApiApiUser, error) {
	res := s.clientv2.ProgrammaticAPIKeysApi.ListGroupApiKeys(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// CreateProjectAPIKey creates an API Keys for a project.
func (s *Store) CreateProjectAPIKey(projectID string, apiKeyInput *atlasv2.CreateAtlasProjectApiKey) (*atlasv2.ApiKeyUserDetails, error) {
	result, _, err := s.clientv2.ProgrammaticAPIKeysApi.CreateGroupApiKey(s.ctx, projectID, apiKeyInput).Execute()
	return result, err
}

// AssignProjectAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) AssignProjectAPIKey(projectID, apiKeyID string, input *atlasv2.UpdateAtlasProjectApiKey) error {
	_, _, err := s.clientv2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(s.ctx, projectID, apiKeyID, input).Execute()
	return err
}

// DeleteProjectAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProjectAPIKey(projectID, id string) error {
	_, err := s.clientv2.ProgrammaticAPIKeysApi.RemoveGroupApiKey(s.ctx, projectID, id).Execute()
	return err
}
