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

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
)

type OrganizationEventLister interface {
	OrganizationEvents(string, *atlas.EventListOptions) (*atlas.EventResponse, error)
}

type ProjectEventLister interface {
	ProjectEvents(string, *atlas.ListOptions) (*atlas.EventResponse, error)
}

type EventsStore interface {
	OrganizationEventLister
	ProjectEventLister
}

// ProjectEvents encapsulate the logic to manage different cloud providers
func (s *Store) ProjectEvents(projectID string, opts *atlas.ListOptions) (*atlas.EventResponse, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Events.ListProjectEvents(context.Background(), projectID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).Events.ListProjectEvents(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// OrganizationEvents encapsulate the logic to manage different cloud providers
func (s *Store) OrganizationEvents(projectID string, opts *atlas.EventListOptions) (*atlas.EventResponse, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Events.ListOrganizationEvents(context.Background(), projectID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).Events.ListOrganizationEvents(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
