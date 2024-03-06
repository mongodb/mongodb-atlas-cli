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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_organizations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store OrganizationLister,OrganizationDeleter,OrganizationDescriber,AtlasOrganizationCreator

type OrganizationLister interface {
	Organizations(*atlas.OrganizationsListOptions) (*admin.PaginatedOrganization, error)
}

type OrganizationDescriber interface {
	Organization(string) (*admin.AtlasOrganization, error)
}

type AtlasOrganizationCreator interface {
	CreateAtlasOrganization(*atlas.CreateOrganizationRequest) (*atlas.CreateOrganizationResponse, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

// Organizations encapsulate the logic to manage different cloud providers.
func (s *Store) Organizations(opts *atlas.OrganizationsListOptions) (*admin.PaginatedOrganization, error) {
	res := s.clientv2.OrganizationsApi.ListOrganizations(s.ctx)
	if opts != nil {
		res = res.Name(opts.Name).PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	return result, err
}

// Organization encapsulate the logic to manage different cloud providers.
func (s *Store) Organization(id string) (*admin.AtlasOrganization, error) {
	result, _, err := s.clientv2.OrganizationsApi.GetOrganization(s.ctx, id).Execute()
	return result, err
}

// CreateAtlasOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAtlasOrganization(o *atlas.CreateOrganizationRequest) (*atlas.CreateOrganizationResponse, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.client.Organizations.Create(s.ctx, o)
	return result, err
}

// DeleteOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOrganization(id string) error {
	_, err := s.client.Organizations.Delete(s.ctx, id)
	return err
}
