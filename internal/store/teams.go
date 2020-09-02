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

//go:generate mockgen -destination=../mocks/mock_teams.go -package=mocks github.com/mongodb/mongocli/internal/store TeamLister,TeamDescriber,TeamCreator,TeamDeleter,TeamAdder

type TeamLister interface {
	Teams(string, *atlas.ListOptions) ([]atlas.Team, error)
}

type TeamDescriber interface {
	TeamByID(string, string) (*atlas.Team, error)
	TeamByName(string, string) (*atlas.Team, error)
}

type TeamCreator interface {
	CreateTeam(string, *atlas.Team) (*atlas.Team, error)
}

type TeamDeleter interface {
	DeleteTeam(string, string) error
}

type TeamAdder interface {
	AddUsersToTeam(string, string, []string) (interface{}, error)
}

// TeamByID encapsulates the logic to manage different cloud providers
func (s *Store) TeamByID(orgID, teamID string) (*atlas.Team, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Teams.Get(context.Background(), orgID, teamID)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Teams.Get(context.Background(), orgID, teamID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// TeamByName encapsulates the logic to manage different cloud providers
func (s *Store) TeamByName(orgID, teamName string) (*atlas.Team, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Teams.GetOneTeamByName(context.Background(), orgID, teamName)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Teams.GetOneTeamByName(context.Background(), orgID, teamName)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Teams encapsulates the logic to manage different cloud providers
func (s *Store) Teams(orgID string, opts *atlas.ListOptions) ([]atlas.Team, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Teams.List(context.Background(), orgID, opts)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Teams.List(context.Background(), orgID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateTeam encapsulates the logic to manage different cloud providers
func (s *Store) CreateTeam(orgID string, team *atlas.Team) (*atlas.Team, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Teams.Create(context.Background(), orgID, team)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Teams.Create(context.Background(), orgID, team)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteTeam encapsulates the logic to manage different cloud providers
func (s *Store) DeleteTeam(orgID, teamID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Teams.RemoveTeamFromOrganization(context.Background(), orgID, teamID)
		return err
	case config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Teams.RemoveTeamFromOrganization(context.Background(), orgID, teamID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// AddUsersToTeam encapsulates the logic to manage different cloud providers
func (s *Store) AddUsersToTeam(orgID, teamID string,  users []string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Teams.AddUsersToTeam(context.Background(), orgID, teamID, users)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Teams.AddUsersToTeam(context.Background(), orgID, teamID, users)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
