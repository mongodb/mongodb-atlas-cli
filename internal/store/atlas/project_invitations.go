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

package atlas

import (
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_project_invitations.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectInvitationLister,ProjectInvitationDescriber,ProjectInvitationDeleter,ProjectInviter,ProjectInvitationUpdater

type ProjectInvitationLister interface {
	ProjectInvitations(string, *atlas.InvitationOptions) ([]mongodbatlasv2.GroupInvitation, error)
}

type ProjectInvitationDescriber interface {
	ProjectInvitation(string, string) (*mongodbatlasv2.GroupInvitation, error)
}

type ProjectInviter interface {
	InviteUserToProject(string, *atlas.Invitation) (*mongodbatlasv2.GroupInvitation, error)
}

type ProjectInvitationDeleter interface {
	DeleteProjectInvitation(string, string) error
}

type ProjectInvitationUpdater interface {
	UpdateProjectInvitation(string, string, *atlas.Invitation) (*mongodbatlasv2.GroupInvitation, error)
}

// ProjectInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitations(groupID string, opts *atlas.InvitationOptions) ([]mongodbatlasv2.GroupInvitation, error) {
	result, _, err := s.clientv2.ProjectsApi.ListProjectInvitations(s.ctx, groupID).Username(opts.Username).Execute()
	return result, err
}

// ProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitation(groupID, invitationID string) (*mongodbatlasv2.GroupInvitation, error) {
	result, _, err := s.clientv2.ProjectsApi.GetProjectInvitation(s.ctx, groupID, invitationID).Execute()
	return result, err
}

// DeleteProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteProjectInvitation(groupID, invitationID string) error {
	_, _, err := s.clientv2.ProjectsApi.DeleteProjectInvitation(s.ctx, groupID, invitationID).Execute()
	return err
}

// InviteUserToProject encapsulate the logic to manage different cloud providers.
func (s *Store) InviteUserToProject(groupID string, invitation *atlas.Invitation) (*mongodbatlasv2.GroupInvitation, error) {
	groupInvitationRequest := mongodbatlasv2.GroupInvitationRequest{
		Username: &invitation.Username,
		Roles:    invitation.Roles,
	}
	result, _, err := s.clientv2.ProjectsApi.CreateProjectInvitation(s.ctx, groupID).GroupInvitationRequest(groupInvitationRequest).Execute()
	return result, err
}

// UpdateProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateProjectInvitation(groupID, invitationID string, invitation *atlas.Invitation) (*mongodbatlasv2.GroupInvitation, error) {
	if invitationID != "" {
		groupInvitationRequest := mongodbatlasv2.GroupInvitationUpdateRequest{
			Roles: invitation.Roles,
		}
		result, _, err := s.clientv2.ProjectsApi.UpdateProjectInvitationById(s.ctx, groupID, invitationID).GroupInvitationUpdateRequest(groupInvitationRequest).Execute()
		return result, err
	}
	groupInvitationRequest := mongodbatlasv2.GroupInvitationRequest{
		Username: &invitation.Username,
		Roles:    invitation.Roles,
	}
	result, _, err := s.clientv2.ProjectsApi.UpdateProjectInvitation(s.ctx, groupID).GroupInvitationRequest(groupInvitationRequest).Execute()
	return result, err
}
