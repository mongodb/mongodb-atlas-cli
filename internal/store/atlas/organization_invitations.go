// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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

//go:generate mockgen -destination=../../mocks/atlas/mock_organization_invitations.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas OrganizationInvitationLister,OrganizationInvitationDeleter,OrganizationInvitationDescriber,OrganizationInvitationUpdater,OrganizationInviter

type OrganizationInvitationLister interface {
	OrganizationInvitations(string, *atlas.InvitationOptions) (interface{}, error)
}

type OrganizationInvitationDescriber interface {
	OrganizationInvitation(string, string) (interface{}, error)
}

type OrganizationInviter interface {
	InviteUser(string, *atlas.Invitation) (interface{}, error)
}

type OrganizationInvitationDeleter interface {
	DeleteInvitation(string, string) error
}

type OrganizationInvitationUpdater interface {
	UpdateOrganizationInvitation(string, string, *atlas.Invitation) (interface{}, error)
}

// OrganizationInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationInvitations(orgID string, opts *atlas.InvitationOptions) (interface{}, error) {
	res := s.clientv2.OrganizationsApi.ListOrganizationInvitations(s.ctx, orgID)
	if opts != nil {
		res = res.Username(opts.Username)
	}
	result, _, err := res.Execute()
	return result, err
}

// OrganizationInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationInvitation(orgID, invitationID string) (interface{}, error) {
	result, _, err := s.clientv2.OrganizationsApi.GetOrganizationInvitation(s.ctx, orgID, invitationID).Execute()
	return result, err
}

// DeleteInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteInvitation(orgID, invitationID string) error {
	_, _, err := s.clientv2.OrganizationsApi.DeleteOrganizationInvitation(s.ctx, orgID, invitationID).Execute()
	return err
}

// UpdateOrganizationInvitation encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateOrganizationInvitation(orgID, invitationID string, invitation *atlas.Invitation) (interface{}, error) {
	if invitationID != "" {
		invitationRequest := atlasv2.OrganizationInvitationUpdateRequest{
			Roles:   invitation.Roles,
			TeamIds: invitation.TeamIDs,
		}

		result, _, err := s.clientv2.OrganizationsApi.UpdateOrganizationInvitationById(s.ctx, orgID, invitationID).OrganizationInvitationUpdateRequest(invitationRequest).Execute()
		return result, err
	}
	invitationRequest := mapInvitation(invitation)
	result, _, err := s.clientv2.OrganizationsApi.UpdateOrganizationInvitation(s.ctx, orgID).OrganizationInvitationRequest(invitationRequest).Execute()

	return result, err
}

// InviteUser encapsulates the logic to manage different cloud providers.
func (s *Store) InviteUser(orgID string, invitation *atlas.Invitation) (interface{}, error) {
	invitationRequest := mapInvitation(invitation)
	result, _, err := s.clientv2.OrganizationsApi.CreateOrganizationInvitation(s.ctx, orgID).OrganizationInvitationRequest(invitationRequest).Execute()

	return result, err
}

func mapInvitation(invitation *atlas.Invitation) atlasv2.OrganizationInvitationRequest {
	return atlasv2.OrganizationInvitationRequest{
		Roles:    invitation.Roles,
		TeamIds:  invitation.TeamIDs,
		Username: &invitation.Username,
	}
}
