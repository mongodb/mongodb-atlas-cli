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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_project_ip_access_lists.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ProjectIPAccessListDescriber,ProjectIPAccessListLister,ProjectIPAccessListCreator,ProjectIPAccessListDeleter

type ProjectIPAccessListDescriber interface {
	IPAccessList(string, string) (*atlasv2.NetworkPermissionEntry, error)
}
type ProjectIPAccessListLister interface {
	ProjectIPAccessLists(string, *atlas.ListOptions) (*atlasv2.PaginatedNetworkAccess, error)
}

type ProjectIPAccessListCreator interface {
	CreateProjectIPAccessList([]*atlasv2.NetworkPermissionEntry) (*atlasv2.PaginatedNetworkAccess, error)
}

type ProjectIPAccessListDeleter interface {
	DeleteProjectIPAccessList(string, string) error
}

// CreateProjectIPAccessList encapsulate the logic to manage different cloud providers.
func (s *Store) CreateProjectIPAccessList(entries []*atlasv2.NetworkPermissionEntry) (*atlasv2.PaginatedNetworkAccess, error) {
	if len(entries) == 0 {
		return nil, errors.New("no entries")
	}
	switch s.service {
	case config.CloudService, config.CloudGovService:
		entry := make([]atlasv2.NetworkPermissionEntry, len(entries))
		for i, ptr := range entries {
			entry[i] = *ptr
		}

		result, _, err := s.clientv2.ProjectIPAccessListApi.CreateProjectIpAccessList(s.ctx, entries[0].GetGroupId(), &entry).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteProjectIPAccessList encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteProjectIPAccessList(projectID, entry string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.ProjectIPAccessListApi.DeleteProjectIpAccessList(s.ctx, projectID, entry).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ProjectIPAccessLists encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectIPAccessLists(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedNetworkAccess, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.ProjectIPAccessListApi.ListProjectIpAccessLists(s.ctx, projectID)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
		}
		result, _, err := res.Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// IPAccessList encapsulate the logic to manage different cloud providers.
func (s *Store) IPAccessList(projectID, name string) (*atlasv2.NetworkPermissionEntry, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.ProjectIPAccessListApi.GetProjectIpList(s.ctx, projectID, name).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
