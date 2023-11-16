// Copyright 2021 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115001/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_serverless_instances.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ServerlessInstanceLister,ServerlessInstanceDescriber

type ServerlessInstanceLister interface {
	ServerlessInstances(string, *ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error)
}

type ServerlessInstanceDescriber interface {
	GetServerlessInstance(string, string) (*atlasv2.ServerlessInstanceDescription, error)
}

// ServerlessInstances encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessInstances(projectID string, listOps *ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.ServerlessInstancesApi.ListServerlessInstances(s.ctx, projectID).
			ItemsPerPage(listOps.ItemsPerPage).
			PageNum(listOps.PageNum).
			Execute()

		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessInstance encapsulates the logic to manage different cloud providers.
func (s *Store) GetServerlessInstance(projectID, clusterName string) (*atlasv2.ServerlessInstanceDescription, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.ServerlessInstancesApi.GetServerlessInstance(s.ctx, projectID, clusterName).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
