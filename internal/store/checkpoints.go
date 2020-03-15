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
	"context"
	"fmt"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
)

type CheckpointsGetter interface {
	Get(string, string, string) (*atlas.Checkpoint, error)
}

type CheckpointsLister interface {
	List(string, string, *atlas.ListOptions) (*atlas.Checkpoints, error)
}

type CheckpointsStore interface {
	CheckpointsGetter
	CheckpointsLister
}

// Get encapsulate the logic to manage different cloud providers
func (s *Store) Get(projectID, clusterID, checkpointID string) (*atlas.Checkpoint, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Checkpoints.Get(context.Background(), projectID, clusterID, checkpointID)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*om.Client).Checkpoints.Get(context.Background(), projectID, clusterID, checkpointID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// List encapsulate the logic to manage different cloud providers
func (s *Store) List(projectID, clusterID string, opts *atlas.ListOptions) (*atlas.Checkpoints, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Checkpoints.List(context.Background(), projectID, clusterID, opts)
		return result, err
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*om.Client).Checkpoints.List(context.Background(), projectID, clusterID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
