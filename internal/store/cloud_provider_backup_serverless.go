// Copyright 2023 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_backup_serverless.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ServerlessSnapshotsLister,ServerlessSnapshotsDescriber,ServerlessRestoreJobsLister,ServerlessRestoreJobsDescriber,ServerlessRestoreJobsCreator

type ServerlessSnapshotsLister interface {
	ServerlessSnapshots(string, string, *atlas.ListOptions) (*atlas.CloudProviderSnapshots, error)
}

type ServerlessSnapshotsDescriber interface {
	ServerlessSnapshot(string, string, string) (*atlas.CloudProviderSnapshot, error)
}

type ServerlessRestoreJobsLister interface {
	ServerlessRestoreJobs(string, string, *atlas.ListOptions) (*atlas.CloudProviderSnapshotRestoreJobs, error)
}

type ServerlessRestoreJobsDescriber interface {
	ServerlessRestoreJob(string, string, string) (*atlas.CloudProviderSnapshotRestoreJob, error)
}

type ServerlessRestoreJobsCreator interface {
	ServerlessCreateRestoreJobs(string, string, *atlas.CloudProviderSnapshotRestoreJob) (*atlas.CloudProviderSnapshotRestoreJob, error)
}

// ServerlessSnapshots encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessSnapshots(projectID, clusterName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshots, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:      projectID,
		InstanceName: clusterName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetAllServerlessSnapshots(s.ctx, o, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessSnapshot(projectID, instanceName, snapshotID string) (*atlas.CloudProviderSnapshot, error) {
	o := &atlas.SnapshotReqPathParameters{
		GroupID:      projectID,
		SnapshotID:   snapshotID,
		InstanceName: instanceName,
	}
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshots.GetOneServerlessSnapshot(s.ctx, o)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessRestoreJobs(projectID, instanceName string, opts *atlas.ListOptions) (*atlas.CloudProviderSnapshotRestoreJobs, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.ListForServerlessBackupRestore(s.ctx, projectID, instanceName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessRestoreJob(projectID, instanceName string, jobID string) (*atlas.CloudProviderSnapshotRestoreJob, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.GetForServerlessBackupRestore(s.ctx, projectID, instanceName, jobID)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessCreateRestoreJobs(projectID, clusterName string, request *atlas.CloudProviderSnapshotRestoreJob) (*atlas.CloudProviderSnapshotRestoreJob, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).CloudProviderSnapshotRestoreJobs.CreateForServerlessBackupRestore(s.ctx, projectID, clusterName, request)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
