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
	CreateUser(*UserView) (interface{}, error)
}

type UserView struct {
	opsmngr.User
	AtlasRoles   []atlas.AtlasRole `json:"roles"`
	MobileNumber string            `json:"mobileNumber"`
	Country      string            `json:"country"`
}

// CreateUser encapsulates the logic to manage different cloud providers
func (s *Store) CreateUser(user *UserView) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		user := &atlas.AtlasUser{
			EmailAddress: user.User.EmailAddress,
			FirstName:    user.User.FirstName,
			LastName:     user.User.LastName,
			Roles:        user.AtlasRoles,
			Username:     user.User.Username,
			MobileNumber: user.MobileNumber,
			Password:     user.User.Password,
			Country:      user.Country,
		}
		result, _, err := s.client.(*atlas.Client).AtlasUsers.Create(context.Background(), user)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		user := &opsmngr.User{
			Username:     user.User.Username,
			Password:     user.User.Password,
			FirstName:    user.User.FirstName,
			LastName:     user.User.LastName,
			EmailAddress: user.User.EmailAddress,
			Roles:        user.User.Roles,
		}
		result, _, err := s.client.(*opsmngr.Client).Users.Create(context.Background(), user)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
