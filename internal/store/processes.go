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
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../mocks/mock_processes.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ProcessLister,ProcessDescriber

type ProcessLister interface {
	Processes(string, *atlas.ProcessesListOptions) (*atlasv2.PaginatedHostViewAtlas, error)
}

type ProcessDescriber interface {
	Process(string, string, int) (*atlasv2.HostViewAtlas, error)
}

// Process encapsulate the logic to manage different cloud providers.
func (s *Store) Process(groupID, hostname string, port int) (*atlasv2.HostViewAtlas, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		processID := hostname + strconv.Itoa(port)
		result, _, err := s.clientv2.MonitoringAndLogsApi.GetAtlasProcess(s.ctx, groupID, processID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Processes encapsulate the logic to manage different cloud providers.
func (s *Store) Processes(groupID string, opts *atlas.ProcessesListOptions) (*atlasv2.PaginatedHostViewAtlas, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.MonitoringAndLogsApi.ListAtlasProcesses(s.ctx, groupID).PageNum(int32(opts.PageNum)).ItemsPerPage(int32(opts.ItemsPerPage)).IncludeCount(opts.IncludeCount).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
