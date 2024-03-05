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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_projects.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ProjectLister,OrgProjectLister,ProjectCreator,ProjectDeleter,ProjectDescriber,ProjectUsersLister,ProjectUserDeleter,ProjectTeamLister,ProjectTeamAdder,ProjectTeamDeleter

type ProjectLister interface {
	Projects(*atlas.ListOptions) (interface{}, error)
	GetOrgProjects(string, *atlas.ProjectsListOptions) (interface{}, error)
}

type OrgProjectLister interface {
	GetOrgProjects(string) (interface{}, error)
}

type ProjectCreator interface {
	CreateProject(string, string, string, *bool, *atlas.CreateProjectOptions) (interface{}, error)
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
	ProjectUsers(string, *atlas.ListOptions) (interface{}, error)
}

type ProjectUserDeleter interface {
	DeleteUserFromProject(string, string) error
}

type ProjectTeamLister interface {
	ProjectTeams(string) (interface{}, error)
}

type ProjectTeamAdder interface {
	AddTeamsToProject(string, []*atlas.ProjectTeam) (*atlas.TeamsAssigned, error)
}

type ProjectTeamDeleter interface {
	DeleteTeamFromProject(string, string) error
}

// Projects encapsulates the logic to manage different cloud providers.
func (s *Store) Projects(opts *atlas.ListOptions) (interface{}, error) {
	result, _, err := s.client.(*atlas.Client).Projects.GetAllProjects(s.ctx, opts)
	return result, err
}

// GetOrgProjects encapsulates the logic to manage different cloud providers.
func (s *Store) GetOrgProjects(orgID string, opts *atlas.ProjectsListOptions) (interface{}, error) {
	result, _, err := s.client.(*atlas.Client).Organizations.Projects(s.ctx, orgID, opts)
	return result, err
}

// Project encapsulates the logic to manage different cloud providers.
func (s *Store) Project(id string) (interface{}, error) {
	result, _, err := s.client.(*atlas.Client).Projects.GetOneProject(s.ctx, id)
	return result, err
}

func (s *Store) ProjectByName(name string) (interface{}, error) {
	result, _, err := s.client.(*atlas.Client).Projects.GetOneProjectByName(s.ctx, name)
	return result, err
}

// CreateProject encapsulates the logic to manage different cloud providers.
func (s *Store) CreateProject(name, orgID, regionUsageRestrictions string, defaultAlertSettings *bool, opts *atlas.CreateProjectOptions) (interface{}, error) {
	project := &atlas.Project{Name: name, OrgID: orgID, RegionUsageRestrictions: regionUsageRestrictions, WithDefaultAlertsSettings: defaultAlertSettings}
	result, _, err := s.client.(*atlas.Client).Projects.Create(s.ctx, project, opts)
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProject(projectID string) error {
	_, err := s.client.(*atlas.Client).Projects.Delete(s.ctx, projectID)
	return err
}

// ProjectUsers lists all IAM users in a project.
func (s *Store) ProjectUsers(projectID string, opts *atlas.ListOptions) (interface{}, error) {
	result, _, err := s.client.(*atlas.Client).AtlasUsers.List(s.ctx, projectID, opts)
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteUserFromProject(projectID, userID string) error {
	_, err := s.client.(*atlas.Client).Projects.RemoveUserFromProject(s.ctx, projectID, userID)
	return err
}

// ProjectTeams encapsulates the logic to manage different cloud providers.
func (s *Store) ProjectTeams(projectID string) (interface{}, error) {
	result, _, err := s.client.(*atlas.Client).Projects.GetProjectTeamsAssigned(s.ctx, projectID)
	return result, err
}

// AddTeamsToProject encapsulates the logic to manage different cloud providers.
func (s *Store) AddTeamsToProject(projectID string, teams []*atlas.ProjectTeam) (*atlas.TeamsAssigned, error) {
	result, _, err := s.client.(*atlas.Client).Projects.AddTeamsToProject(s.ctx, projectID, teams)
	return result, err
}

// DeleteTeamFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeamFromProject(projectID, teamID string) error {
	_, err := s.client.(*atlas.Client).Teams.RemoveTeamFromProject(s.ctx, projectID, teamID)
	return err
}
