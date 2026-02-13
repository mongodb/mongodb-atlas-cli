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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

// TeamByID encapsulates the logic to manage different cloud providers.
func (s *Store) TeamByID(orgID, teamID string) (*atlasv2.TeamResponse, error) {
	result, _, err := s.clientv2.TeamsApi.GetOrgTeam(s.ctx, orgID, teamID).Execute()
	return result, err
}

// TeamByName encapsulates the logic to manage different cloud providers.
func (s *Store) TeamByName(orgID, teamName string) (*atlasv2.TeamResponse, error) {
	result, _, err := s.clientv2.TeamsApi.GetTeamByName(s.ctx, orgID, teamName).Execute()
	return result, err
}

// Teams encapsulates the logic to manage different cloud providers.
func (s *Store) Teams(orgID string, opts *ListOptions) (*atlasv2.PaginatedTeam, error) {
	res := s.clientv2.TeamsApi.ListOrgTeams(s.ctx, orgID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

func (s *Store) CreateTeam(orgID string, team *atlasv2.Team) (*atlasv2.Team, error) {
	result, _, err := s.clientv2.TeamsApi.CreateOrgTeam(s.ctx, orgID, team).Execute()
	return result, err
}

func (s *Store) RenameTeam(orgID, teamID string, team *atlasv2.TeamUpdate) (*atlasv2.TeamResponse, error) {
	result, _, err := s.clientv2.TeamsApi.RenameOrgTeam(s.ctx, orgID, teamID, team).Execute()
	return result, err
}

// DeleteTeam encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeam(orgID, teamID string) error {
	_, err := s.clientv2.TeamsApi.DeleteOrgTeam(s.ctx, orgID, teamID).Execute()
	return err
}

// AddUsersToTeam encapsulates the logic to manage different cloud providers.
func (s *Store) AddUsersToTeam(orgID, teamID string, users []atlasv2.AddUserToTeam) (*atlasv2.PaginatedApiAppUser, error) {
	result, _, err := s.clientv2.TeamsApi.AddTeamUsers(s.ctx, orgID, teamID, &users).Execute()
	return result, err
}

// RemoveUserFromTeam encapsulates the logic to manage different cloud providers.
func (s *Store) RemoveUserFromTeam(orgID, teamID, userID string) error {
	_, err := s.clientv2.TeamsApi.RemoveUserFromTeam(s.ctx, orgID, teamID, userID).Execute()
	return err
}

// UpdateProjectTeamRoles encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateProjectTeamRoles(projectID, teamID string, team *atlasv2.TeamRole) (*atlasv2.PaginatedTeamRole, error) {
	result, _, err := s.clientv2.TeamsApi.UpdateGroupTeam(s.ctx, projectID, teamID, team).Execute()
	return result, err
}
