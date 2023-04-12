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
)

//go:generate mockgen -destination=../mocks/atlas/mock_teams.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas TeamLister,TeamDescriber,TeamCreator,TeamDeleter,TeamAdder,TeamUserRemover,TeamRolesUpdater

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

type TeamUserRemover interface {
	RemoveUserFromTeam(string, string, string) error
}

type TeamRolesUpdater interface {
	UpdateProjectTeamRoles(string, string, *atlas.TeamUpdateRoles) ([]atlas.TeamRoles, error)
}

// TeamByID encapsulates the logic to manage different cloud providers.
func (s *Store) TeamByID(orgID, teamID string) (*atlas.Team, error) {
	result, _, err := s.client.Teams.Get(s.ctx, orgID, teamID)
	return result, err
}

// TeamByName encapsulates the logic to manage different cloud providers.
func (s *Store) TeamByName(orgID, teamName string) (*atlas.Team, error) {
	result, _, err := s.client.Teams.GetOneTeamByName(s.ctx, orgID, teamName)
	return result, err
}

// Teams encapsulates the logic to manage different cloud providers.
func (s *Store) Teams(orgID string, opts *atlas.ListOptions) ([]atlas.Team, error) {
	result, _, err := s.client.Teams.List(s.ctx, orgID, opts)
	return result, err
}

// CreateTeam encapsulates the logic to manage different cloud providers.
func (s *Store) CreateTeam(orgID string, team *atlas.Team) (*atlas.Team, error) {
	result, _, err := s.client.Teams.Create(s.ctx, orgID, team)
	return result, err
}

// DeleteTeam encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeam(orgID, teamID string) error {
	_, err := s.client.Teams.RemoveTeamFromOrganization(s.ctx, orgID, teamID)
	return err
}

// AddUsersToTeam encapsulates the logic to manage different cloud providers.
func (s *Store) AddUsersToTeam(orgID, teamID string, users []string) (interface{}, error) {
	result, _, err := s.client.Teams.AddUsersToTeam(s.ctx, orgID, teamID, users)
	return result, err
}

// RemoveUserFromTeam encapsulates the logic to manage different cloud providers.
func (s *Store) RemoveUserFromTeam(orgID, teamID, userID string) error {
	_, err := s.client.Teams.RemoveUserToTeam(s.ctx, orgID, teamID, userID)
	return err
}

// UpdateProjectTeamRoles encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateProjectTeamRoles(projectID, teamID string, team *atlas.TeamUpdateRoles) ([]atlas.TeamRoles, error) {
	result, _, err := s.client.Teams.UpdateTeamRoles(s.ctx, projectID, teamID, team)
	return result, err
}
