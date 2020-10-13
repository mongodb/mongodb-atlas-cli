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

//go:generate mockgen -destination=../mocks/mock_users.go -package=mocks github.com/mongodb/mongocli/internal/store UserCreator,UserDescriber,UserDeleter,UserLister,TeamUserLister

type UserCreator interface {
	CreateUser(*UserRequest) (interface{}, error)
}

type UserDeleter interface {
	DeleteUser(string) error
}

type UserLister interface {
	OrganizationUsers(string, *atlas.ListOptions) (interface{}, error)
}

type TeamUserLister interface {
	TeamUsers(string, string) (interface{}, error)
}

type UserDescriber interface {
	UserByID(string) (interface{}, error)
	UserByName(string) (interface{}, error)
}

type UserRequest struct {
	*opsmngr.User
	AtlasRoles []atlas.AtlasRole
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

// UserByID encapsulates the logic to manage different cloud providers
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

// UserByName encapsulates the logic to manage different cloud providers
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

// DeleteUser encapsulates the logic to manage different cloud providers
func (s *Store) DeleteUser(userID string) error {
	switch s.service {
	case config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Users.Delete(context.Background(), userID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// OrganizationUsers encapsulates the logic to manage different cloud providers
func (s *Store) OrganizationUsers(organizationID string, opts *atlas.ListOptions) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Organizations.Users(context.Background(), organizationID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.ListUsers(context.Background(), organizationID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// TeamUsers encapsulates the logic to manage different cloud providers
func (s *Store) TeamUsers(orgID, teamID string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Teams.GetTeamUsersAssigned(context.Background(), orgID, teamID)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Teams.GetTeamUsersAssigned(context.Background(), orgID, teamID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
