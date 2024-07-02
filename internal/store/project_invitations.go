// Copyright 2021 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_project_invitations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ProjectInvitationLister,ProjectInvitationDescriber,ProjectInvitationDeleter,ProjectInviter,ProjectInvitationUpdater

type ProjectInvitationLister interface {
	ProjectInvitations(*atlasv2.ListProjectInvitationsApiParams) ([]atlasv2.GroupInvitation, error)
}

type ProjectInvitationDescriber interface {
	ProjectInvitation(string, string) (*atlasv2.GroupInvitation, error)
}

type ProjectInviter interface {
	InviteUserToProject(string, *atlasv2.GroupInvitationRequest) (*atlasv2.GroupInvitation, error)
}

type ProjectInvitationDeleter interface {
	DeleteProjectInvitation(string, string) error
}

type ProjectInvitationUpdater interface {
	UpdateProjectInvitation(string, string, *atlasv2.GroupInvitationRequest) (*atlasv2.GroupInvitation, error)
}

// ProjectInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitations(params *atlasv2.ListProjectInvitationsApiParams) ([]atlasv2.GroupInvitation, error) {
	result, _, err := s.clientv2.ProjectsApi.ListProjectInvitationsWithParams(s.ctx, params).Execute()
	return result, err
}

// ProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitation(groupID, invitationID string) (*atlasv2.GroupInvitation, error) {
	result, _, err := s.clientv2.ProjectsApi.GetProjectInvitation(s.ctx, groupID, invitationID).Execute()
	return result, err
}

// DeleteProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteProjectInvitation(groupID, invitationID string) error {
	_, _, err := s.clientv2.ProjectsApi.DeleteProjectInvitation(s.ctx, groupID, invitationID).Execute()
	return err
}

// InviteUserToProject encapsulate the logic to manage different cloud providers.
func (s *Store) InviteUserToProject(groupID string, invitation *atlasv2.GroupInvitationRequest) (*atlasv2.GroupInvitation, error) {
	result, _, err := s.clientv2.ProjectsApi.CreateProjectInvitation(s.ctx, groupID, invitation).Execute()
	return result, err
}

// UpdateProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateProjectInvitation(groupID, invitationID string, invitation *atlasv2.GroupInvitationRequest) (*atlasv2.GroupInvitation, error) {
	if invitationID != "" {
		groupInvitationRequest := atlasv2.GroupInvitationUpdateRequest{
			Roles: invitation.Roles,
		}
		result, _, err := s.clientv2.ProjectsApi.UpdateProjectInvitationById(s.ctx, groupID, invitationID, &groupInvitationRequest).Execute()
		return result, err
	}

	result, _, err := s.clientv2.ProjectsApi.UpdateProjectInvitation(s.ctx, groupID, invitation).Execute()
	return result, err
}
