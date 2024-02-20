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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_project_ip_access_lists.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectIPAccessListLister

type ProjectIPAccessListLister interface {
	ProjectIPAccessLists(string, *ListOptions) (*atlasv2.PaginatedNetworkAccess, error)
}

// ProjectIPAccessLists encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectIPAccessLists(projectID string, opts *ListOptions) (*atlasv2.PaginatedNetworkAccess, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		params := &atlasv2.ListProjectIpAccessListsApiParams{
			GroupId: projectID,
		}
		if opts != nil {
			params.ItemsPerPage = &opts.ItemsPerPage
			params.PageNum = &opts.PageNum
		}
		result, _, err := s.clientv2.ProjectIPAccessListApi.ListProjectIpAccessListsWithParams(s.ctx, params).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
