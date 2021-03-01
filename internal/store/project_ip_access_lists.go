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

//go:generate mockgen -destination=../mocks/mock_project_ip_access_lists.go -package=mocks github.com/mongodb/mongocli/internal/store ProjectIPAccessListDescriber,ProjectIPAccessListLister,ProjectIPAccessListCreator,ProjectIPAccessListDeleter

type ProjectIPAccessListDescriber interface {
	IPAccessList(string, string) (*atlas.ProjectIPAccessList, error)
}
type ProjectIPAccessListLister interface {
	ProjectIPAccessLists(string, *atlas.ListOptions) (*atlas.ProjectIPAccessLists, error)
}

type ProjectIPAccessListCreator interface {
	CreateProjectIPAccessList(*atlas.ProjectIPAccessList) (*atlas.ProjectIPAccessLists, error)
}

type ProjectIPAccessListDeleter interface {
	DeleteProjectIPAccessList(string, string) error
}

// CreateProjectIPAccessList encapsulate the logic to manage different cloud providers
func (s *Store) CreateProjectIPAccessList(entry *atlas.ProjectIPAccessList) (*atlas.ProjectIPAccessLists, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPAccessList.Create(context.Background(), entry.GroupID, []*atlas.ProjectIPAccessList{entry})
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteProjectIPAccessList encapsulate the logic to manage different cloud providers
func (s *Store) DeleteProjectIPAccessList(projectID, entry string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).ProjectIPAccessList.Delete(context.Background(), projectID, entry)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ProjectIPAccessLists encapsulate the logic to manage different cloud providers
func (s *Store) ProjectIPAccessLists(projectID string, opts *atlas.ListOptions) (*atlas.ProjectIPAccessLists, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPAccessList.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// IPAccessList encapsulate the logic to manage different cloud providers
func (s *Store) IPAccessList(projectID, name string) (*atlas.ProjectIPAccessList, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPAccessList.Get(context.Background(), projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
