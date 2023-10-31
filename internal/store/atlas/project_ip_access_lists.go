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
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_project_ip_access_lists.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ProjectIPAccessListLister

type ProjectIPAccessListLister interface {
	ProjectIPAccessLists(string, *ListOptions) (*atlasv2.PaginatedNetworkAccess, error)
}

// ProjectIPAccessLists encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectIPAccessLists(projectID string, opts *ListOptions) (*atlasv2.PaginatedNetworkAccess, error) {
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
