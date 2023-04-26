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
	request := s.clientv2.EventsApi.ListProjectEvents(s.ctx, opts.GroupId)
	// TODO CLOUDP-173460 - enable using api params directly
	if opts.IncludeCount != nil {
		request = request.IncludeCount(*opts.IncludeCount)
	}
	if opts.PageNum != nil {
		request = request.PageNum(*opts.PageNum)
	}
	if opts.ItemsPerPage != nil {
		request = request.ItemsPerPage(*opts.ItemsPerPage)
	}
	if opts.MaxDate != nil {
		request = request.MaxDate(*opts.MaxDate)
	}
	if opts.MinDate != nil {
		request = request.MinDate(*opts.MinDate)
	}
	if opts.EventType != nil {
		request = request.EventType(*opts.EventType)
	}

	result, _, err := request.Execute()
	return result, err
}

// OrganizationEvents encapsulate the logic to manage different cloud providers.
func (s *Store) OrganizationEvents(opts *atlas.ListOrganizationEventsApiParams) (*atlas.OrgPaginatedEvent, error) {
	request := s.clientv2.EventsApi.ListOrganizationEvents(s.ctx, opts.OrgId)
	if opts.IncludeCount != nil {
		request = request.IncludeCount(*opts.IncludeCount)
	}
	if opts.PageNum != nil {
		request = request.PageNum(*opts.PageNum)
	}
	if opts.ItemsPerPage != nil {
		request = request.ItemsPerPage(*opts.ItemsPerPage)
	}
	if opts.MaxDate != nil {
		request = request.MaxDate(*opts.MaxDate)
	}
	if opts.MinDate != nil {
		request = request.MinDate(*opts.MinDate)
	}
	if opts.EventType != nil {
		event, err := atlas.NewEventTypeForOrgFromValue(string(*opts.EventType))
		if err != nil {
			return nil, err
		}
		request = request.EventType(*event)
	}
	result, _, err := request.Execute()
	return result, err
}
