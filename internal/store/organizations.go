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
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_organizations.go -package=mocks github.com/mongodb/mongocli/internal/store OrganizationLister,OrganizationCreator,OrganizationDeleter,OrganizationDescriber,OrganizationInvitationLister,OrganizationInvitationDescriber,OrganizationInvitationUpdater

type OrganizationLister interface {
	Organizations(*atlas.OrganizationsListOptions) (*atlas.Organizations, error)
}

type OrganizationInvitationLister interface {
	OrganizationInvitations(string, *atlas.InvitationOptions) ([]*atlas.Invitation, error)
}

type OrganizationDescriber interface {
	Organization(string) (*atlas.Organization, error)
}

type OrganizationInvitationDescriber interface {
	OrganizationInvitation(string, string) (*atlas.Invitation, error)
}

type OrganizationCreator interface {
	CreateOrganization(string) (*atlas.Organization, error)
}

type OrganizationDeleter interface {
	DeleteOrganization(string) error
}

type OrganizationInvitationUpdater interface {
	UpdateOrganizationInvitation(string, string, *atlas.Invitation) (*atlas.Invitation, error)
}

// Organizations encapsulate the logic to manage different cloud providers.
func (s *Store) Organizations(opts *atlas.OrganizationsListOptions) (*atlas.Organizations, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Organizations.List(context.Background(), opts)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.List(context.Background(), opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Organization encapsulate the logic to manage different cloud providers.
func (s *Store) Organization(id string) (*atlas.Organization, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Organizations.Get(context.Background(), id)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.Get(context.Background(), id)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) CreateOrganization(name string) (*atlas.Organization, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		org := &atlas.Organization{Name: name}
		result, _, err := s.client.(*opsmngr.Client).Organizations.Create(context.Background(), org)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteOrganization encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteOrganization(id string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Organizations.Delete(context.Background(), id)
		return err
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*opsmngr.Client).Organizations.Delete(context.Background(), id)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OrganizationInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationInvitations(orgID string, opts *atlas.InvitationOptions) ([]*atlas.Invitation, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Organizations.Invitations(context.Background(), orgID, opts)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.Invitations(context.Background(), orgID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OrganizationInvitations encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateOrganizationInvitation(orgID, invitationID string, invitation *atlas.Invitation) (*atlas.Invitation, error) {
	switch s.service {
	case config.CloudService:
		if invitationID != "" {
			result, _, err := s.client.(*atlas.Client).Organizations.UpdateInvitationByID(context.Background(), orgID, invitationID, invitation)
			return result, err
		}
		result, _, err := s.client.(*atlas.Client).Organizations.UpdateInvitation(context.Background(), orgID, invitation)
		return result, err

	case config.CloudManagerService, config.OpsManagerService:
		if invitationID != "" {
			result, _, err := s.client.(*opsmngr.Client).Organizations.UpdateInvitationByID(context.Background(), orgID, invitationID, invitation)
			return result, err
		}
		result, _, err := s.client.(*opsmngr.Client).Organizations.UpdateInvitation(context.Background(), orgID, invitation)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OrganizationInvitation encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationInvitation(orgID, invitationID string) (*atlas.Invitation, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Organizations.Invitation(context.Background(), orgID, invitationID)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*opsmngr.Client).Organizations.Invitation(context.Background(), orgID, invitationID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
