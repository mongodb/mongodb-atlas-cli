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

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

// ProcessDisks encapsulates the logic to manage different cloud providers.
func (s *Store) ProcessDisks(groupID, host string, port int, opts *ListOptions) (*atlasv2.PaginatedDiskPartition, error) {
	processID := host + ":" + strconv.Itoa(port)
	result, _, err := s.clientv2.MonitoringAndLogsApi.ListDiskPartitions(s.ctx, groupID, processID).
		ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum).IncludeCount(opts.IncludeCount).Execute()
	return result, err
}
