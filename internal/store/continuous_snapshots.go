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

type SnapshotsLister interface {
	ContinuousSnapshots(string, string, *atlas.ListOptions) (*atlas.ContinuousSnapshots, error)
}

type SnapshotDescriber interface {
	ContinuousSnapshot(string, string, string) (*atlas.ContinuousSnapshot, error)
}

type SnapshotsStore interface {
	SnapshotsLister
	SnapshotDescriber
}

// ProjectClusters encapsulate the logic to manage different cloud providers
func (s *Store) ContinuousSnapshots(projectID, clusterID string, opts *atlas.ListOptions) (*atlas.ContinuousSnapshots, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ContinuousSnapshots.List(context.Background(), projectID, clusterID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).ContinuousSnapshots.List(context.Background(), projectID, clusterID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Cluster encapsulate the logic to manage different cloud providers
func (s *Store) ContinuousSnapshot(projectID, clusterID, snapshotID string) (*atlas.ContinuousSnapshot, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ContinuousSnapshots.Get(context.Background(), projectID, clusterID, snapshotID)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).ContinuousSnapshots.Get(context.Background(), projectID, clusterID, snapshotID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
