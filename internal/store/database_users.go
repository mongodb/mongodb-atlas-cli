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
	"github.com/mongodb/mcli/internal/config"
)



type DatabaseUserLister interface {
	ProjectDatabaseUser(groupID string,  opts *atlas.ListOptions) ([]atlas.DatabaseUser, error)
}

type DatabaseUserCreator interface {
	CreateDatabaseUser(*atlas.DatabaseUser) (*atlas.DatabaseUser, error)
}

type DatabaseUserDeleter interface {
	DeleteDatabaseUser(string, string) error
}

type DatabaseUserStore interface {
	DatabaseUserCreator
	DatabaseUserDeleter
	DatabaseUserLister
}

// CreateDatabaseUser encapsulate the logic to manage different cloud providers
func (s *Store) CreateDatabaseUser(user *atlas.DatabaseUser) (*atlas.DatabaseUser, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DatabaseUsers.Create(context.Background(), user.GroupID, user)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

func (s *Store) DeleteDatabaseUser(groupID, username string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).DatabaseUsers.Delete(context.Background(), groupID, username)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

func (s *Store) ProjectDatabaseUser(groupID string,  opts *atlas.ListOptions) ([]atlas.DatabaseUser, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DatabaseUsers.List(context.Background(), groupID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
