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

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_organization_invitations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store OrganizationInvitationLister,OrganizationInvitationDeleter,OrganizationInvitationDescriber,OrganizationInvitationUpdater,OrganizationInviter

type OrganizationInvitationLister interface {
	OrganizationInvitations(*atlasv2.ListOrganizationInvitationsApiParams) ([]atlasv2.OrganizationInvitation, error)
}

type OrganizationInvitationDescriber interface {
	OrganizationInvitation(string, string) (*atlasv2.OrganizationInvitation, error)
}

type OrganizationInviter interface {
	InviteUser(string, *atlasv2.OrganizationInvitationRequest) (*atlasv2.OrganizationInvitation, error)
}

type OrganizationInvitationDeleter interface {
	DeleteInvitation(string, string) error
}

type OrganizationInvitationUpdater interface {
	UpdateOrganizationInvitation(string, string, *atlasv2.OrganizationInvitationRequest) (*atlasv2.OrganizationInvitation, error)
}

// OrganizationInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationInvitations(params *atlasv2.ListOrganizationInvitationsApiParams) ([]atlasv2.OrganizationInvitation, error) {
	result, _, err := s.clientv2.OrganizationsApi.ListOrganizationInvitationsWithParams(s.ctx, params).Execute()
	return result, err
}

// OrganizationInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationInvitation(orgID, invitationID string) (*atlasv2.OrganizationInvitation, error) {
	result, _, err := s.clientv2.OrganizationsApi.GetOrganizationInvitation(s.ctx, orgID, invitationID).Execute()
	return result, err
}

// DeleteInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteInvitation(orgID, invitationID string) error {
	_, _, err := s.clientv2.OrganizationsApi.DeleteOrganizationInvitation(s.ctx, orgID, invitationID).Execute()
	return err
}

// UpdateOrganizationInvitation encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateOrganizationInvitation(orgID, invitationID string, invitation *atlasv2.OrganizationInvitationRequest) (*atlasv2.OrganizationInvitation, error) {
	if invitationID != "" {
		invitationRequest := atlasv2.OrganizationInvitationUpdateRequest{
			Roles:   invitation.Roles,
			TeamIds: invitation.TeamIds,
		}

		result, _, err := s.clientv2.OrganizationsApi.UpdateOrganizationInvitationById(s.ctx, orgID,
			invitationID, &invitationRequest).Execute()
		return result, err
	}
	result, _, err := s.clientv2.OrganizationsApi.UpdateOrganizationInvitation(s.ctx, orgID, invitation).Execute()

	return result, err
}

// InviteUser encapsulates the logic to manage different cloud providers.
func (s *Store) InviteUser(orgID string, invitation *atlasv2.OrganizationInvitationRequest) (*atlasv2.OrganizationInvitation, error) {
	result, _, err := s.clientv2.OrganizationsApi.CreateOrganizationInvitation(s.ctx, orgID, invitation).Execute()

	return result, err
}
