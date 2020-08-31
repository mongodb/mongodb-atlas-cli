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

//go:generate mockgen -destination=../mocks/mock_users.go -package=mocks github.com/mongodb/mongocli/internal/store UserCreator

type UserCreator interface {
	CreateUser(string, string, string, string, string, string, string, []atlas.AtlasRole, []*opsmngr.UserRole) (interface{}, error)
}

// CreateUser encapsulates the logic to manage different cloud providers
func (s *Store) CreateUser(username, password, firstName, lastName, emailAddress, mobileNumber, country string, atlasRoles []atlas.AtlasRole, roles []*opsmngr.UserRole) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		user := &atlas.AtlasUser{
			EmailAddress: emailAddress,
			FirstName:    firstName,
			LastName:     lastName,
			Roles:        atlasRoles,
			Username:     username,
			MobileNumber: mobileNumber,
			Password:     password,
			Country:      country,
		}
		result, _, err := s.client.(*atlas.Client).AtlasUsers.Create(context.Background(), user)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		user := &opsmngr.User{
			Username:     username,
			Password:     password,
			FirstName:    firstName,
			LastName:     lastName,
			EmailAddress: emailAddress,
			Roles:        roles,
		}
		result, _, err := s.client.(*opsmngr.Client).Users.Create(context.Background(), user)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
