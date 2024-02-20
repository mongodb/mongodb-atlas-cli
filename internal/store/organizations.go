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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_organizations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store OrganizationLister,OrganizationCreator,OrganizationDeleter,OrganizationDescriber,AtlasOrganizationCreator

type OrganizationLister interface {
	Organizations(*atlas.OrganizationsListOptions) (interface{}, error)
}

type OrganizationDescriber interface {
	Organization(string) (interface{}, error)
}

type OrganizationCreator interface {
	CreateOrganization(string) (*atlas.Organization, error)
}

type AtlasOrganizationCreator interface {
	CreateAtlasOrganization(*atlas.CreateOrganizationRequest) (*atlas.CreateOrganizationResponse, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

// Organizations encapsulate the logic to manage different cloud providers.
func (s *Store) Organizations(opts *atlas.OrganizationsListOptions) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		params := new(atlasv2.ListOrganizationsApiParams)
		if opts != nil {
			params.Name = &opts.Name
			params.PageNum = &opts.PageNum
		}
		result, _, err := s.clientv2.OrganizationsApi.ListOrganizationsWithParams(s.ctx, params).Execute()
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.List(s.ctx, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Organization encapsulate the logic to manage different cloud providers.
func (s *Store) Organization(id string) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.OrganizationsApi.GetOrganization(s.ctx, id).Execute()
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.Get(s.ctx, id)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) CreateOrganization(name string) (*atlas.Organization, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		org := &atlas.Organization{Name: name}
		result, _, err := s.client.(*opsmngr.Client).Organizations.Create(s.ctx, org)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateAtlasOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAtlasOrganization(o *atlas.CreateOrganizationRequest) (*atlas.CreateOrganizationResponse, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Organizations.Create(s.ctx, o)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOrganization(id string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.client.(*atlas.Client).Organizations.Delete(s.ctx, id)
		return err
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Organizations.Delete(s.ctx, id)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
