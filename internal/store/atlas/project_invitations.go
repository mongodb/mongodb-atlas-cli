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
)

//go:generate mockgen -destination=../mocks/atlas/mock_project_invitations.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectInvitationLister,ProjectInvitationDescriber,ProjectInvitationDeleter,ProjectInviter,ProjectInvitationUpdater

type ProjectInvitationLister interface {
	ProjectInvitations(string, *atlas.InvitationOptions) ([]*atlas.Invitation, error)
}

type ProjectInvitationDescriber interface {
	ProjectInvitation(string, string) (*atlas.Invitation, error)
}

type ProjectInviter interface {
	InviteUserToProject(string, *atlas.Invitation) (*atlas.Invitation, error)
}

type ProjectInvitationDeleter interface {
	DeleteProjectInvitation(string, string) error
}

type ProjectInvitationUpdater interface {
	UpdateProjectInvitation(string, string, *atlas.Invitation) (*atlas.Invitation, error)
}

// ProjectInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitations(groupID string, opts *atlas.InvitationOptions) ([]*atlas.Invitation, error) {
	result, _, err := s.client.Projects.Invitations(s.ctx, groupID, opts)
	return result, err
}

// ProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitation(groupID, invitationID string) (*atlas.Invitation, error) {
	result, _, err := s.client.Projects.Invitation(s.ctx, groupID, invitationID)
	return result, err
}

// DeleteProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteProjectInvitation(groupID, invitationID string) error {
	_, err := s.client.Projects.DeleteInvitation(s.ctx, groupID, invitationID)
	return err
}

// InviteUserToProject encapsulate the logic to manage different cloud providers.
func (s *Store) InviteUserToProject(groupID string, invitation *atlas.Invitation) (*atlas.Invitation, error) {
	result, _, err := s.client.Projects.InviteUser(s.ctx, groupID, invitation)
	return result, err
}

// UpdateProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateProjectInvitation(groupID, invitationID string, invitation *atlas.Invitation) (*atlas.Invitation, error) {
	if invitationID != "" {
		result, _, err := s.client.Projects.UpdateInvitationByID(s.ctx, groupID, invitationID, invitation)
		return result, err
	}
	result, _, err := s.client.Projects.UpdateInvitation(s.ctx, groupID, invitation)
	return result, err
}
