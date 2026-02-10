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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
)

// Projects encapsulates the logic to manage different cloud providers.
func (s *Store) Projects(opts *ListOptions) (*atlasv2.PaginatedAtlasGroup, error) {
	res := s.clientv2.ProjectsApi.ListGroups(s.ctx)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	return result, err
}

// GetOrgProjects encapsulates the logic to manage different cloud providers.
func (s *Store) GetOrgProjects(orgID string, opts *ListOptions) (*atlasv2.PaginatedAtlasGroup, error) {
	res := s.clientv2.OrganizationsApi.GetOrgGroups(s.ctx, orgID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	return result, err
}

// Project encapsulates the logic to manage different cloud providers.
func (s *Store) Project(id string) (*atlasv2.Group, error) {
	result, _, err := s.clientv2.ProjectsApi.GetGroup(s.ctx, id).Execute()
	return result, err
}

func (s *Store) ProjectByName(name string) (*atlasv2.Group, error) {
	result, _, err := s.clientv2.ProjectsApi.GetGroupByName(s.ctx, name).Execute()
	return result, err
}

// CreateProject encapsulates the logic to manage different cloud providers.
func (s *Store) CreateProject(params *atlasv2.CreateGroupApiParams) (*atlasv2.Group, error) {
	result, _, err := s.clientv2.ProjectsApi.CreateGroupWithParams(s.ctx, params).Execute()
	return result, err
}

// UpdateProject encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateProject(params *atlasv2.UpdateGroupApiParams) (*atlasv2.Group, error) {
	result, _, err := s.clientv2.ProjectsApi.UpdateGroupWithParams(s.ctx, params).Execute()
	return result, err
}

// DeleteProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteProject(projectID string) error {
	_, err := s.clientv2.ProjectsApi.DeleteGroup(s.ctx, projectID).Execute()
	return err
}

// ProjectUsers lists all IAM users in a project.
func (s *Store) ProjectUsers(projectID string, opts *ListOptions) (*atlasv2.PaginatedGroupUser, error) {
	res := s.clientv2.MongoDBCloudUsersApi.ListGroupUsers(s.ctx, projectID)
	if opts != nil {
		res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// DeleteUserFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteUserFromProject(projectID, userID string) error {
	_, err := s.clientv2.MongoDBCloudUsersApi.RemoveGroupUser(s.ctx, projectID, userID).Execute()
	return err
}

// ProjectTeams encapsulates the logic to manage different cloud providers.
func (s *Store) ProjectTeams(projectID string, opts *ListOptions) (*atlasv2.PaginatedTeamRole, error) {
	res := s.clientv2.TeamsApi.
		ListGroupTeams(s.ctx, projectID)

	if opts != nil {
		res.
			IncludeCount(opts.IncludeCount).
			PageNum(opts.PageNum).
			ItemsPerPage(opts.ItemsPerPage)
	}

	result, _, err := res.Execute()
	return result, err
}

// AddTeamsToProject encapsulates the logic to manage different cloud providers.
func (s *Store) AddTeamsToProject(projectID string, teams []atlasv2.TeamRole) (*atlasv2.PaginatedTeamRole, error) {
	result, _, err := s.clientv2.TeamsApi.AddGroupTeams(s.ctx, projectID, &teams).Execute()
	return result, err
}

// DeleteTeamFromProject encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteTeamFromProject(projectID, teamID string) error {
	_, err := s.clientv2.TeamsApi.RemoveGroupTeam(s.ctx, projectID, teamID).Execute()
	return err
}

type MDBVersionListOptions struct {
	ListOptions
	CloudProvider *string
	InstanceSize  *string
	DefaultStatus *string
}

// MDBVersions encapsulates the logic to manage different cloud providers.
func (s *Store) MDBVersions(projectID string, opt *MDBVersionListOptions) (*atlasv2.PaginatedAvailableVersion, error) {
	req := s.clientv2.ProjectsApi.GetMongoDbVersions(s.ctx, projectID)

	if opt != nil {
		req = req.
			PageNum(opt.PageNum).
			ItemsPerPage(int64(opt.ItemsPerPage))
		if opt.CloudProvider != nil {
			req = req.CloudProvider(*opt.CloudProvider)
		}
		if opt.DefaultStatus != nil {
			req = req.DefaultStatus(*opt.DefaultStatus)
		}
		if opt.InstanceSize != nil {
			req = req.InstanceSize(*opt.InstanceSize)
		}
	}

	res, _, err := req.Execute()

	return res, err
}
