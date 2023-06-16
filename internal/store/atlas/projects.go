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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_projects.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectLister,OrgProjectLister,ProjectCreator,ProjectDeleter,ProjectDescriber,ProjectUsersLister,ProjectUserDeleter,ProjectTeamLister,ProjectTeamAdder,ProjectTeamDeleter

type ProjectLister interface {
	Projects(*atlas.ListOptions) (*atlasv2.PaginatedAtlasGroup, error)
}

type OrgProjectLister interface {
	GetOrgProjects(string) (*atlasv2.PaginatedAtlasGroup, error)
}

type ProjectCreator interface {
	CreateProject(*atlasv2.CreateProjectApiParams) (*atlasv2.Group, error)
	ServiceVersionDescriber
}

type ProjectDeleter interface {
	DeleteProject(string) error
}

type ProjectDescriber interface {
	Project(string) (interface{}, error)
	ProjectByName(string) (interface{}, error)
}

type ProjectUsersLister interface {
	ProjectUsers(string, *atlas.ListOptions) (*atlasv2.PaginatedApiAppUser, error)
}

type ProjectUserDeleter interface {
	DeleteUserFromProject(string, string) error
}

type ProjectTeamLister interface {
	ProjectTeams(string) (*atlasv2.PaginatedTeamRole, error)
}

type ProjectTeamAdder interface {
	AddTeamsToProject(string, []atlasv2.TeamRole) (*atlasv2.PaginatedTeamRole, error)
}

type ProjectTeamDeleter interface {
	DeleteTeamFromProject(string, string) error
}

// Projects encapsulates the logic to manage different cloud providers.
func (s *Store) Projects(opts *atlas.ListOptions) (*atlasv2.PaginatedAtlasGroup, error) {
	res := s.clientv2.ProjectsApi.ListProjects(s.ctx)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	return result, err
}

// GetOrgProjects encapsulates the logic to manage different cloud providers.
func (s *Store) GetOrgProjects(orgID string) (*atlasv2.PaginatedAtlasGroup, error) {
	result, _, err := s.clientv2.OrganizationsApi.ListOrganizationProjects(s.ctx, orgID).Execute()
	return result, err
}

// Project encapsulates the logic to manage different cloud providers.
func (s *Store) Project(id string) (interface{}, error) {
	result, _, err := s.clientv2.ProjectsApi.GetProject(s.ctx, id).Execute()
	return result, err
}

func (s *Store) ProjectByName(name string) (interface{}, error) {
	result, _, err := s.clientv2.ProjectsApi.GetProjectByName(s.ctx, name).Execute()
	return result, err
}

// CreateProject encapsulates the logic to manage different cloud providers.
func (s *Store) CreateProject(params *atlasv2.CreateProjectApiParams) (*atlasv2.Group, error) {
	result, _, err := s.clientv2.ProjectsApi.CreateProjectWithParams(s.ctx, params).Execute()
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProject(projectID string) error {
	_, _, err := s.clientv2.ProjectsApi.DeleteProject(s.ctx, projectID).Execute()
	return err
}

// ProjectUsers lists all IAM users in a project.
func (s *Store) ProjectUsers(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedApiAppUser, error) {
	res := s.clientv2.ProjectsApi.ListProjectUsers(s.ctx, projectID)
	if opts != nil {
		res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum)
	}
	result, _, err := res.Execute()
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
func (s *Store) AddTeamsToProject(projectID string, teams []atlasv2.TeamRole) (*atlasv2.PaginatedTeamRole, error) {
	result, _, err := s.clientv2.TeamsApi.AddAllTeamsToProject(s.ctx, projectID, &teams).Execute()
	return result, err
}

// DeleteTeamFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeamFromProject(projectID, teamID string) error {
	_, err := s.clientv2.TeamsApi.RemoveProjectTeam(s.ctx, projectID, teamID).Execute()
	return err
}
