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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_project_invitations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store ProjectInvitationLister,ProjectInvitationDescriber,ProjectInvitationDeleter,ProjectInviter,ProjectInvitationUpdater

type ProjectInvitationLister interface {
	ProjectInvitations(string, *opsmngr.InvitationOptions) ([]*opsmngr.Invitation, error)
}

type ProjectInvitationDescriber interface {
	ProjectInvitation(string, string) (*opsmngr.Invitation, error)
}

type ProjectInviter interface {
	InviteUserToProject(string, *opsmngr.Invitation) (*opsmngr.Invitation, error)
}

type ProjectInvitationDeleter interface {
	DeleteProjectInvitation(string, string) error
}

type ProjectInvitationUpdater interface {
	UpdateProjectInvitation(string, string, *opsmngr.Invitation) (*opsmngr.Invitation, error)
}

// ProjectInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitations(groupID string, opts *opsmngr.InvitationOptions) ([]*opsmngr.Invitation, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.Projects.Invitations(s.ctx, groupID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectInvitation(groupID, invitationID string) (*opsmngr.Invitation, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.Projects.Invitation(s.ctx, groupID, invitationID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteProjectInvitation(groupID, invitationID string) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.Projects.DeleteInvitation(s.ctx, groupID, invitationID)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// InviteUserToProject encapsulate the logic to manage different cloud providers.
func (s *Store) InviteUserToProject(groupID string, invitation *opsmngr.Invitation) (*opsmngr.Invitation, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.Projects.InviteUser(s.ctx, groupID, invitation)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UpdateProjectInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateProjectInvitation(groupID, invitationID string, invitation *opsmngr.Invitation) (*opsmngr.Invitation, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		if invitationID != "" {
			result, _, err := s.client.Projects.UpdateInvitationByID(s.ctx, groupID, invitationID, invitation)
			return result, err
		}
		result, _, err := s.client.Projects.UpdateInvitation(s.ctx, groupID, invitation)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
