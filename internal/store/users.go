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

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_users.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store UserCreator,UserDescriber,UserDeleter,UserLister,TeamUserLister

type UserCreator interface {
	CreateUser(*opsmngr.User) (*opsmngr.User, error)
}

type UserDeleter interface {
	DeleteUser(string) error
}

type UserLister interface {
	OrganizationUsers(string, *opsmngr.ListOptions) (*opsmngr.UsersResponse, error)
}

type TeamUserLister interface {
	TeamUsers(string, string) ([]*opsmngr.User, error)
}

type UserDescriber interface {
	UserByID(string) (*opsmngr.User, error)
	UserByName(string) (*opsmngr.User, error)
}

// CreateUser encapsulates the logic to manage different cloud providers.
func (s *Store) CreateUser(user *opsmngr.User) (*opsmngr.User, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Users.Create(s.ctx, user)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UserByID encapsulates the logic to manage different cloud providers.
func (s *Store) UserByID(userID string) (*opsmngr.User, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Users.Get(s.ctx, userID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UserByName encapsulates the logic to manage different cloud providers.
func (s *Store) UserByName(username string) (*opsmngr.User, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Users.GetByName(s.ctx, username)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteUser encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteUser(userID string) error {
	switch s.service {
	case config.OpsManagerService:
		_, err := s.client.Users.Delete(s.ctx, userID)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OrganizationUsers encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationUsers(organizationID string, opts *opsmngr.ListOptions) (*opsmngr.UsersResponse, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Organizations.ListUsers(s.ctx, organizationID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// TeamUsers encapsulates the logic to manage different cloud providers.
func (s *Store) TeamUsers(orgID, teamID string) ([]*opsmngr.User, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Teams.GetTeamUsersAssigned(s.ctx, orgID, teamID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
