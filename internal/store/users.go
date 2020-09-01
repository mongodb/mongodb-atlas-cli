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
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_users.go -package=mocks github.com/mongodb/mongocli/internal/store UserCreator,UserDescriber,UserDeleter

type UserCreator interface {
	CreateUser(*UserRequest) (interface{}, error)
}

type UserDeleter interface {
	DeleteUser(string) error
}

type UserDescriber interface {
	UserByID(string) (interface{}, error)
	UserByName(string) (interface{}, error)
}

type UserRequest struct {
	*opsmngr.User
	AtlasRoles []atlas.AtlasRole
	Country    string
}

// CreateUser encapsulates the logic to manage different cloud providers
func (s *Store) CreateUser(user *UserRequest) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		atlasUser := &atlas.AtlasUser{
			EmailAddress: user.EmailAddress,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Roles:        user.AtlasRoles,
			Username:     user.Username,
			MobileNumber: user.MobileNumber,
			Password:     user.Password,
			Country:      user.Country,
		}
		result, _, err := s.client.(*atlas.Client).AtlasUsers.Create(context.Background(), atlasUser)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Users.Create(context.Background(), user.User)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UserByID gets an IAM user by ID
func (s *Store) UserByID(userID string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).AtlasUsers.Get(context.Background(), userID)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Users.Get(context.Background(), userID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UserByName gets an IAM user by name
func (s *Store) UserByName(username string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).AtlasUsers.GetByName(context.Background(), username)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Users.GetByName(context.Background(), username)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

func (s *Store) DeleteUser(userID string) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Users.Delete(context.Background(), userID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
