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
	"strconv"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

// ProcessDatabases encapsulate the logic to manage different cloud providers.
func (s *Store) ProcessDatabases(groupID, host string, port int, opts *ListOptions) (*atlasv2.PaginatedDatabase, error) {
	process := host + ":" + strconv.Itoa(port)
	result, _, err := s.clientv2.MonitoringAndLogsApi.ListDatabases(s.ctx, groupID, process).
		PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount).Execute()
	return result, err
}
