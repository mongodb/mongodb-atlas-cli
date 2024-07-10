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
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_events.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store OrganizationEventLister,ProjectEventLister,EventLister

type OrganizationEventLister interface {
	OrganizationEvents(opts *admin.ListOrganizationEventsApiParams) (*admin.OrgPaginatedEvent, error)
}

type ProjectEventLister interface {
	ProjectEvents(opts *admin.ListProjectEventsApiParams) (*admin.GroupPaginatedEvent, error)
}

type EventLister interface {
	OrganizationEventLister
	ProjectEventLister
}

// ProjectEvents encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectEvents(opts *admin.ListProjectEventsApiParams) (*admin.GroupPaginatedEvent, error) {
	result, _, err := s.clientv2.EventsApi.ListProjectEventsWithParams(s.ctx, opts).Execute()
	return result, err
}

// OrganizationEvents encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationEvents(opts *admin.ListOrganizationEventsApiParams) (*admin.OrgPaginatedEvent, error) {
	result, _, err := s.clientv2.EventsApi.ListOrganizationEventsWithParams(s.ctx, opts).Execute()
	return result, err
}
