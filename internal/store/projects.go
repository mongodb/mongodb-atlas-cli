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

type ProjectLister interface {
	GetAllProjects() (interface{}, error)
	GetOrgProjects(string) (interface{}, error)
}

type OrgProjectLister interface {
	GetOrgProjects(string) (interface{}, error)
}

type ProjectCreator interface {
	CreateProject(string, string) (interface{}, error)
}

type ProjectDeleter interface {
	DeleteProject(string) error
}

type ProjectStore interface {
	ProjectLister
	ProjectCreator
	ProjectDeleter
}

// GetAllProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetAllProjects() (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Projects.GetAllProjects(context.Background(), nil)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Projects.List(context.Background(), nil)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// GetOrgProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetOrgProjects(orgID string) (interface{}, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.GetProjects(context.Background(), orgID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateProject encapsulate the logic to manage different cloud providers
func (s *Store) CreateProject(name, orgID string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		project := &atlas.Project{Name: name, OrgID: orgID}
		result, _, err := s.client.(*atlas.Client).Projects.Create(context.Background(), project)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		project := &opsmngr.Project{Name: name, OrgID: orgID}
		result, _, err := s.client.(*opsmngr.Client).Projects.Create(context.Background(), project)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteProject encapsulate the logic to manage different cloud providers
func (s *Store) DeleteProject(projectID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Projects.Delete(context.Background(), projectID)
		return err
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Projects.Delete(context.Background(), projectID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
