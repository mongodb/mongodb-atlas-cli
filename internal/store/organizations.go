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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_organizations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store OrganizationLister,OrganizationDeleter,OrganizationDescriber,OrganizationCreator

type OrganizationLister interface {
	Organizations(*atlasv2.ListOrganizationsApiParams) (*atlasv2.PaginatedOrganization, error)
}

type OrganizationDescriber interface {
	Organization(string) (*atlasv2.AtlasOrganization, error)
}

type OrganizationCreator interface {
	CreateAtlasOrganization(*atlasv2.CreateOrganizationRequest) (*atlasv2.CreateOrganizationResponse, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

// Organizations encapsulate the logic to manage different cloud providers.
func (s *Store) Organizations(params *atlasv2.ListOrganizationsApiParams) (*atlasv2.PaginatedOrganization, error) {
	result, _, err := s.clientv2.OrganizationsApi.ListOrganizationsWithParams(s.ctx, params).Execute()
	return result, err
}

// Organization encapsulate the logic to manage different cloud providers.
func (s *Store) Organization(id string) (*atlasv2.AtlasOrganization, error) {
	result, _, err := s.clientv2.OrganizationsApi.GetOrganization(s.ctx, id).Execute()
	return result, err
}

// CreateAtlasOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAtlasOrganization(o *atlasv2.CreateOrganizationRequest) (*atlasv2.CreateOrganizationResponse, error) {
	result, _, err := s.clientv2.OrganizationsApi.CreateOrganization(s.ctx, o).Execute()
	return result, err
}

// DeleteOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOrganization(id string) error {
	_, _, err := s.clientv2.OrganizationsApi.DeleteOrganization(s.ctx, id).Execute()
	return err
}
