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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_api_keys.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectAPIKeyLister,ProjectAPIKeyCreator,OrganizationAPIKeyLister,OrganizationAPIKeyDescriber,OrganizationAPIKeyUpdater,OrganizationAPIKeyCreator,OrganizationAPIKeyDeleter,ProjectAPIKeyDeleter,ProjectAPIKeyAssigner

type ProjectAPIKeyLister interface {
	ProjectAPIKeys(string, *atlas.ListOptions) ([]atlas.APIKey, error)
}

type ProjectAPIKeyCreator interface {
	CreateProjectAPIKey(string, *atlas.APIKeyInput) (*atlas.APIKey, error)
}

type ProjectAPIKeyDeleter interface {
	DeleteProjectAPIKey(string, string) error
}

type ProjectAPIKeyAssigner interface {
	AssignProjectAPIKey(string, string, *atlas.AssignAPIKey) error
}

type OrganizationAPIKeyLister interface {
	OrganizationAPIKeys(string, *atlas.ListOptions) ([]atlas.APIKey, error)
}

type OrganizationAPIKeyDescriber interface {
	OrganizationAPIKey(string, string) (*atlas.APIKey, error)
}

type OrganizationAPIKeyUpdater interface {
	UpdateOrganizationAPIKey(string, string, *atlas.APIKeyInput) (*atlas.APIKey, error)
}

type OrganizationAPIKeyCreator interface {
	CreateOrganizationAPIKey(string, *atlas.APIKeyInput) (*atlas.APIKey, error)
}

type OrganizationAPIKeyDeleter interface {
	DeleteOrganizationAPIKey(string, string) error
}

// OrganizationAPIKeys encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKeys(orgID string, opts *atlas.ListOptions) ([]atlas.APIKey, error) {
	result, _, err := s.client.APIKeys.List(s.ctx, orgID, opts)
	return result, err
}

// OrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationAPIKey(orgID, apiKeyID string) (*atlas.APIKey, error) {
	result, _, err := s.client.APIKeys.Get(s.ctx, orgID, apiKeyID)
	return result, err
}

// UpdateOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateOrganizationAPIKey(orgID, apiKeyID string, input *atlas.APIKeyInput) (*atlas.APIKey, error) {
	result, _, err := s.client.APIKeys.Update(s.ctx, orgID, apiKeyID, input)
	return result, err
}

// CreateOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOrganizationAPIKey(orgID string, input *atlas.APIKeyInput) (*atlas.APIKey, error) {
	result, _, err := s.client.APIKeys.Create(s.ctx, orgID, input)
	return result, err
}

// DeleteOrganizationAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOrganizationAPIKey(orgID, id string) error {
	_, err := s.client.APIKeys.Delete(s.ctx, orgID, id)
	return err
}

// ProjectAPIKeys returns the API Keys for a specific project.
func (s *Store) ProjectAPIKeys(projectID string, opts *atlas.ListOptions) ([]atlas.APIKey, error) {
	result, _, err := s.client.ProjectAPIKeys.List(s.ctx, projectID, opts)
	return result, err
}

// CreateProjectAPIKey creates an API Keys for a project.
func (s *Store) CreateProjectAPIKey(projectID string, apiKeyInput *atlas.APIKeyInput) (*atlas.APIKey, error) {
	result, _, err := s.client.ProjectAPIKeys.Create(s.ctx, projectID, apiKeyInput)
	return result, err
}

// AssignProjectAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) AssignProjectAPIKey(projectID, apiKeyID string, input *atlas.AssignAPIKey) error {
	_, err := s.client.ProjectAPIKeys.Assign(s.ctx, projectID, apiKeyID, input)
	return err
}

// DeleteProjectAPIKey encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProjectAPIKey(projectID, id string) error {
	_, err := s.client.ProjectAPIKeys.Unassign(s.ctx, projectID, id)
	return err
}
