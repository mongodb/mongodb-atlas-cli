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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_organizations.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas OrganizationLister,OrganizationDeleter,OrganizationDescriber,OrganizationCreator

type OrganizationLister interface {
	Organizations(*atlas.OrganizationsListOptions) (*atlasv2.PaginatedOrganization, error)
}

type OrganizationDescriber interface {
	Organization(string) (*atlasv2.Organization, error)
}

type OrganizationCreator interface {
	CreateAtlasOrganization(*atlasv2.CreateOrganizationRequest) (*atlasv2.CreateOrganizationResponse, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

// Organizations encapsulate the logic to manage different cloud providers.
func (s *Store) Organizations(opts *atlas.OrganizationsListOptions) (*atlasv2.PaginatedOrganization, error) {
	res := s.clientv2.OrganizationsApi.ListOrganizations(s.ctx)
	if opts != nil {
		res = res.Name(opts.Name).PageNum(opts.PageNum)
	}
	result, _, err := res.Execute()
	return result, err
}

// Organization encapsulate the logic to manage different cloud providers.
func (s *Store) Organization(id string) (*atlasv2.Organization, error) {
	result, _, err := s.clientv2.OrganizationsApi.GetOrganization(s.ctx, id).Execute()
	return result, err
}

// CreateAtlasOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAtlasOrganization(o *atlasv2.CreateOrganizationRequest) (*atlasv2.CreateOrganizationResponse, error) {
	result, _, err := s.clientv2.OrganizationsApi.CreateOrganization(s.ctx).CreateOrganizationRequest(o).Execute()
	return result, err
}

// DeleteOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOrganization(id string) error {
	_, _, err := s.clientv2.OrganizationsApi.DeleteOrganization(s.ctx, id).Execute()
	return err
}
