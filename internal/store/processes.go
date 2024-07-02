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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_processes.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ProcessLister,ProcessDescriber

type ProcessLister interface {
	Processes(*atlasv2.ListAtlasProcessesApiParams) (*atlasv2.PaginatedHostViewAtlas, error)
}

type ProcessDescriber interface {
	Process(*atlasv2.GetAtlasProcessApiParams) (*atlasv2.ApiHostViewAtlas, error)
}

// Process encapsulate the logic to manage different cloud providers.
func (s *Store) Process(params *atlasv2.GetAtlasProcessApiParams) (*atlasv2.ApiHostViewAtlas, error) {
	result, _, err := s.clientv2.MonitoringAndLogsApi.GetAtlasProcessWithParams(s.ctx, params).Execute()
	return result, err
}

// Processes encapsulate the logic to manage different cloud providers.
func (s *Store) Processes(params *atlasv2.ListAtlasProcessesApiParams) (*atlasv2.PaginatedHostViewAtlas, error) {
	result, _, err := s.clientv2.MonitoringAndLogsApi.ListAtlasProcessesWithParams(s.ctx, params).Execute()
	return result, err
}
