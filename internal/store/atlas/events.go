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

package atlas

import (
	atlas "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_events.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas OrganizationEventLister,ProjectEventLister,EventLister

type OrganizationEventLister interface {
	OrganizationEvents(opts *atlas.ListOrganizationEventsApiParams) (*atlas.OrgPaginatedEvent, error)
}

type ProjectEventLister interface {
	ProjectEvents(opts *atlas.ListProjectEventsApiParams) (*atlas.GroupPaginatedEvent, error)
}

type EventLister interface {
	OrganizationEventLister
	ProjectEventLister
}

// ProjectEvents encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectEvents(opts *atlas.ListProjectEventsApiParams) (*atlas.GroupPaginatedEvent, error) {
	event, err := atlas.NewEventTypeForNdsGroupFromValue(string(*opts.EventType))
	if err != nil {
		return nil, err
	}
	result, _, err := s.clientv2.EventsApi.ListProjectEvents(s.ctx, opts.GroupId).
		IncludeCount(*opts.IncludeCount).
		PageNum(*opts.PageNum).
		ItemsPerPage(*opts.ItemsPerPage).
		MaxDate(*opts.MaxDate).MinDate(*opts.MinDate).EventType(*event).Execute()
	return result, err
}

// OrganizationEvents encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationEvents(opts *atlas.ListOrganizationEventsApiParams) (*atlas.OrgPaginatedEvent, error) {
	event, err := atlas.NewEventTypeForOrgFromValue(string(*opts.EventType))
	if err != nil {
		return nil, err
	}
	result, _, err := s.clientv2.EventsApi.ListOrganizationEvents(s.ctx, opts.OrgId).
		IncludeCount(*opts.IncludeCount).
		PageNum(*opts.PageNum).
		ItemsPerPage(*opts.ItemsPerPage).
		MaxDate(*opts.MaxDate).MinDate(*opts.MinDate).EventType(*event).Execute()
	return result, err
}
