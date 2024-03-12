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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_events.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store OrganizationEventLister,ProjectEventLister,EventLister

type OrganizationEventLister interface {
	OrganizationEvents(string, *opsmngr.EventListOptions) (*opsmngr.EventResponse, error)
}

type ProjectEventLister interface {
	ProjectEvents(string, *opsmngr.EventListOptions) (*opsmngr.EventResponse, error)
}

type EventLister interface {
	OrganizationEventLister
	ProjectEventLister
}

// ProjectEvents encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectEvents(projectID string, opts *opsmngr.EventListOptions) (*opsmngr.EventResponse, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Events.ListProjectEvents(s.ctx, projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OrganizationEvents encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationEvents(orgID string, opts *opsmngr.EventListOptions) (*opsmngr.EventResponse, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Events.ListOrganizationEvents(s.ctx, orgID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
