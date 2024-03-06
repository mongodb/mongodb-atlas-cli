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

//go:generate mockgen -destination=../mocks/mock_projects.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ProjectLister,ProjectCreator,ProjectDeleter,ProjectDescriber,ProjectUsersLister,ProjectUserDeleter,ProjectTeamLister,ProjectTeamAdder,ProjectTeamDeleter

type ProjectLister interface {
	Projects(*atlas.ListOptions) (*atlas.Projects, error)
	GetOrgProjects(string, *atlas.ProjectsListOptions) (*atlas.Projects, error)
}

type ProjectCreator interface {
	CreateProject(string, string, string, *bool, *atlas.CreateProjectOptions) (*atlas.Project, error)
}

type ProjectDeleter interface {
	DeleteProject(string) error
}

type ProjectDescriber interface {
	Project(string) (*atlas.Project, error)
	ProjectByName(string) (*atlas.Project, error)
}

type ProjectUsersLister interface {
	ProjectUsers(string, *atlas.ListOptions) ([]atlas.AtlasUser, error)
}

type ProjectUserDeleter interface {
	DeleteUserFromProject(string, string) error
}

type ProjectTeamLister interface {
	ProjectTeams(string) (*atlas.TeamsAssigned, error)
}

type ProjectTeamAdder interface {
	AddTeamsToProject(string, []*atlas.ProjectTeam) (*atlas.TeamsAssigned, error)
}

type ProjectTeamDeleter interface {
	DeleteTeamFromProject(string, string) error
}

// Projects encapsulates the logic to manage different cloud providers.
func (s *Store) Projects(opts *atlas.ListOptions) (*atlas.Projects, error) {
	result, _, err := s.client.Projects.GetAllProjects(s.ctx, opts)
	return result, err
}

// GetOrgProjects encapsulates the logic to manage different cloud providers.
func (s *Store) GetOrgProjects(orgID string, opts *atlas.ProjectsListOptions) (*atlas.Projects, error) {
	result, _, err := s.client.Organizations.Projects(s.ctx, orgID, opts)
	return result, err
}

// Project encapsulates the logic to manage different cloud providers.
func (s *Store) Project(id string) (*atlas.Project, error) {
	result, _, err := s.client.Projects.GetOneProject(s.ctx, id)
	return result, err
}

func (s *Store) ProjectByName(name string) (*atlas.Project, error) {
	result, _, err := s.client.Projects.GetOneProjectByName(s.ctx, name)
	return result, err
}

// CreateProject encapsulates the logic to manage different cloud providers.
func (s *Store) CreateProject(name, orgID, regionUsageRestrictions string, defaultAlertSettings *bool, opts *atlas.CreateProjectOptions) (*atlas.Project, error) {
	project := &atlas.Project{Name: name, OrgID: orgID, RegionUsageRestrictions: regionUsageRestrictions, WithDefaultAlertsSettings: defaultAlertSettings}
	result, _, err := s.client.Projects.Create(s.ctx, project, opts)
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProject(projectID string) error {
	_, err := s.client.Projects.Delete(s.ctx, projectID)
	return err
}

// ProjectUsers lists all IAM users in a project.
func (s *Store) ProjectUsers(projectID string, opts *atlas.ListOptions) ([]atlas.AtlasUser, error) {
	result, _, err := s.client.AtlasUsers.List(s.ctx, projectID, opts)
	return result, err
}

// DeleteUserFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteUserFromProject(projectID, userID string) error {
	_, err := s.client.Projects.RemoveUserFromProject(s.ctx, projectID, userID)
	return err
}

// ProjectTeams encapsulates the logic to manage different cloud providers.
func (s *Store) ProjectTeams(projectID string) (*atlas.TeamsAssigned, error) {
	result, _, err := s.client.Projects.GetProjectTeamsAssigned(s.ctx, projectID)
	return result, err
}

// AddTeamsToProject encapsulates the logic to manage different cloud providers.
func (s *Store) AddTeamsToProject(projectID string, teams []*atlas.ProjectTeam) (*atlas.TeamsAssigned, error) {
	result, _, err := s.client.Projects.AddTeamsToProject(s.ctx, projectID, teams)
	return result, err
}

// DeleteTeamFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeamFromProject(projectID, teamID string) error {
	_, err := s.client.Teams.RemoveTeamFromProject(s.ctx, projectID, teamID)
	return err
}
