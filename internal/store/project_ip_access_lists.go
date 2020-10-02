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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_project_ip_access_lists.go -package=mocks github.com/mongodb/mongocli/internal/store ProjectIPWhitelistDescriber,ProjectIPWhitelistLister,ProjectIPWhitelistCreator,ProjectIPWhitelistDeleter

type ProjectIPWhitelistDescriber interface {
	IPWhitelist(string, string) (*atlas.ProjectIPWhitelist, error)
}
type ProjectIPWhitelistLister interface {
	ProjectIPWhitelists(string, *atlas.ListOptions) ([]atlas.ProjectIPWhitelist, error)
}

type ProjectIPWhitelistCreator interface {
	CreateProjectIPWhitelist(*atlas.ProjectIPWhitelist) ([]atlas.ProjectIPWhitelist, error)
}

type ProjectIPWhitelistDeleter interface {
	DeleteProjectIPWhitelist(string, string) error
}

// CreateProjectIPWhitelist encapsulate the logic to manage different cloud providers
func (s *Store) CreateProjectIPWhitelist(entry *atlas.ProjectIPWhitelist) ([]atlas.ProjectIPWhitelist, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPWhitelist.Create(context.Background(), entry.GroupID, []*atlas.ProjectIPWhitelist{entry})
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteProjectIPWhitelist encapsulate the logic to manage different cloud providers
func (s *Store) DeleteProjectIPWhitelist(projectID, entry string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).ProjectIPWhitelist.Delete(context.Background(), projectID, entry)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ProjectIPWhitelists encapsulate the logic to manage different cloud providers
func (s *Store) ProjectIPWhitelists(projectID string, opts *atlas.ListOptions) ([]atlas.ProjectIPWhitelist, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPWhitelist.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// IPWhitelist encapsulate the logic to manage different cloud providers
func (s *Store) IPWhitelist(projectID, name string) (*atlas.ProjectIPWhitelist, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPWhitelist.Get(context.Background(), projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
