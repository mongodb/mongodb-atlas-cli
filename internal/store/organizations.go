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
	"go.mongodb.org/ops-manager/opsmngr"
)

type OrganizationLister interface {
	GetAllOrganizations() (interface{}, error)
}

type OrganizationCreator interface {
	CreateOrganization(string) (interface{}, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

type OrganizationStore interface {
	OrganizationLister
	OrganizationCreator
	OrganizationDeleter
}

// GetAllProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetAllOrganizations() (interface{}, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.List(context.Background(), nil)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateOrganization encapsulate the logic to manage different cloud providers
func (s *Store) CreateOrganization(name string) (interface{}, error) {
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
func (s *Store) DeleteOrganization(ID string) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Organizations.Delete(context.Background(), ID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
