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
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_projects.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectLister,OrgProjectLister,ProjectCreator,ProjectDeleter,ProjectDescriber,ProjectUsersLister,ProjectUserDeleter,ProjectTeamLister,ProjectTeamAdder,ProjectTeamDeleter

type ProjectLister interface {
	Projects(*atlas.ListOptions) (*atlasv2.PaginatedAtlasGroup, error)
	GetOrgProjects(string, *atlas.ProjectsListOptions) (*atlasv2.PaginatedAtlasGroup, error)
}

type OrgProjectLister interface {
	GetOrgProjects(string) (*atlasv2.PaginatedAtlasGroup, error)
}

type ProjectCreator interface {
	CreateProject(string, string, *bool, *atlas.CreateProjectOptions) (*atlasv2.Group, error)
	ServiceVersionDescriber
}

type ProjectDeleter interface {
	DeleteProject(string) error
}

type ProjectDescriber interface {
	Project(string) (interface{}, error)
}

type ProjectUsersLister interface {
	ProjectUsers(string, *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasDatabaseUser, error)
}

type ProjectUserDeleter interface {
	DeleteUserFromProject(string, string) error
}

type ProjectTeamLister interface {
	ProjectTeams(string) (*atlasv2.PaginatedTeamRole, error)
}

type ProjectTeamAdder interface {
	AddTeamsToProject(string, []*atlas.ProjectTeam) (*atlasv2.PaginatedTeamRole, error)
}

type ProjectTeamDeleter interface {
	DeleteTeamFromProject(string, string) error
}

// Projects encapsulates the logic to manage different cloud providers.
func (s *Store) Projects(opts *atlas.ListOptions) (*atlasv2.PaginatedAtlasGroup, error) {
	result, _, err := s.clientv2.ProjectsApi.ListProjects(s.ctx).PageNum(int32(opts.PageNum)).
		ItemsPerPage(int32(opts.ItemsPerPage)).IncludeCount(opts.IncludeCount).Execute()
	return result, err
}

// GetOrgProjects encapsulates the logic to manage different cloud providers.
func (s *Store) GetOrgProjects(orgID string, opts *atlas.ProjectsListOptions) (*atlasv2.PaginatedAtlasGroup, error) {
	result, _, err := s.clientv2.OrganizationsApi.ListOrganizationProjects(s.ctx, orgID).IncludeCount(opts.IncludeCount).
		PageNum(int32(opts.PageNum)).Name(opts.Name).ItemsPerPage(int32(opts.ItemsPerPage)).Execute()
	return result, err
}

// Project encapsulates the logic to manage different cloud providers.
func (s *Store) Project(id string) (interface{}, error) {
	result, _, err := s.clientv2.ProjectsApi.GetProject(s.ctx, id).Execute()
	return result, err
}

// CreateProject encapsulates the logic to manage different cloud providers.
func (s *Store) CreateProject(name, orgID string, defaultAlertSettings *bool, opts *atlas.CreateProjectOptions) (*atlasv2.Group, error) {
	group := &atlasv2.Group{Name: name, OrgId: orgID, WithDefaultAlertsSettings: defaultAlertSettings}
	result, _, err := s.clientv2.ProjectsApi.CreateProject(s.ctx).ProjectOwnerId(opts.ProjectOwnerID).Group(*group).Execute()
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProject(projectID string) error {
	_, err := s.clientv2.ProjectsApi.DeleteProject(s.ctx, projectID).Execute()
	return err
}

// ProjectUsers lists all IAM users in a project.
func (s *Store) ProjectUsers(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAtlasDatabaseUser, error) {
	result, _, err := s.clientv2.DatabaseUsersApi.ListDatabaseUsers(s.ctx, projectID).IncludeCount(opts.IncludeCount).
		ItemsPerPage(int32(opts.ItemsPerPage)).PageNum(int32(opts.PageNum)).Execute()
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteUserFromProject(projectID, userID string) error {
	_, err := s.clientv2.ProjectsApi.RemoveProjectUser(s.ctx, projectID, userID).Execute()
	return err
}

// ProjectTeams encapsulates the logic to manage different cloud providers.
func (s *Store) ProjectTeams(projectID string) (*atlasv2.PaginatedTeamRole, error) {
	result, _, err := s.clientv2.TeamsApi.ListProjectTeams(s.ctx, projectID).Execute()
	return result, err
}

// AddTeamsToProject encapsulates the logic to manage different cloud providers.
func (s *Store) AddTeamsToProject(projectID string, teams []*atlas.ProjectTeam) (*atlasv2.PaginatedTeamRole, error) {
	teamRole := make([]atlasv2.TeamRole, len(teams))
	for i, team := range teams {
		teamRole[i] = atlasv2.TeamRole{
			TeamId:    &team.TeamID,
			RoleNames: team.RoleNames,
		}
	}

	result, _, err := s.clientv2.TeamsApi.AddAllTeamsToProject(s.ctx, projectID).TeamRole(teamRole).Execute()
	return result, err
}

// DeleteTeamFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeamFromProject(projectID, teamID string) error {
	_, _, err := s.clientv2.TeamsApi.RemoveProjectTeam(s.ctx, projectID, teamID).Execute()
	return err
}
