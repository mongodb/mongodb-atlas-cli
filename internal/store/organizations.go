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

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_organizations.go -package=mocks github.com/mongodb/mongocli/internal/store OrganizationLister,OrganizationCreator,OrganizationDeleter,OrganizationDescriber

type OrganizationLister interface {
	Organizations(*atlas.ListOptions) (*opsmngr.Organizations, error)
}

type OrganizationDescriber interface {
	Organization(string) (*opsmngr.Organization, error)
}

type OrganizationCreator interface {
	CreateOrganization(string) (*opsmngr.Organization, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

// Organizations encapsulate the logic to manage different cloud providers
func (s *Store) Organizations(opts *atlas.ListOptions) (*opsmngr.Organizations, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.List(context.Background(), opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Organization encapsulate the logic to manage different cloud providers
func (s *Store) Organization(id string) (*opsmngr.Organization, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.Get(context.Background(), id)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateOrganization encapsulate the logic to manage different cloud providers
func (s *Store) CreateOrganization(name string) (*opsmngr.Organization, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		org := &opsmngr.Organization{Name: name}
		result, _, err := s.client.(*opsmngr.Client).Organizations.Create(context.Background(), org)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteOrganization encapsulate the logic to manage different cloud providers
func (s *Store) DeleteOrganization(id string) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Organizations.Delete(context.Background(), id)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
