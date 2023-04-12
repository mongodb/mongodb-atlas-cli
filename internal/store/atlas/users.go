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

package atlas

import (
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_users.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas UserCreator,UserDescriber,UserLister,TeamUserLister

type UserCreator interface {
	CreateUser(*UserRequest) (interface{}, error)
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

// CreateUser encapsulates the logic to manage different cloud providers.
func (s *Store) CreateUser(user *UserRequest) (interface{}, error) {

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
	result, _, err := s.client.AtlasUsers.Create(s.ctx, atlasUser)
	return result, err
}

// UserByID encapsulates the logic to manage different cloud providers.
func (s *Store) UserByID(userID string) (interface{}, error) {
	result, _, err := s.client.AtlasUsers.Get(s.ctx, userID)
	return result, err
}

// UserByName encapsulates the logic to manage different cloud providers.
func (s *Store) UserByName(username string) (interface{}, error) {
	result, _, err := s.client.AtlasUsers.GetByName(s.ctx, username)
	return result, err
}

// OrganizationUsers encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationUsers(organizationID string, opts *atlas.ListOptions) (interface{}, error) {
	result, _, err := s.client.Organizations.Users(s.ctx, organizationID, opts)
	return result, err
}

// TeamUsers encapsulates the logic to manage different cloud providers.
func (s *Store) TeamUsers(orgID, teamID string) (interface{}, error) {
	result, _, err := s.client.Teams.GetTeamUsersAssigned(s.ctx, orgID, teamID)
	return result, err
}
