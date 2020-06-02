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
	"github.com/mongodb/mongocli/internal/config"
)

// SnapshotRestoreJobs encapsulate the logic to manage different cloud providers
func (s *Store) SnapshotRestoreJobs(projectID, clusterName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshotRestoreJobs, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.List(context.Background(), o, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateSnapshotRestoreJobs encapsulate the logic to manage different cloud providers
func (s *Store) CreateSnapshotRestoreJobs(projectID, clusterName string, request *atlas.CloudProviderSnapshotRestoreJob) (*atlas.CloudProviderSnapshotRestoreJob, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.Create(context.Background(), o, request)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ContinuousSnapshots encapsulate the logic to manage different cloud providers
func (s *Store) CloudProviderSnapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshots, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), o, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CloudProviderSnapshot encapsulate the logic to manage different cloud providers
func (s *Store) CloudProviderSnapshot(projectID, clusterName, snapshotID string) (*atlas.CloudProviderSnapshot, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
		SnapshotID:  snapshotID,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), o)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateCloudProviderSnapshot encapsulate the logic to manage different cloud providers
func (s *Store) CreateCloudProviderSnapshot(projectID, clusterName string, req *atlas.CloudProviderSnapshot) (*atlas.CloudProviderSnapshot, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.Create(context.Background(), o, req)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteCloudProviderSnapshot encapsulate the logic to manage different cloud providers
func (s *Store) DeleteCloudProviderSnapshot(projectID, clusterName, snapshotID string) error {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
		SnapshotID:  snapshotID,
	}
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).CloudProviderSnapshots.Delete(context.Background(), o)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
