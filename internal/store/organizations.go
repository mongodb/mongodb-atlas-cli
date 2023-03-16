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
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_organizations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store OrganizationLister,OrganizationCreator,OrganizationDeleter,OrganizationDescriber

type OrganizationLister interface {
	Organizations(*atlas.OrganizationsListOptions) (*atlas.Organizations, error)
}

type OrganizationDescriber interface {
	Organization(string) (*atlas.Organization, error)
}

type OrganizationCreator interface {
	CreateOrganization(string) (*atlas.Organization, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

// Organizations encapsulate the logic to manage different cloud providers.
func (s *Store) Organizations(opts *atlas.OrganizationsListOptions) (*atlas.Organizations, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		//TODO: Migrate once OrganizationsListOptions.IncludeDeletedOrgs property is generated
		result, _, err := s.client.(*atlas.Client).Organizations.List(s.ctx, opts)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.List(s.ctx, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Organization encapsulate the logic to manage different cloud providers.
func (s *Store) Organization(id string) (*atlas.Organization, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.OrganizationsApi.GetOrganization(s.ctx, id).Execute()
		newOrg := atlas.Organization{ID: *result.Id, IsDeleted: result.IsDeleted, Name: result.Name, Links: mapLinks(result.Links)}
		return &newOrg, err
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

// DeleteOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOrganization(id string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		// TODO: migrate once 406 response is fixed
		// _, err := s.clientv2.OrganizationsApi.DeleteOrganization(s.ctx, id).Execute()
		_, err := s.client.(*atlas.Client).Organizations.Delete(s.ctx, id)
		return err
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Organizations.Delete(s.ctx, id)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func mapLinks(v2Links []atlasv2.Link) []*atlas.Link {
	atlasLinks := make([]*atlas.Link, len(v2Links))
	for i, v2Link := range v2Links {
		atlasLinks[i] = &atlas.Link{Rel: *v2Link.Rel, Href: *v2Link.Href}
	}
	return atlasLinks
}
